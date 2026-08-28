package rtctokenbuilder2

import (
	"testing"

	accesstoken "github.com/AgoraIO-Community/go-tokenbuilder/accesstoken2"
	"github.com/AgoraIO-Community/go-tokenbuilder/internal/testutil"
)

const (
	DataMockAccount                      = "2882341273"
	DataMockAppCertificate               = "5CFd2fd1755d40ecb72977518be15d3b"
	DataMockAppId                        = "970CA35de60c44645bbae8a215061b33"
	DataMockChannelName                  = "7d72365eb983485397e3e3f9d460bdda"
	DataMockExpire                       = uint32(600)
	DataMockJoinChannelPrivilegeExpire   = uint32(600)
	DataMockPubAudioPrivilegeExpire      = uint32(600)
	DataMockPubVideoPrivilegeExpire      = uint32(600)
	DataMockPubDataStreamPrivilegeExpire = uint32(600)
	DataMockUid                          = uint32(2882341273)
	DataMockUidStr                       = "2882341273"
)

func Test_BuildTokenWithUid_RolePublisher(t *testing.T) {
	token, err := BuildTokenWithUid(DataMockAppId, DataMockAppCertificate, DataMockChannelName, DataMockUid, RolePublisher, DataMockExpire, DataMockExpire)
	testutil.Nil(t, err)

	accessToken := accesstoken.CreateAccessToken()
	parsed, err := accessToken.Parse(token)
	testutil.Nil(t, err)
	testutil.Equal(t, true, parsed)

	testutil.Equal(t, DataMockAppId, accessToken.AppId)
	testutil.Equal(t, DataMockExpire, accessToken.Expire)
	testutil.Equal(t, true, accessToken.Services[accesstoken.ServiceTypeRtc] != nil)
	testutil.Equal(t, DataMockChannelName, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).ChannelName)
	testutil.Equal(t, DataMockUidStr, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Uid)
	testutil.Equal(t, uint16(accesstoken.ServiceTypeRtc), accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Type)
	testutil.Equal(t, DataMockExpire, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegeJoinChannel])
	testutil.Equal(t, DataMockExpire, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegePublishAudioStream])
	testutil.Equal(t, DataMockExpire, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegePublishVideoStream])
	testutil.Equal(t, DataMockExpire, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegePublishDataStream])
}

func Test_BuildTokenWithUid_RoleSubscriber(t *testing.T) {
	token, err := BuildTokenWithUid(DataMockAppId, DataMockAppCertificate, DataMockChannelName, DataMockUid, RoleSubscriber, DataMockExpire, DataMockExpire)
	testutil.Nil(t, err)

	accessToken := accesstoken.CreateAccessToken()
	parsed, err := accessToken.Parse(token)
	testutil.Nil(t, err)
	testutil.Equal(t, true, parsed)

	testutil.Equal(t, DataMockAppId, accessToken.AppId)
	testutil.Equal(t, DataMockExpire, accessToken.Expire)
	testutil.Equal(t, true, accessToken.Services[accesstoken.ServiceTypeRtc] != nil)
	testutil.Equal(t, DataMockChannelName, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).ChannelName)
	testutil.Equal(t, DataMockUidStr, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Uid)
	testutil.Equal(t, uint16(accesstoken.ServiceTypeRtc), accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Type)
	testutil.Equal(t, DataMockExpire, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegeJoinChannel])
	testutil.Equal(t, uint32(0), accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegePublishAudioStream])
	testutil.Equal(t, uint32(0), accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegePublishVideoStream])
	testutil.Equal(t, uint32(0), accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegePublishDataStream])
}

func Test_BuildTokenWithUserAccount_RolePublisher(t *testing.T) {
	token, err := BuildTokenWithUserAccount(DataMockAppId, DataMockAppCertificate, DataMockChannelName, DataMockAccount, RolePublisher, DataMockExpire, DataMockExpire)
	testutil.Nil(t, err)

	accessToken := accesstoken.CreateAccessToken()
	parsed, err := accessToken.Parse(token)
	testutil.Nil(t, err)
	testutil.Equal(t, true, parsed)

	testutil.Equal(t, DataMockAppId, accessToken.AppId)
	testutil.Equal(t, DataMockExpire, accessToken.Expire)
	testutil.Equal(t, true, accessToken.Services[accesstoken.ServiceTypeRtc] != nil)
	testutil.Equal(t, DataMockChannelName, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).ChannelName)
	testutil.Equal(t, DataMockAccount, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Uid)
	testutil.Equal(t, uint16(accesstoken.ServiceTypeRtc), accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Type)
	testutil.Equal(t, DataMockExpire, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegeJoinChannel])
	testutil.Equal(t, DataMockExpire, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegePublishAudioStream])
	testutil.Equal(t, DataMockExpire, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegePublishVideoStream])
	testutil.Equal(t, DataMockExpire, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegePublishDataStream])
}

func Test_BuildTokenWithUserAccount_RoleSubscriber(t *testing.T) {
	token, err := BuildTokenWithUserAccount(DataMockAppId, DataMockAppCertificate, DataMockChannelName, DataMockAccount, RoleSubscriber, DataMockExpire, DataMockExpire)
	testutil.Nil(t, err)

	accessToken := accesstoken.CreateAccessToken()
	parsed, err := accessToken.Parse(token)
	testutil.Nil(t, err)
	testutil.Equal(t, true, parsed)

	testutil.Equal(t, DataMockAppId, accessToken.AppId)
	testutil.Equal(t, DataMockExpire, accessToken.Expire)
	testutil.Equal(t, true, accessToken.Services[accesstoken.ServiceTypeRtc] != nil)
	testutil.Equal(t, DataMockChannelName, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).ChannelName)
	testutil.Equal(t, DataMockAccount, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Uid)
	testutil.Equal(t, uint16(accesstoken.ServiceTypeRtc), accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Type)
	testutil.Equal(t, DataMockExpire, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegeJoinChannel])
	testutil.Equal(t, uint32(0), accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegePublishAudioStream])
	testutil.Equal(t, uint32(0), accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegePublishVideoStream])
	testutil.Equal(t, uint32(0), accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegePublishDataStream])
}

