package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name           string
		headers        http.Header
		expectedKey    string
		expectedErr    error
		expectedErrMsg string // for non-named errors
	}{
		{
			name:        "Successful Retrieval",
			headers:     http.Header{"Authorization": {"ApiKey valid_api_key"}},
			expectedKey: "valid_api_key",
			expectedErr: nil,
		},
		{
			name:        "Missing API Key",
			headers:     http.Header{},
			expectedKey: "",
			expectedErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:           "Invalid API Key Format",
			headers:        http.Header{"Authorization": {"InvalidHeaderFormat"}},
			expectedKey:    "",
			expectedErrMsg: "malformed authorization header",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetAPIKey(tt.headers)

			// Assert key matches
			if result != tt.expectedKey {
				t.Errorf("expected API key %q, got: %q", tt.expectedKey, result)
			}

			// Assert error type or message
			if tt.expectedErr != nil {
				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("expected error %v, got: %v", tt.expectedErr, err)
				}
			} else if tt.expectedErrMsg != "" {
				if err == nil || err.Error() != tt.expectedErrMsg {
					t.Errorf("expected error message %q, got: %v", tt.expectedErrMsg, err)
				}
			} else if err != nil {
				t.Errorf("expected no error, got: %v", err)
			}
		})
	}
}
