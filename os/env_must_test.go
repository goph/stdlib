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
			t.Fatal("MustEnv is not supposed to panic for key 'key'", err)
		}
	}()

	value := MustEnv("key")

	if value != "value" {
		t.Errorf("MustEnv is supposed to return 'value', got: %s", value)
	}
}

func TestMustEnv_Panics(t *testing.T) {
	syscall.Clearenv()

	defer func() {
		err := recover()

		if err == nil {
			t.Fatal("MustEnv is supposed to panic for key 'test_key', but it did not")
		}
	}()

	MustEnv("test_key")
}
