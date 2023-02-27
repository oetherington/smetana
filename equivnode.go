package smetana

import "fmt"

// A "meta" [Node] in an HTML document to associate some "content" value to a
// particular "http-equiv" (see [MetaNode] for a more generic node using "name"
// instead of "http-equiv").
type EquivNode struct {
	Equiv   string
	Content string
}

// Convert an [EquivNode] to HTML
func (node EquivNode) ToHtml(builder *Builder) {
	// `meta` is a void tag so we only need the opening tag
	builder.writeOpeningTag("meta", Attrs{
		"http-equiv": node.Equiv,
		"content":    node.Content,
	})
}

// Create a generic [EquivNode] with the given http-equiv and content
func Equiv(equiv string, content string) EquivNode {
	return EquivNode{equiv, content}
}

// Create an [EquivNode] with "http-equiv" set to "refresh" and "content"
// set to the provided value in seconds.
func Refresh(value uint) EquivNode {
	return EquivNode{"refresh", fmt.Sprintf("%d", value)}
}

// Create an [EquivNode] with "http-equiv" set to "x-ua-compatible" and
// "content" set to the given value
func XUaCompatible(value string) EquivNode {
	return EquivNode{"x-ua-compatible", value}
}
