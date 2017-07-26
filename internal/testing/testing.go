// Package testing extends the testing package in the stdlib.
package testing

// T is an interface wrapper around *testing.T.
type T interface {
	Errorf(format string, args ...interface{})
}
