package errors

// ErrorCollection holds a list of errors.
type ErrorCollection interface {
	Errors() []error
}

// ContextualError represents an error which holds a context.
type ContextualError interface {
	Context() []interface{}
}

// Causer is the interface defined in github.com/pkg/errors for specifying a parent error.
type Causer interface {
	Cause() error
}
