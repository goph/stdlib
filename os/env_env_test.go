package os

import (
	"testing"

	"syscall"
)

func TestEnv(t *testing.T) {
	syscall.Clearenv()

	err := syscall.Setenv("key", "value")
	if err != nil {
		t.Fatal(err)
	}

	env := Env()

	value, ok := env["key"]

	if !ok {
		t.Fatal("variable 'key' cannot be found in the environment")
	}

	if value != "value" {
		t.Error("Env is expected to return a map containing value: value, under the key: key")
	}
}
