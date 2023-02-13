package smetana

import "fmt"

type DomNode struct {
	Tag      Tag
	Attrs    Attrs
	Children Children
	errors   []error
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

func (node *DomNode) AssignAttrs(attrs Attrs) {
	if len(node.Attrs) < 1 {
		node.Attrs = attrs
	} else {
		mergeMaps(node.Attrs, attrs)
	}
}

func (node *DomNode) AssignChildren(children Children) {
	if len(node.Children) < 1 {
		node.Children = children
	} else {
		node.Children = append(node.Children, children...)
	}
}

func buildDomNode(tag Tag, args []any) DomNode {
	node := DomNode{tag, Attrs{}, Children{}, nil}
	for _, arg := range args {
		switch value := arg.(type) {
		case Attrs:
			node.AssignAttrs(value)
		case Children:
			node.AssignChildren(value)
		case Attr:
			node.Attrs[value.Key] = value.Value
		case Node:
			node.Children = append(node.Children, value)
		case Classes:
			node.Attrs["class"] = Class(node.Attrs["class"], value)
		default:
			err := fmt.Errorf("Invalid DomNode argument: %v", arg)
			node.errors = append(node.errors, err)
		}
	}
	return node
}

func Div(args ...any) DomNode {
	return buildDomNode("div", args)
}

func Span(args ...any) DomNode {
	return buildDomNode("span", args)
}

func Head(args ...any) DomNode {
	return buildDomNode("head", args)
}

func Body(args ...any) DomNode {
	return buildDomNode("body", args)
}

func Title(title string) DomNode {
	return DomNode{"title", Attrs{}, Children{Text(title)}, nil}
}

func Link(rel string, href string) DomNode {
	attrs := Attrs{
		"rel":  rel,
		"href": href,
	}
	return DomNode{"link", attrs, Children{}, nil}
}

func Charset(value string) DomNode {
	if len(value) < 1 {
		value = "UTF-8"
	}
	return DomNode{"meta", Attrs{"charset": value}, Children{}, nil}
}

func Refresh(value uint) DomNode {
	return DomNode{"meta", Attrs{
		"http-equiv": "refresh",
		"content":    fmt.Sprintf("%d", value),
	}, Children{}, nil}
}

func Base(href string) DomNode {
	return DomNode{"base", Attrs{
		"href":   href,
		"target": "_blank",
	}, Children{}, nil}
}

func Script(src string) DomNode {
	return DomNode{"script", Attrs{"src": src}, Children{}, nil}
}

func InlineScript(value string) DomNode {
	return DomNode{"script", Attrs{}, Children{Text(value)}, nil}
}

func Style(value string) DomNode {
	return DomNode{"style", Attrs{}, Children{Text(value)}, nil}
}
