//+build experimental

package types

// LookupString retrieves an argument of type string from the list stored under the specified the index.
//
// If the index is present in the list and it is of type string the value is returned and the boolean is true.
//
// Otherwise the type's zero value and false are returned.
func (a Arguments) LookupString(index int) (string, bool) {
	arg, ok := a.Lookup(index)
	if !ok {
		return "", false
	}
	if v, ok := arg.(string); ok {
		return v, true
	}
	return "", false
}

// String retrieves an argument of type string from the list stored under the specified the index.
//
// If the index is present in the list and it is of type string the value is returned.
//
// Otherwise the type's zero value is returned. To distinguish between an empty value and an unset value, use LookupString.
func (a Arguments) String(index int) string {
	if arg, ok := a.LookupString(index); ok {
		return arg
	}
	return ""
}

// DefaultString retrieves an argument of type string from the list stored under the specified the index.
//
// If the index is present in the list and it is of type string the value is returned.
//
// Otherwise the specified default value is returned.
func (a Arguments) DefaultString(index int, def string) string {
	if arg, ok := a.LookupString(index); ok {
		return arg
	}
	return def
}

// LookupBool retrieves an argument of type bool from the list stored under the specified the index.
//
// If the index is present in the list and it is of type bool the value is returned and the boolean is true.
//
// Otherwise the type's zero value and false are returned.
func (a Arguments) LookupBool(index int) (bool, bool) {
	arg, ok := a.Lookup(index)
	if !ok {
		return false, false
	}
	if v, ok := arg.(bool); ok {
		return v, true
	}
	return false, false
}

// Bool retrieves an argument of type bool from the list stored under the specified the index.
//
// If the index is present in the list and it is of type bool the value is returned.
//
// Otherwise the type's zero value is returned. To distinguish between an empty value and an unset value, use LookupBool.
func (a Arguments) Bool(index int) bool {
	if arg, ok := a.LookupBool(index); ok {
		return arg
	}
	return false
}

// DefaultBool retrieves an argument of type bool from the list stored under the specified the index.
//
// If the index is present in the list and it is of type bool the value is returned.
//
// Otherwise the specified default value is returned.
func (a Arguments) DefaultBool(index int, def bool) bool {
	if arg, ok := a.LookupBool(index); ok {
		return arg
	}
	return def
}

// LookupInt retrieves an argument of type int from the list stored under the specified the index.
//
// If the index is present in the list and it is of type int the value is returned and the boolean is true.
//
// Otherwise the type's zero value and false are returned.
func (a Arguments) LookupInt(index int) (int, bool) {
	arg, ok := a.Lookup(index)
	if !ok {
		return 0, false
	}
	if v, ok := arg.(int); ok {
		return v, true
	}
	return 0, false
}

// Int retrieves an argument of type int from the list stored under the specified the index.
//
// If the index is present in the list and it is of type int the value is returned.
//
// Otherwise the type's zero value is returned. To distinguish between an empty value and an unset value, use LookupInt.
func (a Arguments) Int(index int) int {
	if arg, ok := a.LookupInt(index); ok {
		return arg
	}
	return 0
}

// DefaultInt retrieves an argument of type int from the list stored under the specified the index.
//
// If the index is present in the list and it is of type int the value is returned.
//
// Otherwise the specified default value is returned.
func (a Arguments) DefaultInt(index int, def int) int {
	if arg, ok := a.LookupInt(index); ok {
		return arg
	}
	return def
}

// LookupInt32 retrieves an argument of type int32 from the list stored under the specified the index.
//
// If the index is present in the list and it is of type int32 the value is returned and the boolean is true.
//
// Otherwise the type's zero value and false are returned.
func (a Arguments) LookupInt32(index int) (int32, bool) {
	arg, ok := a.Lookup(index)
	if !ok {
		return 0, false
	}
	if v, ok := arg.(int32); ok {
		return v, true
	}
	return 0, false
}

// Int32 retrieves an argument of type int32 from the list stored under the specified the index.
//
// If the index is present in the list and it is of type int32 the value is returned.
//
// Otherwise the type's zero value is returned. To distinguish between an empty value and an unset value, use LookupInt32.
func (a Arguments) Int32(index int) int32 {
	if arg, ok := a.LookupInt32(index); ok {
		return arg
	}
	return 0
}

