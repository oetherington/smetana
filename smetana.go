package smetana

import "strings"

type Node interface {
	ToHtml(b *Builder)
}

type Tag = string

type Attrs map[string]string

type Children []Node

func Render(node Node) string {
	return RenderOpts(node, false)
}

func RenderOpts(node Node, deterministicAttrs bool) string {
	builder := Builder{strings.Builder{}, deterministicAttrs}
	node.ToHtml(&builder)
	return builder.Buf.String()
}
