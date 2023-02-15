package smetana

import (
	"log"
	"sort"
	"strings"
)

// Struct for tracking internal state during HTML compilation.
//   - `Buf` is the string buffer being written to.
//   - By default, the order of HTML tag attributes is undefined and
//     non-deterministic. It can be changed to be deterministic by
//     setting `deterministicAttributes` to true. Note that this has
//     a significant performance cost.
//   - `logger` is used for reporting warnings and errors during
//     compilation.
type Builder struct {
	Buf                     strings.Builder
	DeterministicAttributes bool
	Logger                  *log.Logger
}

func (builder *Builder) writeAttr(key string, value string) {
	builder.Buf.WriteByte(' ')
	builder.Buf.WriteString(key)
	builder.Buf.WriteString("=\"")
	builder.Buf.WriteString(value)
	builder.Buf.WriteByte('"')
}

func (builder *Builder) writeAttrs(attrs Attrs) {
	if builder.DeterministicAttributes {
		keys := make([]string, 0, len(attrs))
		for k := range attrs {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		for _, key := range keys {
			builder.writeAttr(key, attrs[key])
		}
	} else {
		for key, value := range attrs {
			builder.writeAttr(key, value)
		}
	}
}

func (builder *Builder) writeOpeningTag(tag Tag, attrs Attrs) {
	builder.Buf.WriteByte('<')
	builder.Buf.WriteString(tag)
	builder.writeAttrs(attrs)
	builder.Buf.WriteByte('>')
}

func (builder *Builder) writeClosingTag(tag Tag) {
	builder.Buf.WriteString("</")
	builder.Buf.WriteString(tag)
	builder.Buf.WriteByte('>')
}

func (builder *Builder) writeShortTag(tag Tag, attrs Attrs) {
	builder.Buf.WriteByte('<')
	builder.Buf.WriteString(tag)
	builder.writeAttrs(attrs)
	builder.Buf.WriteString(" />")
}

func (builder *Builder) writeChildren(children Children) {
	for _, child := range children {
		child.ToHtml(builder)
	}
}
