package errors_test

import (
	"testing"

	"fmt"

	"github.com/goph/stdlib/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMultiErrorBuilder_ErrOrNil(t *testing.T) {
	builder := errors.NewMultiErrorBuilder()

	err := fmt.Errorf("error")

	builder.Add(err)

	merr := builder.ErrOrNil()

	require.Implements(t, (*errors.ErrorCollection)(nil), merr)
	assert.Equal(t, err, merr.(errors.ErrorCollection).Errors()[0])
}

func TestMultiErrorBuilder_ErrOrNil_NilWhenEmpty(t *testing.T) {
	builder := errors.NewMultiErrorBuilder()

	assert.NoError(t, builder.ErrOrNil())
}

func TestMultiErrorBuilder_ErrOrNil_Single(t *testing.T) {
	builder := &errors.MultiErrorBuilder{
		SingleWrapMode: errors.ReturnSingle,
	}

	err := fmt.Errorf("error")

	builder.Add(err)

	assert.Equal(t, err, builder.ErrOrNil())
}

func TestMultiErrorBuilder_Message(t *testing.T) {
	want := "Multiple errors happened during action"

	builder := &errors.MultiErrorBuilder{
		Message: want,
	}

	err := fmt.Errorf("error")

	builder.Add(err)

	assert.Equal(t, want, builder.ErrOrNil().Error())
}
