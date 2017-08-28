package types

//go:generate go run ./_codegen/args.go

// Arguments holds a list of arguments.
type Arguments []interface{}

// Lookup retrieves an argument from the list stored under the specified the index.
//
// If the index is present in the list the value is returned and the boolean is true.
//
// Otherwise nil and false are returned.
func (a Arguments) Lookup(index int) (interface{}, bool) {
	if l := len(a); index >= l {
		return nil, false
	}

	return a[index], true
}

// Get retrieves an argument from the list stored under the specified the index.
//
// If the index is present in the list the value is returned.
//
// Otherwise nil is returned. To distinguish between an empty value and an unset value, use Lookup.
func (a Arguments) Get(index int) interface{} {
	if v, ok := a.Lookup(index); ok {
		return v
	}

	return nil
}

// Default retrieves an argument from the list stored under the specified the index.
//
// If the index is present in the list the value is returned.
//
// Otherwise the specified default value is returned.
func (a Arguments) Default(index int, def interface{}) interface{} {
	if v, ok := a.Lookup(index); ok {
		return v
	}

	return def
}
