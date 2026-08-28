package accesstoken

import (
	"bytes"
	"testing"

	"github.com/AgoraIO-Community/go-tokenbuilder/internal/testutil"
)

const (
	DataMockAppId          = "970CA35de60c44645bbae8a215061b33"
	DataMockAppCertificate = "5CFd2fd1755d40ecb72977518be15d3b"
	DataMockChannelName    = "7d72365eb983485397e3e3f9d460bdda"
	DataMockExpire         = uint32(600)
	DataMockIssueTs        = uint32(1111111)
	DataMockSalt           = uint32(1)
	DataMockUid            = uint32(2882341273)
	DataMockUidStr         = "2882341273"
	DataMockUserId         = "test_user"
)

func Test_AccessToken_Build(t *testing.T) {
	accessToken := NewAccessToken(DataMockAppId, DataMockAppCertificate, DataMockExpire)
	accessToken.IssueTs = DataMockIssueTs
	accessToken.Salt = DataMockSalt

	testutil.Equal(t, DataMockAppCertificate, accessToken.AppCert)
	testutil.Equal(t, DataMockAppId, accessToken.AppId)
	testutil.Equal(t, DataMockExpire, accessToken.Expire)
	testutil.Equal(t, DataMockIssueTs, accessToken.IssueTs)
	testutil.Equal(t, DataMockSalt, accessToken.Salt)
	testutil.Equal(t, 0, len(accessToken.Services))

	token, err := accessToken.Build()
	testutil.Nil(t, err)
	testutil.Equal(t, "007eJxSYEiJ9+zw7Gb1viNuGtMfy3JriuZNp+1h1iLu/rOePHlS91WBwdLcwNnR2DQl1cwg2cTEzMQ0KSkx1SLRyNDUwMwwydjY/YsAQwQTAwMjAwgAAgAA//+rZxiv", token)
}

func Test_AccessToken_Build_Error_AppId(t *testing.T) {
	accessToken := NewAccessToken("", DataMockAppCertificate, DataMockExpire)
	token, err := accessToken.Build()
	testutil.Equal(t, "check appId or appCertificate", err.Error())
	testutil.Equal(t, "", token)

	accessToken = NewAccessToken("abc", DataMockAppCertificate, DataMockExpire)
	token, err = accessToken.Build()
	testutil.Equal(t, "check appId or appCertificate", err.Error())
	testutil.Equal(t, "", token)

	accessToken = NewAccessToken("Z70CA35de60c44645bbae8a215061b33", DataMockAppCertificate, DataMockExpire)
	token, err = accessToken.Build()
	testutil.Equal(t, "check appId or appCertificate", err.Error())
	testutil.Equal(t, "", token)
}

func Test_AccessToken_Build_Error_AppCertificate(t *testing.T) {
	accessToken := NewAccessToken(DataMockAppId, "", DataMockExpire)
	token, err := accessToken.Build()
	testutil.Equal(t, "check appId or appCertificate", err.Error())
	testutil.Equal(t, "", token)

	accessToken = NewAccessToken(DataMockAppId, "abc", DataMockExpire)
	token, err = accessToken.Build()
	testutil.Equal(t, "check appId or appCertificate", err.Error())
	testutil.Equal(t, "", token)

	accessToken = NewAccessToken(DataMockAppId, "5CFd2fd1755d40ecb72977518be15d3Z", DataMockExpire)
	token, err = accessToken.Build()
	testutil.Equal(t, "check appId or appCertificate", err.Error())
	testutil.Equal(t, "", token)
}

func Test_AccessToken_Build_ServiceRtc(t *testing.T) {
	accessToken := NewAccessToken(DataMockAppId, DataMockAppCertificate, DataMockExpire)
	accessToken.IssueTs = DataMockIssueTs
	accessToken.Salt = DataMockSalt

	serviceRtc := NewServiceRtc(DataMockChannelName, DataMockUidStr)
	serviceRtc.AddPrivilege(PrivilegeJoinChannel, DataMockExpire)
	accessToken.AddService(serviceRtc)

	testutil.Equal(t, DataMockChannelName, serviceRtc.ChannelName)
	testutil.Equal(t, DataMockUidStr, serviceRtc.Uid)

	token, err := accessToken.Build()
	testutil.Nil(t, err)
	testutil.Equal(t, "007eJxSYBBbsMMnKq7p9Hf/HcIX5kce9b518kCiQgSr5Zrp4X1Tu6UUGCzNDZwdjU1TUs0Mkk1MzExMk5ISUy0SjQxNDcwMk4yN3b8IMEQwMTAwMoAwBIL4CgzmKeZGxmamqUmWFsYmFqbGluapxqnGaZYpJmYGSSkpiVwMRhYWRsYmhkbmxoAAAAD//8JqJOM=", token)
}

