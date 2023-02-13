package smetana

import "fmt"

type DomNode struct {
	Tag      Tag
	Attrs    Attrs
	Children Children
	errors   []error
}

func (node DomNode) ToHtml(builder *Builder) {
	if node.errors != nil {
		for _, err := range node.errors {
			builder.logger.Println(err)
		}
	}

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

func (node *DomNode) appendError(err error) {
	if node.errors == nil {
		node.errors = []error{}
	}
	node.errors = append(node.errors, err)
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
			node.Attrs["class"] = ClassNames(node.Attrs["class"], value)
		case string:
			node.Children = append(node.Children, Text(value))
		default:
			node.appendError(fmt.Errorf("Invalid DomNode argument: %v", arg))
		}
	}
	return node
}

func A(args ...any) DomNode {
	return buildDomNode("a", args)
}

func AHref(href string, args ...any) DomNode {
	node := buildDomNode("a", args)
	node.Attrs["href"] = href
	return node
}

func Abbr(args ...any) DomNode {
	return buildDomNode("abbr", args)
}

func Address(args ...any) DomNode {
	return buildDomNode("address", args)
}

func Area(args ...any) DomNode {
	return buildDomNode("area", args)
}

func Article(args ...any) DomNode {
	return buildDomNode("article", args)
}

func Aside(args ...any) DomNode {
	return buildDomNode("aside", args)
}

func Audio(args ...any) DomNode {
	return buildDomNode("audio", args)
}

func B(args ...any) DomNode {
	return buildDomNode("b", args)
}

func Base(args ...any) DomNode {
	return buildDomNode("base", args)
}

func BaseHref(href string) DomNode {
	return DomNode{"base", Attrs{
		"href":   href,
		"target": "_blank",
	}, Children{}, nil}
}

func Bdi(args ...any) DomNode {
	return buildDomNode("bdi", args)
}

func Bdo(args ...any) DomNode {
	return buildDomNode("bdo", args)
}

func Blockquote(args ...any) DomNode {
	return buildDomNode("blockquote", args)
}

func Body(args ...any) DomNode {
	return buildDomNode("body", args)
}

func Br(args ...any) DomNode {
	return buildDomNode("br", args)
}

func Button(args ...any) DomNode {
	return buildDomNode("button", args)
}

func Canvas(args ...any) DomNode {
	return buildDomNode("canvas", args)
}

func Caption(args ...any) DomNode {
	return buildDomNode("caption", args)
}

func Charset(value string) DomNode {
	if len(value) < 1 {
		value = "UTF-8"
	}
	return DomNode{"meta", Attrs{"charset": value}, Children{}, nil}
}

func Cite(args ...any) DomNode {
	return buildDomNode("cite", args)
}

func Code(args ...any) DomNode {
	return buildDomNode("code", args)
}

func Col(args ...any) DomNode {
	return buildDomNode("col", args)
}

func Colgroup(args ...any) DomNode {
	return buildDomNode("colgroup", args)
}

func Data(args ...any) DomNode {
	return buildDomNode("data", args)
}

func Datalist(args ...any) DomNode {
	return buildDomNode("datalist", args)
}

func Dd(args ...any) DomNode {
	return buildDomNode("dd", args)
}

func Del(args ...any) DomNode {
	return buildDomNode("del", args)
}

func Details(args ...any) DomNode {
	return buildDomNode("details", args)
}

func Dfn(args ...any) DomNode {
	return buildDomNode("dfn", args)
}

func Dialog(args ...any) DomNode {
	return buildDomNode("dialog", args)
}

func Div(args ...any) DomNode {
	return buildDomNode("div", args)
}

func Dl(args ...any) DomNode {
	return buildDomNode("dl", args)
}

func Dt(args ...any) DomNode {
	return buildDomNode("dt", args)
}

func Em(args ...any) DomNode {
	return buildDomNode("em", args)
}

func Embed(args ...any) DomNode {
	return buildDomNode("embed", args)
}

func Fieldset(args ...any) DomNode {
	return buildDomNode("fieldset", args)
}

func Figcaption(args ...any) DomNode {
	return buildDomNode("figcaption", args)
}

func Figure(args ...any) DomNode {
	return buildDomNode("figure", args)
}

func Footer(args ...any) DomNode {
	return buildDomNode("footer", args)
}

func Form(args ...any) DomNode {
	return buildDomNode("form", args)
}

func H(level int, args ...any) DomNode {
	tag := fmt.Sprintf("h%d", level)
	return buildDomNode(tag, args)
}

func H1(args ...any) DomNode {
	return buildDomNode("h1", args)
}

func H2(args ...any) DomNode {
	return buildDomNode("h2", args)
}

func H3(args ...any) DomNode {
	return buildDomNode("h3", args)
}

func H4(args ...any) DomNode {
	return buildDomNode("h4", args)
}

func H5(args ...any) DomNode {
	return buildDomNode("h5", args)
}

func H6(args ...any) DomNode {
	return buildDomNode("h6", args)
}

func Head(args ...any) DomNode {
	return buildDomNode("head", args)
}

func Header(args ...any) DomNode {
	return buildDomNode("header", args)
}

func Hr(args ...any) DomNode {
	return buildDomNode("hr", args)
}

func I(args ...any) DomNode {
	return buildDomNode("i", args)
}

