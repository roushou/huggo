package huggo

import (
	"net/http"
	"testing"
)

func TestNewHttpClient(t *testing.T) {
	client, _ := NewHttpClient("apiKey")
	if client.baseURL != DefaultAPIBaseURL {
		t.Errorf("Expected base URL %s, got %s", DefaultAPIBaseURL, client.baseURL)
	}
	if client.httpClient == nil {
		t.Errorf("Expected HTTP client to not be nil")
	}

	customClient, _ := NewHttpClient("apiKey", WithBaseURL("customURL"), WithAPIKey("customApiKey"))
	if customClient.baseURL != "customURL" {
		t.Errorf("Expected custom base URL, got %s", customClient.baseURL)
	}
	if customClient.apiKey != "customApiKey" {
		t.Errorf("Expected custom API key, got %s", customClient.apiKey)
	}
}

func TestNewRequest(t *testing.T) {
	client, _ := NewHttpClient("apiKey")
	req, err := client.newRequest(http.MethodGet, "/hello", nil)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if req.URL.String() != DefaultAPIBaseURL+"/hello" {
		t.Errorf("Expected URL %s, got %s", DefaultAPIBaseURL+"/hello", req.URL.String())
	}
	if req.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Expected 'Content-Type' header to be set to 'application/json', got %s", req.Header.Get("Content-Type"))
	}
	if req.Header.Get("Authorization") != "Bearer apiKey" {
		t.Errorf("Expected 'Authorization' header to be set to 'Bearer apiKey', got %s", req.Header.Get("Authorization"))
	}
}
