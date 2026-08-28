package rtmtokenbuilder2

import (
	"testing"

	accesstoken "github.com/AgoraIO-Community/go-tokenbuilder/accesstoken"
	"github.com/AgoraIO-Community/go-tokenbuilder/internal/testutil"
)

const (
	DataMockAppCertificate = "5CFd2fd1755d40ecb72977518be15d3b"
	DataMockAppId          = "970CA35de60c44645bbae8a215061b33"
	DataMockExpire         = uint32(900)
	DataMockUserId         = "test_user"
)

func Test_BuildToken(t *testing.T) {
	token, err := BuildToken(DataMockAppId, DataMockAppCertificate, DataMockUserId, DataMockExpire, "")
	testutil.Nil(t, err)

	accessToken := accesstoken.CreateAccessToken()
	parsed, err := accessToken.Parse(token)
	testutil.Nil(t, err)
	testutil.Equal(t, true, parsed)

	testutil.Equal(t, DataMockAppId, accessToken.AppId)
	testutil.Equal(t, DataMockExpire, accessToken.Expire)
	testutil.Equal(t, true, accessToken.Services[accesstoken.ServiceTypeRtm] != nil)
	testutil.Equal(t, DataMockUserId, accessToken.Services[accesstoken.ServiceTypeRtm].(*accesstoken.ServiceRtm).UserId)
	testutil.Equal(t, uint16(accesstoken.ServiceTypeRtm), accessToken.Services[accesstoken.ServiceTypeRtm].(*accesstoken.ServiceRtm).Type)
	testutil.Equal(t, DataMockExpire, accessToken.Services[accesstoken.ServiceTypeRtm].(*accesstoken.ServiceRtm).Privileges[accesstoken.PrivilegeLogin])
}

func Test_BuildTokenWithStream(t *testing.T) {
	token, err := BuildToken(DataMockAppId, DataMockAppCertificate, DataMockUserId, DataMockExpire, "*")
	testutil.Nil(t, err)

	accessToken := accesstoken.CreateAccessToken()
	parsed, err := accessToken.Parse(token)
	testutil.Nil(t, err)
	testutil.Equal(t, true, parsed)

	testutil.Equal(t, DataMockAppId, accessToken.AppId)
	testutil.Equal(t, DataMockExpire, accessToken.Expire)
	testutil.Equal(t, true, accessToken.Services[accesstoken.ServiceTypeRtm] != nil)
	testutil.Equal(t, DataMockUserId, accessToken.Services[accesstoken.ServiceTypeRtm].(*accesstoken.ServiceRtm).UserId)
	testutil.Equal(t, uint16(accesstoken.ServiceTypeRtm), accessToken.Services[accesstoken.ServiceTypeRtm].(*accesstoken.ServiceRtm).Type)
	testutil.Equal(t, DataMockExpire, accessToken.Services[accesstoken.ServiceTypeRtm].(*accesstoken.ServiceRtm).Privileges[accesstoken.PrivilegeLogin])
	testutil.Equal(t, DataMockExpire, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegePublishDataStream])
	testutil.Equal(t, DataMockExpire, accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc).Privileges[accesstoken.PrivilegeJoinChannel])
}
