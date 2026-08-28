# Changelog

All notable changes to this project are documented in this file.

The project follows [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.5.0] - 2026-08-28

### Added

- Added AccessToken2 service type `9` and `convoaitokenbuilder.BuildToken` for generating combined RTC, RTM, and ConvoAI tokens.
- Added AccessToken2 service type `10` and `stttokenbuilder.BuildToken` for generating combined RTC, RTM, and Speech-to-Text tokens.
- Added AccessToken2 support for Streaming (`3`), FCDN (`6`), and resource-permission RTM (`8`) services, bringing the Go implementation in sync with the authoritative Tools implementation.
- Added support for multiple services of the same type through `AccessToken.AddService` and `AccessToken.GetServices` while preserving the v1 `Services` map view.
- Added `rtmtokenbuilder2.BuildTokenWithPermissions` for generating resource-permission RTM tokens.
- Added validated, named `Config` APIs for ConvoAI and STT token generation while retaining the positional builders.
- Added `VerifySignature` to both token implementations for authenticating parsed tokens.
- Added parsing, round-trip, duplicate-service, malformed-input, combined-builder, and Python interoperability coverage.
- Added runnable ConvoAI and STT examples using environment-provided Agora credentials.

### Changed

- Expanded the README with installation, security, supported-service, and builder usage guidance.
- Moved assertion helpers out of production packages and into an internal test-only package.
- Changed AccessToken2 builds without any services to return a `no service added` error.
- Corrected the exported library version to match the module release series.
- Strengthened CI with Go 1.14 and current stable Go coverage, formatting, vet, build, race tests, a coverage floor, and release-tag checks.

### Fixed

- Changed malformed token parsing to return errors instead of panicking and to clear any previous parsed state before reuse.
- Made AccessToken2 service packing deterministic for every supported service type.

## [1.4.0] - 2025-11-12

### Added

- Added AccessToken2 support through the `accesstoken2`, `rtctokenbuilder2`, and `rtmtokenbuilder2` packages.
- Added RTC token generation with separate expiration values for join, audio, video, and data-stream privileges.
- Added combined RTC and RTM AccessToken2 generation.

## [1.3.0] - 2023-07-26

### Added

- Added optional RTC join and data-stream privileges to legacy RTM tokens through the `streamName` argument.

## [1.2.0] - 2023-07-11

### Added

- Added Chat token generation.
- Added automated Go build and test checks.

## [1.1.0] - 2020-08-05

### Changed

- Updated the token implementation to Token007 generation.

## [1.0.0] - 2020-08-03

### Added

- Published the initial stable Go token-builder packages.

[Unreleased]: https://github.com/AgoraIO-Community/go-tokenbuilder/compare/v1.5.0...HEAD
[1.5.0]: https://github.com/AgoraIO-Community/go-tokenbuilder/compare/v1.4.0...v1.5.0
[1.4.0]: https://github.com/AgoraIO-Community/go-tokenbuilder/compare/v1.3.0...v1.4.0
[1.3.0]: https://github.com/AgoraIO-Community/go-tokenbuilder/compare/v1.2.0...v1.3.0
[1.2.0]: https://github.com/AgoraIO-Community/go-tokenbuilder/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/AgoraIO-Community/go-tokenbuilder/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/AgoraIO-Community/go-tokenbuilder/releases/tag/v1.0.0
