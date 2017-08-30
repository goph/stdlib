package errors_test

import (
	"testing"

	"github.com/goph/stdlib/errors"
	"github.com/stretchr/testify/assert"
)

func TestHandleRecovery(t *testing.T) {
	handler := new(errors.TestHandler)
	err := errors.New("error")

	defer func() {
		assert.EqualError(t, handler.Last(), "error")
	}()
	defer errors.HandleRecover(handler)

	panic(err)
}

func TestHandleIfErr(t *testing.T) {
	handler := new(errors.TestHandler)
	err := errors.New("error")

	errors.HandleIfErr(handler, err)

	assert.Equal(t, err, handler.Last())
}

func TestHandleIfErr_Nil(t *testing.T) {
	handler := new(errors.TestHandler)

	errors.HandleIfErr(handler, nil)

	assert.NoError(t, handler.Last())
}
