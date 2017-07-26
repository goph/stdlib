package assert

import (
	"github.com/goph/stdlib/internal/testing"
	"github.com/goph/stdlib/internal/testing/assertion"
)

// Panics asserts that the code inside the specified PanicTestFunc panics.
//
//   assert.Panics(t, func(){ GoCrazy() })
//
// Returns whether the function did panic (true) or not (false).
func Panics(t testing.T, f assertion.PanicTestFunc) bool {
	if p := assertion.Panics(f); !p {
		t.Errorf("func %#v should panic", f)

		return false
	}

	return true
}

// NotPanics asserts that the code inside the specified PanicTestFunc does not panic.
//
//   assert.NotPanics(t, func(){ GoCrazy() })
//
// Returns whether the function did not panic (false) or not (true).
func NotPanics(t testing.T, f assertion.PanicTestFunc) bool {
	if p := assertion.NotPanics(f); !p {
		t.Errorf("func %#v should panic", f)

		return false
	}

	return true
}
