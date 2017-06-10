package net

import (
	"testing"
)

func TestResolveVirtualAddr(t *testing.T) {
	addr := ResolveVirtualAddr("network", "addr")

	if got, want := addr.Network(), "network"; got != want {
		t.Errorf("expected %s, received %s", want, got)
	}

	if got, want := addr.String(), "addr"; got != want {
		t.Errorf("expected %s, received %s", want, got)
	}
}
