package smetana

/*
 * HtmlNode is the top level DomNode of an HTML document. For almost
 * all intents and purposes it functions identically to a DomNode,
 * but it needs some special handling internally to render the
 * initial `DOCTYPE` specifier.
 */
type HtmlNode struct {
	node DomNode
}

// Convert an HtmlNode to HTML
func (node HtmlNode) ToHtml(builder *Builder) {
	builder.Buf.WriteString("<!DOCTYPE html>\n")
	node.node.ToHtml(builder)
}

// Create a `html` DOM node. Arguments follow the semantics of NewDomNode.
func Html(args ...any) HtmlNode {
	return HtmlNode{NewDomNode("html", args)}
}
