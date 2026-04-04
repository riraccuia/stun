# stun

STUN is a library that allows for complete control over the STUN protocol and some of its ICE-specific extensions.
It can be used to build custom STUN clients and servers, as well as to implement binding checks for services such as WebRTC and QUIC.

A client, server and binding agent implementations are provided in this package, but that should not discourage from using the underlying primitives to build personalized solutions.

Both IPv4 and IPv6 are fully supported, as well as UDP, TCP and TLS protocols.

This package is the base for further work I am doing around STUN and related protocols. I am releasing it as open source so others can reuse the same building blocks and/or contribute to it.

## Installation

```bash
go get github.com/riraccuia/stun
```

## Examples

### Querying a STUN server

There are several ways to query a STUN server, the first is to use one of the convenience functions: QueryServerUDP, QueryServerTCP, QueryServerTLS.

```go
randomPort := rand.Intn(65535-1024) + 1024
result, err := stun.QueryServerUDP("stun.example.com:3478", &net.UDPAddr{IP: nil, Port: randomPort})

if err != nil {
    // handle error
}

fmt.Printf("STUN server returned mapped <ip:port>: %s:%d\n", result.IP, result.Port)


result, err = stun.QueryServerTCP("stun.example.com:3478", &net.TCPAddr{IP: nil, Port: 0})

if err != nil {
	// handle error
}

fmt.Printf("STUN server returned mapped <ip:port>: %s:%d\n", result.IP, result.Port)
```

A generic `QueryServer` method is also available on `Client` and can be used to query a STUN server over any supported protocol. Optional logging uses `WithLoggerFunc`:

```go
logger := func(level string, args ...any) {
	// plug into your logger; level is informational for the client
	_ = level
	_ = args
}
client := stun.NewClient("stun.example.com:3478").WithLoggerFunc(logger)
randomPort := rand.Intn(65535-1024) + 1024
result, err := client.QueryServer("stun.example.com:3478", randomPort, "tcp", nil)

if err != nil {
	// handle error
}

fmt.Printf("STUN server returned mapped <ip:port>: %s:%d\n", result.IP, result.Port)
```

### Build your own STUN client

```go
// connect to the STUN server using the chosen protocol
conn, err := net.Dial("tcp", "stun.example.com:3478")
if err != nil {
	// handle error
}
defer conn.Close()
// build the request
req, err := stun.BuildBindingRequest(nil)
if err != nil {
	// handle error
}
// send the request to the connection
_, err = req.WriteTo(conn)
if err != nil {
	// handle error
}
// receive the response
hdr := make([]byte, 20)
if _, err = io.ReadFull(conn, hdr); err != nil {
	// handle error
}
bodyLen := int(binary.BigEndian.Uint16(hdr[2:4]))
body := make([]byte, bodyLen)
if _, err = io.ReadFull(conn, body); err != nil {
	// handle error
}
raw := append(hdr, body...)
msg, err := stun.DecodeMessage(raw)
if err != nil {
	// handle error
}
result, err := stun.ProcessBindingResponse(msg, req.Header.TransactionID, nil)
if err != nil {
	// handle error
}
fmt.Printf("STUN server returned mapped <ip:port>: %s:%d\n", result.IP, result.Port)
```

## Implemented RFCs

- RFC 3489 - Session Traversal Utilities for NAT (STUN)
- RFC 5389 - Session Traversal Utilities for NAT (STUN)
- RFC 8489 - Session Traversal Utilities for NAT (STUN)
- RFC 8445 - Interactive Connectivity Establishment (ICE) (only STUN extensions are implemented)
- RFC 8265 - Preparation, Enforcement, and Comparison of Internationalized Strings Representing Usernames and Passwords (only OpaqueString profile is implemented)
