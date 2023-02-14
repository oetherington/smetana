package smetana

import "fmt"

/*
 * DomNode is a basic Node that renders into a particular HTML tag
 * with optional attributes and/or children. Any errors that are
 * generated are stored in `errors`.
 */
type DomNode struct {
	Tag      Tag
	Attrs    Attrs
	Children Children
	errors   []error
}

// Convert a DomNode to HTML
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

/*
 * Assign new attributes to a DomNode. These values are merged
 * together with any existing attributes. If a value exists in
 * both the old attributes map and the new attributes map then
 * the new value is used.
 */
func (node *DomNode) AssignAttrs(attrs Attrs) {
	if len(node.Attrs) < 1 {
		node.Attrs = attrs
	} else {
		mergeMaps(node.Attrs, attrs)
	}
}

// Append more children to the end of a DomNode
func (node *DomNode) AssignChildren(children Children) {
	if len(node.Children) < 1 {
		node.Children = children
	} else {
		node.Children = append(node.Children, children...)
	}
}

// Record a compilation error for a DomNode
func (node *DomNode) appendError(err error) {
	if node.errors == nil {
		node.errors = []error{}
	}
	node.errors = append(node.errors, err)
}

func NewDomNode(tag Tag, args []any) DomNode {
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
		case ClassName, Classes:
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
	return NewDomNode("a", args)
}

func AHref(href string, args ...any) DomNode {
	node := NewDomNode("a", args)
	node.Attrs["href"] = href
	return node
}

func Abbr(args ...any) DomNode {
	return NewDomNode("abbr", args)
}

func Address(args ...any) DomNode {
	return NewDomNode("address", args)
}

func Area(args ...any) DomNode {
	return NewDomNode("area", args)
}

func Article(args ...any) DomNode {
	return NewDomNode("article", args)
}

func Aside(args ...any) DomNode {
	return NewDomNode("aside", args)
}

func Audio(args ...any) DomNode {
	return NewDomNode("audio", args)
}

func B(args ...any) DomNode {
	return NewDomNode("b", args)
}

func Base(args ...any) DomNode {
	return NewDomNode("base", args)
}

func BaseHref(href string) DomNode {
	return DomNode{"base", Attrs{
		"href":   href,
		"target": "_blank",
	}, Children{}, nil}
}

func Bdi(args ...any) DomNode {
	return NewDomNode("bdi", args)
}

func Bdo(args ...any) DomNode {
	return NewDomNode("bdo", args)
}

func Blockquote(args ...any) DomNode {
	return NewDomNode("blockquote", args)
}

func Body(args ...any) DomNode {
	return NewDomNode("body", args)
}

func Br(args ...any) DomNode {
	return NewDomNode("br", args)
}

func Button(args ...any) DomNode {
	return NewDomNode("button", args)
}

func Canvas(args ...any) DomNode {
	return NewDomNode("canvas", args)
}

func Caption(args ...any) DomNode {
	return NewDomNode("caption", args)
}

func Charset(value string) DomNode {
	if len(value) < 1 {
		value = "UTF-8"
	}
	return DomNode{"meta", Attrs{"charset": value}, Children{}, nil}
}

func Cite(args ...any) DomNode {
	return NewDomNode("cite", args)
}

func Code(args ...any) DomNode {
	return NewDomNode("code", args)
}

func Col(args ...any) DomNode {
	return NewDomNode("col", args)
}

func Colgroup(args ...any) DomNode {
	return NewDomNode("colgroup", args)
}

func Data(args ...any) DomNode {
	return NewDomNode("data", args)
}

func Datalist(args ...any) DomNode {
	return NewDomNode("datalist", args)
}

func Dd(args ...any) DomNode {
	return NewDomNode("dd", args)
}

func Del(args ...any) DomNode {
	return NewDomNode("del", args)
}

func Details(args ...any) DomNode {
	return NewDomNode("details", args)
}

func Dfn(args ...any) DomNode {
	return NewDomNode("dfn", args)
}

func Dialog(args ...any) DomNode {
	return NewDomNode("dialog", args)
}

func Div(args ...any) DomNode {
	return NewDomNode("div", args)
}

func Dl(args ...any) DomNode {
	return NewDomNode("dl", args)
}

func Dt(args ...any) DomNode {
	return NewDomNode("dt", args)
}

func Em(args ...any) DomNode {
	return NewDomNode("em", args)
}

func Embed(args ...any) DomNode {
	return NewDomNode("embed", args)
}

func Fieldset(args ...any) DomNode {
	return NewDomNode("fieldset", args)
}

func Figcaption(args ...any) DomNode {
	return NewDomNode("figcaption", args)
}

func Figure(args ...any) DomNode {
	return NewDomNode("figure", args)
}

func Footer(args ...any) DomNode {
	return NewDomNode("footer", args)
}

func Form(args ...any) DomNode {
	return NewDomNode("form", args)
}

func H(level int, args ...any) DomNode {
	tag := fmt.Sprintf("h%d", level)
	return NewDomNode(tag, args)
}

func H1(args ...any) DomNode {
	return NewDomNode("h1", args)
}

func H2(args ...any) DomNode {
	return NewDomNode("h2", args)
}

func H3(args ...any) DomNode {
	return NewDomNode("h3", args)
}

func H4(args ...any) DomNode {
	return NewDomNode("h4", args)
}

func H5(args ...any) DomNode {
	return NewDomNode("h5", args)
}

func H6(args ...any) DomNode {
	return NewDomNode("h6", args)
}

func Head(args ...any) DomNode {
	return NewDomNode("head", args)
}

