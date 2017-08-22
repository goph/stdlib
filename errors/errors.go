// Package errors extends the errors package in the stdlib.
//
// Despite the implicit nature of interface satisfication in Go
// this package exports a number of interfaces to avoid defining them over and over again.
// Although it means coupling between the consumer code and this package,
// the purpose of this library (being a stdlib extension) justifies that.
package errors

import "errors"

// New returns an error that formats as the given text.
//
// This is an alias to the stdlib errors.New function.
func New(text string) error {
	return errors.New(text)
}
