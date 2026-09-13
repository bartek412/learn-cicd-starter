package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	test := []struct {
		name        string
		header      http.Header
		expectedKey string
		expectedErr error
	}{
		{
			name:        "ValidApiKey",
			header:      http.Header{"Authorization": []string{"ApiKey my-secret-api-key"}},
			expectedKey: "my-secret-api-key",
			expectedErr: nil,
		},
		{
			name:        "EmptyHeader",
			header:      http.Header{},
			expectedKey: "",
			expectedErr: ErrNoAuthHeaderIncluded,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, err := GetAPIKey(tt.header)
			if err != tt.expectedErr {
				t.Fatalf("failed error, error = %s expectedErr= %s", err, tt.expectedErr)
			}
			if gotKey != tt.expectedKey {
				t.Fatalf("failed wrong key, key = %s, expectedKey = %s", gotKey, tt.expectedKey)
			}
		})
	}
}
