package smetana

import "fmt"

// [DomNode] is a basic [Node] that renders into a particular HTML tag
// with optional attributes and/or children.
type DomNode struct {
	Tag      Tag
	Attrs    Attrs
	Children Children
	errors   []error
}

// Convert a [DomNode] to HTML.
func (node DomNode) ToHtml(builder *Builder) {
	if node.errors != nil {
		for _, err := range node.errors {
			builder.Logger.Println(err)
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

// Assign new attributes to a [DomNode]. These values are merged
// together with any existing attributes. If a value exists in
// both the old attributes map and the new attributes map then
// the new value is used.
func (node *DomNode) AssignAttrs(attrs Attrs) {
	if len(node.Attrs) < 1 {
		node.Attrs = attrs
	} else {
		mergeMaps(node.Attrs, attrs)
	}
}

// Append more children to the end of a [DomNode].
func (node *DomNode) AssignChildren(children Children) {
	if len(node.Children) < 1 {
		node.Children = children
	} else {
		node.Children = append(node.Children, children...)
	}
}

// Record a compilation error for a [DomNode].
func (node *DomNode) appendError(err error) {
	if node.errors == nil {
		node.errors = []error{}
	}
	node.errors = append(node.errors, err)
}

// This is the base constructor for building a new [DomNode].
// "tag" is the name of the HTML tag to use.
// "args" is a variadic array of arguments, each of which can be of
// several different types:
//   - [Attrs] defines node attributes (multiple instances are merged
//     together)
//   - [Children] appends children to the end of the node
//   - [Attr] sets a single attribute value
//   - [Node] appends a single child
//   - [ClassName] adds a single class
//   - [ClassNames] adds multiple classes at once
//   - `string` appends a [Text] child with the given content
//
// Passing multiple instances of each of the above is supported.
// Any other type is ignored and logs an error message when compiled.
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
			node.Attrs["class"] = string(ClassNames(node.Attrs["class"], value))
		case string:
			node.Children = append(node.Children, Text(value))
		default:
			node.appendError(fmt.Errorf("Invalid DomNode argument: %v", arg))
		}
	}
	return node
}

// Create an `a` DOM node. Arguments follow the semantics of [NewDomNode].
func A(args ...any) DomNode {
	return NewDomNode("a", args)
}

// Create an `a` DOM node with the given href value. Arguments follow the
// semantics of [NewDomNode].
func AHref(href string, args ...any) DomNode {
	node := NewDomNode("a", args)
	node.Attrs["href"] = href
	return node
}

// Create an `abbr` DOM node. Arguments follow the semantics of [NewDomNode].
func Abbr(args ...any) DomNode {
	return NewDomNode("abbr", args)
}

// Create an `address` DOM node. Arguments follow the semantics of [NewDomNode].
func Address(args ...any) DomNode {
	return NewDomNode("address", args)
}

// Create an `area` DOM node. Arguments follow the semantics of [NewDomNode].
func Area(args ...any) DomNode {
	return NewDomNode("area", args)
}

// Create an `article` DOM node. Arguments follow the semantics of [NewDomNode].
func Article(args ...any) DomNode {
	return NewDomNode("article", args)
}

// Create an `aside` DOM node. Arguments follow the semantics of [NewDomNode].
func Aside(args ...any) DomNode {
	return NewDomNode("aside", args)
}

// Create an `audio` DOM node. Arguments follow the semantics of [NewDomNode].
func Audio(args ...any) DomNode {
	return NewDomNode("audio", args)
}

// Create a `b` DOM node. Arguments follow the semantics of [NewDomNode].
func B(args ...any) DomNode {
	return NewDomNode("b", args)
}

// Create a `base` DOM node. Arguments follow the semantics of [NewDomNode].
func Base(args ...any) DomNode {
	return NewDomNode("base", args)
}

// Create a `base` DOM node with the given href value.
func BaseHref(href string) DomNode {
	return DomNode{"base", Attrs{
		"href":   href,
		"target": "_blank",
	}, Children{}, nil}
}

// Create a `bdi` DOM node. Arguments follow the semantics of [NewDomNode].
func Bdi(args ...any) DomNode {
	return NewDomNode("bdi", args)
}

