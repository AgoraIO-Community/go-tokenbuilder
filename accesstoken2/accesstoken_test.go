package accesstoken2

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
	testutil.Equal(t, "no service added", err.Error())
	testutil.Equal(t, "", token)
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

func Test_AccessToken_BuildAndParse_ConvoAIAndStt(t *testing.T) {
	accessToken := NewAccessToken(DataMockAppId, DataMockAppCertificate, DataMockExpire)
	accessToken.IssueTs = DataMockIssueTs
	accessToken.Salt = DataMockSalt

	accessToken.AddService(NewServiceStt())
	accessToken.AddService(NewServiceConvoAI())

	token, err := accessToken.Build()
	testutil.Nil(t, err)

	parsed := CreateAccessToken()
	res, err := parsed.Parse(token)
	testutil.Nil(t, err)
	testutil.Equal(t, true, res)
	testutil.Equal(t, 2, len(parsed.Services))

	convoAI := parsed.Services[ServiceTypeConvoAI].(*ServiceConvoAI)
	testutil.Equal(t, uint16(ServiceTypeConvoAI), convoAI.Type)
	testutil.Equal(t, 0, len(convoAI.Privileges))

	stt := parsed.Services[ServiceTypeStt].(*ServiceStt)
	testutil.Equal(t, uint16(ServiceTypeStt), stt.Type)
	testutil.Equal(t, 0, len(stt.Privileges))
}

func Test_AccessToken_BuildAndParse_AllServices(t *testing.T) {
	accessToken := NewAccessToken(DataMockAppId, DataMockAppCertificate, DataMockExpire)
	accessToken.IssueTs = DataMockIssueTs
	accessToken.Salt = DataMockSalt

	rtc := NewServiceRtc(DataMockChannelName, DataMockUidStr)
	rtc.AddPrivilege(PrivilegeJoinChannel, DataMockExpire)
	accessToken.AddService(rtc)

	rtm := NewServiceRtm(DataMockUserId)
	rtm.AddPrivilege(PrivilegeLogin, DataMockExpire)
	accessToken.AddService(rtm)

	streaming := NewServiceStreaming(DataMockChannelName, DataMockUserId)
	streaming.AddPrivilege(PrivilegeStreamingPublishMixStream, DataMockExpire)
	accessToken.AddService(streaming)

	fpa := NewServiceFpa()
	fpa.AddPrivilege(PrivilegeLogin, DataMockExpire)
	accessToken.AddService(fpa)

	chat := NewServiceChat(DataMockUserId)
	chat.AddPrivilege(PrivilegeChatUser, DataMockExpire)
	accessToken.AddService(chat)

	fcdn := NewServiceFCdn(DataMockChannelName, DataMockUserId)
	fcdn.AddPrivilege(PrivilegeFCdnPublish, DataMockExpire)
	accessToken.AddService(fcdn)

	apaas := NewServiceApaas("room", "user", 1)
	apaas.AddPrivilege(PrivilegeApaasRoomUser, DataMockExpire)
	accessToken.AddService(apaas)

	permissions := NewRtm2Permissions()
	permissions.Add(Rtm2ResourceMessageChannels, Rtm2PermissionRead, []string{"channel-a", "channel-b"})
	rtm2 := NewServiceRtm2(DataMockUserId, permissions)
	rtm2.AddPrivilege(PrivilegeLogin, DataMockExpire)
	accessToken.AddService(rtm2)

	accessToken.AddService(NewServiceConvoAI())
	accessToken.AddService(NewServiceStt())

	token, err := accessToken.Build()
	testutil.Nil(t, err)

	parsed := CreateAccessToken()
	res, err := parsed.Parse(token)
	testutil.Nil(t, err)
	testutil.Equal(t, true, res)
	testutil.Equal(t, 10, len(parsed.Services))
	testutil.Equal(t, DataMockChannelName, parsed.Services[ServiceTypeStreaming].(*ServiceStreaming).ChannelName)
	testutil.Equal(t, DataMockUserId, parsed.Services[ServiceTypeFCdn].(*ServiceFCdn).Account)
	testutil.Equal(t, int16(1), parsed.Services[ServiceTypeApaas].(*ServiceApaas).Role)
	testutil.Equal(t, permissions, parsed.Services[ServiceTypeRtm2].(*ServiceRtm2).Permissions)
	testutil.Equal(t, uint16(ServiceTypeConvoAI), parsed.Services[ServiceTypeConvoAI].(*ServiceConvoAI).Type)
	testutil.Equal(t, uint16(ServiceTypeStt), parsed.Services[ServiceTypeStt].(*ServiceStt).Type)
}

