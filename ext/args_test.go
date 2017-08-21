//+build experimental

package ext_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/goph/stdlib/ext"
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

func TestArguments_GetE(t *testing.T) {
	args := ext.Arguments{"arg"}

	arg, err := args.GetE(0)

	require.NoError(t, err)
	assert.Equal(t, "arg", arg)
}

func TestArguments_GetE_NotFound(t *testing.T) {
	args := ext.Arguments{"arg"}

	arg, err := args.GetE(1)

	require.Error(t, err)
	assert.EqualError(t, err, "no such index (1) in the argument list: there are only 1 item(s)")
	assert.Nil(t, arg)
}

func TestArguments_Get(t *testing.T) {
	args := ext.Arguments{"arg"}

	arg := args.Get(0)

	assert.Equal(t, "arg", arg)
}

func TestArguments_Get_NotFound(t *testing.T) {
	args := ext.Arguments{"arg"}

	assert.Panics(
		t,
		//errors.New("no such index (1) in the argument list: there are only 1 item(s)"),
		func() {
			args.Get(1)
		},
	)
}

func TestArguments_TypeE(t *testing.T) {
	for typ, test := range argumentTests {
		t.Logf("Testing %s getter without error", typ)

		args := ext.Arguments{test.example}

		arg := reflect.ValueOf(args).MethodByName(fmt.Sprintf("%sE", strings.ToCamel(typ))).Call([]reflect.Value{reflect.ValueOf(0)})

		assert.Equal(t, test.example, arg[0].Interface())
		assert.True(t, arg[1].IsNil())
	}
}

func TestArguments_TypeE_NotFound(t *testing.T) {
	for typ, test := range argumentTests {
		t.Logf("Testing %s getter with not found error", typ)

		args := ext.Arguments{test.example}

		arg := reflect.ValueOf(args).MethodByName(fmt.Sprintf("%sE", strings.ToCamel(typ))).Call([]reflect.Value{reflect.ValueOf(1)})

		err := arg[1].Interface().(error)

		assert.Equal(t, test.def, arg[0].Interface())
		assert.Error(t, err)
		assert.EqualError(t, err, "no such index (1) in the argument list: there are only 1 item(s)")
	}
}

func TestArguments_TypeE_InvalidType(t *testing.T) {
	for typ, test := range argumentTests {
		t.Logf("Testing %s getter with invalid type error", typ)

		args := ext.Arguments{nil}

		arg := reflect.ValueOf(args).MethodByName(fmt.Sprintf("%sE", strings.ToCamel(typ))).Call([]reflect.Value{reflect.ValueOf(0)})

		err := arg[1].Interface().(error)

		assert.Equal(t, test.def, arg[0].Interface())
		assert.Error(t, err)
		assert.EqualError(t, err, fmt.Sprintf("cannot return argument (0) as %s because it is of type <nil>", typ))
	}
}

func TestArguments_Type(t *testing.T) {
	for typ, test := range argumentTests {
		t.Logf("Testing %s getter", typ)

		args := ext.Arguments{test.example}

		arg := reflect.ValueOf(args).MethodByName(fmt.Sprintf("%s", strings.ToCamel(typ))).Call([]reflect.Value{reflect.ValueOf(0)})

		assert.Equal(t, test.example, arg[0].Interface())
	}
}

func TestArguments_Type_Panic(t *testing.T) {
	for typ, _ := range argumentTests {
		t.Logf("Testing %s getter with panic", typ)

		args := ext.Arguments{nil}

		assert.Panics(t, func() {
			reflect.ValueOf(args).MethodByName(fmt.Sprintf("%s", strings.ToCamel(typ))).Call([]reflect.Value{reflect.ValueOf(0)})
		})
	}
}