func Test_AccessToken_Build_ServiceRtc_Uid0(t *testing.T) {
	accessToken := NewAccessToken(DataMockAppId, DataMockAppCertificate, DataMockExpire)
	accessToken.IssueTs = DataMockIssueTs
	accessToken.Salt = DataMockSalt

	serviceRtc := NewServiceRtc(DataMockChannelName, "")
	serviceRtc.AddPrivilege(PrivilegeJoinChannel, DataMockExpire)
	accessToken.AddService(serviceRtc)

	testutil.Equal(t, DataMockChannelName, serviceRtc.ChannelName)
	testutil.Equal(t, "", serviceRtc.Uid)

	token, err := accessToken.Build()
	testutil.Nil(t, err)
	testutil.Equal(t, "007eJxSYLhzZP08Lxa1Pg57+TcXb/3cZ3wi4V6kbpbOog0G2dOYk20UGCzNDZwdjU1TUs0Mkk1MzExMk5ISUy0SjQxNDcwMk4yN3b8IMEQwMTAwMoAwBIL4CgzmKeZGxmamqUmWFsYmFqbGluapxqnGaZYpJmYGSSkpiQwMgAAAAP//Npwiag==", token)
}

func Test_AccessToken_Build_ServiceRtm(t *testing.T) {
	accessToken := NewAccessToken(DataMockAppId, DataMockAppCertificate, DataMockExpire)
	accessToken.IssueTs = DataMockIssueTs
	accessToken.Salt = DataMockSalt

	serviceRtm := NewServiceRtm(DataMockUserId)
	serviceRtm.AddPrivilege(PrivilegeLogin, DataMockExpire)
	accessToken.AddService(serviceRtm)

	testutil.Equal(t, DataMockUserId, serviceRtm.UserId)

	token, err := accessToken.Build()
	testutil.Nil(t, err)
	testutil.Equal(t, "007eJxSYOCdJftjyTM2zxW6Xhm/5T0j5LdcUt/xYVt48fb5Mp3PX9coMFiaGzg7GpumpJoZJJuYmJmYJiUlplokGhmaGpgZJhkbu38RYIhgYmBgZABhJgZGBkYwn5OhJLW4JL60OLUIEAAA//9ZVh6A", token)
}

func Test_AccessToken_Build_ServiceChatUser(t *testing.T) {
	accessToken := NewAccessToken(DataMockAppId, DataMockAppCertificate, DataMockExpire)
	accessToken.IssueTs = DataMockIssueTs
	accessToken.Salt = DataMockSalt

	serviceChat := NewServiceChat(DataMockUidStr)
	serviceChat.AddPrivilege(PrivilegeChatUser, DataMockExpire)
	accessToken.AddService(serviceChat)

	testutil.Equal(t, DataMockUidStr, serviceChat.UserId)

	token, err := accessToken.Build()
	testutil.Nil(t, err)
	testutil.Equal(t, "007eJxSYNAIsnbS3v/A5t2TC6feR15r+6cq8bqAvfaW+tk/Vzz+p6xTYLA0N3B2NDZNSTUzSDYxMTMxTUpKTLVINDI0NTAzTDI2dv8iwBDBxMDAyADCrAyMDIxgPheDkYWFkbGJoZG5MSAAAP//H6UeuA==", token)
}

