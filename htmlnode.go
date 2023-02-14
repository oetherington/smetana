package smetana

type HtmlNode struct {
	node DomNode
}

func (node HtmlNode) ToHtml(builder *Builder) {
	builder.Buf.WriteString("<!DOCTYPE html>\n")
	node.node.ToHtml(builder)
}

func Html(args ...any) HtmlNode {
	return HtmlNode{NewDomNode("html", args)}
}
