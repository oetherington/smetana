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
//	smetana := s.NewSmetanaWithPalettes(s.Palettes{
//		"light": {"color": s.Hex("#222")},
//		"dark": {"color": s.Hex("#ddd")},
//	})
//	myClass := smetana.Styles.AddAnonClass(s.CssProps{
//		"color": s.PaletteValue("color"),
//	})
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
//	htmlToServe := s.RenderHtml(page)
//	cssToServe := smetana.RenderStyles()

package smetana

import "log"

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

// A map from a palette name to a [Palette]
type Palettes map[string]Palette

// The [Smetana] struct is an overarching compilation context for tying
// together different parts of the application.
type Smetana struct {
	Palettes Palettes
	Styles   StyleSheet
}

// Create a new [Smetana] instance with default values.
func NewSmetana() Smetana {
	return Smetana{
		Palettes: Palettes{},
		Styles:   NewStyleSheet(),
	}
}

// Create a new [Smetana] instance with specific [Palettes].
func NewSmetanaWithPalettes(palettes Palettes) Smetana {
	return Smetana{
		Palettes: palettes,
		Styles:   NewStyleSheet(),
	}
}

// Add a new [Palette] to a [Smetana] context with the given name.
func (s *Smetana) AddPalette(name string, palette Palette) {
	s.Palettes[name] = palette
}

// Render the styles from the [Smetana] context into CSS strings. One CSS
// stylesheet will be created for each palette added with [AddPalette]. The
// return value is a map from palette names to rendered CSS strings.
// See [RenderStylesOpts] for more fine-grained control.
func (s Smetana) RenderStyles() map[string]string {
	return s.RenderStylesOpts(nil)
}

// Render the styles from the [Smetana] context into CSS strings. One CSS
// stylesheet will be created for each palette added with [AddPalette]. The
// return value is a map from palette names to rendered CSS strings.
// See [RenderStyles] for a simple interface with default values.
func (s Smetana) RenderStylesOpts(logger *log.Logger) map[string]string {
	result := map[string]string{}
	for name, palette := range s.Palettes {
		result[name] = RenderCssOpts(s.Styles, palette, logger)
	}
	return result
}
