// Copyright 2026 Riccardo Raccuia
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
# Querying a STUN server

There are several ways to do this, the first is to use one of the convenience functions: QueryServerUDP, QueryServerTCP, QueryServerTLS.

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

A generic QueryServer method on Client can be used to query a STUN server over any supported protocol. Optional logging uses WithLoggerFunc.

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

# Build your own STUN client

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

# Serializing STUN messages

	// prepare attributes for the message
	attrs := []stun.Attribute{
	    stun.NewMappedAddressAttr(net.ParseIP("192.168.1.1"), 12345)
		stun.NewXorMappedAddressAttr(net.ParseIP("192.168.1.1"), 12345)
	}
	// build the message
	msg, err := stun.NewMessage(stun.BindingRequest, attrs...)
	if err != nil {
		// handle error
	}
	// write the message to a connection
	_, err = msg.WithFingerprint().WriteTo(conn)
	if err != nil {
		// handle error
	}

# Decoding STUN messages and consuming attributes

	msg, err := stun.DecodeMessage(data)
	if err != nil {
		// handle error
	}
	// find the MAPPED-ADDRESS attribute
	mappedAddr, err := msg.FindAttribute(stun.AttrMappedAddress)
	if err != nil {
		// handle error
	}
	fmt.Printf("MAPPED-ADDRESS: %+v\n", mappedAddr)
	// print all attributes names for attributes carried by the message
	attrs := msg.Attributes
	for _, attr := range attrs {
		fmt.Printf("Attribute: %s\n", stun.AttributeToString[attr.GetType()])
	}

# NAT behavior discovery

In order for NAT behavior discovery to work, it is important to pick a stun server that can dispose of a true alternate address and port.
Even though some servers might respond with an OTHER-ADDRESS attribute, and therefore seem suitable, the IP address carried by it matches
in many cases the one of the currently connected socket. See RFC 5780 Section 3 for more details.
The following repository is a good starting point to find suitable servers for NAT behavior discovery:
https://github.com/pradt2/always-online-stun/blob/master/valid_nat_testing_hosts.txt

	result, err := stun.DiscoverNATBehavior("stun.example.com:3478", nil)
	if err != nil {
		// handle error
	}
	fmt.Printf("NAT behavior discovery result: %+v\n", result)
*/
package stun
