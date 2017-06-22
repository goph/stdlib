package ext

import "github.com/goph/stdlib/errors"

// Closer is an alias interface to io.Closer.
type Closer interface {
	Close() error
}

// CloserFunc makes any function a Closer.
type CloserFunc func()

// Close calls the underlying function and converts any panic to an error.
func (f CloserFunc) Close() (err error) {
	defer func() {
		err = errors.Recover(recover())
	}()

	f()

	return err
}

// Closers is a collection of Closer instances.
type Closers []Closer

// Close calls the underlying Closer instances and returns all their errors as a single value.
func (c Closers) Close() error {
	errBuilder := errors.NewMultiErrorBuilder()

	for _, closer := range c {
		err := closer.Close()

		errBuilder.Add(err)
	}

	return errBuilder.ErrOrNil()
}
