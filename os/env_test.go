package os_test

import (
	"syscall"
	"testing"

	"github.com/goph/stdlib/os"
	"github.com/stretchr/testify/assert"
)

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
