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
