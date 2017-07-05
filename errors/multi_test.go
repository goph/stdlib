package errors_test

import (
	"testing"

	"fmt"

	"github.com/goph/stdlib/errors"
)

func TestMultiErrorBuilder_ErrOrNil(t *testing.T) {
	builder := errors.NewMultiErrorBuilder()

	err := fmt.Errorf("error")

	builder.Add(err)

	merr := builder.ErrOrNil()

	if _, ok := (merr).(errors.ErrorCollection); !ok {
		t.Fatalf("expected ErrorCollection, received %t", merr)
	}

	if got := merr.(errors.ErrorCollection).Errors(); got[0] != err {
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
