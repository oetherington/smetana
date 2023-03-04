// Smetana is a library for programatically generating HTML and CSS.
// This can be pre-compiled for a static site, on-demand for dynamic
// pages, or somewhere in the middle with caching.
//
// The basic idea is that pages are constructed from a tree of structs
// that implement [Node]. These can be simple structures such as
// [DomNode] that correspond 1-to-1 with HTML tags (Smetana defines all
// standard HTML5 tag nodes for you) or [TextNode], or you can use
// these basic primitives to build more complex, reusable, React-style
// components.
//
// Typical usage looks something like this:
//
//	styles := NewStyleSheet()
//	myClass := styles.AddClass({"background": "red", "padding": PX(10)})
//	page := Html(
//		Head(
//		  Title("My HTML Document"),
//		),
//		Body(
//			myClass,
//			Div(
//				H(1, "Hello world"),
//				P("Foo bar"),
//			),
//		),
//	)
//	htmlToServe := RenderHtml(page)
//	cssToServe := RenderCss(styles)
package smetana

// All structural elements of an HTML document are implementers of
// the [Node] interface for converting to HTML. This is primarily
// different types of HTML tags and text.
type Node interface {
	ToHtml(b *Builder)
}

// Type alias for an HTML tag name (ie; "div", "span", etc.)
type Tag = string

// A single HTML attribute. For example,
//
//	{Key: "href", Value: "https://duckduckgo.com"}
type Attr struct {
	Key   string
	Value string
}

// A map of multiple HTML attributes.
type Attrs map[string]string

// Many types of [Node] have children to create a tree.
type Children []Node