func Test_AccessToken_Build_ServiceChatApp(t *testing.T) {
	accessToken := NewAccessToken(DataMockAppId, DataMockAppCertificate, DataMockExpire)
	accessToken.IssueTs = DataMockIssueTs
	accessToken.Salt = DataMockSalt

	serviceChat := NewServiceChat("")
	serviceChat.AddPrivilege(PrivilegeChatApp, DataMockExpire)
	accessToken.AddService(serviceChat)

	testutil.Equal(t, "", serviceChat.UserId)

	token, err := accessToken.Build()
	testutil.Nil(t, err)
	testutil.Equal(t, "007eJxSYNDNaz3snC8huEfHWdz6s98qltq4zqy9fl99Uh0FDvy6F6DAYGlu4OxobJqSamaQbGJiZmKalJSYapFoZGhqYGaYZGzs/kWAIYKJgYGRAYRZGRgZmMB8BgZAAAAA//+t8hhr", token)
}

func Test_AccessToken_Build_MultipleServices(t *testing.T) {
	accessToken := NewAccessToken(DataMockAppId, DataMockAppCertificate, DataMockExpire)
	accessToken.IssueTs = DataMockIssueTs
	accessToken.Salt = DataMockSalt

	// RTC
	serviceRtc := NewServiceRtc(DataMockChannelName, DataMockUidStr)
	serviceRtc.AddPrivilege(PrivilegeJoinChannel, DataMockExpire)
	accessToken.AddService(serviceRtc)

	// RTM
	serviceRtm := NewServiceRtm(DataMockUserId)
	serviceRtm.AddPrivilege(PrivilegeLogin, DataMockExpire)
	accessToken.AddService(serviceRtm)

	// CHAT
	serviceChat := NewServiceChat(DataMockUidStr)
	serviceChat.AddPrivilege(PrivilegeChatUser, DataMockExpire)
	accessToken.AddService(serviceChat)

	token, err := accessToken.Build()
	testutil.Nil(t, err)
	testutil.Equal(t, "007eJxSYJjqLJBlM239wwWvmBZ7tW619coNnPKSXaHayfKzZODswxMVGCzNDZwdjU1TUs0Mkk1MzExMk5ISUy0SjQxNDcwMk4yN3b8IMEQwMTAwMjAwMDMwgiGIr8BgnmJuZGxmmppkaWFsYmFqbGmeapxqnGaZYmJmkJSSksjFYGRhYWRsYmhkbswE18fJUJJaXBJfWpxaxAoXRFYKCAAA///aoiqr", token)
}

func Test_AccessToken_Parse_TokenRtc(t *testing.T) {
	accessToken := CreateAccessToken()
	res, err := accessToken.Parse("007eJxSYBBbsMMnKq7p9Hf/HcIX5kce9b518kCiQgSr5Zrp4X1Tu6UUGCzNDZwdjU1TUs0Mkk1MzExMk5ISUy0SjQxNDcwMk4yN3b8IMEQwMTAwMoAwBIL4CgzmKeZGxmamqUmWFsYmFqbGluapxqnGaZYpJmYGSSkpiVwMRhYWRsYmhkbmxoAAAAD//8JqJOM=")

	testutil.Nil(t, err)
	testutil.Equal(t, res, true)
	testutil.Equal(t, DataMockAppId, accessToken.AppId)
	testutil.Equal(t, DataMockExpire, accessToken.Expire)
	testutil.Equal(t, DataMockIssueTs, accessToken.IssueTs)
	testutil.Equal(t, DataMockSalt, accessToken.Salt)
	testutil.Equal(t, 1, len(accessToken.Services))
	testutil.Equal(t, true, accessToken.Services[ServiceTypeRtc] != nil)
	testutil.Equal(t, DataMockChannelName, accessToken.Services[ServiceTypeRtc].(*ServiceRtc).ChannelName)
	testutil.Equal(t, DataMockUidStr, accessToken.Services[ServiceTypeRtc].(*ServiceRtc).Uid)
	testutil.Equal(t, uint16(ServiceTypeRtc), accessToken.Services[ServiceTypeRtc].(*ServiceRtc).Type)
	testutil.Equal(t, DataMockExpire, accessToken.Services[ServiceTypeRtc].(*ServiceRtc).Privileges[PrivilegeJoinChannel])
	testutil.Equal(t, uint32(0), accessToken.Services[ServiceTypeRtc].(*ServiceRtc).Privileges[PrivilegePublishAudioStream])
	testutil.Equal(t, uint32(0), accessToken.Services[ServiceTypeRtc].(*ServiceRtc).Privileges[PrivilegePublishVideoStream])
	testutil.Equal(t, uint32(0), accessToken.Services[ServiceTypeRtc].(*ServiceRtc).Privileges[PrivilegePublishDataStream])
}

