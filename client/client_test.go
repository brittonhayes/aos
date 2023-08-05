package client

import (
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient(&http.Client{}, "http://localhost:8080", nil)
	if c == nil {
		t.Error("Expected client to be not nil")
	}
}
