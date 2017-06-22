package errors

// ErrorCollection holds a list of errors.
type ErrorCollection interface {
	Errors() []error
}
