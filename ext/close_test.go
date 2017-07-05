package ext_test

import (
	"fmt"
	"testing"

	"github.com/goph/stdlib/errors"
	"github.com/goph/stdlib/ext"
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
	closer1 := &testCloser{}
	closer2 := &testCloser{}

	closer := ext.Closers{closer1, closer2}

	var err error

	if got, want := closer.Close(), err; got != want {
		t.Errorf("expected to close successfully, received: %v", got)
	}

	if closer1.called != true || closer2.called != true {
		t.Error("expected closer to be called")
	}
}

func TestClosers_Empty(t *testing.T) {
	closer := ext.Closers{}

	var err error

	if got, want := closer.Close(), err; got != want {
		t.Errorf("expected to close successfully, received: %v", got)
	}
}

func TestClosers_Error(t *testing.T) {
	closer1 := &testCloser{}

	err := fmt.Errorf("error")
	closer2 := &testCloser{
		err: err,
	}

	closer := ext.Closers{closer1, closer2}

	merr := closer.Close()
	if merr == nil {
		t.Fatalf("expected an error, received: %v", merr)
	}

	if merr, ok := merr.(errors.ErrorCollection); !ok {
		e := merr.Errors()
		if err != e[0] {
			t.Fatalf("expected: %v, received: %v", err, e[0])
		}
	}

	if closer1.called != true || closer2.called != true {
		t.Error("expected closer to be called")
	}
}

func TestClose(t *testing.T) {
	err := fmt.Errorf("error")
	closer := ext.CloserFunc(func() { panic(err) })

	if got := ext.Close(closer); got != err {
		t.Errorf("expected: %v, received: %v", err, got)
	}
}
