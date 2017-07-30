package os_test

import (
	"syscall"
	"testing"

	"github.com/goph/stdlib/os"
	"github.com/stretchr/testify/assert"
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

func TestDefaultEnv(t *testing.T) {
	syscall.Clearenv()

	err := syscall.Setenv("key", "value")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "value", os.DefaultEnv("key", "default"))
}

func TestDefaultEnv_NotFound(t *testing.T) {
	syscall.Clearenv()

	assert.Equal(t, "default", os.DefaultEnv("key", "default"))
}