func Test_AccessToken_Parse_TokenRtc_FromPython(t *testing.T) {
	accessToken := CreateAccessToken()
	res, err := accessToken.Parse("007eJxTYBBbsMMnKq7p9Hf/HcIX5kce9b518kCiQgSr5Zrp4X1Tu6UUGCzNDZwdjU1TUs0Mkk1MzExMk5ISUy0SjQxNDcwMk4yN3b8IMEQwMTAwMoAwBIL4CgzmKeZGxmamqUmWFsYmFqbGluapxqnGaZYpJmYGSSkpiVwMRhYWRsYmhkbmxgDCaiTj")

	testutil.Nil(t, err)
	testutil.Equal(t, true, res)
	testutil.Equal(t, DataMockAppId, accessToken.AppId)
	testutil.Equal(t, DataMockExpire, accessToken.Expire)
	testutil.Equal(t, DataMockIssueTs, accessToken.IssueTs)
	testutil.Equal(t, DataMockSalt, accessToken.Salt)
	testutil.Equal(t, 1, len(accessToken.Services))
	testutil.Equal(t, true, accessToken.Services[ServiceTypeRtc] != nil)
	testutil.Equal(t, DataMockChannelName, accessToken.Services[ServiceTypeRtc].(*ServiceRtc).ChannelName)
	testutil.Equal(t, DataMockUidStr, accessToken.Services[ServiceTypeRtc].(*ServiceRtc).Uid)
	testutil.Equal(t, uint16(ServiceTypeRtc), accessToken.Services[ServiceTypeRtc].(*ServiceRtc).Type)
	testutil.Equal(t, DataMockExpire, accessToken.Services[ServiceTypeRtc].(*ServiceRtc).Privileges[PrivilegeJoinChannel])
	testutil.Equal(t, uint32(0), accessToken.Services[ServiceTypeRtc].(*ServiceRtc).Privileges[PrivilegePublishAudioStream])
	testutil.Equal(t, uint32(0), accessToken.Services[ServiceTypeRtc].(*ServiceRtc).Privileges[PrivilegePublishVideoStream])
	testutil.Equal(t, uint32(0), accessToken.Services[ServiceTypeRtc].(*ServiceRtc).Privileges[PrivilegePublishDataStream])
}

func Test_AccessToken_Parse_TokenRtc_Rtm_MultiService_FromPython(t *testing.T) {
	accessToken := CreateAccessToken()
	res, err := accessToken.Parse("007eJxTYOAQsrQ5s3TfH+1tvy8zZZ46EpCc0V43JXdGd2jS8porKo4KDJbmBs6OxqYpqWYGySYmZiamSUmJqRaJRoamBmaGScbG7l8EGCKYGBgYGRgYmIAkCxCD+ExgkhlMsoBJBQbzFHMjYzPT1CRLC2MTC1NjS/NU41TjNMsUEzODpJSURC4GIwsLI2MTQyNzY5BZEJM4GUpSi0viS4tTiwAipyp4")

	testutil.Nil(t, err)
	testutil.Equal(t, true, res)
	testutil.Equal(t, DataMockAppId, accessToken.AppId)
	testutil.Equal(t, DataMockExpire, accessToken.Expire)
	testutil.Equal(t, DataMockIssueTs, accessToken.IssueTs)
	testutil.Equal(t, DataMockSalt, accessToken.Salt)
	testutil.Equal(t, 2, len(accessToken.Services))
	testutil.Equal(t, true, accessToken.Services[ServiceTypeRtc] != nil)
	testutil.Equal(t, DataMockChannelName, accessToken.Services[ServiceTypeRtc].(*ServiceRtc).ChannelName)
	testutil.Equal(t, DataMockUidStr, accessToken.Services[ServiceTypeRtc].(*ServiceRtc).Uid)
	testutil.Equal(t, uint16(ServiceTypeRtc), accessToken.Services[ServiceTypeRtc].(*ServiceRtc).Type)
	testutil.Equal(t, DataMockExpire, accessToken.Services[ServiceTypeRtc].(*ServiceRtc).Privileges[PrivilegeJoinChannel])
	testutil.Equal(t, DataMockExpire, accessToken.Services[ServiceTypeRtc].(*ServiceRtc).Privileges[PrivilegePublishAudioStream])
	testutil.Equal(t, DataMockExpire, accessToken.Services[ServiceTypeRtc].(*ServiceRtc).Privileges[PrivilegePublishVideoStream])
	testutil.Equal(t, DataMockExpire, accessToken.Services[ServiceTypeRtc].(*ServiceRtc).Privileges[PrivilegePublishDataStream])
	testutil.Equal(t, true, accessToken.Services[ServiceTypeRtm] != nil)
	testutil.Equal(t, DataMockUserId, accessToken.Services[ServiceTypeRtm].(*ServiceRtm).UserId)
	testutil.Equal(t, uint16(ServiceTypeRtm), accessToken.Services[ServiceTypeRtm].(*ServiceRtm).Type)
	testutil.Equal(t, DataMockExpire, accessToken.Services[ServiceTypeRtm].(*ServiceRtm).Privileges[PrivilegeLogin])
}

