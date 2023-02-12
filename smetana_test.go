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
	assert(t, expected, result, true)
}

func TestWriteClosingTag(t *testing.T) {
	var builder Builder
	tag := "span"
	writeClosingTag(&builder, tag)
	result := builder.String()
	assert(t, "</span>", result, true)
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
	assert(t, expected, result, true)
}