func Test_AccessToken_PreservesDuplicateServiceTypes(t *testing.T) {
	accessToken := NewAccessToken(DataMockAppId, DataMockAppCertificate, DataMockExpire)
	accessToken.IssueTs = DataMockIssueTs
	accessToken.Salt = DataMockSalt

	first := NewServiceRtc("channel-a", "user-a")
	first.AddPrivilege(PrivilegeJoinChannel, 100)
	second := NewServiceRtc("channel-b", "user-b")
	second.AddPrivilege(PrivilegeJoinChannel, 200)
	accessToken.AddService(first)
	accessToken.AddService(second)

	testutil.Equal(t, 2, len(accessToken.GetServices(ServiceTypeRtc)))
	testutil.Equal(t, second, accessToken.Services[ServiceTypeRtc])

	token, err := accessToken.Build()
	testutil.Nil(t, err)

	parsed := CreateAccessToken()
	res, err := parsed.Parse(token)
	testutil.Nil(t, err)
	testutil.Equal(t, true, res)
	services := parsed.GetServices(ServiceTypeRtc)
	testutil.Equal(t, 2, len(services))
	testutil.Equal(t, "channel-a", services[0].(*ServiceRtc).ChannelName)
	testutil.Equal(t, "channel-b", services[1].(*ServiceRtc).ChannelName)
	testutil.Equal(t, "channel-b", parsed.Services[ServiceTypeRtc].(*ServiceRtc).ChannelName)
}

func Test_AccessToken_DirectServicesMapCompatibility(t *testing.T) {
	accessToken := NewAccessToken(DataMockAppId, DataMockAppCertificate, DataMockExpire)
	accessToken.IssueTs = DataMockIssueTs
	accessToken.Salt = DataMockSalt
	accessToken.Services[ServiceTypeRtm] = NewServiceRtm(DataMockUserId)

	token, err := accessToken.Build()
	testutil.Nil(t, err)

	parsed := CreateAccessToken()
	res, err := parsed.Parse(token)
	testutil.Nil(t, err)
	testutil.Equal(t, true, res)
	testutil.Equal(t, DataMockUserId, parsed.Services[ServiceTypeRtm].(*ServiceRtm).UserId)
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
	accessToken.AddService(NewServiceConvoAI())

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

func Test_AccessToken_Parse_ProductTokens_FromPython(t *testing.T) {
	// These fixtures were generated with the authoritative Python AccessToken2
	// implementation at feature/access-token2-convoai-stt-services using fixed
	// issue timestamp 1111111 and salt 1.
	tests := []struct {
		name        string
		token       string
		serviceType uint16
	}{
		{
			name:        "ConvoAI",
			token:       "007eJxTYHir8iZGjD/y66czascYefIetUc+3Rf4yYdrefmMudN5TmcrMFiaGzg7GpumpJoZJJuYmJmYJiUlplokGhmaGpgZJhkbu38RYIhgYmBgZGBgYAaSLEAM4jOBSWYwyQImFRjMU8yNjM1MU5MsLYxNLEyNLc1TjVON0yxTTMwMklJSErkYjCwsjIxNDI3MjZmA5kBM4mQoSS0uiS8tTi3iBFoCAA59KuQ=",
			serviceType: ServiceTypeConvoAI,
		},
		{
			name:        "STT",
			token:       "007eJxTYHDgnv+xdU7Dq1Qt68/XGRdO0l530Mbz0lG1tfvdX+8+/l5DgcHS3MDZ0dg0JdXMINnExMzENCkpMdUi0cjQ1MDMMMnY2P2LAEMEEwMDIwMDAzOQZAFiEJ8JTDKDSRYwqcBgnmJuZGxmmppkaWFsYmFqbGmeapxqnGaZYmJmkJSSksjFYGRhYWRsYmhkbswENAdiEidDSWpxSXxpcWoRF9ASALuALAs=",
			serviceType: ServiceTypeStt,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			parsed := CreateAccessToken()
			res, err := parsed.Parse(test.token)
			testutil.Nil(t, err)
			testutil.Equal(t, true, res)
			testutil.Equal(t, 3, len(parsed.Services))
			testutil.Equal(t, 1, len(parsed.GetServices(test.serviceType)))

			valid, err := parsed.VerifySignature(DataMockAppCertificate)
			testutil.Nil(t, err)
			testutil.Equal(t, true, valid)
		})
	}
}

func Test_GetUidStr(t *testing.T) {
	testutil.Equal(t, "", GetUidStr(0))
	testutil.Equal(t, DataMockUidStr, GetUidStr(DataMockUid))
}

func Test_getVersion(t *testing.T) {
	testutil.Equal(t, "007", getVersion())
}

func Test_isUuid(t *testing.T) {
	testutil.Equal(t, true, isUuid(DataMockAppId))
	testutil.Equal(t, true, isUuid(DataMockAppCertificate))
	testutil.Equal(t, false, isUuid(""))
	testutil.Equal(t, false, isUuid("abc"))
	testutil.Equal(t, false, isUuid("Z70CA35de60c44645bbae8a215061b33"))
}
