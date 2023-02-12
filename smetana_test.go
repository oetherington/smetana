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

func TestWriteOpeningTag(t *testing.T) {
	var builder Builder
	tag := "div"
	attrs := Attrs{
		"foo":   "bar",
		"hello": "world",
	}
	writeOpeningTag(&builder, tag, attrs)
	result := builder.String()
	var expected string
	if result[5] == 'f' {
		expected = "<div foo=\"bar\" hello=\"world\">"
	} else {
		expected = "<div hello=\"world\" foo=\"bar\">"
	}
	assertEqual(t, expected, result)
}

func TestWriteClosingTag(t *testing.T) {
	var builder Builder
	tag := "span"
	writeClosingTag(&builder, tag)
	result := builder.String()
	assertEqual(t, "</span>", result)
}

func TestWriteShortTag(t *testing.T) {
	var builder Builder
	tag := "div"
	attrs := Attrs{
		"foo":   "bar",
		"hello": "world",
	}
	writeShortTag(&builder, tag, attrs)
	result := builder.String()
	var expected string
	if result[5] == 'f' {
		expected = "<div foo=\"bar\" hello=\"world\" />"
	} else {
		expected = "<div hello=\"world\" foo=\"bar\" />"
	}
	assertEqual(t, expected, result)
}
