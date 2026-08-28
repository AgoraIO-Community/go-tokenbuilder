package rtmtokenbuilder2

import (
	"testing"

	accesstoken "github.com/AgoraIO-Community/go-tokenbuilder/accesstoken2"
	"github.com/AgoraIO-Community/go-tokenbuilder/internal/testutil"
)

const (
	DataMockAppCertificate = "5CFd2fd1755d40ecb72977518be15d3b"
	DataMockAppId          = "970CA35de60c44645bbae8a215061b33"
	DataMockExpire         = uint32(900)
	DataMockUserId         = "test_user"
)

func Test_BuildToken(t *testing.T) {
	token, err := BuildToken(DataMockAppId, DataMockAppCertificate, DataMockUserId, DataMockExpire)
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

func TestBuildTokenWithPermissions(t *testing.T) {
	permissions := accesstoken.NewRtm2Permissions()
	permissions.Add(accesstoken.Rtm2ResourceMessageChannels, accesstoken.Rtm2PermissionRead, []string{"channel-a"})
	permissions.Add(accesstoken.Rtm2ResourceUsers, accesstoken.Rtm2PermissionWrite, []string{"user-a", "user-b"})

	token, err := BuildTokenWithPermissions(DataMockAppId, DataMockAppCertificate, DataMockUserId, permissions, DataMockExpire)
	testutil.Nil(t, err)

	parsed := accesstoken.CreateAccessToken()
	ok, err := parsed.Parse(token)
	testutil.Nil(t, err)
	testutil.Equal(t, true, ok)

	services := parsed.GetServices(accesstoken.ServiceTypeRtm2)
	testutil.Equal(t, 1, len(services))
	service := services[0].(*accesstoken.ServiceRtm2)
	testutil.Equal(t, DataMockUserId, service.UserId)
	testutil.Equal(t, permissions.Details, service.Permissions.Details)
}
