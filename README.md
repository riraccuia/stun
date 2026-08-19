# stun

Stun is a Go library that provides primitives for the STUN protocol and its ICE extension.

Client and server implementations are provided as convenience layers. This should not discourage from using the underlying primitives to build personalized solutions.

Both IPv4 and IPv6 are fully supported, as well as UDP, TCP and TLS protocols.

This repository is the base for further work I am doing around STUN and related protocols. I am releasing it as open source so others can reuse the same building blocks and help making it better if they want to.

Also see the Go package documentation at [https://pkg.go.dev/github.com/riraccuia/stun](https://pkg.go.dev/github.com/riraccuia/stun) and examples in the `doc.go` file.

## Usage

Import the package in your project:

```bash
go get github.com/riraccuia/stun@latest
```

## Features

- UDP, TCP and TLS support.
- IPv4 and IPv6 support.
- One shot query functions for UDP, TCP and TLS.
- Client implementation with logging and binding options customization.
- Server implementation with configuration.
- Binding agent with ICE support.
- NAT behavior discovery.

## Supported Attributes

**RFC 8489**

MAPPED-ADDRESS
USERNAME
MESSAGE-INTEGRITY
ERROR-CODE
REALM
NONCE
MESSAGE-INTEGRITY-SHA256
XOR-MAPPED-ADDRESS
FINGERPRINT

**RFC 5780**

CHANGE-REQUEST
OTHER-ADDRESS

**RFC 8445**

PRIORITY
USE-CANDIDATE
ICE-CONTROLLED
ICE-CONTROLLING

## Supported RFCs

- RFC 3489 - Session Traversal Utilities for NAT (STUN)
- RFC 5389 - Session Traversal Utilities for NAT (STUN)
- RFC 8489 - Session Traversal Utilities for NAT (STUN)
- RFC 8445 - Interactive Connectivity Establishment (ICE) (only STUN extensions are implemented)
- RFC 8265 - Preparation, Enforcement, and Comparison of Internationalized Strings Representing Usernames and Passwords (only OpaqueString profile is implemented)
- RFC 5780 - NAT Behavior Discovery Using Session Traversal Utilities for NAT
