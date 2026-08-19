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

package stun

import (
	"net"
)

// mockStunServer wraps a net.Listener. For each connection returned by Accept,
// it creates a BindingAgent with the given config, attaches the conn, and
// starts Receive() to handle incoming binding requests.
type mockStunServer struct {
	listener net.Listener
	config   *BindingAgentConfig
}

func newMockStunServer(auth *IceAuth) *mockStunServer {
	config := NewControlledICEBindingAgentConfig(NoopLoggerFunc, auth, 0)
	return &mockStunServer{config: config}
}

func (s *mockStunServer) start() error {
	var err error
	s.listener, err = net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	go func() {
		for {
			conn, err := s.listener.Accept()
			if err != nil {
				return
			}
			agent := NewBindingAgent(s.config)
			agent.SetConn(conn)
			agent.Receive()
		}
	}()
	return nil
}

func (s *mockStunServer) stop() {
	if s.listener != nil {
		s.listener.Close()
	}
}

func (s *mockStunServer) address() string {
	return s.listener.Addr().String()
}
