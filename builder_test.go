package smetana

import "testing"

func TestWriteOpeningTag(t *testing.T) {
	tag := "div"
	attrs := Attrs{
		"foo":   "bar",
		"hello": "world",
	}
	builder := Builder{}
	builder.writeOpeningTag(tag, attrs)
	result := builder.Buf.String()
	var expected string
	if result[5] == 'f' {
		expected = "<div foo=\"bar\" hello=\"world\">"
	} else {
		expected = "<div hello=\"world\" foo=\"bar\">"
	}
	assertEqual(t, expected, result)
}

func TestWriteClosingTag(t *testing.T) {
	tag := "span"
	builder := Builder{}
	builder.writeClosingTag(tag)
	result := builder.Buf.String()
	assertEqual(t, "</span>", result)
}

func TestWriteShortTag(t *testing.T) {
	tag := "div"
	attrs := Attrs{
		"foo":   "bar",
		"hello": "world",
	}
	builder := Builder{}
	builder.writeShortTag(tag, attrs)
	result := builder.Buf.String()
	var expected string
	if result[5] == 'f' {
		expected = "<div foo=\"bar\" hello=\"world\" />"
	} else {
		expected = "<div hello=\"world\" foo=\"bar\" />"
	}
	assertEqual(t, expected, result)
}
