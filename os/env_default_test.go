package os

import (
	"testing"

	"syscall"
)

func TestDefaultEnv(t *testing.T) {
	syscall.Clearenv()

	err := syscall.Setenv("key", "value")
	if err != nil {
		t.Fatal(err)
	}

	value := DefaultEnv("key", "default")

	if value != "value" {
		t.Errorf("DefaultEnv is expcted to return: value, got: %s", value)
	}
}

func TestDefaultEnv_NotFound(t *testing.T) {
	syscall.Clearenv()

	value := DefaultEnv("key", "default")

	if value != "default" {
		t.Errorf("DefaultEnv is expected to return: default, got: %s", value)
	}
}
