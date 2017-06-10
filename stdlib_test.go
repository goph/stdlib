package stdlib

import (
	"errors"
	"testing"
)

func TestMust_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	Must(errors.New("should panic"))
}

func TestMust_DoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code should not panic")
		}
	}()

	Must(nil)
}
