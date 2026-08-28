# Agora Go Token Builder

`go-tokenbuilder` generates server-side authentication tokens for Agora services. It supports the current AccessToken2 (`Token007`) format as well as the repository's legacy token packages.

## Installation

Install the latest release:

```bash
go get github.com/AgoraIO-Community/go-tokenbuilder@v1.5.0
```

The module declares Go 1.14 compatibility.

## Security

Generate tokens on a trusted server. Never embed or expose your Agora App Certificate in browser, mobile, desktop, or other client-side code.

Use short expiration periods and require the channel name and user identity in the token to match the values used by the Agora client.

## Supported token builders

| Service | Package | AccessToken2 service type |
| --- | --- | ---: |
| RTC | `rtctokenbuilder2` | 1 |
| RTM | `rtmtokenbuilder2` | 2 |
| Streaming | `accesstoken2` | 3 |
| FPA | `accesstoken2` | 4 |
| Chat | `accesstoken2` | 5 |
| FCDN | `accesstoken2` | 6 |
| APaaS | `accesstoken2` | 7 |
| RTM with resource permissions | `rtmtokenbuilder2` | 8 |
| ConvoAI | `convoaitokenbuilder` | 9 |
| Speech-to-Text (STT) | `stttokenbuilder` | 10 |

ConvoAI and STT tokens contain RTC and RTM services plus their respective product service. The ConvoAI and STT services have no privilege payload; their presence in the signed token grants access to that service.

The `accesstoken2` package supports multiple services of the same type through `AddService` and `GetServices`. Its public `Services` map remains available in v1 as a backward-compatible view of the last service for each type.

## ConvoAI token

```go
package main

import (
	"log"
	"os"

	"github.com/AgoraIO-Community/go-tokenbuilder/convoaitokenbuilder"
	"github.com/AgoraIO-Community/go-tokenbuilder/rtctokenbuilder2"
)

func main() {
	const expire = uint32(3600)

	token, err := convoaitokenbuilder.BuildTokenWithConfig(convoaitokenbuilder.Config{
		AppID:                            os.Getenv("AGORA_APP_ID"),
		AppCertificate:                   os.Getenv("AGORA_APP_CERTIFICATE"),
		ChannelName:                      "my-channel",
		RTCAccount:                       "rtc-user",
		RTCRole:                          rtctokenbuilder2.RolePublisher,
		RTCTokenExpire:                   expire,
		JoinChannelPrivilegeExpire:       expire,
		PublishAudioPrivilegeExpire:      expire,
		PublishVideoPrivilegeExpire:      expire,
		PublishDataStreamPrivilegeExpire: expire,
		RTMUserID:                        "rtm-user",
		RTMTokenExpire:                   expire,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Print(token)
}
```

## STT token

STT uses the same validated configuration and expiration model:

```go
token, err := stttokenbuilder.BuildTokenWithConfig(stttokenbuilder.Config{
	AppID:                            appID,
	AppCertificate:                   appCertificate,
	ChannelName:                      channelName,
	RTCAccount:                       rtcAccount,
	RTCRole:                          rtctokenbuilder2.RolePublisher,
	RTCTokenExpire:                   rtcTokenExpire,
	JoinChannelPrivilegeExpire:       joinChannelPrivilegeExpire,
	PublishAudioPrivilegeExpire:      pubAudioPrivilegeExpire,
	PublishVideoPrivilegeExpire:      pubVideoPrivilegeExpire,
	PublishDataStreamPrivilegeExpire: pubDataStreamPrivilegeExpire,
	RTMUserID:                        rtmUserID,
	RTMTokenExpire:                   rtmTokenExpire,
})
```

The config builders validate identifier lengths, credentials, roles, and expiration relationships before building. The positional `BuildToken` functions remain available for compatibility.

See the complete runnable [ConvoAI](examples/convoaitokenbuilder/main.go) and [STT](examples/stttokenbuilder/main.go) examples.

## Existing builders

- `rtctokenbuilder2` generates RTC AccessToken2 tokens for numeric UIDs or user accounts, with either shared or individual privilege expirations.
- `rtmtokenbuilder2` generates RTM AccessToken2 tokens.
- `rtmtokenbuilder2.BuildTokenWithPermissions` generates service type 8 RTM tokens with resource-level permissions; coordinate with Agora before using this specialized interface.
- `rtctokenbuilder2.BuildTokenWithRtm` and `BuildTokenWithRtm2` generate combined RTC and RTM tokens.
- Legacy packages remain available for backward compatibility.

For new integrations, prefer the packages ending in `2` and the ConvoAI or STT builders introduced in `v1.5.0`.

## Parsing and signature verification

`Parse` decodes a token but does not establish that it is authentic. Call `VerifySignature` with the App Certificate after parsing when accepting a token from an untrusted source:

```go
parsed := accesstoken2.CreateAccessToken()
if ok, err := parsed.Parse(token); err != nil || !ok {
	return err
}
valid, err := parsed.VerifySignature(appCertificate)
```

Malformed tokens return errors and clear previously parsed signature state.

## Development

Run all checks from the repository root:

```bash
go fmt ./...
go vet ./...
go test ./...
```

## Releases

See [CHANGELOG.md](CHANGELOG.md) for release notes. Releases follow semantic versioning and are published as Git tags such as `v1.5.0`.

Breaking naming and API normalization is intentionally deferred to the next major release; see [V2_MIGRATION.md](V2_MIGRATION.md).

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE).
