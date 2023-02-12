package smetana

import (
	"log"
	"os"
	"strings"
)

type Node interface {
	ToHtml(b *Builder)
}

type Tag = string

type Attrs map[string]string

type Children []Node

func Render(node Node) string {
	return RenderOpts(node, false, nil)
}

func RenderOpts(node Node, deterministicAttrs bool, logger *log.Logger) string {
	if logger == nil {
		logger = log.New(os.Stderr, "", 0)
	}
	builder := Builder{strings.Builder{}, deterministicAttrs, logger}
	node.ToHtml(&builder)
	return builder.Buf.String()
}