func Header(args ...any) DomNode {
	return NewDomNode("header", args)
}

func Hr(args ...any) DomNode {
	return NewDomNode("hr", args)
}

func I(args ...any) DomNode {
	return NewDomNode("i", args)
}

func Iframe(args ...any) DomNode {
	return NewDomNode("iframe", args)
}

func Img(args ...any) DomNode {
	return NewDomNode("img", args)
}

func Input(args ...any) DomNode {
	return NewDomNode("input", args)
}

func Ins(args ...any) DomNode {
	return NewDomNode("ins", args)
}

func Kbd(args ...any) DomNode {
	return NewDomNode("kbd", args)
}

func Label(args ...any) DomNode {
	return NewDomNode("label", args)
}

func Legend(args ...any) DomNode {
	return NewDomNode("legend", args)
}

func Li(args ...any) DomNode {
	return NewDomNode("li", args)
}

func Link(args ...any) DomNode {
	return NewDomNode("link", args)
}

func LinkHref(rel string, href string) DomNode {
	attrs := Attrs{
		"rel":  rel,
		"href": href,
	}
	return DomNode{"link", attrs, Children{}, nil}
}

func Main(args ...any) DomNode {
	return NewDomNode("main", args)
}

func Map(args ...any) DomNode {
	return NewDomNode("map", args)
}

func Mark(args ...any) DomNode {
	return NewDomNode("mark", args)
}

func Meter(args ...any) DomNode {
	return NewDomNode("meter", args)
}

func Nav(args ...any) DomNode {
	return NewDomNode("nav", args)
}

func Noscript(args ...any) DomNode {
	return NewDomNode("noscript", args)
}

func Object(args ...any) DomNode {
	return NewDomNode("object", args)
}

func Ol(args ...any) DomNode {
	return NewDomNode("ol", args)
}

func Optgroup(args ...any) DomNode {
	return NewDomNode("optgroup", args)
}

func Option(args ...any) DomNode {
	return NewDomNode("option", args)
}

func Output(args ...any) DomNode {
	return NewDomNode("output", args)
}

func P(args ...any) DomNode {
	return NewDomNode("p", args)
}

func Param(args ...any) DomNode {
	return NewDomNode("param", args)
}

func Picture(args ...any) DomNode {
	return NewDomNode("picture", args)
}

func Pre(args ...any) DomNode {
	return NewDomNode("pre", args)
}

func Progress(args ...any) DomNode {
	return NewDomNode("progress", args)
}

func Q(args ...any) DomNode {
	return NewDomNode("q", args)
}

func Refresh(value uint) DomNode {
	return DomNode{"meta", Attrs{
		"http-equiv": "refresh",
		"content":    fmt.Sprintf("%d", value),
	}, Children{}, nil}
}

func Rp(args ...any) DomNode {
	return NewDomNode("rp", args)
}

func Rt(args ...any) DomNode {
	return NewDomNode("rt", args)
}

func Ruby(args ...any) DomNode {
	return NewDomNode("ruby", args)
}

func S(args ...any) DomNode {
	return NewDomNode("s", args)
}

func Samp(args ...any) DomNode {
	return NewDomNode("samp", args)
}

func Script(args ...any) DomNode {
	return NewDomNode("script", args)
}

func ScriptSrc(src string) DomNode {
	return DomNode{"script", Attrs{"src": src}, Children{}, nil}
}

func Section(args ...any) DomNode {
	return NewDomNode("section", args)
}

func Select(args ...any) DomNode {
	return NewDomNode("select", args)
}

func Small(args ...any) DomNode {
	return NewDomNode("small", args)
}

func Source(args ...any) DomNode {
	return NewDomNode("source", args)
}

func Span(args ...any) DomNode {
	return NewDomNode("span", args)
}

func Strong(args ...any) DomNode {
	return NewDomNode("strong", args)
}

func Style(args ...any) DomNode {
	return NewDomNode("style", args)
}

func Sub(args ...any) DomNode {
	return NewDomNode("sub", args)
}

func Summary(args ...any) DomNode {
	return NewDomNode("summary", args)
}

func Sup(args ...any) DomNode {
	return NewDomNode("sup", args)
}

func Svg(args ...any) DomNode {
	return NewDomNode("svg", args)
}

func Table(args ...any) DomNode {
	return NewDomNode("table", args)
}

func Tbody(args ...any) DomNode {
	return NewDomNode("tbody", args)
}

func Td(args ...any) DomNode {
	return NewDomNode("td", args)
}

func Template(args ...any) DomNode {
	return NewDomNode("template", args)
}

func Textarea(args ...any) DomNode {
	return NewDomNode("textarea", args)
}

func Tfoot(args ...any) DomNode {
	return NewDomNode("tfoot", args)
}

func Th(args ...any) DomNode {
	return NewDomNode("th", args)
}

func Thead(args ...any) DomNode {
	return NewDomNode("thead", args)
}

func Time(args ...any) DomNode {
	return NewDomNode("time", args)
}

func Title(args ...any) DomNode {
	return NewDomNode("title", args)
}

func Tr(args ...any) DomNode {
	return NewDomNode("tr", args)
}

func Track(args ...any) DomNode {
	return NewDomNode("track", args)
}

func U(args ...any) DomNode {
	return NewDomNode("u", args)
}

func Ul(args ...any) DomNode {
	return NewDomNode("ul", args)
}

func Var(args ...any) DomNode {
	return NewDomNode("var", args)
}

func Video(args ...any) DomNode {
	return NewDomNode("video", args)
}

func Wbr(args ...any) DomNode {
	return NewDomNode("wbr", args)
}
