package ext_test

import (
	"fmt"
	"testing"

	"github.com/goph/stdlib/errors"
	"github.com/goph/stdlib/ext"
	"github.com/goph/stdlib/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCloser struct {
	called bool
	err    error
}

func (c *testCloser) Close() error {
	c.called = true

	return c.err
}

func TestCloserFunc_CallsUnderlyingFunc(t *testing.T) {
	var called bool

	closer := ext.CloserFunc(func() {
		called = true
	})

	var err error

	if got, want := closer.Close(), err; got != want {
		t.Errorf("wrapped functions are expected to return nil, error received: %v", got)
	}

	if called != true {
		t.Error("the wrapped function is expected to be called")
	}
}

func TestCloserFunc_RecoversErrorPanic(t *testing.T) {
	err := fmt.Errorf("internal error")

	closer := ext.CloserFunc(func() {
		panic(err)
	})

	if got, want := closer.Close(), err; got != want {
		t.Errorf("expected to recover a specific error, received: %v", got)
	}
}

func TestClosers(t *testing.T) {
	closer1 := &mocks.Closer{}
	closer1.On("Close").Return(nil)

	closer2 := &mocks.Closer{}
	closer2.On("Close").Return(nil)

	closer := ext.Closers{closer1, closer2}

	assert.NoError(t, closer.Close())
	closer1.AssertExpectations(t)
	closer2.AssertExpectations(t)
}

func TestClosers_Empty(t *testing.T) {
	closer := ext.Closers{}

	assert.NoError(t, closer.Close())
}

func TestClosers_Error(t *testing.T) {
	closer1 := &mocks.Closer{}
	closer1.On("Close").Return(nil)

	err := fmt.Errorf("error")
	closer2 := &mocks.Closer{}
	closer2.On("Close").Return(err)

	closer := ext.Closers{closer1, closer2}

	merr := closer.Close()

	require.Error(t, merr)
	require.Implements(t, (*errors.ErrorCollection)(nil), merr)
	assert.Contains(t, merr.(errors.ErrorCollection).Errors(), err)

	closer1.AssertExpectations(t)
	closer2.AssertExpectations(t)
}

func TestClose(t *testing.T) {
	err := fmt.Errorf("error")
	closer := ext.CloserFunc(func() { panic(err) })

	assert.Equal(t, err, closer.Close())
}
