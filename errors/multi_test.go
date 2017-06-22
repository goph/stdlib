package errors_test

import (
	"testing"

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

func TestMultiError_ErrorOrNil_NilWhenNil(t *testing.T) {
	var err *errors.MultiError
	err = nil

	var want error

	if got := err.ErrorOrNil(); got != want {
		t.Errorf(`expected nil, received: %v`, got)
	}
}

func TestMultiError_ErrorOrNil_NilWhenEmpty(t *testing.T) {
	err := &errors.MultiError{}

	var want error

	if got := err.ErrorOrNil(); got != want {
		t.Errorf(`expected nil, received: %v`, got)
	}
}
