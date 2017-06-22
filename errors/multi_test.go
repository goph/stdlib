package errors_test

import (
	"testing"

	"fmt"

	"github.com/goph/stdlib/errors"
)

func TestMultiErrorIsAnError(t *testing.T) {
	var i interface{} = new(errors.MultiError)

	if _, ok := i.(error); !ok {
		t.Errorf("expected %t to implement error", i)
	}
}

func TestMultiErrorIsAnErrorCollection(t *testing.T) {
	var i interface{} = new(errors.MultiError)

	if _, ok := i.(errors.ErrorCollection); !ok {
		t.Errorf("expected %t to implement errors.ErrorCollection", i)
	}
}

func TestMultiError_Error(t *testing.T) {
	err := &errors.MultiError{}

	if got, want := err.Error(), "Multiple errors happened"; got != want {
		t.Errorf(`expected error "%s", received "%s"`, want, got)
	}
}

func TestMultiErrorBuilder_ErrOrNil(t *testing.T) {
	builder := errors.NewMultiErrorBuilder()

	err := fmt.Errorf("error")

	builder.Add(err)

	merr := builder.ErrOrNil()

	if _, ok := (merr).(*errors.MultiError); !ok {
		t.Fatalf("expected MultiError, received %t", merr)
	}

	if got := merr.(*errors.MultiError).Errors(); got[0] != err {
		t.Errorf(`expected %v, received: %v`, err, got[0])
	}
}

func TestMultiErrorBuilder_ErrOrNil_NilWhenEmpty(t *testing.T) {
	builder := errors.NewMultiErrorBuilder()

	var want error

	if got := builder.ErrOrNil(); got != want {
		t.Errorf(`expected nil, received: %v`, got)
	}
}

func TestMultiErrorBuilder_ErrOrNil_Single(t *testing.T) {
	builder := &errors.MultiErrorBuilder{
		SingleWrapMode: errors.ReturnSingle,
	}

	err := fmt.Errorf("error")

	builder.Add(err)

	if got := builder.ErrOrNil(); got != err {
		t.Errorf(`expected %v, received: %v`, err, got)
	}
}
