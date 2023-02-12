package smetana

type DomNode struct {
	Tag      Tag
	Attrs    Attrs
	Children Children
}

func (node DomNode) ToHtml(builder *Builder) {
	if len(node.Children) > 0 {
		writeOpeningTag(builder, node.Tag, node.Attrs)
		writeChildren(builder, node.Children)
		writeClosingTag(builder, node.Tag)
	} else {
		writeShortTag(builder, node.Tag, node.Attrs)
	}
}

func Div(attrs Attrs, children Children) DomNode {
	return DomNode{"div", attrs, children}
}

func Span(attrs Attrs, children Children) DomNode {
	return DomNode{"span", attrs, children}
}

func Head(attrs Attrs, children Children) DomNode {
	return DomNode{"head", attrs, children}
}

func Body(attrs Attrs, children Children) DomNode {
	return DomNode{"body", attrs, children}
}

func Title(title string) DomNode {
	return DomNode{"title", Attrs{}, Children{Text(title)}}
}

func Link(rel string, href string) DomNode {
	attrs := Attrs{
		"rel":  rel,
		"href": href,
	}
	return DomNode{"link", attrs, Children{}}
}
