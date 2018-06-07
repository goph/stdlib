package time

import "time"

// Abs returns the absolute value of a duration.
func Abs(d time.Duration) time.Duration {
	if d < 0 {
		return -d
	}

	return d
}
