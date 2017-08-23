// Package stdlib contains all kinds of utilities and extensions of the standard library.
package stdlib

// Must checks if a parameter is an error and panics if so.
// Useful when you want to force a call to succeed.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
