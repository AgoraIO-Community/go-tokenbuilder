# Version 2 migration plan

Version 1.5 keeps existing import paths and exported field names stable so the AccessToken2 service update can ship without an unrelated migration.

The next major release is the appropriate boundary for these breaking cleanups:

- Add `/v2` to the module path.
- Normalize package paths to all-lowercase Go naming, including `chatTokenBuilder` to `chattokenbuilder`.
- Normalize exported initialisms such as `AppId`, `UserId`, and `Uid` to `AppID`, `UserID`, and `UID`.
- Replace the compatibility `AccessToken.Services` map with the ordered service collection used by `AddService` and `GetServices`.
- Remove legacy token packages and other deprecated v1 compatibility surfaces only after publishing equivalent migration examples.

These changes are intentionally deferred and are not release blockers for v1.5.0.
