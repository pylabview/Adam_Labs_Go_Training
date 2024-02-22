package client

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

type mockTripper struct{}

func (m mockTripper) RoundTrip(*http.Request) (*http.Response, error) {
	log.Printf("mockTripper was here")
	return nil, fmt.Errorf("bad health")
}

func TestClient_HealthError(t *testing.T) {
	c := Client{baseURL: "http://localhost:9999"}
	c.c.Transport = mockTripper{}
	err := c.Health()
	if err == nil {
		t.Fatal("no error")
	}
}
