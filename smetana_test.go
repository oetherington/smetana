package smetana

import (
	"reflect"
	"runtime/debug"
	"testing"
)

func assert(t *testing.T, exp, got interface{}, equal bool) {
	if reflect.DeepEqual(exp, got) != equal {
		debug.PrintStack()
		t.Fatalf("Expecting '%v' got '%v'\n", exp, got)
	}
}

func assertEqual(t *testing.T, exp, got interface{}) {
	assert(t, exp, got, true)
}

func assertNotEqual(t *testing.T, exp, got interface{}) {
	assert(t, exp, got, false)
}
