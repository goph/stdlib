package strings

import "testing"

func TestToCamel(t *testing.T) {
	data := map[string]string{
		"foo": "Foo",

		// Snake
		"_foo":        "_Foo",
		"__foo":       "_Foo",
		"___foo":      "__Foo",
		"foo_":        "Foo_",
		"foo__":       "Foo_",
		"foo___":      "Foo__", // TODO: good idea?
		"foo_bar":     "FooBar",
		"foo_bar_baz": "FooBarBaz",
		"foo__bar":    "Foo_Bar",

		// Spinal
		"-foo":        "-Foo",
		"--foo":       "-Foo",
		"foo-":        "Foo-",
		"foo--":       "Foo-",
		"foo---":      "Foo--", // TODO: good idea?
		"foo-bar":     "FooBar",
		"foo-bar-baz": "FooBarBaz",
		"foo--bar":    "Foo-Bar",

		// Train
		"-Foo":        "-Foo",
		"--Foo":       "-Foo",
		"---Foo":      "--Foo",
		"Foo-":        "Foo-",
		"Foo--":       "Foo-",
		"Foo---":      "Foo--", // TODO: good idea?
		"Foo-Bar":     "FooBar",
		"Foo-Bar-Baz": "FooBarBaz",
		"Foo--Bar":    "Foo-Bar",
	}

	for in, expected := range data {
		actual := ToCamel(in)

		if actual != expected {
			t.Errorf("converting %q to camel case failed, expected: %q, actual: %q", in, expected, actual)
		}
	}
}

func TestToSnake(t *testing.T) {
	data := map[string]string{
		"foo":       "foo",
		"FooBar":    "foo_bar",
		"fooBar":    "foo_bar",
		"Foo_Bar":   "foo_bar",
		"Foo-Bar":   "foo_bar",
		"Foo Bar":   "foo_bar",
		"FOOBar":    "foo_bar",
		"FOOBarBaz": "foo_bar_baz",
		"FOOBarBAZ": "foo_bar_baz",
		"Foo_-Bar":  "foo__bar",
	}

	for in, expected := range data {
		actual := ToSnake(in)

		if actual != expected {
			t.Errorf("converting %q to snake case failed, expected: %q, actual: %q", in, expected, actual)
		}
	}
}

func TestToSpinal(t *testing.T) {
	data := map[string]string{
		"foo":       "foo",
		"FooBar":    "foo-bar",
		"fooBar":    "foo-bar",
		"Foo_Bar":   "foo-bar",
		"Foo-Bar":   "foo-bar",
		"Foo Bar":   "foo-bar",
		"FOOBar":    "foo-bar",
		"FOOBarBaz": "foo-bar-baz",
		"FOOBarBAZ": "foo-bar-baz",
		"Foo_-Bar":  "foo--bar",
	}

	for in, expected := range data {
		actual := ToSpinal(in)

		if actual != expected {
			t.Errorf("converting %q to spinal case failed, expected: %q, actual: %q", in, expected, actual)
		}
	}
}

func TestToTrain(t *testing.T) {
	data := map[string]string{
		"foo":       "Foo",
		"FooBar":    "Foo-Bar",
		"fooBar":    "Foo-Bar",
		"Foo_Bar":   "Foo-Bar",
		"Foo-Bar":   "Foo-Bar",
		"Foo Bar":   "Foo-Bar",
		"FOOBar":    "Foo-Bar",
		"FOOBarBaz": "Foo-Bar-Baz",
		"FOOBarBAZ": "Foo-Bar-Baz",
		"Foo_-Bar":  "Foo--Bar",
	}

	for in, expected := range data {
		actual := ToTrain(in)

		if actual != expected {
			t.Errorf("converting %q to train case failed, expected: %q, actual: %q", in, expected, actual)
		}
	}
}
