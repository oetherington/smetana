package smetana

// A Node representing raw text without any surrounding tag
type TextNode struct {
	Text string
}

// Convert a TextNode to HTML
func (node TextNode) ToHtml(builder *Builder) {
	builder.Buf.WriteString(node.Text)
}

// Create a TextNode from the given string
func Text(text string) TextNode {
	return TextNode{text}
}
