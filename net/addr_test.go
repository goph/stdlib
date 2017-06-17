package net_test

import (
	"testing"

	"github.com/goph/stdlib/net"
)

func TestResolveVirtualAddr(t *testing.T) {
	addr := net.ResolveVirtualAddr("network", "addr")

	if got, want := addr.Network(), "network"; got != want {
		t.Errorf("expected %s, received %s", want, got)
	}

	if got, want := addr.String(), "addr"; got != want {
		t.Errorf("expected %s, received %s", want, got)
	}
}
