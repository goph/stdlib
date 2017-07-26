package net_test

import (
	"testing"

	stdnet "net"
	"sync"

	"github.com/goph/stdlib/net"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPipeListen(t *testing.T) {
	addr := net.ResolveVirtualAddr("network", "addr")

	listener, dialer := net.PipeListen(addr)

	var wg sync.WaitGroup

	var clientConn, serverConn stdnet.Conn

	writtenBytes := []byte("piped")
	var readBytes = make([]byte, len(writtenBytes))
	var written, read int

	wg.Add(2)

	go func() {
		defer wg.Done()

		var err error

		clientConn, err = dialer.Dial()

		require.NoError(t, err, "cannot dial: %v")

		written, err = clientConn.Write(writtenBytes)

		require.NoError(t, err, "cannot write: %v")
	}()

	go func() {
		defer wg.Done()

		var err error

		serverConn, err = listener.Accept()

		require.NoError(t, err, "cannot accept: %v")

		read, err = serverConn.Read(readBytes)

		require.NoError(t, err, "cannot write: %v")
	}()

	wg.Wait()

	assert.Equal(t, written, read, "data size mismatch, written %v bytes, read %v bytes")
	assert.Equal(t, writtenBytes, readBytes, "data mismatch, written %v, read %v")
}
