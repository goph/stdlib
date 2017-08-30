package errors_test

import (
	"testing"

	"github.com/goph/stdlib/errors"
	"github.com/stretchr/testify/assert"
)

func TestHandlerFunc(t *testing.T) {
	var actual error
	log := func(err error) {
		actual = err
	}

	fn := errors.HandlerFunc(log)

	expected := errors.New("error")

	fn.Handle(expected)

	assert.Equal(t, expected, actual)
}

func TestHandlerLogFunc(t *testing.T) {
	var actual error
	log := func(args ...interface{}) {
		actual = args[0].(error)
	}

	fn := errors.HandlerLogFunc(log)

	expected := errors.New("error")

	fn.Handle(expected)

	assert.Equal(t, expected, actual)
}
