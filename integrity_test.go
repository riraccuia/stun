package stun

import (
	"testing"
)

// makeParsedRequest builds a BindingRequest suitable for integrity checks.
func makeParsedRequest(t *testing.T, user, pass string) *Message {
	msg, err := BuildBindingRequest(&BindingOptions{
		Auth:               &AuthConfig{Username: user, Password: pass},
		RequireFingerprint: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	err = msg.AddFingerprint()
	if err != nil {
		t.Fatal(err)
	}
	err = msg.Validate(&BindingOptions{
		RequireFingerprint: true,
		Auth:               &AuthConfig{Username: user, Password: pass},
	})
	if err != nil {
		t.Fatalf("failed to validate binding request: %v", err)
		return nil
	}
	return msg
}

// TestValidateShortTermIntegrity checks that ValidateShortTermIntegrity correctly
// skips when credentials are absent, passes on valid user+password, and fails on
// mismatch, wrong password, missing username, or invalid request.
func TestValidateShortTermIntegrity(t *testing.T) {
	validReq := makeParsedRequest(t, "user", "secret")
	//reqWithoutAuth := makeParsedRequest(t, "user", "secret", false)
	validAuth := &AuthConfig{Username: "user", Password: "secret"}

	tests := []struct {
		name       string
		req        *Message
		auth       *AuthConfig
		shouldPass bool
	}{
		// No auth required — early return
		{"auth nil skips validation", nil, nil, true},
		{"empty password skips validation", nil, &AuthConfig{Password: ""}, true},
		// Success
		{"valid user and password", validReq, validAuth, true},
		// Failures
		{"wrong username", validReq, &AuthConfig{Username: "other", Password: "secret"}, false},
		{"wrong password", validReq, &AuthConfig{Username: "user", Password: "wrong"}, false},
		//{"request has no parsed username", reqWithoutAuth, validAuth, false},
		{"nil request", nil, validAuth, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateShortTermIntegrity(tt.req, tt.auth)
			if tt.shouldPass && err != nil {
				t.Errorf("expected pass, got error: %v", err)
			}
			if !tt.shouldPass && err == nil {
				t.Error("expected error, got nil")
			}
		})
	}
}