func Iframe(args ...any) DomNode {
	return buildDomNode("iframe", args)
}

func Img(args ...any) DomNode {
	return buildDomNode("img", args)
}

func Input(args ...any) DomNode {
	return buildDomNode("input", args)
}

func Ins(args ...any) DomNode {
	return buildDomNode("ins", args)
}

func Kbd(args ...any) DomNode {
	return buildDomNode("kbd", args)
}

func Label(args ...any) DomNode {
	return buildDomNode("label", args)
}

func Legend(args ...any) DomNode {
	return buildDomNode("legend", args)
}

func Li(args ...any) DomNode {
	return buildDomNode("li", args)
}

func Link(args ...any) DomNode {
	return buildDomNode("link", args)
}

func LinkHref(rel string, href string) DomNode {
	attrs := Attrs{
		"rel":  rel,
		"href": href,
	}
	return DomNode{"link", attrs, Children{}, nil}
}

func Main(args ...any) DomNode {
	return buildDomNode("main", args)
}

func Map(args ...any) DomNode {
	return buildDomNode("map", args)
}

func Mark(args ...any) DomNode {
	return buildDomNode("mark", args)
}

func Meter(args ...any) DomNode {
	return buildDomNode("meter", args)
}

func Nav(args ...any) DomNode {
	return buildDomNode("nav", args)
}

func Noscript(args ...any) DomNode {
	return buildDomNode("noscript", args)
}

func Object(args ...any) DomNode {
	return buildDomNode("object", args)
}

func Ol(args ...any) DomNode {
	return buildDomNode("ol", args)
}

func Optgroup(args ...any) DomNode {
	return buildDomNode("optgroup", args)
}

func Option(args ...any) DomNode {
	return buildDomNode("option", args)
}

func Output(args ...any) DomNode {
	return buildDomNode("output", args)
}

func P(args ...any) DomNode {
	return buildDomNode("p", args)
}

func Param(args ...any) DomNode {
	return buildDomNode("param", args)
}

func Picture(args ...any) DomNode {
	return buildDomNode("picture", args)
}

func Pre(args ...any) DomNode {
	return buildDomNode("pre", args)
}

func Progress(args ...any) DomNode {
	return buildDomNode("progress", args)
}

func Q(args ...any) DomNode {
	return buildDomNode("q", args)
}

func Refresh(value uint) DomNode {
	return DomNode{"meta", Attrs{
		"http-equiv": "refresh",
		"content":    fmt.Sprintf("%d", value),
	}, Children{}, nil}
}

func Rp(args ...any) DomNode {
	return buildDomNode("rp", args)
}

func Rt(args ...any) DomNode {
	return buildDomNode("rt", args)
}

func Ruby(args ...any) DomNode {
	return buildDomNode("ruby", args)
}

func S(args ...any) DomNode {
	return buildDomNode("s", args)
}

func Samp(args ...any) DomNode {
	return buildDomNode("samp", args)
}

func Script(args ...any) DomNode {
	return buildDomNode("script", args)
}

func ScriptSrc(src string) DomNode {
	return DomNode{"script", Attrs{"src": src}, Children{}, nil}
}

func Section(args ...any) DomNode {
	return buildDomNode("section", args)
}

func Select(args ...any) DomNode {
	return buildDomNode("select", args)
}

func Small(args ...any) DomNode {
	return buildDomNode("small", args)
}

func Source(args ...any) DomNode {
	return buildDomNode("source", args)
}

func Span(args ...any) DomNode {
	return buildDomNode("span", args)
}

func Strong(args ...any) DomNode {
	return buildDomNode("strong", args)
}

func Style(args ...any) DomNode {
	return buildDomNode("style", args)
}

func Sub(args ...any) DomNode {
	return buildDomNode("sub", args)
}

func Summary(args ...any) DomNode {
	return buildDomNode("summary", args)
}

func Sup(args ...any) DomNode {
	return buildDomNode("sup", args)
}

func Svg(args ...any) DomNode {
	return buildDomNode("svg", args)
}

func Table(args ...any) DomNode {
	return buildDomNode("table", args)
}

func Tbody(args ...any) DomNode {
	return buildDomNode("tbody", args)
}

func Td(args ...any) DomNode {
	return buildDomNode("td", args)
}

func Template(args ...any) DomNode {
	return buildDomNode("template", args)
}

func Textarea(args ...any) DomNode {
	return buildDomNode("textarea", args)
}

func Tfoot(args ...any) DomNode {
	return buildDomNode("tfoot", args)
}

func Th(args ...any) DomNode {
	return buildDomNode("th", args)
}

func Thead(args ...any) DomNode {
	return buildDomNode("thead", args)
}

func Time(args ...any) DomNode {
	return buildDomNode("time", args)
}

func Title(args ...any) DomNode {
	return buildDomNode("title", args)
}

func Tr(args ...any) DomNode {
	return buildDomNode("tr", args)
}

func Track(args ...any) DomNode {
	return buildDomNode("track", args)
}

func U(args ...any) DomNode {
	return buildDomNode("u", args)
}

func Ul(args ...any) DomNode {
	return buildDomNode("ul", args)
}

func Var(args ...any) DomNode {
	return buildDomNode("var", args)
}

func Video(args ...any) DomNode {
	return buildDomNode("video", args)
}

func Wbr(args ...any) DomNode {
	return buildDomNode("wbr", args)
}
