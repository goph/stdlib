package time

import (
	"testing"

	"time"
)

func TestStoppedClock_Now(t *testing.T) {
	ti := time.Date(2017, time.May, 10, 22, 52, 0, 0, time.UTC)

	clock := NewStoppedClock(ti)

	if ti != clock.Now() {
		t.Errorf("expected clock's current time to be %v", ti)
	}
}

func TestClock_Now(t *testing.T) {
	ti := time.Now()

	time.Sleep(time.Nanosecond)

	clock := NewClock()

	if ti = ti.Add(time.Second); clock.Now().After(ti) {
		t.Errorf("expected clock's current time to be before %v", ti)
	}
}
