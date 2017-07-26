package assert

import (
	"github.com/goph/stdlib/internal/testing"
	"github.com/goph/stdlib/internal/testing/assertion"
)

// Equal asserts that two values are equal.
//
//		assert.Equal(t, 123, 123)
//
// Returns whether the two values are equal (true) or not (false).
func Equal(t testing.T, expected interface{}, actual interface{}) bool {
	// TODO: check function type

	if ok := assertion.Equal(actual, expected); !ok {
		t.Errorf("\nwant: %v\nhave: %v", expected, actual)

		return false
	}

	return true
}
