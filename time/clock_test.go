package time

import (
	"fmt"
	"testing"

	"time"
)

func TestSystemClock(t *testing.T) {
	ti := time.Now()

	time.Sleep(time.Nanosecond)

	if ti = ti.Add(time.Second); SystemClock.Now().After(ti) {
		t.Errorf("expected clock's current time to be before %v", ti)
	}
}

func TestStoppedClock(t *testing.T) {
	ti := time.Date(2017, time.May, 10, 22, 52, 0, 0, time.UTC)

	clock := StoppedAt(ti)

	if ti != clock.Now() {
		t.Errorf("expected clock's current time to be %v", ti)
	}
}

func ExampleStoppedAt() {
	t := time.Date(2017, time.May, 10, 22, 52, 0, 0, time.UTC)
	clock := StoppedAt(t)

	fmt.Println(clock.Now())
	// Output: 2017-05-10 22:52:00 +0000 UTC
}
