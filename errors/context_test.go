package errors_test

import (
	"testing"

	"github.com/goph/stdlib/errors"
)

func TestContext(t *testing.T) {
	t.Parallel()

	err := errors.New("")

	kvs := []interface{}{"a", 123}
	err = errors.With(err, kvs...)
	kvs[1] = 0 // With should copy its key values

	cerr, ok := err.(errors.ContextualError)

	if !ok {
		t.Fatal("expected error to implement ContextualError")
	}

	ctx := cerr.Context()

	if want, have := "a", ctx[0]; want != have {
		t.Errorf("\nwant: %s\nhave: %v", want, have)
	}

	if want, have := 123, ctx[1]; want != have {
		t.Errorf("\nwant: %d\nhave: %v", want, have)
	}
}

func TestContext_Multi(t *testing.T) {
	t.Parallel()

	err := errors.New("")

	err = errors.With(errors.With(err, "a", 123), "b", 321)

	cerr, ok := err.(errors.ContextualError)

	if !ok {
		t.Fatal("expected error to implement ContextualError")
	}

	ctx := cerr.Context()

	if want, have := "a", ctx[0]; want != have {
		t.Errorf("\nwant: %s\nhave: %v", want, have)
	}

	if want, have := 123, ctx[1]; want != have {
		t.Errorf("\nwant: %d\nhave: %v", want, have)
	}

	if want, have := "b", ctx[2]; want != have {
		t.Errorf("\nwant: %s\nhave: %v", want, have)
	}

	if want, have := 321, ctx[3]; want != have {
		t.Errorf("\nwant: %d\nhave: %v", want, have)
	}
}

func TestContext_MultiPrefix(t *testing.T) {
	t.Parallel()

	err := errors.New("")

	err = errors.WithPrefix(errors.With(err, "a", 123), "b", 321)

	cerr, ok := err.(errors.ContextualError)

	if !ok {
		t.Fatal("expected error to implement ContextualError")
	}

	ctx := cerr.Context()

	if want, have := "a", ctx[2]; want != have {
		t.Errorf("\nwant: %s\nhave: %v", want, have)
	}

	if want, have := 123, ctx[3]; want != have {
		t.Errorf("\nwant: %d\nhave: %v", want, have)
	}

	if want, have := "b", ctx[0]; want != have {
		t.Errorf("\nwant: %s\nhave: %v", want, have)
	}

	if want, have := 321, ctx[1]; want != have {
		t.Errorf("\nwant: %d\nhave: %v", want, have)
	}
}

func TestContext_MissingValue(t *testing.T) {
	t.Parallel()

	err := errors.New("")

	err = errors.WithPrefix(errors.With(err, "k0"), "k1")

	cerr, ok := err.(errors.ContextualError)

	if !ok {
		t.Fatal("expected error to implement ContextualError")
	}

	ctx := cerr.Context()

	if want, have := 4, len(ctx); want != have {
		t.Errorf("want len(output) == %v, have %v", want, have)
	}

	for i := 1; i < 4; i += 2 {
		if want, have := errors.ErrMissingValue, ctx[i]; want != have {
			t.Errorf("want output[%d] == %#v, have %#v", i, want, have)
		}
	}
}
