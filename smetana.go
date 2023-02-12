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
