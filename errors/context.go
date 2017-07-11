package errors

import "errors"

// ContextualError represents an error which holds a context.
type ContextualError interface {
	Context() []interface{}
}

// The implementation bellow is heavily influenced by go-kit's log context.

// ErrMissingValue is appended to keyvals slices with odd length to substitute
// the missing value.
var ErrMissingValue = errors.New("(MISSING)")

// With returns a new error with keyvals context appended to it.
// If the wrapped error is already a contextual error created by With
// keyvals is appended to the existing context, but a new error is returned.
func With(err error, keyvals ...interface{}) error {
	if len(keyvals) == 0 {
		return err
	}

	var kvs []interface{}

	if c, ok := err.(*contextualError); ok {
		err = c.err
		kvs = c.keyvals
	} else if c, ok := err.(ContextualError); ok {
		kvs = c.Context()
	}

	kvs = append(kvs, keyvals...)

	if len(kvs)%2 != 0 {
		kvs = append(kvs, ErrMissingValue)
	}
	return &contextualError{
		err: err,
		// Limiting the capacity of the stored keyvals ensures that a new
		// backing array is created if the slice must grow in With.
		// Using the extra capacity without copying risks a data race.
		keyvals: kvs[:len(kvs):len(kvs)],
	}
}

// contextualError is the ContextualError implementation returned by With.
//
// It wraps an error and a holds keyvals as the context.
type contextualError struct {
	err     error
	keyvals []interface{}
}

// Error calls the underlying error and returns it's message.
func (e *contextualError) Error() string {
	return e.err.Error()
}

// Context returns the appended keyvals.
func (e *contextualError) Context() []interface{} {
	return e.keyvals
}
