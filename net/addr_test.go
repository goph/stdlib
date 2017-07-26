package net_test

import (
	"testing"

	"github.com/goph/stdlib/net"
	"github.com/stretchr/testify/assert"
)

func TestResolveVirtualAddr(t *testing.T) {
	addr := net.ResolveVirtualAddr("network", "addr")

	assert.Equal(t, "network", addr.Network())
	assert.Equal(t, "addr", addr.String())
}
