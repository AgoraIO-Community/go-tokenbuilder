package accesstoken

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/AgoraIO-Community/go-tokenbuilder/internal/testutil"
)

var errInjectedWriter = errors.New("injected writer failure")

type countingWriter struct {
	writes int
}

func (writer *countingWriter) Write(data []byte) (int, error) {
	writer.writes++
	return len(data), nil
}

type failAfterWriter struct {
	failAfter int
	writes    int
}

func (writer *failAfterWriter) Write(data []byte) (int, error) {
	if writer.writes == writer.failAfter {
		return 0, errInjectedWriter
	}
	writer.writes++
	return len(data), nil
}

func TestLegacyServicesRoundTrip(t *testing.T) {
	token := NewAccessToken(DataMockAppId, DataMockAppCertificate, DataMockExpire)
	token.IssueTs = DataMockIssueTs
	token.Salt = DataMockSalt

	rtc := NewServiceRtc(DataMockChannelName, DataMockUidStr)
	rtc.AddPrivilege(PrivilegeJoinChannel, DataMockExpire)
	token.AddService(rtc)

	rtm := NewServiceRtm(DataMockUserId)
	rtm.AddPrivilege(PrivilegeLogin, DataMockExpire)
	token.AddService(rtm)

	fpa := NewServiceFpa()
	fpa.AddPrivilege(PrivilegeLogin, DataMockExpire)
	token.AddService(fpa)

	chat := NewServiceChat(DataMockUserId)
	chat.AddPrivilege(PrivilegeChatUser, DataMockExpire)
	token.AddService(chat)

	education := NewServiceEducation("room", "user", 1)
	education.AddPrivilege(PrivilegeEducationRoomUser, DataMockExpire)
	token.AddService(education)

	built, err := token.Build()
	testutil.Nil(t, err)

	parsed := CreateAccessToken()
	ok, err := parsed.Parse(built)
	testutil.Nil(t, err)
	testutil.Equal(t, true, ok)
	testutil.Equal(t, 5, len(parsed.Services))
	testutil.Equal(t, DataMockExpire, parsed.Services[ServiceTypeFpa].(*ServiceFpa).Privileges[PrivilegeLogin])
	parsedEducation := parsed.Services[ServiceTypeEducation].(*ServiceEducation)
	testutil.Equal(t, "room", parsedEducation.RoomUuid)
	testutil.Equal(t, "user", parsedEducation.UserUuid)
	testutil.Equal(t, int16(1), parsedEducation.Role)
}

func TestLegacyServiceSerializationFailures(t *testing.T) {
	tests := []struct {
		name    string
		service IService
		fresh   func() IService
	}{
		{"RTC", privileged(NewServiceRtc("channel", "uid")), func() IService { return NewServiceRtc("", "") }},
		{"RTM", privileged(NewServiceRtm("user")), func() IService { return NewServiceRtm("") }},
		{"FPA", privileged(NewServiceFpa()), func() IService { return NewServiceFpa() }},
		{"Chat", privileged(NewServiceChat("user")), func() IService { return NewServiceChat("") }},
		{"Education", privileged(NewServiceEducation("room", "user", 1)), func() IService { return NewServiceEducation("", "", -1) }},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			counter := new(countingWriter)
			testutil.Nil(t, test.service.Pack(counter))
			for failAfter := 0; failAfter < counter.writes; failAfter++ {
				err := test.service.Pack(&failAfterWriter{failAfter: failAfter})
				if !errors.Is(err, errInjectedWriter) {
					t.Fatalf("Pack failure after write %d = %v", failAfter, err)
				}
			}

			var packed bytes.Buffer
			testutil.Nil(t, test.service.Pack(&packed))
			body := packed.Bytes()[2:]
			for cut := 0; cut < len(body); cut++ {
				if err := test.fresh().UnPack(bytes.NewReader(body[:cut])); err == nil {
					t.Fatalf("UnPack unexpectedly accepted %d of %d bytes", cut, len(body))
				}
			}
		})
	}
}

func TestLegacyTokenRejectsEveryTruncatedPayload(t *testing.T) {
	token := NewAccessToken(DataMockAppId, DataMockAppCertificate, DataMockExpire)
	token.IssueTs = DataMockIssueTs
	token.Salt = DataMockSalt
	token.AddService(privileged(NewServiceRtc(DataMockChannelName, DataMockUidStr)))
	token.AddService(privileged(NewServiceRtm(DataMockUserId)))
	token.AddService(privileged(NewServiceFpa()))
	token.AddService(privileged(NewServiceChat(DataMockUserId)))
	token.AddService(privileged(NewServiceEducation("room", "user", 1)))

	built, err := token.Build()
	testutil.Nil(t, err)
	compressed, err := base64DecodeStr(built[VersionLength:])
	testutil.Nil(t, err)
	raw, err := decompressZlibWithError(compressed)
	testutil.Nil(t, err)

	for cut := 0; cut < len(raw); cut++ {
		truncated := Version + base64EncodeStr(compressZlib(raw[:cut]))
		if ok, err := CreateAccessToken().Parse(truncated); err == nil || ok {
			t.Fatalf("Parse unexpectedly accepted %d of %d payload bytes", cut, len(raw))
		}
	}
}

func TestLegacyErrorAndUtilityEdges(t *testing.T) {
	token := NewAccessToken(DataMockAppId, DataMockAppCertificate, DataMockExpire)
	token.Services[ServiceTypeRtc] = &failingService{Service: NewService(ServiceTypeRtc)}
	if _, err := token.Build(); !errors.Is(err, errInjectedWriter) {
		t.Fatalf("Build error = %v", err)
	}

	parsed := CreateAccessToken()
	if valid, err := parsed.VerifySignature(DataMockAppCertificate); err == nil || valid {
		t.Fatal("VerifySignature unexpectedly accepted an unparsed token")
	}
	testutil.Equal(t, IService(nil), parsed.newService(99))

	var unknownPayload bytes.Buffer
	testutil.Nil(t, packString(&unknownPayload, "signature"))
	testutil.Nil(t, packString(&unknownPayload, DataMockAppId))
	testutil.Nil(t, packUint32(&unknownPayload, DataMockIssueTs))
	testutil.Nil(t, packUint32(&unknownPayload, DataMockExpire))
	testutil.Nil(t, packUint32(&unknownPayload, DataMockSalt))
	testutil.Nil(t, packUint16(&unknownPayload, 1))
	testutil.Nil(t, packUint16(&unknownPayload, 99))
	unknownToken := Version + base64EncodeStr(compressZlib(unknownPayload.Bytes()))
	if ok, err := CreateAccessToken().Parse(unknownToken); err != nil || !ok {
		t.Fatalf("Parse unknown service = %t, %v", ok, err)
	}

	testutil.Equal(t, "5d41402abc4b2a76b9719d911017c592", Md5("hello"))

	corrupt := compressZlib([]byte("hello"))
	corrupt[len(corrupt)-1] ^= 0xff
	if _, err := decompressZlibWithError(corrupt); err == nil {
		t.Fatal("expected corrupt zlib checksum to fail")
	}
}

type failingService struct {
	*Service
}

func (service *failingService) Pack(io.Writer) error {
	return errInjectedWriter
}

func privileged(service interface {
	IService
	AddPrivilege(uint16, uint32)
}) IService {
	service.AddPrivilege(1, DataMockExpire)
	return service
}