func Test_BuildTokenWithUidAndPrivilege(t *testing.T) {
	token, err := BuildTokenWithUidAndPrivilege(DataMockAppId, DataMockAppCertificate, DataMockChannelName, DataMockUid, DataMockExpire, DataMockJoinChannelPrivilegeExpire, DataMockPubAudioPrivilegeExpire, DataMockPubVideoPrivilegeExpire, DataMockPubDataStreamPrivilegeExpire)
	testutil.Nil(t, err)

	accessToken := accesstoken.CreateAccessToken()
	parsed, err := accessToken.Parse(token)
	testutil.Nil(t, err)
	testutil.Equal(t, true, parsed)

	testutil.Equal(t, DataMockAppId, accessToken.AppId)
	testutil.Equal(t, DataMockExpire, accessToken.Expire)
	testutil.Equal(t, true, accessToken.Services[accesstoken.ServiceTypeRtc] != nil)
	testutil.Equal(t, DataMockChannelName, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).ChannelName)
	testutil.Equal(t, DataMockUidStr, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Uid)
	testutil.Equal(t, uint16(accesstoken.ServiceTypeRtc), accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Type)
	testutil.Equal(t, DataMockJoinChannelPrivilegeExpire, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegeJoinChannel])
	testutil.Equal(t, DataMockPubAudioPrivilegeExpire, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegePublishAudioStream])
	testutil.Equal(t, DataMockPubVideoPrivilegeExpire, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegePublishVideoStream])
	testutil.Equal(t, DataMockPubDataStreamPrivilegeExpire, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegePublishDataStream])
}

func Test_BuildTokenWithUserAccountAndPrivilege(t *testing.T) {
	token, err := BuildTokenWithUserAccountAndPrivilege(DataMockAppId, DataMockAppCertificate, DataMockChannelName, DataMockAccount, DataMockExpire, DataMockJoinChannelPrivilegeExpire, DataMockPubAudioPrivilegeExpire, DataMockPubVideoPrivilegeExpire, DataMockPubDataStreamPrivilegeExpire)
	testutil.Nil(t, err)

	accessToken := accesstoken.CreateAccessToken()
	parsed, err := accessToken.Parse(token)
	testutil.Nil(t, err)
	testutil.Equal(t, true, parsed)

	testutil.Equal(t, DataMockAppId, accessToken.AppId)
	testutil.Equal(t, DataMockExpire, accessToken.Expire)
	testutil.Equal(t, true, accessToken.Services[accesstoken.ServiceTypeRtc] != nil)
	testutil.Equal(t, DataMockChannelName, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).ChannelName)
	testutil.Equal(t, DataMockAccount, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Uid)
	testutil.Equal(t, uint16(accesstoken.ServiceTypeRtc), accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Type)
	testutil.Equal(t, DataMockJoinChannelPrivilegeExpire, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegeJoinChannel])
	testutil.Equal(t, DataMockPubAudioPrivilegeExpire, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegePublishAudioStream])
	testutil.Equal(t, DataMockPubVideoPrivilegeExpire, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegePublishVideoStream])
	testutil.Equal(t, DataMockPubDataStreamPrivilegeExpire, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegePublishDataStream])
}

func Test_BuildTokenWithRtm(t *testing.T) {
	token, err := BuildTokenWithRtm(DataMockAppId, DataMockAppCertificate, DataMockChannelName, DataMockAccount, RolePublisher, DataMockExpire, DataMockExpire)
	testutil.Nil(t, err)

	parsed := accesstoken.CreateAccessToken()
	ok, err := parsed.Parse(token)
	testutil.Nil(t, err)
	testutil.Equal(t, true, ok)
	testutil.Equal(t, 2, len(parsed.Services))
	testutil.Equal(t, DataMockAccount, parsed.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Uid)
	testutil.Equal(t, DataMockAccount, parsed.Services[accesstoken.ServiceTypeRtm].(*accesstoken.ServiceRtm).UserId)
}

func Test_BuildTokenWithRtm2(t *testing.T) {
	const rtmUserID = "rtm-user"
	const rtmExpire = uint32(300)

	token, err := BuildTokenWithRtm2(
		DataMockAppId,
		DataMockAppCertificate,
		DataMockChannelName,
		DataMockAccount,
		RolePublisher,
		DataMockExpire,
		DataMockJoinChannelPrivilegeExpire,
		DataMockPubAudioPrivilegeExpire,
		DataMockPubVideoPrivilegeExpire,
		DataMockPubDataStreamPrivilegeExpire,
		rtmUserID,
		rtmExpire,
	)
	testutil.Nil(t, err)

	parsed := accesstoken.CreateAccessToken()
	ok, err := parsed.Parse(token)
	testutil.Nil(t, err)
	testutil.Equal(t, true, ok)
	testutil.Equal(t, 2, len(parsed.Services))
	testutil.Equal(t, DataMockAccount, parsed.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Uid)
	rtm := parsed.Services[accesstoken.ServiceTypeRtm].(*accesstoken.ServiceRtm)
	testutil.Equal(t, rtmUserID, rtm.UserId)
	testutil.Equal(t, rtmExpire, rtm.Privileges[accesstoken.PrivilegeLogin])
}