// Create a `bdo` DOM node. Arguments follow the semantics of [NewDomNode].
func Bdo(args ...any) DomNode {
	return NewDomNode("bdo", args)
}

// Create a `blockquote` DOM node. Arguments follow the semantics of [NewDomNode].
func Blockquote(args ...any) DomNode {
	return NewDomNode("blockquote", args)
}

// Create a `body` DOM node. Arguments follow the semantics of [NewDomNode].
func Body(args ...any) DomNode {
	return NewDomNode("body", args)
}

// Create a `br` DOM node. Arguments follow the semantics of [NewDomNode].
func Br(args ...any) DomNode {
	return NewDomNode("br", args)
}

// Create a `button` DOM node. Arguments follow the semantics of [NewDomNode].
func Button(args ...any) DomNode {
	return NewDomNode("button", args)
}

// Create a `canvas` DOM node. Arguments follow the semantics of [NewDomNode].
func Canvas(args ...any) DomNode {
	return NewDomNode("canvas", args)
}

// Create a `caption` DOM node. Arguments follow the semantics of [NewDomNode].
func Caption(args ...any) DomNode {
	return NewDomNode("caption", args)
}

// Create a `charset` DOM node. If value is the empty string then the default
// value of "UTF-8" is used.
func Charset(value string) DomNode {
	if len(value) < 1 {
		value = "UTF-8"
	}
	return DomNode{"meta", Attrs{"charset": value}, Children{}, nil}
}

// Create a `cite` DOM node. Arguments follow the semantics of [NewDomNode].
func Cite(args ...any) DomNode {
	return NewDomNode("cite", args)
}

// Create a `code` DOM node. Arguments follow the semantics of [NewDomNode].
func Code(args ...any) DomNode {
	return NewDomNode("code", args)
}

// Create a `col` DOM node. Arguments follow the semantics of [NewDomNode].
func Col(args ...any) DomNode {
	return NewDomNode("col", args)
}

// Create a `colgroup` DOM node. Arguments follow the semantics of [NewDomNode].
func Colgroup(args ...any) DomNode {
	return NewDomNode("colgroup", args)
}

// Create a `data` DOM node. Arguments follow the semantics of [NewDomNode].
func Data(args ...any) DomNode {
	return NewDomNode("data", args)
}

// Create a `datalist` DOM node. Arguments follow the semantics of [NewDomNode].
func Datalist(args ...any) DomNode {
	return NewDomNode("datalist", args)
}

// Create a `dd` DOM node. Arguments follow the semantics of [NewDomNode].
func Dd(args ...any) DomNode {
	return NewDomNode("dd", args)
}

// Create a `del` DOM node. Arguments follow the semantics of [NewDomNode].
func Del(args ...any) DomNode {
	return NewDomNode("del", args)
}

// Create a `details` DOM node. Arguments follow the semantics of [NewDomNode].
func Details(args ...any) DomNode {
	return NewDomNode("details", args)
}

// Create a `dfn` DOM node. Arguments follow the semantics of [NewDomNode].
func Dfn(args ...any) DomNode {
	return NewDomNode("dfn", args)
}

// Create a `dialog` DOM node. Arguments follow the semantics of [NewDomNode].
func Dialog(args ...any) DomNode {
	return NewDomNode("dialog", args)
}

// Create a `div` DOM node. Arguments follow the semantics of [NewDomNode].
func Div(args ...any) DomNode {
	return NewDomNode("div", args)
}

// Create a `dl` DOM node. Arguments follow the semantics of [NewDomNode].
func Dl(args ...any) DomNode {
	return NewDomNode("dl", args)
}

// Create a `dt` DOM node. Arguments follow the semantics of [NewDomNode].
func Dt(args ...any) DomNode {
	return NewDomNode("dt", args)
}

// Create a `em` DOM node. Arguments follow the semantics of [NewDomNode].
func Em(args ...any) DomNode {
	return NewDomNode("em", args)
}

// Create a `embed` DOM node. Arguments follow the semantics of [NewDomNode].
func Embed(args ...any) DomNode {
	return NewDomNode("embed", args)
}

// Create a `fieldset` DOM node. Arguments follow the semantics of [NewDomNode].
func Fieldset(args ...any) DomNode {
	return NewDomNode("fieldset", args)
}

