package smetana

type HtmlNode struct {
	Attrs Attrs
	Head  DomNode
	Body  DomNode
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
}

func Html(attrs Attrs, head any, body any) HtmlNode {
	var node HtmlNode
	node.Attrs = attrs

	switch value := head.(type) {
	case DomNode:
		node.Head = value
	case Children:
		node.Head = Head(Attrs{}, value)
	default:
		logger.Println("Invalid head value in html: ", head)
		node.Head = Head(Attrs{}, Children{})
	}

	switch value := body.(type) {
	case DomNode:
		node.Body = value
	case Children:
		node.Body = Body(Attrs{}, value)
	default:
		logger.Println("Invalid body value in html: ", head)
		node.Body = Body(Attrs{}, Children{})
	}

	return node
}
