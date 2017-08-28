package types_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/goph/stdlib/ext/types"
	"github.com/goph/stdlib/strings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var argumentTests = map[string]struct {
	example interface{}
	def     interface{}
}{
	"string":  {"arg", ""},
	"bool":    {true, false},
	"int":     {1, 0},
	"int32":   {int32(1), int32(0)},
	"int64":   {int64(1), int64(0)},
	"float32": {float32(1.0), float32(0.0)},
	"float64": {float64(1.0), float64(0.0)},
}

func TestArguments_Lookup(t *testing.T) {
	args := types.Arguments{"arg"}

	arg, ok := args.Lookup(0)

	require.True(t, ok)
	assert.Equal(t, "arg", arg)
}

func TestArguments_Lookup_NotFound(t *testing.T) {
	args := types.Arguments{"arg"}

	arg, ok := args.Lookup(1)

	require.False(t, ok)
	assert.Nil(t, arg)
}

func TestArguments_Get(t *testing.T) {
	args := types.Arguments{"arg"}

	arg := args.Get(0)

	assert.Equal(t, "arg", arg)
}

func TestArguments_Get_NotFound(t *testing.T) {
	args := types.Arguments{"arg"}

	arg := args.Get(1)

	assert.Nil(t, arg)
}

func TestArguments_Default(t *testing.T) {
	args := types.Arguments{"arg"}

	arg := args.Default(0, "another_arg")

	assert.Equal(t, "arg", arg)
}

func TestArguments_Default_NotFound(t *testing.T) {
	args := types.Arguments{"arg"}

	arg := args.Default(1, "another_arg")

	assert.Equal(t, "another_arg", arg)
}

func TestArguments_TypeLookup(t *testing.T) {
	for typ, test := range argumentTests {
		t.Logf("Testing %s lookup", typ)

		args := types.Arguments{test.example}

		arg := reflect.ValueOf(args).MethodByName(fmt.Sprintf("Lookup%s", strings.ToCamel(typ))).Call([]reflect.Value{reflect.ValueOf(0)})

		assert.Equal(t, test.example, arg[0].Interface())
		assert.True(t, arg[1].Interface().(bool))
	}
}

func TestArguments_TypeLookup_NotFound(t *testing.T) {
	for typ, test := range argumentTests {
		t.Logf("Testing %s lookup when not found", typ)

		args := types.Arguments{test.example}

		arg := reflect.ValueOf(args).MethodByName(fmt.Sprintf("Lookup%s", strings.ToCamel(typ))).Call([]reflect.Value{reflect.ValueOf(1)})

		assert.Equal(t, test.def, arg[0].Interface())
		assert.False(t, arg[1].Interface().(bool))
	}
}

func TestArguments_TypeLookup_InvalidType(t *testing.T) {
	for typ, test := range argumentTests {
		t.Logf("Testing %s lookup when type is invalid", typ)

		args := types.Arguments{nil}

		arg := reflect.ValueOf(args).MethodByName(fmt.Sprintf("Lookup%s", strings.ToCamel(typ))).Call([]reflect.Value{reflect.ValueOf(0)})

		assert.Equal(t, test.def, arg[0].Interface())
		assert.False(t, arg[1].Interface().(bool))
	}
}

func TestArguments_Type(t *testing.T) {
	for typ, test := range argumentTests {
		t.Logf("Testing %s getter", typ)

		args := types.Arguments{test.example}

		arg := reflect.ValueOf(args).MethodByName(fmt.Sprintf("%s", strings.ToCamel(typ))).Call([]reflect.Value{reflect.ValueOf(0)})

		assert.Equal(t, test.example, arg[0].Interface())
	}
}

func TestArguments_Type_NotFound(t *testing.T) {
	for typ, test := range argumentTests {
		t.Logf("Testing %s getter when not found", typ)

		args := types.Arguments{nil}

		arg := reflect.ValueOf(args).MethodByName(fmt.Sprintf("%s", strings.ToCamel(typ))).Call([]reflect.Value{reflect.ValueOf(0)})

		assert.Equal(t, test.def, arg[0].Interface())
	}
}

func TestArguments_TypeDefault(t *testing.T) {
	for typ, test := range argumentTests {
		t.Logf("Testing %s getter with default", typ)

		args := types.Arguments{test.example}

		arg := reflect.ValueOf(args).MethodByName(fmt.Sprintf("Default%s", strings.ToCamel(typ))).Call([]reflect.Value{reflect.ValueOf(0), reflect.ValueOf(test.example)})

		assert.Equal(t, test.example, arg[0].Interface())
	}
}

func TestArguments_TypeDefault_NotFound(t *testing.T) {
	for typ, test := range argumentTests {
		t.Logf("Testing %s getter when not found", typ)

		args := types.Arguments{nil}

		arg := reflect.ValueOf(args).MethodByName(fmt.Sprintf("Default%s", strings.ToCamel(typ))).Call([]reflect.Value{reflect.ValueOf(0), reflect.ValueOf(test.example)})

		assert.Equal(t, test.example, arg[0].Interface())
	}
}
