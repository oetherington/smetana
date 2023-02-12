package smetana

type Node interface {
	ToHtml(b *Builder)
}

func Render(node Node) string {
	builder := Builder{}
	node.ToHtml(&builder)
	return builder.Buf.String()
}

type Tag = string

type Attrs map[string]string

type Children []Node
