package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		header      string
		expectedKey string
		expectError bool
	}{
		{
			name:        "valid api key",
			header:      "ApiKey test-key-123",
			expectedKey: "test-key-123",
			expectError: false,
		},
		{
			name:        "no authorization header",
			header:      "",
			expectedKey: "",
			expectError: true,
		},
		{
			name:        "malformed header - no space",
			header:      "ApiKeytest-key-123",
			expectedKey: "",
			expectError: true,
		},
		{
			name:        "malformed header - wrong prefix",
			header:      "Bearer test-key-123",
			expectedKey: "",
			expectError: true,
		},
		{
			name:        "malformed header - only prefix",
			header:      "ApiKey",
			expectedKey: "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}
			if tt.header != "" {
				headers.Set("Authorization", tt.header)
			}

			key, err := GetAPIKey(headers)

			if tt.expectError && err == nil {
				t.Error("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if key != tt.expectedKey {
				t.Errorf("expected key %s, got %s", tt.expectedKey, key)
			}
		})
	}
}
