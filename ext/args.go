//+build experimental

package ext

import "fmt"

//go:generate go run ./_codegen/args.go

// Arguments represents a list of arguments with a set of type getters.
type Arguments []interface{}

// GetE returns an argument from the list or an error if it cannot be found.
func (a Arguments) GetE(index int) (interface{}, error) {
	if l := len(a); index >= l {
		return nil, fmt.Errorf("no such index (%d) in the argument list: there are only %d item(s)", index, l)
	}

	return a[index], nil
}

// Get returns an argument from the list.
//
// It panics if the argument with such index cannot be found.
func (a Arguments) Get(index int) interface{} {
	v, err := a.GetE(index)
	if err != nil {
		panic(err)
	}

	return v
}
