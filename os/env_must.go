package os

import (
	"fmt"
	"syscall"

	"github.com/pkg/errors"
)

// MustEnv retrieves the value of the environment variable named by the key.
//
// If the variable is not present in the environment, it panics.
func MustEnv(key string) string {
	if v, ok := syscall.Getenv(key); ok {
		return v
	}

	panic(errors.New(fmt.Sprintf("variable '%s' cannot be found in the environment", key)))
}
