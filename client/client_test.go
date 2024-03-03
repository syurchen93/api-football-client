package client

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	apiKey := "test"
	client := NewClient(apiKey)
	if client.apiKey != apiKey {
		t.Errorf("Expected %s, got %s", apiKey, client.apiKey)
	}
}