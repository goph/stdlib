package errors

// MultiError aggregates multiple errors into a single value.
//
// While ErrorCollection is only an interface for listing errors,
// MultiError actually implements the error interface so it can be returned as an error.
type MultiError struct {
	errors []error
}

// Error implements the error interface.
func (e *MultiError) Error() string {
	return "Multiple errors happened"
}

// Errors returns the list of wrapped errors.
//
// Since MultiError is mutable, calling this method concurrently is not safe.
func (e *MultiError) Errors() []error {
	return e.errors
}

// ErrorOrNil returns an error if this MultiError aggregates a list of errors,
// or returns nil if the list of errors is empty.
//
// It is useful to avoid checking if there are any errors added to the list.
func (e *MultiError) ErrorOrNil() error {
	// MultiError typed nil is possible, return nil.
	if e == nil {
		return nil
	}

	// No errors added, return nil.
	if len(e.errors) == 0 {
		return nil
	}

	return e
}
