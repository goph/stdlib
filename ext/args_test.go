package ext_test

import (
	"testing"

	"github.com/goph/stdlib/ext"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArguments_GetE(t *testing.T) {
	args := ext.Arguments{"arg"}

	arg, err := args.GetE(0)

	require.NoError(t, err)
	assert.Equal(t, "arg", arg)
}

func TestArguments_GetE_NotFound(t *testing.T) {
	args := ext.Arguments{"arg"}

	arg, err := args.GetE(1)

	require.Error(t, err)
	assert.EqualError(t, err, "no such index (1) in the argument list: there are only 1 item(s)")
	assert.Nil(t, arg)
}

func TestArguments_Get(t *testing.T) {
	args := ext.Arguments{"arg"}

	arg := args.Get(0)

	assert.Equal(t, "arg", arg)
}

func TestArguments_Get_NotFound(t *testing.T) {
	args := ext.Arguments{"arg"}

	assert.Panics(
		t,
		//errors.New("no such index (1) in the argument list: there are only 1 item(s)"),
		func() {
			args.Get(1)
		},
	)
}