// Create a `figcaption` DOM node. Arguments follow the semantics of [NewDomNode].
func Figcaption(args ...any) DomNode {
	return NewDomNode("figcaption", args)
}

// Create a `figure` DOM node. Arguments follow the semantics of [NewDomNode].
func Figure(args ...any) DomNode {
	return NewDomNode("figure", args)
}

// Create a `footer` DOM node. Arguments follow the semantics of [NewDomNode].
func Footer(args ...any) DomNode {
	return NewDomNode("footer", args)
}

// Create a `form` DOM node. Arguments follow the semantics of [NewDomNode].
func Form(args ...any) DomNode {
	return NewDomNode("form", args)
}

// Create a heading DOM node (ie; `h1`-`h6`). The given `level` should be an
// an integer between 1-6 inclusive. Arguments follow the semantics of
// [NewDomNode].
func H(level int, args ...any) DomNode {
	tag := fmt.Sprintf("h%d", level)
	return NewDomNode(tag, args)
}

// Create a `h1` DOM node. Arguments follow the semantics of [NewDomNode].
func H1(args ...any) DomNode {
	return NewDomNode("h1", args)
}

// Create a `h2` DOM node. Arguments follow the semantics of [NewDomNode].
func H2(args ...any) DomNode {
	return NewDomNode("h2", args)
}

// Create a `h3` DOM node. Arguments follow the semantics of [NewDomNode].
func H3(args ...any) DomNode {
	return NewDomNode("h3", args)
}

// Create a `h4` DOM node. Arguments follow the semantics of [NewDomNode].
func H4(args ...any) DomNode {
	return NewDomNode("h4", args)
}

// Create a `h5` DOM node. Arguments follow the semantics of [NewDomNode].
func H5(args ...any) DomNode {
	return NewDomNode("h5", args)
}

// Create a `h6` DOM node. Arguments follow the semantics of [NewDomNode].
func H6(args ...any) DomNode {
	return NewDomNode("h6", args)
}

// Create a `head` DOM node. Arguments follow the semantics of [NewDomNode].
func Head(args ...any) DomNode {
	return NewDomNode("head", args)
}

// Create a `header` DOM node. Arguments follow the semantics of [NewDomNode].
func Header(args ...any) DomNode {
	return NewDomNode("header", args)
}

// Create a `hr` DOM node. Arguments follow the semantics of [NewDomNode].
func Hr(args ...any) DomNode {
	return NewDomNode("hr", args)
}

// Create an `i` DOM node. Arguments follow the semantics of [NewDomNode].
func I(args ...any) DomNode {
	return NewDomNode("i", args)
}

// Create an `iframe` DOM node. Arguments follow the semantics of [NewDomNode].
func Iframe(args ...any) DomNode {
	return NewDomNode("iframe", args)
}

// Create an `img` DOM node. Arguments follow the semantics of [NewDomNode].
func Img(args ...any) DomNode {
	return NewDomNode("img", args)
}

// Create an `input` DOM node. Arguments follow the semantics of [NewDomNode].
func Input(args ...any) DomNode {
	return NewDomNode("input", args)
}

// Create an `ins` DOM node. Arguments follow the semantics of [NewDomNode].
func Ins(args ...any) DomNode {
	return NewDomNode("ins", args)
}

// Create a `kbd` DOM node. Arguments follow the semantics of [NewDomNode].
func Kbd(args ...any) DomNode {
	return NewDomNode("kbd", args)
}

// Create a `level` DOM node. Arguments follow the semantics of [NewDomNode].
func Label(args ...any) DomNode {
	return NewDomNode("label", args)
}

// Create a `legend` DOM node. Arguments follow the semantics of [NewDomNode].
func Legend(args ...any) DomNode {
	return NewDomNode("legend", args)
}

// Create a `li` DOM node. Arguments follow the semantics of [NewDomNode].
func Li(args ...any) DomNode {
	return NewDomNode("li", args)
}

// Create a `link` DOM node. Arguments follow the semantics of [NewDomNode].
func Link(args ...any) DomNode {
	return NewDomNode("link", args)
}

// Create a `link` DOM node with the given values for the `rel` and `href`
// attributes.
func LinkHref(rel string, href string) DomNode {
	attrs := Attrs{
		"rel":  rel,
		"href": href,
	}
	return DomNode{"link", attrs, Children{}, nil}
}

