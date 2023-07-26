package rtmtokenbuilder2

import (
	accesstoken "github.com/AgoraIO-Community/go-tokenbuilder/accesstoken"
)

// BuildToken creates an access token with the specified parameters.
// If streamName is provided, it adds the RTC service with the given streamName.
//
// Parameters:
//   - appId: The application ID for the access token.
//   - appCertificate: The application certificate for the access token.
//   - userId: The user ID for the access token.
//   - expire: The expiration time in seconds for the access token.
//   - streamName: The name of the stream (optional).
//
// Returns:
//   - The generated access token as a string.
//   - An error if there was an issue generating the token.
func BuildToken(appId string, appCertificate string, userId string, expire uint32, streamName string) (string, error) {
	token := accesstoken.NewAccessToken(appId, appCertificate, expire)

	serviceRtm := accesstoken.NewServiceRtm(userId)
	serviceRtm.AddPrivilege(accesstoken.PrivilegeLogin, expire)
	if streamName != "" {
		// Adds stream data priviliges if streamName != "". Can be "*" for all channels.
		rtcStreamService := accesstoken.NewServiceRtc(streamName, userId)
		rtcStreamService.AddPrivilege(accesstoken.PrivilegeJoinChannel, expire)
		rtcStreamService.AddPrivilege(accesstoken.PrivilegePublishDataStream, expire)
		token.AddService(rtcStreamService)
	}
	token.AddService(serviceRtm)

	return token.Build()
}
