package stttokenbuilder

import (
	"testing"

	accesstoken "github.com/AgoraIO-Community/go-tokenbuilder/accesstoken2"
	"github.com/AgoraIO-Community/go-tokenbuilder/internal/testutil"
	rtctokenbuilder "github.com/AgoraIO-Community/go-tokenbuilder/rtctokenbuilder2"
)

const (
	dataMockAppCertificate = "5CFd2fd1755d40ecb72977518be15d3b"
	dataMockAppID          = "970CA35de60c44645bbae8a215061b33"
	dataMockChannelName    = "7d72365eb983485397e3e3f9d460bdda"
	dataMockRTCAccount     = "2882341273"
	dataMockRTMUserID      = "2882341273"
	dataMockExpire         = uint32(600)
)

func TestBuildToken(t *testing.T) {
	token, err := BuildToken(
		dataMockAppID, dataMockAppCertificate, dataMockChannelName, dataMockRTCAccount, rtctokenbuilder.RolePublisher,
		dataMockExpire, dataMockExpire, dataMockExpire, dataMockExpire, dataMockExpire, dataMockRTMUserID, dataMockExpire)
	testutil.Nil(t, err)

	accessToken := accesstoken.CreateAccessToken()
	parsed, err := accessToken.Parse(token)
	testutil.Nil(t, err)
	testutil.Equal(t, true, parsed)
	testutil.Equal(t, dataMockAppID, accessToken.AppId)
	testutil.Equal(t, dataMockExpire, accessToken.Expire)

	rtc := accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc)
	testutil.Equal(t, dataMockChannelName, rtc.ChannelName)
	testutil.Equal(t, dataMockRTCAccount, rtc.Uid)
	testutil.Equal(t, dataMockExpire, rtc.Privileges[accesstoken.PrivilegeJoinChannel])
	testutil.Equal(t, dataMockExpire, rtc.Privileges[accesstoken.PrivilegePublishAudioStream])

	rtm := accessToken.Services[accesstoken.ServiceTypeRtm].(*accesstoken.ServiceRtm)
	testutil.Equal(t, dataMockRTMUserID, rtm.UserId)
	testutil.Equal(t, dataMockExpire, rtm.Privileges[accesstoken.PrivilegeLogin])

	stt := accessToken.Services[accesstoken.ServiceTypeStt].(*accesstoken.ServiceStt)
	testutil.Equal(t, uint16(accesstoken.ServiceTypeStt), stt.Type)
	testutil.Equal(t, 0, len(stt.Privileges))
}

func TestBuildTokenSubscriberDoesNotReceivePublishingPrivileges(t *testing.T) {
	token, err := BuildToken(
		dataMockAppID, dataMockAppCertificate, dataMockChannelName, dataMockRTCAccount, rtctokenbuilder.RoleSubscriber,
		dataMockExpire, dataMockExpire, dataMockExpire, dataMockExpire, dataMockExpire, dataMockRTMUserID, dataMockExpire)
	testutil.Nil(t, err)

	accessToken := accesstoken.CreateAccessToken()
	parsed, err := accessToken.Parse(token)
	testutil.Nil(t, err)
	testutil.Equal(t, true, parsed)
	rtc := accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc)
	testutil.Equal(t, dataMockExpire, rtc.Privileges[accesstoken.PrivilegeJoinChannel])
	testutil.Equal(t, uint32(0), rtc.Privileges[accesstoken.PrivilegePublishAudioStream])
	testutil.Equal(t, uint32(0), rtc.Privileges[accesstoken.PrivilegePublishVideoStream])
	testutil.Equal(t, uint32(0), rtc.Privileges[accesstoken.PrivilegePublishDataStream])
	testutil.Equal(t, true, accessToken.Services[accesstoken.ServiceTypeStt] != nil)
}

func TestBuildTokenInvalidCredentials(t *testing.T) {
	token, err := BuildToken(
		"invalid", dataMockAppCertificate, dataMockChannelName, dataMockRTCAccount, rtctokenbuilder.RolePublisher,
		dataMockExpire, dataMockExpire, dataMockExpire, dataMockExpire, dataMockExpire, dataMockRTMUserID, dataMockExpire)
	testutil.Equal(t, "check appId or appCertificate", err.Error())
	testutil.Equal(t, "", token)

	token, err = BuildToken(
		dataMockAppID, "invalid", dataMockChannelName, dataMockRTCAccount, rtctokenbuilder.RolePublisher,
		dataMockExpire, dataMockExpire, dataMockExpire, dataMockExpire, dataMockExpire, dataMockRTMUserID, dataMockExpire)
	testutil.Equal(t, "check appId or appCertificate", err.Error())
	testutil.Equal(t, "", token)
}

func TestBuildTokenWithConfig(t *testing.T) {
	token, err := BuildTokenWithConfig(Config{
		AppID:                            dataMockAppID,
		AppCertificate:                   dataMockAppCertificate,
		ChannelName:                      dataMockChannelName,
		RTCAccount:                       dataMockRTCAccount,
		RTCRole:                          rtctokenbuilder.RolePublisher,
		RTCTokenExpire:                   dataMockExpire,
		JoinChannelPrivilegeExpire:       dataMockExpire,
		PublishAudioPrivilegeExpire:      dataMockExpire,
		PublishVideoPrivilegeExpire:      dataMockExpire,
		PublishDataStreamPrivilegeExpire: dataMockExpire,
		RTMUserID:                        dataMockRTMUserID,
		RTMTokenExpire:                   dataMockExpire,
	})
	testutil.Nil(t, err)

	parsed := accesstoken.CreateAccessToken()
	ok, err := parsed.Parse(token)
	testutil.Nil(t, err)
	testutil.Equal(t, true, ok)
	testutil.Equal(t, 1, len(parsed.GetServices(accesstoken.ServiceTypeStt)))
}
