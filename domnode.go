package smetana

import "fmt"

type DomNode struct {
	Tag      Tag
	Attrs    Attrs
	Children Children
}

func (node DomNode) ToHtml(builder *Builder) {
	if len(node.Children) > 0 {
		builder.writeOpeningTag(node.Tag, node.Attrs)
		builder.writeChildren(node.Children)
		builder.writeClosingTag(node.Tag)
	} else {
		builder.writeShortTag(node.Tag, node.Attrs)
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

func Charset(value string) DomNode {
	if len(value) < 1 {
		value = "UTF-8"
	}
	return DomNode{"meta", Attrs{"charset": value}, Children{}}
}

func Refresh(value uint) DomNode {
	return DomNode{"meta", Attrs{
		"http-equiv": "refresh",
		"content":    fmt.Sprintf("%d", value),
	}, Children{}}
}

func Base(href string) DomNode {
	return DomNode{"base", Attrs{
		"href":   href,
		"target": "_blank",
	}, Children{}}
}

func Script(value string) DomNode {
	return DomNode{"script", Attrs{}, Children{Text(value)}}
}

func Style(value string) DomNode {
	return DomNode{"style", Attrs{}, Children{Text(value)}}
}
