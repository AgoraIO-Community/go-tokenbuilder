package accesstoken2

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

func TestAccessToken2ServiceSerializationFailures(t *testing.T) {
	permissions := NewRtm2Permissions()
	permissions.Add(Rtm2ResourceMessageChannels, Rtm2PermissionRead, []string{"channel"})

	tests := []struct {
		name    string
		service IService
		fresh   func() IService
	}{
		{"RTC", privileged(NewServiceRtc("channel", "uid")), func() IService { return NewServiceRtc("", "") }},
		{"RTM", privileged(NewServiceRtm("user")), func() IService { return NewServiceRtm("") }},
		{"Streaming", privileged(NewServiceStreaming("channel", "account")), func() IService { return NewServiceStreaming("", "") }},
		{"FPA", privileged(NewServiceFpa()), func() IService { return NewServiceFpa() }},
		{"Chat", privileged(NewServiceChat("user")), func() IService { return NewServiceChat("") }},
		{"FCDN", privileged(NewServiceFCdn("channel", "account")), func() IService { return NewServiceFCdn("", "") }},
		{"APaaS", privileged(NewServiceApaas("room", "user", 1)), func() IService { return NewServiceApaas("", "", -1) }},
		{"RTM2", privileged(NewServiceRtm2("user", permissions)), func() IService { return NewServiceRtm2("", nil) }},
		{"ConvoAI", privileged(NewServiceConvoAI()), func() IService { return NewServiceConvoAI() }},
		{"STT", privileged(NewServiceStt()), func() IService { return NewServiceStt() }},
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

func TestAccessToken2RejectsEveryTruncatedPayload(t *testing.T) {
	token := fullAccessToken2()
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

func TestAccessToken2CompatibilityEdges(t *testing.T) {
	testutil.Equal(t, DataMockUidStr, NewServiceStreamingWithUid(DataMockChannelName, DataMockUid).Account)
	testutil.Equal(t, DataMockUidStr, NewServiceFCdnWithUid(DataMockChannelName, DataMockUid).Account)
	testutil.Equal(t, "", NewServiceStreamingWithUid(DataMockChannelName, 0).Account)
	testutil.Equal(t, "", NewServiceFCdnWithUid(DataMockChannelName, 0).Account)

	token := &AccessToken{}
	token.AddService(NewServiceConvoAI())
	testutil.Equal(t, 1, len(token.Services))

	direct := NewAccessToken(DataMockAppId, DataMockAppCertificate, DataMockExpire)
	direct.Services[ServiceTypeRtm] = NewServiceRtm(DataMockUserId)
	testutil.Equal(t, 1, len(direct.GetServices(ServiceTypeRtm)))
	testutil.Equal(t, 0, len(direct.GetServices(ServiceTypeRtc)))

	mixed := NewAccessToken(DataMockAppId, DataMockAppCertificate, DataMockExpire)
	mixed.AddService(NewServiceRtc("first", "user"))
	mixed.Services[ServiceTypeRtc] = NewServiceRtc("replacement", "user")
	mixed.Services[ServiceTypeRtm] = NewServiceRtm("rtm-user")
	services := mixed.servicesForPacking()
	testutil.Equal(t, 2, len(services))
	testutil.Equal(t, "replacement", services[0].(*ServiceRtc).ChannelName)
	testutil.Equal(t, "rtm-user", services[1].(*ServiceRtm).UserId)

	unknown := NewAccessToken(DataMockAppId, DataMockAppCertificate, DataMockExpire)
	unknown.IssueTs = DataMockIssueTs
	unknown.Salt = DataMockSalt
	unknown.AddService(NewService(99))
	built, err := unknown.Build()
	testutil.Nil(t, err)
	parsed := CreateAccessToken()
	ok, err := parsed.Parse(built)
	testutil.Nil(t, err)
	testutil.Equal(t, true, ok)
	testutil.Equal(t, 0, len(parsed.Services))
}

func TestAccessToken2ErrorAndUtilityEdges(t *testing.T) {
	token := NewAccessToken(DataMockAppId, DataMockAppCertificate, DataMockExpire)
	token.AddService(&failingService{Service: NewService(ServiceTypeRtc)})
	if _, err := token.Build(); !errors.Is(err, errInjectedWriter) {
		t.Fatalf("Build error = %v", err)
	}

	parsed := CreateAccessToken()
	if valid, err := parsed.VerifySignature(DataMockAppCertificate); err == nil || valid {
		t.Fatal("VerifySignature unexpectedly accepted an unparsed token")
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

func fullAccessToken2() *AccessToken {
	token := NewAccessToken(DataMockAppId, DataMockAppCertificate, DataMockExpire)
	token.IssueTs = DataMockIssueTs
	token.Salt = DataMockSalt
	token.AddService(privileged(NewServiceRtc(DataMockChannelName, DataMockUidStr)))
	token.AddService(privileged(NewServiceRtm(DataMockUserId)))
	token.AddService(privileged(NewServiceStreaming(DataMockChannelName, DataMockUserId)))
	token.AddService(privileged(NewServiceFpa()))
	token.AddService(privileged(NewServiceChat(DataMockUserId)))
	token.AddService(privileged(NewServiceFCdn(DataMockChannelName, DataMockUserId)))
	token.AddService(privileged(NewServiceApaas("room", "user", 1)))
	permissions := NewRtm2Permissions()
	permissions.Add(Rtm2ResourceMessageChannels, Rtm2PermissionRead, []string{"channel"})
	token.AddService(privileged(NewServiceRtm2(DataMockUserId, permissions)))
	token.AddService(privileged(NewServiceConvoAI()))
	token.AddService(privileged(NewServiceStt()))
	return token
}
