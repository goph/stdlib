package assertion

// Inspired by https://github.com/stretchr/testify/blob/05e8a0eda380579888eb53c394909df027f06991/assert/assertions.go#L774-L812

// PanicTestFunc represents a simple func that takes no arguments, and returns nothing.
type PanicTestFunc func()

// didPanic returns true if the function passed to it panics. Otherwise, it returns false.
// It also returns the recovered value.
func didPanic(f PanicTestFunc) (p bool, v interface{}) {
	func() {
		defer func() {
			if v = recover(); v != nil {
				p = true
			}
		}()

		// call the target function
		f()
	}()

	return
}

// Panics checks that the code inside the specified PanicTestFunc panics.
//
// 		assert.Panics(t, func(){ GoCrazy() })
//
// Returns whether the function did panic (true) or not (false).
func Panics(f PanicTestFunc) bool {
	p, _ := didPanic(f)

	return p
}

// NotPanics checks that the code inside the specified PanicTestFunc does not panic.
//
// 		assert.NotPanics(t, func(){ DontGoCrazy() })
//
// Returns whether the function did panic (false) or not (true).
func NotPanics(f PanicTestFunc) bool {
	return !Panics(f)
}
