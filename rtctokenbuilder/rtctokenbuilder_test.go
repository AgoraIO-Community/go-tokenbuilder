package rtctokenbuilder2

import (
	"testing"

	accesstoken "github.com/AgoraIO-Community/go-tokenbuilder/accesstoken"
	"github.com/AgoraIO-Community/go-tokenbuilder/internal/testutil"
)

const (
	DataMockAccount        = "^ZSgT<%q:Fj*@`92>#OHL?\"hkm~nGYiP"
	DataMockAppCertificate = "5CFd2fd1755d40ecb72977518be15d3b"
	DataMockAppId          = "970CA35de60c44645bbae8a215061b33"
	DataMockChannelName    = "7d72365eb983485397e3e3f9d460bdda"
	DataMockExpire         = uint32(600)
	DataMockUid            = uint32(2882341273)
	DataMockUidStr         = "2882341273"
)

func Test_BuildTokenWithUid_RolePublisher(t *testing.T) {
	token, err := BuildTokenWithUid(DataMockAppId, DataMockAppCertificate, DataMockChannelName, DataMockUid, RolePublisher, DataMockExpire)
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
	token, err := BuildTokenWithUid(DataMockAppId, DataMockAppCertificate, DataMockChannelName, DataMockUid, RoleSubscriber, DataMockExpire)
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

func Test_BuildTokenWithAccount_RolePublisher(t *testing.T) {
	token, err := BuildTokenWithAccount(DataMockAppId, DataMockAppCertificate, DataMockChannelName, DataMockAccount, RolePublisher, DataMockExpire)
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

func Test_BuildTokenWithAccount_RoleSubscriber(t *testing.T) {
	token, err := BuildTokenWithAccount(DataMockAppId, DataMockAppCertificate, DataMockChannelName, DataMockAccount, RoleSubscriber, DataMockExpire)
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
