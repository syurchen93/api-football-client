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

func TestSetBaseURL(t *testing.T) {
	baseURL := "https://test.com"
	client := NewClient("test")
	client.SetBaseURL(baseURL)
	if client.baseURL != baseURL {
		t.Errorf("Expected %s, got %s", baseURL, client.baseURL)
	}
}

func TestSetApiHost(t *testing.T) {
	apiHost := "test.com"
	client := NewClient("test")
	client.SetApiHost(apiHost)
	if client.apiHost != apiHost {
		t.Errorf("Expected %s, got %s", apiHost, client.apiHost)
	}
}