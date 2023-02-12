package smetana

import "fmt"

type HtmlNode struct {
	Attrs  Attrs
	Head   DomNode
	Body   DomNode
	errors []error
}

func (node HtmlNode) ToHtml(builder *Builder) {
	builder.Buf.WriteString("<!DOCTYPE html>\n")
	builder.writeOpeningTag("html", node.Attrs)
	builder.Buf.WriteByte('\n')
	node.Head.ToHtml(builder)
	builder.Buf.WriteByte('\n')
	node.Body.ToHtml(builder)
	builder.Buf.WriteByte('\n')
	builder.writeClosingTag("html")

	for err := range node.errors {
		builder.logger.Println(err)
	}
}

func Html(attrs Attrs, head any, body any) HtmlNode {
	var node HtmlNode
	node.Attrs = attrs
	node.errors = []error{}

	switch value := head.(type) {
	case DomNode:
		node.Head = value
	case Children:
		node.Head = Head(Attrs{}, value)
	default:
		err := fmt.Errorf("Invalid head value in html: %v", head)
		node.errors = append(node.errors, err)
		node.Head = Head(Attrs{}, Children{})
	}

	switch value := body.(type) {
	case DomNode:
		node.Body = value
	case Children:
		node.Body = Body(Attrs{}, value)
	default:
		err := fmt.Errorf("Invalid body value in html: %v", head)
		node.errors = append(node.errors, err)
		node.Body = Body(Attrs{}, Children{})
	}

	return node
}
