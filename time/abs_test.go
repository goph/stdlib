package time

import (
	"testing"
	"time"
)

func TestAbs(t *testing.T) {
	tests := map[string]struct {
		duration time.Duration
		expected time.Duration
	}{
		"negative": {
			-time.Second,
			time.Second,
		},
		"positive": {
			time.Second,
			time.Second,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := Abs(test.duration)

			if actual != test.expected {
				t.Errorf("expected Abs to return `%s` for `%s`, got `%s`", test.expected, test.duration, actual)
			}
		})
	}
}
