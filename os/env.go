package os

import (
	"syscall"
)

// DefaultEnv retrieves the value of the environment variable named by the key.
//
// If the variable is not present in the environment a default value is returned.
func DefaultEnv(key string, def string) string {
	if v, ok := syscall.Getenv(key); ok {
		return v
	}

	return def
}
