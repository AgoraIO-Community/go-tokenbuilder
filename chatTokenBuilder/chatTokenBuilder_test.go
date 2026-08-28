package chatTokenBuilder

import (
	"testing"

	accesstoken "github.com/AgoraIO-Community/go-tokenbuilder/accesstoken"
	"github.com/AgoraIO-Community/go-tokenbuilder/internal/testutil"
)

const (
	DataMockAppCertificate = "5CFd2fd1755d40ecb72977518be15d3b"
	DataMockAppId          = "970CA35de60c44645bbae8a215061b33"
	DataMockUserUuid       = "2882341273"
	DataMockExpire         = uint32(600)
)

func Test_BuildChatUserToken(t *testing.T) {
	token, err := BuildChatUserToken(DataMockAppId, DataMockAppCertificate, DataMockUserUuid, DataMockExpire)
	testutil.Nil(t, err)

	accessToken := accesstoken.CreateAccessToken()
	parsed, err := accessToken.Parse(token)
	testutil.Nil(t, err)
	testutil.Equal(t, true, parsed)

	testutil.Equal(t, DataMockAppId, accessToken.AppId)
	testutil.Equal(t, DataMockExpire, accessToken.Expire)
	testutil.Equal(t, DataMockUserUuid, accessToken.Services[accesstoken.ServiceTypeChat].(*accesstoken.ServiceChat).UserId)
	testutil.Equal(t, DataMockExpire, accessToken.Services[accesstoken.ServiceTypeChat].(*accesstoken.ServiceChat).Privileges[accesstoken.PrivilegeChatUser])
}

func Test_BuildChatAppToken(t *testing.T) {
	token, err := BuildChatAppToken(DataMockAppId, DataMockAppCertificate, DataMockExpire)
	testutil.Nil(t, err)

	accessToken := accesstoken.CreateAccessToken()
	parsed, err := accessToken.Parse(token)
	testutil.Nil(t, err)
	testutil.Equal(t, true, parsed)

	testutil.Equal(t, DataMockAppId, accessToken.AppId)
	testutil.Equal(t, DataMockExpire, accessToken.Expire)
	testutil.Equal(t, DataMockExpire, accessToken.Services[accesstoken.ServiceTypeChat].(*accesstoken.ServiceChat).Privileges[accesstoken.PrivilegeChatApp])

}