func Test_AccessToken_Parse_TokenRtm(t *testing.T) {
	accessToken := CreateAccessToken()
	res, err := accessToken.Parse("007eJxSYOCdJftjyTM2zxW6Xhm/5T0j5LdcUt/xYVt48fb5Mp3PX9coMFiaGzg7GpumpJoZJJuYmJmYJiUlplokGhmaGpgZJhkbu38RYIhgYmBgZABhJgZGBkYwn5OhJLW4JL60OLUIEAAA//9ZVh6A")

	testutil.Nil(t, err)
	testutil.Equal(t, true, res)
	testutil.Equal(t, DataMockAppId, accessToken.AppId)
	testutil.Equal(t, DataMockExpire, accessToken.Expire)
	testutil.Equal(t, DataMockIssueTs, accessToken.IssueTs)
	testutil.Equal(t, DataMockSalt, accessToken.Salt)
	testutil.Equal(t, 1, len(accessToken.Services))
	testutil.Equal(t, true, accessToken.Services[ServiceTypeRtm] != nil)
	testutil.Equal(t, DataMockUserId, accessToken.Services[ServiceTypeRtm].(*ServiceRtm).UserId)
	testutil.Equal(t, uint16(ServiceTypeRtm), accessToken.Services[ServiceTypeRtm].(*ServiceRtm).Type)
	testutil.Equal(t, DataMockExpire, accessToken.Services[ServiceTypeRtm].(*ServiceRtm).Privileges[PrivilegeLogin])
}

func Test_AccessToken_Parse_TokenChatUser(t *testing.T) {
	accessToken := CreateAccessToken()
	res, err := accessToken.Parse("007eJxTYNAIsnbS3v/A5t2TC6feR15r+6cq8bqAvfaW+tk/Vzz+p6xTYLA0N3B2NDZNSTUzSDYxMTMxTUpKTLVINDI0NTAzTDI2dv8iwBDBxMDAyADCrEDMCOZzMRhZWBgZmxgamRsDAB+lHrg=")

	testutil.Nil(t, err)
	testutil.Equal(t, true, res)
	testutil.Equal(t, DataMockAppId, accessToken.AppId)
	testutil.Equal(t, DataMockExpire, accessToken.Expire)
	testutil.Equal(t, DataMockIssueTs, accessToken.IssueTs)
	testutil.Equal(t, DataMockSalt, accessToken.Salt)
	testutil.Equal(t, 1, len(accessToken.Services))
	testutil.Equal(t, true, accessToken.Services[ServiceTypeChat] != nil)
	testutil.Equal(t, DataMockUidStr, accessToken.Services[ServiceTypeChat].(*ServiceChat).UserId)
	testutil.Equal(t, uint16(ServiceTypeChat), accessToken.Services[ServiceTypeChat].(*ServiceChat).Type)
	testutil.Equal(t, DataMockExpire, accessToken.Services[ServiceTypeChat].(*ServiceChat).Privileges[PrivilegeChatUser])
}

