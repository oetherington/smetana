package smetana

/*
 * Sometimes we want to combine multiple nodes at the same level of a
 * document to treat them as a single unit. In some cases it may be
 * acceptable to wrap them in another node such as a `div` or `span`
 * but this is often undesirable and it alters the generated markup.
 * In these cases, the children nodes can instead be wrapped in a
 * `FragmentNode` to treat them as a single entity but without adding
 * an extra layer to the generated markup.
 */
type FragmentNode struct {
	Children Children
}

// Convert a FragmentNode to HTML
func (node FragmentNode) ToHtml(builder *Builder) {
	builder.writeChildren(node.Children)
}

// Create a FragmentNode with the given children
func Fragment(children ...Node) FragmentNode {
	return FragmentNode{children}
}

// Append more children to the end of a FragmentNode
func (node *FragmentNode) AssignChildren(children Children) {
	if len(node.Children) < 1 {
		node.Children = children
	} else {
		node.Children = append(node.Children, children...)
	}
}