// Create a `main` DOM node. Arguments follow the semantics of [NewDomNode].
func Main(args ...any) DomNode {
	return NewDomNode("main", args)
}

// Create a `map` DOM node. Arguments follow the semantics of [NewDomNode].
func Map(args ...any) DomNode {
	return NewDomNode("map", args)
}

// Create a `mark` DOM node. Arguments follow the semantics of [NewDomNode].
func Mark(args ...any) DomNode {
	return NewDomNode("mark", args)
}

// Create a `meter` DOM node. Arguments follow the semantics of [NewDomNode].
func Meter(args ...any) DomNode {
	return NewDomNode("meter", args)
}

// Create a `nav` DOM node. Arguments follow the semantics of [NewDomNode].
func Nav(args ...any) DomNode {
	return NewDomNode("nav", args)
}

// Create a `noscript` DOM node. Arguments follow the semantics of [NewDomNode].
func Noscript(args ...any) DomNode {
	return NewDomNode("noscript", args)
}

// Create an `object` DOM node. Arguments follow the semantics of [NewDomNode].
func Object(args ...any) DomNode {
	return NewDomNode("object", args)
}

// Create an `ol` DOM node. Arguments follow the semantics of [NewDomNode].
func Ol(args ...any) DomNode {
	return NewDomNode("ol", args)
}

// Create an `optgroup` DOM node. Arguments follow the semantics of [NewDomNode].
func Optgroup(args ...any) DomNode {
	return NewDomNode("optgroup", args)
}

// Create an `option` DOM node. Arguments follow the semantics of [NewDomNode].
func Option(args ...any) DomNode {
	return NewDomNode("option", args)
}

// Create an `output` DOM node. Arguments follow the semantics of [NewDomNode].
func Output(args ...any) DomNode {
	return NewDomNode("output", args)
}

// Create a `p` DOM node. Arguments follow the semantics of [NewDomNode].
func P(args ...any) DomNode {
	return NewDomNode("p", args)
}

// Create a `param` DOM node. Arguments follow the semantics of [NewDomNode].
func Param(args ...any) DomNode {
	return NewDomNode("param", args)
}

// Create a `picture` DOM node. Arguments follow the semantics of [NewDomNode].
func Picture(args ...any) DomNode {
	return NewDomNode("picture", args)
}

// Create a `pre` DOM node. Arguments follow the semantics of [NewDomNode].
func Pre(args ...any) DomNode {
	return NewDomNode("pre", args)
}

// Create a `progress` DOM node. Arguments follow the semantics of [NewDomNode].
func Progress(args ...any) DomNode {
	return NewDomNode("progress", args)
}

// Create a `q` DOM node. Arguments follow the semantics of [NewDomNode].
func Q(args ...any) DomNode {
	return NewDomNode("q", args)
}

// Create an `rp` DOM node. Arguments follow the semantics of [NewDomNode].
func Rp(args ...any) DomNode {
	return NewDomNode("rp", args)
}

// Create an `rt` DOM node. Arguments follow the semantics of [NewDomNode].
func Rt(args ...any) DomNode {
	return NewDomNode("rt", args)
}

// Create a `ruby` DOM node. Arguments follow the semantics of [NewDomNode].
func Ruby(args ...any) DomNode {
	return NewDomNode("ruby", args)
}

// Create an `s` DOM node. Arguments follow the semantics of [NewDomNode].
func S(args ...any) DomNode {
	return NewDomNode("s", args)
}

// Create a `samp` DOM node. Arguments follow the semantics of [NewDomNode].
func Samp(args ...any) DomNode {
	return NewDomNode("samp", args)
}

// Create a `script` DOM node. Arguments follow the semantics of [NewDomNode].
func Script(args ...any) DomNode {
	return NewDomNode("script", args)
}

// Create a `script` DOM node with the given "src" URL.
func ScriptSrc(src string) DomNode {
	return DomNode{"script", Attrs{"src": src}, Children{}, nil}
}

// Create a `section` DOM node. Arguments follow the semantics of [NewDomNode].
func Section(args ...any) DomNode {
	return NewDomNode("section", args)
}

// Create a `select` DOM node. Arguments follow the semantics of [NewDomNode].
func Select(args ...any) DomNode {
	return NewDomNode("select", args)
}

