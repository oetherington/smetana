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

import (
	"log"
	"os"
	"strings"
)

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

// Render a [Node] to an HTML string with the default settings. See
// [RenderHtmlOpts] for more fine-grained control.
func RenderHtml(node Node) string {
	return RenderHtmlOpts(node, false, nil)
}

// Render a [Node] to an HTML string specifying particular settings for the
// internal [Builder]. See the [Builder] struct for the available
// configuration values.
func RenderHtmlOpts(
	node Node,
	deterministicAttrs bool,
	logger *log.Logger,
) string {
	if logger == nil {
		logger = log.New(os.Stderr, "", 0)
	}
	builder := Builder{strings.Builder{}, deterministicAttrs, logger}
	node.ToHtml(&builder)
	return builder.Buf.String()
}

// Render a [StyleSheet] into a CSS string.
func RenderCss(styles StyleSheet) string {
	return RenderCssOpts(styles, nil)
}

// Render a [StyleSheet] into a CSS string specifying particular settings for
// the internal [Builder]. See the [Builder] struct for the available
// configuration values.
func RenderCssOpts(styles StyleSheet, logger *log.Logger) string {
	if logger == nil {
		logger = log.New(os.Stderr, "", 0)
	}
	builder := Builder{strings.Builder{}, false, logger}
	styles.Compile(&builder)
	return builder.Buf.String()
}

// Render a [Sitemap] into an XML string.
func RenderSitemap(sitemap Sitemap) string {
	return RenderSitemapOpts(sitemap, nil)
}

// Render a [Sitemap] into an XML string specifying particular settings for the
// internal [Builder]. See the [Builder] struct for the available
// configuration values.
func RenderSitemapOpts(sitemap Sitemap, logger *log.Logger) string {
	if logger == nil {
		logger = log.New(os.Stderr, "", 0)
	}
	builder := Builder{strings.Builder{}, false, logger}
	sitemap.ToXml(&builder)
	return builder.Buf.String()
}
