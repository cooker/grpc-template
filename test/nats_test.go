package test

import (
	"github.com/nats-io/nats.go"
	"testing"
)

func TestClient(t *testing.T) {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))

	// Close connection
	nc.Close()
}
