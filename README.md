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
| ConvoAI | `convoaitokenbuilder` | 9 |
| Speech-to-Text (STT) | `stttokenbuilder` | 10 |

ConvoAI and STT tokens contain RTC and RTM services plus their respective product service. The ConvoAI and STT services have no privilege payload; their presence in the signed token grants access to that service.

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

	token, err := convoaitokenbuilder.BuildToken(
		os.Getenv("AGORA_APP_ID"),
		os.Getenv("AGORA_APP_CERTIFICATE"),
		"my-channel",
		"rtc-user",
		rtctokenbuilder2.RolePublisher,
		expire,
		expire,
		expire,
		expire,
		expire,
		"rtm-user",
		expire,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(token)
}
```

## STT token

STT uses the same arguments and expiration model:

```go
token, err := stttokenbuilder.BuildToken(
	appID,
	appCertificate,
	channelName,
	rtcAccount,
	rtctokenbuilder2.RolePublisher,
	rtcTokenExpire,
	joinChannelPrivilegeExpire,
	pubAudioPrivilegeExpire,
	pubVideoPrivilegeExpire,
	pubDataStreamPrivilegeExpire,
	rtmUserID,
	rtmTokenExpire,
)
```

See the complete runnable [ConvoAI](examples/convoaitokenbuilder/main.go) and [STT](examples/stttokenbuilder/main.go) examples.

## Existing builders

- `rtctokenbuilder2` generates RTC AccessToken2 tokens for numeric UIDs or user accounts, with either shared or individual privilege expirations.
- `rtmtokenbuilder2` generates RTM AccessToken2 tokens.
- `rtctokenbuilder2.BuildTokenWithRtm` and `BuildTokenWithRtm2` generate combined RTC and RTM tokens.
- Legacy packages remain available for backward compatibility.

For new integrations, prefer the packages ending in `2` and the ConvoAI or STT builders introduced in `v1.5.0`.

## Development

Run all checks from the repository root:

```bash
go fmt ./...
go vet ./...
go test ./...
```

## Releases

See [CHANGELOG.md](CHANGELOG.md) for release notes. Releases follow semantic versioning and are published as Git tags such as `v1.5.0`.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE).
