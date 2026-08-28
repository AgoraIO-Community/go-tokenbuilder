package tokenbuilder

import (
	"strings"
	"testing"

	"github.com/AgoraIO-Community/go-tokenbuilder/rtctokenbuilder2"
)

func validProductTokenConfig() ProductTokenConfig {
	return ProductTokenConfig{
		AppID:                            "970CA35de60c44645bbae8a215061b33",
		AppCertificate:                   "5CFd2fd1755d40ecb72977518be15d3b",
		ChannelName:                      "test-channel",
		RTCAccount:                       "rtc-user",
		RTCRole:                          rtctokenbuilder2.RolePublisher,
		RTCTokenExpire:                   600,
		JoinChannelPrivilegeExpire:       600,
		PublishAudioPrivilegeExpire:      600,
		PublishVideoPrivilegeExpire:      600,
		PublishDataStreamPrivilegeExpire: 600,
		RTMUserID:                        "rtm-user",
		RTMTokenExpire:                   600,
	}
}

func TestProductTokenConfigValidate(t *testing.T) {
	if err := validProductTokenConfig().Validate(); err != nil {
		t.Fatalf("valid config returned error: %v", err)
	}
}

func TestProductTokenConfigValidateErrors(t *testing.T) {
	tests := []struct {
		name    string
		update  func(*ProductTokenConfig)
		message string
	}{
		{"credentials", func(config *ProductTokenConfig) { config.AppID = "invalid" }, "appId"},
		{"channel", func(config *ProductTokenConfig) { config.ChannelName = "" }, "channel name"},
		{"RTC account", func(config *ProductTokenConfig) { config.RTCAccount = "" }, "RTC account"},
		{"RTM user", func(config *ProductTokenConfig) { config.RTMUserID = "" }, "RTM user ID"},
		{"role", func(config *ProductTokenConfig) { config.RTCRole = 99 }, "RTC role"},
		{"token expiration", func(config *ProductTokenConfig) { config.RTCTokenExpire = 0 }, "RTC token expiration"},
		{"privilege expiration", func(config *ProductTokenConfig) { config.RTMTokenExpire = 601 }, "must not exceed"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config := validProductTokenConfig()
			test.update(&config)
			err := config.Validate()
			if err == nil || !strings.Contains(err.Error(), test.message) {
				t.Fatalf("Validate() error = %v, want error containing %q", err, test.message)
			}
		})
	}
}
