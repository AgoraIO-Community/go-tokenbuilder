package stttokenbuilder

import (
	"testing"

	accesstoken "github.com/AgoraIO-Community/go-tokenbuilder/accesstoken2"
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
	accesstoken.AssertNil(t, err)

	accessToken := accesstoken.CreateAccessToken()
	parsed, err := accessToken.Parse(token)
	accesstoken.AssertNil(t, err)
	accesstoken.AssertEqual(t, true, parsed)
	accesstoken.AssertEqual(t, dataMockAppID, accessToken.AppId)
	accesstoken.AssertEqual(t, dataMockExpire, accessToken.Expire)

	rtc := accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc)
	accesstoken.AssertEqual(t, dataMockChannelName, rtc.ChannelName)
	accesstoken.AssertEqual(t, dataMockRTCAccount, rtc.Uid)
	accesstoken.AssertEqual(t, dataMockExpire, rtc.Privileges[accesstoken.PrivilegeJoinChannel])
	accesstoken.AssertEqual(t, dataMockExpire, rtc.Privileges[accesstoken.PrivilegePublishAudioStream])

	rtm := accessToken.Services[accesstoken.ServiceTypeRtm].(*accesstoken.ServiceRtm)
	accesstoken.AssertEqual(t, dataMockRTMUserID, rtm.UserId)
	accesstoken.AssertEqual(t, dataMockExpire, rtm.Privileges[accesstoken.PrivilegeLogin])

	stt := accessToken.Services[accesstoken.ServiceTypeStt].(*accesstoken.ServiceStt)
	accesstoken.AssertEqual(t, uint16(accesstoken.ServiceTypeStt), stt.Type)
	accesstoken.AssertEqual(t, 0, len(stt.Privileges))
}

func TestBuildTokenSubscriberDoesNotReceivePublishingPrivileges(t *testing.T) {
	token, err := BuildToken(
		dataMockAppID, dataMockAppCertificate, dataMockChannelName, dataMockRTCAccount, rtctokenbuilder.RoleSubscriber,
		dataMockExpire, dataMockExpire, dataMockExpire, dataMockExpire, dataMockExpire, dataMockRTMUserID, dataMockExpire)
	accesstoken.AssertNil(t, err)

	accessToken := accesstoken.CreateAccessToken()
	parsed, err := accessToken.Parse(token)
	accesstoken.AssertNil(t, err)
	accesstoken.AssertEqual(t, true, parsed)
	rtc := accessToken.Services[accesstoken.ServiceTypeRtc].(*accesstoken.ServiceRtc)
	accesstoken.AssertEqual(t, dataMockExpire, rtc.Privileges[accesstoken.PrivilegeJoinChannel])
	accesstoken.AssertEqual(t, uint32(0), rtc.Privileges[accesstoken.PrivilegePublishAudioStream])
	accesstoken.AssertEqual(t, uint32(0), rtc.Privileges[accesstoken.PrivilegePublishVideoStream])
	accesstoken.AssertEqual(t, uint32(0), rtc.Privileges[accesstoken.PrivilegePublishDataStream])
	accesstoken.AssertEqual(t, true, accessToken.Services[accesstoken.ServiceTypeStt] != nil)
}

func TestBuildTokenInvalidCredentials(t *testing.T) {
	token, err := BuildToken(
		"invalid", dataMockAppCertificate, dataMockChannelName, dataMockRTCAccount, rtctokenbuilder.RolePublisher,
		dataMockExpire, dataMockExpire, dataMockExpire, dataMockExpire, dataMockExpire, dataMockRTMUserID, dataMockExpire)
	accesstoken.AssertEqual(t, "check appId or appCertificate", err.Error())
	accesstoken.AssertEqual(t, "", token)

	token, err = BuildToken(
		dataMockAppID, "invalid", dataMockChannelName, dataMockRTCAccount, rtctokenbuilder.RolePublisher,
		dataMockExpire, dataMockExpire, dataMockExpire, dataMockExpire, dataMockExpire, dataMockRTMUserID, dataMockExpire)
	accesstoken.AssertEqual(t, "check appId or appCertificate", err.Error())
	accesstoken.AssertEqual(t, "", token)
}
