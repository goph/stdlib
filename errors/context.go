package errors

// The implementation bellow is heavily influenced by go-kit's log context.

// ErrMissingValue is appended to keyvals slices with odd length to substitute
// the missing value.
var ErrMissingValue = New("(MISSING)")

// With returns a new error with keyvals context appended to it.
// If the wrapped error is already a contextual error created by With or WithPrefix
// keyvals is appended to the existing context, but a new error is returned.
func With(err error, keyvals ...interface{}) error {
	if len(keyvals) == 0 {
		return err
	}

	var kvs []interface{}

	if c, ok := err.(*contextualError); ok {
		err = c.error
		kvs = c.keyvals
	} else if c, ok := err.(ContextualError); ok {
		kvs = c.Context()
	}

	kvs = append(kvs, keyvals...)

	if len(kvs)%2 != 0 {
		kvs = append(kvs, ErrMissingValue)
	}
	return &contextualError{
		error: err,

		// Limiting the capacity of the stored keyvals ensures that a new
		// backing array is created if the slice must grow in With.
		// Using the extra capacity without copying risks a data race.
		keyvals: kvs[:len(kvs):len(kvs)],
	}
}

// WithPrefix returns a new error with keyvals context appended to it.
// If the wrapped error is already a contextual error created by With or WithPrefix
// keyvals is prepended to the existing context, but a new error is returned.
func WithPrefix(err error, keyvals ...interface{}) error {
	if len(keyvals) == 0 {
		return err
	}

	var prevkvs []interface{}

	if c, ok := err.(*contextualError); ok {
		err = c.error
		prevkvs = c.keyvals
	} else if c, ok := err.(ContextualError); ok {
		prevkvs = c.Context()
	}

	n := len(prevkvs) + len(keyvals)
	if len(keyvals)%2 != 0 {
		n++
	}

	kvs := make([]interface{}, 0, n)
	kvs = append(kvs, keyvals...)

	if len(kvs)%2 != 0 {
		kvs = append(kvs, ErrMissingValue)
	}

	kvs = append(kvs, prevkvs...)

	return &contextualError{
		error:   err,
		keyvals: kvs,
	}
}

// contextualError is the ContextualError implementation returned by With.
//
// It wraps an error and a holds keyvals as the context.
type contextualError struct {
	error
	keyvals []interface{}
}

// Context returns the appended keyvals.
func (e *contextualError) Context() []interface{} {
	return e.keyvals
}

// Cause returns the underlying error.
//
// This method fulfills the causer interface described in github.com/pkg/errors.
func (e *contextualError) Cause() error {
	return e.error
}
