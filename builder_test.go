package smetana

import (
	"strings"
	"testing"
)

func TestWriteOpeningTag(t *testing.T) {
	tag := "div"
	attrs := Attrs{
		"foo":   "bar",
		"hello": "world",
	}
	builder := Builder{strings.Builder{}, true}
	builder.writeOpeningTag(tag, attrs)
	result := builder.Buf.String()
	assertEqual(t, "<div foo=\"bar\" hello=\"world\">", result)
}

func TestWriteClosingTag(t *testing.T) {
	tag := "span"
	builder := Builder{strings.Builder{}, true}
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
	builder := Builder{strings.Builder{}, true}
	builder.writeShortTag(tag, attrs)
	result := builder.Buf.String()
	assertEqual(t, "<div foo=\"bar\" hello=\"world\" />", result)
}
