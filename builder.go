package smetana

import "strings"

type Builder struct {
	Buf strings.Builder
}

func (builder *Builder) writeAttrs(attrs Attrs) {
	for key, value := range attrs {
		builder.Buf.WriteByte(' ')
		builder.Buf.WriteString(key)
		builder.Buf.WriteString("=\"")
		builder.Buf.WriteString(value)
		builder.Buf.WriteByte('"')
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
