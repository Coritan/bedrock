# Universal Protocol Compatibility

This proxy has been modified to support **any Minecraft Bedrock Edition protocol version** for passthrough proxying.

## Changes Made

### 1. Universal Protocol Handler (`universal_protocol.go`)
A new protocol implementation that:
- Accepts any protocol version ID
- Passes packets through without modification
- Doesn't require protocol-specific packet conversion

### 2. Main Configuration (`main.go`)
The listener now accepts a wide range of protocol versions:

```go
AcceptedProtocols:    generateProtocolRange(400, 800)
AllowUnknownPackets:  true
AllowInvalidPackets:  true
```

**Protocol Range Coverage:**
- Protocol 400-800 covers approximately Minecraft versions 1.16.0 through 1.22.0+
- Adjust the range in `main.go` if you need to support older or newer versions

### 3. Packet Forwarding
The proxy forwards packets at the packet level without inspecting or modifying contents:
- Uses `ReadPacket()` and `WritePacket()` for efficient forwarding
- All game packets pass through transparently
- Authentication is still handled for security

## Adjusting Protocol Support

To change which Minecraft versions are supported, edit the range in `main.go`:

```go
// Support versions 1.16 to 1.25+
AcceptedProtocols: generateProtocolRange(400, 900)

// Support only modern versions (1.20+)
AcceptedProtocols: generateProtocolRange(600, 800)

// Support all versions including legacy
AcceptedProtocols: generateProtocolRange(200, 1000)
```

### Common Protocol IDs

| Minecraft Version | Protocol ID (approx) |
|-------------------|---------------------|
| 1.16.x            | 400-450             |
| 1.17.x            | 450-480             |
| 1.18.x            | 480-520             |
| 1.19.x            | 520-560             |
| 1.20.x            | 560-650             |
| 1.21.x            | 650-700             |

## How It Works

1. Client connects with their protocol version
2. Proxy accepts any version in the configured range
3. Proxy authenticates client (validates XBOX Live if enabled)
4. Proxy forwards all packets to remote server unchanged
5. Server responds with packets that proxy forwards back to client

The proxy acts as a transparent passthrough for all game packets while handling the initial connection handshake.

## Security Notes

- Authentication is still enforced by default (`AuthenticationDisabled: false`)
- XBOX Live validation ensures only legitimate clients connect
- Set `AuthenticationDisabled: true` if you need to bypass authentication (not recommended)

## Testing

To test with different Minecraft versions:
1. Start the proxy
2. Connect from any Minecraft client version in the supported range
3. The proxy will accept the connection and forward to the remote server
4. Check console logs for protocol version information
