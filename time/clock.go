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

// StoppedAt shows the moment it has been stopped at.
type StoppedAt time.Time

// Now tells the time when it has been stopped.
func (c StoppedAt) Now() time.Time {
	return time.Time(c)
}
