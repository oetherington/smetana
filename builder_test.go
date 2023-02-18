package smetana

import (
	"log"
	"strings"
	"testing"
)

func TestWriteOpeningTag(t *testing.T) {
	tag := "div"
	attrs := Attrs{
		"foo":   "bar",
		"hello": "world",
	}
	builder := Builder{strings.Builder{}, true, nil}
	builder.writeOpeningTag(tag, attrs)
	result := builder.Buf.String()
	assertEqual(t, "<div foo=\"bar\" hello=\"world\">", result)
}

func TestWriteOpeningTagWithNonDeterministicAttributes(t *testing.T) {
	tag := "div"
	attrs := Attrs{
		"foo":   "bar",
		"hello": "world",
	}
	builder := Builder{strings.Builder{}, false, nil}
	builder.writeOpeningTag(tag, attrs)
	result := builder.Buf.String()
	if result[5] == 'f' {
		assertEqual(t, "<div foo=\"bar\" hello=\"world\">", result)
	} else {
		assertEqual(t, "<div hello=\"world\" foo=\"bar\">", result)
	}
}

func TestWriteClosingTag(t *testing.T) {
	tag := "span"
	builder := Builder{strings.Builder{}, true, nil}
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
	builder := Builder{strings.Builder{}, true, nil}
	builder.writeShortTag(tag, attrs)
	result := builder.Buf.String()
	assertEqual(t, "<div foo=\"bar\" hello=\"world\" />", result)
}

func TestCustomLogger(t *testing.T) {
	var target strings.Builder
	logger := log.New(&target, "", 0)
	builder := Builder{strings.Builder{}, true, logger}
	builder.Logger.Print("Hello world")
	result := strings.TrimSpace(target.String())
	assertEqual(t, "Hello world", result)
}
