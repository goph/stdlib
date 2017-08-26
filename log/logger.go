package log

// Logger is a structured logger.
//
// See github.com/go-kit/kit/log package for details.
type Logger interface {
	// Log creates a log event from keyvals, a variadic sequence of alternating keys and values.
	Log(keyvals ...interface{}) error
}
