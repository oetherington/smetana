package smetana

import (
	"reflect"
	"runtime/debug"
	"testing"
)

func assert[T any](t *testing.T, exp T, got T, equal bool) {
	if reflect.DeepEqual(exp, got) != equal {
		debug.PrintStack()
		t.Fatalf("Expecting '%v' got '%v'\n", exp, got)
	}
}

func assertEqual[T any](t *testing.T, exp T, got T) {
	assert(t, exp, got, true)
}

func assertNotEqual[T any](t *testing.T, exp T, got T) {
	assert(t, exp, got, false)
}

func assertOneOf[T any](t *testing.T, exp []T, got T) {
	for _, option := range exp {
		if reflect.DeepEqual(option, got) {
			return
		}
	}

	debug.PrintStack()
	t.Fatalf("Expecting one of '%v' got '%v'\n", exp, got)
}

func TestMergeMaps(t *testing.T) {
	m1 := map[string]int{"foo": 1, "bar": 2}
	m2 := map[string]int{"baz": 3}
	mergeMaps(m1, m2)
	assertEqual(t, map[string]int{"foo": 1, "bar": 2, "baz": 3}, m1)
	assertEqual(t, map[string]int{"baz": 3}, m2)
}

func TestMergeMapsOverwritesDuplicates(t *testing.T) {
	m1 := map[string]int{"foo": 1, "bar": 2, "baz": 4}
	m2 := map[string]int{"baz": 3}
	mergeMaps(m1, m2)
	assertEqual(t, map[string]int{"foo": 1, "bar": 2, "baz": 3}, m1)
	assertEqual(t, map[string]int{"baz": 3}, m2)
}

func TestId(t *testing.T) {
	result := Id("foo")
	assertEqual(t, Attr{Key: "id", Value: "foo"}, result)
}
