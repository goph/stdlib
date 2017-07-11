// Package errors extends the errors package in the stdlib.
package errors

import "errors"

// New returns an error that formats as the given text.
//
// This is an alias to the stdlib errors.New function.
func New(text string) error {
	return errors.New(text)
}