func Test_AccessToken_Parse_TokenChatApp(t *testing.T) {
	accessToken := CreateAccessToken()
	res, err := accessToken.Parse("007eJxTYNDNaz3snC8huEfHWdz6s98qltq4zqy9fl99Uh0FDvy6F6DAYGlu4OxobJqSamaQbGJiZmKalJSYapFoZGhqYGaYZGzs/kWAIYKJgYGRAYRZgZgJzGdgAACt8hhr")

	testutil.Nil(t, err)
	testutil.Equal(t, true, res)
	testutil.Equal(t, DataMockAppId, accessToken.AppId)
	testutil.Equal(t, DataMockExpire, accessToken.Expire)
	testutil.Equal(t, DataMockIssueTs, accessToken.IssueTs)
	testutil.Equal(t, DataMockSalt, accessToken.Salt)
	testutil.Equal(t, 1, len(accessToken.Services))
	testutil.Equal(t, true, accessToken.Services[ServiceTypeChat] != nil)
	testutil.Equal(t, "", accessToken.Services[ServiceTypeChat].(*ServiceChat).UserId)
	testutil.Equal(t, uint16(ServiceTypeChat), accessToken.Services[ServiceTypeChat].(*ServiceChat).Type)
	testutil.Equal(t, DataMockExpire, accessToken.Services[ServiceTypeChat].(*ServiceChat).Privileges[PrivilegeChatApp])
}

func Test_GetUidStr(t *testing.T) {
	testutil.Equal(t, "", GetUidStr(0))
	testutil.Equal(t, DataMockUidStr, GetUidStr(DataMockUid))
}

func Test_getVersion(t *testing.T) {
	testutil.Equal(t, "007", getVersion())
}

func Test_AccessToken_ParseMalformedTokens(t *testing.T) {
	truncated := new(bytes.Buffer)
	testutil.Nil(t, packString(truncated, "signature"))
	testutil.Nil(t, packUint16(truncated, 32))
	testutil.Nil(t, truncated.WriteByte('a'))

	tests := []string{
		"",
		"0",
		"007",
		"006not-a-token",
		"007not-base64",
		"007bm90LXpsaWI=",
		Version + base64EncodeStr(compressZlib(truncated.Bytes())),
	}

	for _, token := range tests {
		t.Run(token, func(t *testing.T) {
			parsed := CreateAccessToken()
			res, err := parsed.Parse(token)
			testutil.Equal(t, false, res)
			if err == nil {
				t.Errorf("expected malformed token %q to return an error", token)
			}
		})
	}
}

func Test_AccessToken_VerifySignature(t *testing.T) {
	accessToken := NewAccessToken(DataMockAppId, DataMockAppCertificate, DataMockExpire)
	accessToken.IssueTs = DataMockIssueTs
	accessToken.Salt = DataMockSalt
	service := NewServiceRtc(DataMockChannelName, DataMockUidStr)
	service.AddPrivilege(PrivilegeJoinChannel, DataMockExpire)
	accessToken.AddService(service)

	token, err := accessToken.Build()
	testutil.Nil(t, err)

	parsed := CreateAccessToken()
	res, err := parsed.Parse(token)
	testutil.Nil(t, err)
	testutil.Equal(t, true, res)

	valid, err := parsed.VerifySignature(DataMockAppCertificate)
	testutil.Nil(t, err)
	testutil.Equal(t, true, valid)

	valid, err = parsed.VerifySignature("6CFd2fd1755d40ecb72977518be15d3b")
	testutil.Nil(t, err)
	testutil.Equal(t, false, valid)

	_, err = parsed.Parse("007not-base64")
	if err == nil {
		t.Fatal("expected malformed token to fail")
	}
	valid, err = parsed.VerifySignature(DataMockAppCertificate)
	testutil.Equal(t, false, valid)
	if err == nil {
		t.Fatal("expected failed parse to clear signature state")
	}
}

func Test_isUuid(t *testing.T) {
	testutil.Equal(t, true, isUuid(DataMockAppId))
	testutil.Equal(t, true, isUuid(DataMockAppCertificate))
	testutil.Equal(t, false, isUuid(""))
	testutil.Equal(t, false, isUuid("abc"))
	testutil.Equal(t, false, isUuid("Z70CA35de60c44645bbae8a215061b33"))
}
