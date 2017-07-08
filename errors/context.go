package errors

// ContextualError represents an error which holds a context.
type ContextualError interface {
	Context() map[string]interface{}
}
