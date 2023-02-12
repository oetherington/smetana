package smetana

type TextNode struct {
	Text string
}

func (node TextNode) ToHtml(builder *Builder) {
	builder.WriteString(node.Text)
}

func Text(text string) TextNode {
	return TextNode{text}
}
