package stdlib_test

import (
	"errors"
	"testing"

	"github.com/goph/stdlib"
	"github.com/goph/stdlib/internal/testing/assert"
)

func TestMust_Panics(t *testing.T) {
	assert.Panics(t, func() { stdlib.Must(errors.New("should panic")) })
}

func TestMust_DoesNotPanic(t *testing.T) {
	assert.NotPanics(t, func() { stdlib.Must(nil) })
}
