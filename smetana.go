package smetana

import "strings"

type Builder = strings.Builder

type Node interface {
	ToHtml(sb *Builder)
}

func Render(node Node) string {
	var builder Builder
	node.ToHtml(&builder)
	return builder.String()
}

type Tag = string

type Attrs map[string]string

type Children []Node

func writeOpeningTag(builder *Builder, tag Tag, attrs Attrs) {
	builder.WriteByte('<')
	builder.WriteString(tag)

	for key, value := range attrs {
		builder.WriteByte(' ')
		builder.WriteString(key)
		builder.WriteString("=\"")
		builder.WriteString(value)
		builder.WriteByte('"')
	}

	builder.WriteByte('>')
}

func writeClosingTag(builder *Builder, tag Tag) {
	builder.WriteString("</")
	builder.WriteString(tag)
	builder.WriteByte('>')
}

func writeShortTag(builder *Builder, tag Tag, attrs Attrs) {
	builder.WriteByte('<')
	builder.WriteString(tag)

	for key, value := range attrs {
		builder.WriteByte(' ')
		builder.WriteString(key)
		builder.WriteString("=\"")
		builder.WriteString(value)
		builder.WriteByte('"')
	}

	builder.WriteString(" />")
}

func writeChildren(builder *Builder, children Children) {
	for _, child := range children {
		child.ToHtml(builder)
	}
}
