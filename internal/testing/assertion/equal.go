package assertion

import (
	"bytes"
	"reflect"
)

// Equal checks that two values are equal.
//
//		assertion.Equal(123, 123)
//
// Returns whether the two values are equal (true) or not (false).
//
// Borrowed from https://github.com/stretchr/testify/blob/05e8a0eda380579888eb53c394909df027f06991/assert/assertions.go#L36-L55
func Equal(expected interface{}, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	if exp, ok := expected.([]byte); ok {
		act, ok := actual.([]byte)
		if !ok {
			return false
		} else if exp == nil || act == nil {
			return exp == nil && act == nil
		}

		return bytes.Equal(exp, act)
	}

	return reflect.DeepEqual(expected, actual)
}
