package os

import (
	"os"
	"strings"
)

// Env returns the complete environment as a map of strings.
func Env() map[string]string {
	env := map[string]string{}
	environ := os.Environ()

	for _, value := range environ {
		v := strings.SplitN(value, "=", 2)

		env[v[0]] = v[1]
	}
	return env
}
