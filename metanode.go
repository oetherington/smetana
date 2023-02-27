package smetana

// A "meta" [Node] in an HTML document to associate some "content" value to a
// particular "name".
type MetaNode struct {
	Name    string
	Content string
}

// Convert a [MetaNode] to HTML
func (node MetaNode) ToHtml(builder *Builder) {
	// `meta` is a void tag so we only need the opening tag
	builder.writeOpeningTag("meta", Attrs{
		"name":    node.Name,
		"content": node.Content,
	})
}

// Create a generic [MetaNode] with the given name and content
func Meta(name string, content string) MetaNode {
	return MetaNode{name, content}
}

// Create a "keywords" [MetaNode] with the given value
func Keywords(value string) MetaNode {
	return MetaNode{"keywords", value}
}

// Create a "description" [MetaNode] with the given value
func Description(value string) MetaNode {
	return MetaNode{"description", value}
}

// Create an "author" [MetaNode] with the given value
func Author(value string) MetaNode {
	return MetaNode{"author", value}
}

// Create a "viewport" [MetaNode] with the given value, or, if the
// empty string is provided, the default "width=device-width, initial-scale=1.0"
func Viewport(value string) MetaNode {
	if len(value) < 1 {
		value = "width=device-width, initial-scale=1.0"
	}
	return MetaNode{"viewport", value}
}