// DefaultInt32 retrieves an argument of type int32 from the list stored under the specified the index.
//
// If the index is present in the list and it is of type int32 the value is returned.
//
// Otherwise the specified default value is returned.
func (a Arguments) DefaultInt32(index int, def int32) int32 {
	if arg, ok := a.LookupInt32(index); ok {
		return arg
	}
	return def
}

// LookupInt64 retrieves an argument of type int64 from the list stored under the specified the index.
//
// If the index is present in the list and it is of type int64 the value is returned and the boolean is true.
//
// Otherwise the type's zero value and false are returned.
func (a Arguments) LookupInt64(index int) (int64, bool) {
	arg, ok := a.Lookup(index)
	if !ok {
		return 0, false
	}
	if v, ok := arg.(int64); ok {
		return v, true
	}
	return 0, false
}

// Int64 retrieves an argument of type int64 from the list stored under the specified the index.
//
// If the index is present in the list and it is of type int64 the value is returned.
//
// Otherwise the type's zero value is returned. To distinguish between an empty value and an unset value, use LookupInt64.
func (a Arguments) Int64(index int) int64 {
	if arg, ok := a.LookupInt64(index); ok {
		return arg
	}
	return 0
}

// DefaultInt64 retrieves an argument of type int64 from the list stored under the specified the index.
//
// If the index is present in the list and it is of type int64 the value is returned.
//
// Otherwise the specified default value is returned.
func (a Arguments) DefaultInt64(index int, def int64) int64 {
	if arg, ok := a.LookupInt64(index); ok {
		return arg
	}
	return def
}

// LookupFloat32 retrieves an argument of type float32 from the list stored under the specified the index.
//
// If the index is present in the list and it is of type float32 the value is returned and the boolean is true.
//
// Otherwise the type's zero value and false are returned.
func (a Arguments) LookupFloat32(index int) (float32, bool) {
	arg, ok := a.Lookup(index)
	if !ok {
		return 0.0, false
	}
	if v, ok := arg.(float32); ok {
		return v, true
	}
	return 0.0, false
}

// Float32 retrieves an argument of type float32 from the list stored under the specified the index.
//
// If the index is present in the list and it is of type float32 the value is returned.
//
// Otherwise the type's zero value is returned. To distinguish between an empty value and an unset value, use LookupFloat32.
func (a Arguments) Float32(index int) float32 {
	if arg, ok := a.LookupFloat32(index); ok {
		return arg
	}
	return 0.0
}

// DefaultFloat32 retrieves an argument of type float32 from the list stored under the specified the index.
//
// If the index is present in the list and it is of type float32 the value is returned.
//
// Otherwise the specified default value is returned.
func (a Arguments) DefaultFloat32(index int, def float32) float32 {
	if arg, ok := a.LookupFloat32(index); ok {
		return arg
	}
	return def
}

// LookupFloat64 retrieves an argument of type float64 from the list stored under the specified the index.
//
// If the index is present in the list and it is of type float64 the value is returned and the boolean is true.
//
// Otherwise the type's zero value and false are returned.
func (a Arguments) LookupFloat64(index int) (float64, bool) {
	arg, ok := a.Lookup(index)
	if !ok {
		return 0.0, false
	}
	if v, ok := arg.(float64); ok {
		return v, true
	}
	return 0.0, false
}

// Float64 retrieves an argument of type float64 from the list stored under the specified the index.
//
// If the index is present in the list and it is of type float64 the value is returned.
//
// Otherwise the type's zero value is returned. To distinguish between an empty value and an unset value, use LookupFloat64.
func (a Arguments) Float64(index int) float64 {
	if arg, ok := a.LookupFloat64(index); ok {
		return arg
	}
	return 0.0
}

// DefaultFloat64 retrieves an argument of type float64 from the list stored under the specified the index.
//
// If the index is present in the list and it is of type float64 the value is returned.
//
// Otherwise the specified default value is returned.
func (a Arguments) DefaultFloat64(index int, def float64) float64 {
	if arg, ok := a.LookupFloat64(index); ok {
		return arg
	}
	return def
}
