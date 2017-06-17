package time

import "time"

// Clock tells the time.
type Clock interface {
	// Now tells the actual time.
	Now() time.Time
}

// ClockFunc wraps an ordinary function and makes it a Clock.
type ClockFunc func() time.Time

// Now calls the underlying function and returns the result.
func (f ClockFunc) Now() time.Time {
	return f()
}

// SystemClock returns the current system time.
var SystemClock = ClockFunc(time.Now)

// StoppedClock shows the moment it has been stopped.
type StoppedClock struct {
	t time.Time
}

// NewStoppedClock returns a new StoppedClock.
func NewStoppedClock(t time.Time) Clock {
	return &StoppedClock{t}
}

// Now tells the time when it has been stopped.
func (c *StoppedClock) Now() time.Time {
	return c.t
}
