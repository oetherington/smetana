package smetana

type MetaNode struct {
	Name    string
	Content string
}

func (node MetaNode) ToHtml(builder *Builder) {
	writeShortTag(builder, "meta", Attrs{
		"name":    node.Name,
		"content": node.Content,
	})
}

func Meta(name string, content string) MetaNode {
	return MetaNode{name, content}
}

func Keywords(value string) MetaNode {
	return MetaNode{"keywords", value}
}

func Description(value string) MetaNode {
	return MetaNode{"description", value}
}

func Author(value string) MetaNode {
	return MetaNode{"author", value}
}

func Viewport(value string) MetaNode {
	if len(value) < 1 {
		value = "width=device-width, initial-scale=1.0"
	}
	return MetaNode{"viewport", value}
}
