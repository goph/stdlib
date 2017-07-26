package errors_test

import (
	"testing"

	"fmt"

	"github.com/goph/stdlib/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRecoverFunc(p interface{}) func() error {
	return func() (err error) {
		defer func() {
			err = errors.Recover(recover())
		}()

		panic(p)
	}
}

func TestRecover_ErrorPanic(t *testing.T) {
	err := fmt.Errorf("internal error")

	f := createRecoverFunc(err)

	require.NotPanics(t, func() { f() })
	assert.Equal(t, err, f())
}

func TestRecover_StringPanic(t *testing.T) {
	f := createRecoverFunc("internal error")

	require.NotPanics(t, func() { f() })
	assert.Equal(t, "internal error", f().Error())
}

func TestRecover_AnyPanic(t *testing.T) {
	f := createRecoverFunc(123)

	require.NotPanics(t, func() { f() })
	assert.Equal(t, "Unknown panic, received: 123", f().Error())
}