// Create a `small` DOM node. Arguments follow the semantics of [NewDomNode].
func Small(args ...any) DomNode {
	return NewDomNode("small", args)
}

// Create a `source` DOM node. Arguments follow the semantics of [NewDomNode].
func Source(args ...any) DomNode {
	return NewDomNode("source", args)
}

// Create a `span` DOM node. Arguments follow the semantics of [NewDomNode].
func Span(args ...any) DomNode {
	return NewDomNode("span", args)
}

// Create a `strong` DOM node. Arguments follow the semantics of [NewDomNode].
func Strong(args ...any) DomNode {
	return NewDomNode("strong", args)
}

// Create a `style` DOM node. Arguments follow the semantics of [NewDomNode].
func Style(args ...any) DomNode {
	return NewDomNode("style", args)
}

// Create a `sub` DOM node. Arguments follow the semantics of [NewDomNode].
func Sub(args ...any) DomNode {
	return NewDomNode("sub", args)
}

// Create a `summary` DOM node. Arguments follow the semantics of [NewDomNode].
func Summary(args ...any) DomNode {
	return NewDomNode("summary", args)
}

// Create a `sup` DOM node. Arguments follow the semantics of [NewDomNode].
func Sup(args ...any) DomNode {
	return NewDomNode("sup", args)
}

// Create a `svg` DOM node. Arguments follow the semantics of [NewDomNode].
func Svg(args ...any) DomNode {
	return NewDomNode("svg", args)
}

// Create a `table` DOM node. Arguments follow the semantics of [NewDomNode].
func Table(args ...any) DomNode {
	return NewDomNode("table", args)
}

// Create a `tbody` DOM node. Arguments follow the semantics of [NewDomNode].
func Tbody(args ...any) DomNode {
	return NewDomNode("tbody", args)
}

// Create a `td` DOM node. Arguments follow the semantics of [NewDomNode].
func Td(args ...any) DomNode {
	return NewDomNode("td", args)
}

// Create a `template` DOM node. Arguments follow the semantics of [NewDomNode].
func Template(args ...any) DomNode {
	return NewDomNode("template", args)
}

// Create a `textarea` DOM node. Arguments follow the semantics of [NewDomNode].
func Textarea(args ...any) DomNode {
	return NewDomNode("textarea", args)
}

// Create a `tfoot` DOM node. Arguments follow the semantics of [NewDomNode].
func Tfoot(args ...any) DomNode {
	return NewDomNode("tfoot", args)
}

// Create a `th` DOM node. Arguments follow the semantics of [NewDomNode].
func Th(args ...any) DomNode {
	return NewDomNode("th", args)
}

// Create a `thead` DOM node. Arguments follow the semantics of [NewDomNode].
func Thead(args ...any) DomNode {
	return NewDomNode("thead", args)
}

// Create a `time` DOM node. Arguments follow the semantics of [NewDomNode].
func Time(args ...any) DomNode {
	return NewDomNode("time", args)
}

// Create a `title` DOM node. Arguments follow the semantics of [NewDomNode].
func Title(args ...any) DomNode {
	return NewDomNode("title", args)
}

// Create a `tr` DOM node. Arguments follow the semantics of [NewDomNode].
func Tr(args ...any) DomNode {
	return NewDomNode("tr", args)
}

// Create a `track` DOM node. Arguments follow the semantics of [NewDomNode].
func Track(args ...any) DomNode {
	return NewDomNode("track", args)
}

// Create a `u` DOM node. Arguments follow the semantics of [NewDomNode].
func U(args ...any) DomNode {
	return NewDomNode("u", args)
}

// Create a `ul` DOM node. Arguments follow the semantics of [NewDomNode].
func Ul(args ...any) DomNode {
	return NewDomNode("ul", args)
}

// Create a `var` DOM node. Arguments follow the semantics of [NewDomNode].
func Var(args ...any) DomNode {
	return NewDomNode("var", args)
}

// Create a `video` DOM node. Arguments follow the semantics of [NewDomNode].
func Video(args ...any) DomNode {
	return NewDomNode("video", args)
}

// Create a `wbr` DOM node. Arguments follow the semantics of [NewDomNode].
func Wbr(args ...any) DomNode {
	return NewDomNode("wbr", args)
}
