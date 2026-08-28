package tokenbuilder

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/AgoraIO-Community/go-tokenbuilder/rtctokenbuilder2"
)

// MaxTokenLifetimeSeconds is the maximum lifetime accepted by the validated product-token builders.
const MaxTokenLifetimeSeconds = uint32(24 * 60 * 60)

// ProductTokenConfig contains the shared RTC and RTM settings used by product tokens.
type ProductTokenConfig struct {
	AppID                            string
	AppCertificate                   string
	ChannelName                      string
	RTCAccount                       string
	RTCRole                          rtctokenbuilder2.Role
	RTCTokenExpire                   uint32
	JoinChannelPrivilegeExpire       uint32
	PublishAudioPrivilegeExpire      uint32
	PublishVideoPrivilegeExpire      uint32
	PublishDataStreamPrivilegeExpire uint32
	RTMUserID                        string
	RTMTokenExpire                   uint32
}

// Validate checks credentials, identities, roles, and expiration relationships.
func (config ProductTokenConfig) Validate() error {
	if !isHexIdentifier(config.AppID) || !isHexIdentifier(config.AppCertificate) {
		return errors.New("check appId or appCertificate")
	}
	if len([]byte(config.ChannelName)) == 0 || len([]byte(config.ChannelName)) > 64 {
		return errors.New("channel name must contain between 1 and 64 bytes")
	}
	if len([]byte(config.RTCAccount)) == 0 || len([]byte(config.RTCAccount)) > 255 {
		return errors.New("RTC account must contain between 1 and 255 bytes")
	}
	if len([]byte(config.RTMUserID)) == 0 || len([]byte(config.RTMUserID)) > 255 {
		return errors.New("RTM user ID must contain between 1 and 255 bytes")
	}
	if config.RTCRole != rtctokenbuilder2.RolePublisher && config.RTCRole != rtctokenbuilder2.RoleSubscriber {
		return errors.New("RTC role must be RolePublisher or RoleSubscriber")
	}
	if config.RTCTokenExpire == 0 || config.RTCTokenExpire > MaxTokenLifetimeSeconds {
		return fmt.Errorf("RTC token expiration must contain between 1 and %d seconds", MaxTokenLifetimeSeconds)
	}

	expirations := []struct {
		name  string
		value uint32
	}{
		{"join-channel privilege", config.JoinChannelPrivilegeExpire},
		{"publish-audio privilege", config.PublishAudioPrivilegeExpire},
		{"publish-video privilege", config.PublishVideoPrivilegeExpire},
		{"publish-data-stream privilege", config.PublishDataStreamPrivilegeExpire},
		{"RTM token", config.RTMTokenExpire},
	}
	for _, expiration := range expirations {
		if expiration.value > config.RTCTokenExpire {
			return fmt.Errorf("%s expiration must not exceed RTC token expiration", expiration.name)
		}
	}
	return nil
}

func isHexIdentifier(value string) bool {
	if len(value) != 32 {
		return false
	}
	_, err := hex.DecodeString(value)
	return err == nil
}
