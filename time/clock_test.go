package time

import (
	"testing"

	"fmt"
	"time"
)

func TestStoppedClock(t *testing.T) {
	now := time.Date(2017, time.May, 10, 22, 52, 0, 0, time.UTC)

	clock := StoppedAt(now)

	realNow := clock.Now()

	if now != realNow {
		t.Errorf("StoppedAt is expected to return the exact time it is stopped at, got: %s", realNow)
	}
}

func ExampleStoppedAt() {
	now := time.Date(2017, time.May, 10, 22, 52, 0, 0, time.UTC)
	clock := StoppedAt(now)

	fmt.Println(clock.Now())
	// Output: 2017-05-10 22:52:00 +0000 UTC
}
