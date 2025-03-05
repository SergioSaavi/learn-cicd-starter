package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey test-key")
	key, err := GetAPIKey(headers)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if key != "test-key" {
		t.Errorf("expected 'test-key', got '%s'", key)
	}
}
