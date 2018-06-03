package os

import (
	"testing"

	"syscall"
)

func TestMustEnv(t *testing.T) {
	syscall.Clearenv()

	err := syscall.Setenv("key", "value")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		err := recover()

		if err != nil {
			t.Fatal("MustEnv is not expected to panic for key: key", err)
		}
	}()

	value := MustEnv("key")

	if value != "value" {
		t.Errorf("MustEnv is expected to return: value, got: %s", value)
	}
}

func TestMustEnv_Panics(t *testing.T) {
	syscall.Clearenv()

	defer func() {
		err := recover()

		if err == nil {
			t.Fatal("MustEnv is expected to panic for key: test_key")
		}
	}()

	MustEnv("test_key")
}
