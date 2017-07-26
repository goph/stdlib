package time_test

import (
	"testing"

	"fmt"
	stdtime "time"

	"github.com/goph/stdlib/time"
	"github.com/stretchr/testify/assert"
)

func TestSystemClock(t *testing.T) {
	ti := stdtime.Now()

	stdtime.Sleep(stdtime.Nanosecond)

	ti = ti.Add(stdtime.Second)
	assert.True(t, time.SystemClock.Now().Before(ti), "expected clock's current time to be before %v", ti)
}

func TestStoppedClock(t *testing.T) {
	ti := stdtime.Date(2017, stdtime.May, 10, 22, 52, 0, 0, stdtime.UTC)

	clock := time.StoppedAt(ti)

	assert.Equal(t, ti, clock.Now())
}

func ExampleStoppedAt() {
	t := stdtime.Date(2017, stdtime.May, 10, 22, 52, 0, 0, stdtime.UTC)
	clock := time.StoppedAt(t)

	fmt.Println(clock.Now())
	// Output: 2017-05-10 22:52:00 +0000 UTC
}
