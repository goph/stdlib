package os_test

import (
	"syscall"
	"testing"

	"github.com/goph/stdlib/internal/testing/assert"
	"github.com/goph/stdlib/os"
)

func TestMustEnv(t *testing.T) {
	syscall.Clearenv()

	err := syscall.Setenv("key", "value")
	if err != nil {
		t.Fatal(err)
	}

	var value string

	assert.NotPanics(t, func() { value = os.MustEnv("key") })
	assert.Equal(t, "value", value)
}

func TestMustEnv_Panics(t *testing.T) {
	syscall.Clearenv()

	assert.Panics(t, func() { os.MustEnv("test_key") })
}
