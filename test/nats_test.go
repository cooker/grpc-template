package test

import (
	"github.com/nats-io/nats.go"
	"testing"
)

func TestClient(t *testing.T) {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)
	t.Log("Nats 消息发送")
	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))
	nc.Subscribe("foo", func(msg *nats.Msg) {
		t.Log("收到消息：", string(msg.Data))
	})
	select {}
	// Close connection
	nc.Close()
}
