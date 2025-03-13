package main

import (
	"net"
	"testing"
)

func TestListen(t *testing.T) {
	listner, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		t.Fatalf("Failed to listen: %v", err)
	}
	defer func() { _ = listner.Close() }()
	t.Logf("Listening on %s", listner.Addr())
}
