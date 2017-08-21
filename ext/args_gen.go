//+build experimental

package ext

import "fmt"
// StringE returns a(n) string argument from the list or an error if it cannot be found or not string.
func (a Arguments) StringE(index int) (string, error) {
	arg, err := a.GetE(index)
	if err != nil {
		return "", err
	}
	v, ok := arg.(string)
	if !ok {
		return "", fmt.Errorf("cannot return argument (%d) as string because it is of type %T", index, arg)
	}
	return v, nil
}
// String returns a(n) string argument from the list.
//
// It panics if the argument with such index cannot be found or it is not string.
func (a Arguments) String(index int) string {
	arg, err := a.StringE(index)
	if err != nil {
		panic(err)
	}
	return arg
}
// BoolE returns a(n) bool argument from the list or an error if it cannot be found or not bool.
func (a Arguments) BoolE(index int) (bool, error) {
	arg, err := a.GetE(index)
	if err != nil {
		return false, err
	}
	v, ok := arg.(bool)
	if !ok {
		return false, fmt.Errorf("cannot return argument (%d) as bool because it is of type %T", index, arg)
	}
	return v, nil
}
// Bool returns a(n) bool argument from the list.
//
// It panics if the argument with such index cannot be found or it is not bool.
func (a Arguments) Bool(index int) bool {
	arg, err := a.BoolE(index)
	if err != nil {
		panic(err)
	}
	return arg
}
// IntE returns a(n) int argument from the list or an error if it cannot be found or not int.
func (a Arguments) IntE(index int) (int, error) {
	arg, err := a.GetE(index)
	if err != nil {
		return 0, err
	}
	v, ok := arg.(int)
	if !ok {
		return 0, fmt.Errorf("cannot return argument (%d) as int because it is of type %T", index, arg)
	}
	return v, nil
}
// Int returns a(n) int argument from the list.
//
// It panics if the argument with such index cannot be found or it is not int.
func (a Arguments) Int(index int) int {
	arg, err := a.IntE(index)
	if err != nil {
		panic(err)
	}
	return arg
}
// Int32E returns a(n) int32 argument from the list or an error if it cannot be found or not int32.
func (a Arguments) Int32E(index int) (int32, error) {
	arg, err := a.GetE(index)
	if err != nil {
		return 0, err
	}
	v, ok := arg.(int32)
	if !ok {
		return 0, fmt.Errorf("cannot return argument (%d) as int32 because it is of type %T", index, arg)
	}
	return v, nil
}
// Int32 returns a(n) int32 argument from the list.
//
// It panics if the argument with such index cannot be found or it is not int32.
func (a Arguments) Int32(index int) int32 {
	arg, err := a.Int32E(index)
	if err != nil {
		panic(err)
	}
	return arg
}
// Int64E returns a(n) int64 argument from the list or an error if it cannot be found or not int64.
func (a Arguments) Int64E(index int) (int64, error) {
	arg, err := a.GetE(index)
	if err != nil {
		return 0, err
	}
	v, ok := arg.(int64)
	if !ok {
		return 0, fmt.Errorf("cannot return argument (%d) as int64 because it is of type %T", index, arg)
	}
	return v, nil
}
// Int64 returns a(n) int64 argument from the list.
//
// It panics if the argument with such index cannot be found or it is not int64.
func (a Arguments) Int64(index int) int64 {
	arg, err := a.Int64E(index)
	if err != nil {
		panic(err)
	}
	return arg
}
// Float32E returns a(n) float32 argument from the list or an error if it cannot be found or not float32.
func (a Arguments) Float32E(index int) (float32, error) {
	arg, err := a.GetE(index)
	if err != nil {
		return 0.0, err
	}
	v, ok := arg.(float32)
	if !ok {
		return 0.0, fmt.Errorf("cannot return argument (%d) as float32 because it is of type %T", index, arg)
	}
	return v, nil
}
// Float32 returns a(n) float32 argument from the list.
//
// It panics if the argument with such index cannot be found or it is not float32.
func (a Arguments) Float32(index int) float32 {
	arg, err := a.Float32E(index)
	if err != nil {
		panic(err)
	}
	return arg
}
// Float64E returns a(n) float64 argument from the list or an error if it cannot be found or not float64.
func (a Arguments) Float64E(index int) (float64, error) {
	arg, err := a.GetE(index)
	if err != nil {
		return 0.0, err
	}
	v, ok := arg.(float64)
	if !ok {
		return 0.0, fmt.Errorf("cannot return argument (%d) as float64 because it is of type %T", index, arg)
	}
	return v, nil
}
// Float64 returns a(n) float64 argument from the list.
//
// It panics if the argument with such index cannot be found or it is not float64.
func (a Arguments) Float64(index int) float64 {
	arg, err := a.Float64E(index)
	if err != nil {
		panic(err)
	}
	return arg
}
