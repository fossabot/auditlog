# 0.9.3: Audit log configuration now has an `Enable` flag (November 27, 2020)

We have added a new `Enable` flag to the [configuration](config.go). If the flag is set to `false` (default) the `New()` function creates an empty logger to save CPU cycles. Callers are encouraged to use this flag to completely disable audit logging.

# 0.9.2: Message for handshake success / failure (November 26, 2020)

There are now two additional message types: `TypeHandshakeFailed` (opcode 198) and `TypeHandshakeSuccessful` (opcode 199).

The `NewEncoder()` method for the `asciinema` and `binary` and the `New()` method for the `auditlog` package now have an added dependency to the [GeoIP library](https://github.com/containerssh/geoip).

# 0.9.1: Channel ID, Request ID, and better data formatting (November 15, 2020)

This release changes the audit log format to better fit with the [sshserver library 0.9.2 release](https://github.com/ContainerSSH/sshserver/releases/tag/v0.9.2).

The following changes have been made:

- All payloads for channels now contain the `RequestID` field of the type `uint64` as a first parameter.
- The `ChannelID` field has been changed to `*uint64` (pointer to an `uint64`). Previously, this field was `-1` if no channel ID was set.
- The `PayloadChannelRequestUnknownType` and `PayloadChannelRequestDecodeFailed` payloads  now contain a `Payload` field that contains the request payload for later analysis.
- The `PayloadChannelRequestPty` payload now contains additional fields for `Term`, `Width`, `Height`, and `Modelist`.
- The `PayloadChannelRequestWindow` payload now contains additional fields for `Width` and `Height`.
- The dimension fields in `PayloadChannelRequestPty` and `PayloadChannelRequestWindow` have been changed to `uint32` to match the SSH specification.
- A new message type `TypeRequestFailed` with the payload `PayloadRequestFailed` have been introduced to indicate that a channel-specific or global request has failed. The payload contains the request ID that allows for identifying the request that failed in case of multiple, parallel requests.
- A new message type `TypeExit` with the payload `TypeRequestFailed` has been introduced to indicate the program exit code that has been sent to the client.

# 0.9.0: Initial release (November 7, 2020)

This is the initial release of the audit log library.