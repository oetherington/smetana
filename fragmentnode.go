package smetana

type FragmentNode struct {
	Children Children
}

func (node FragmentNode) ToHtml(builder *Builder) {
	builder.writeChildren(node.Children)
}

func Fragment(children ...Node) FragmentNode {
	return FragmentNode{children}
}
