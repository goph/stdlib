package log

// The following code has been inspired by github.com/go-kit/kit/log.

// nopLogger doesn't do anything.
type nopLogger struct{}

// NewNopLogger returns a logger that doesn't do anything.
func NewNopLogger() Logger { return nopLogger{} }

// Log implements the Logger interface and it doesn't do anything.
func (nopLogger) Log(...interface{}) error { return nil }
