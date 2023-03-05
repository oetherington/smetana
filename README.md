# Smetana

[![https://pkg.go.dev/github.com/oetherington/smetana](https://pkg.go.dev/badge/github.com/oetherington/smetana.svg)](https://pkg.go.dev/github.com/oetherington/smetana)
[![https://github.com/oetherington/smetana/actions/workflows/ci.yml](https://github.com/oetherington/smetana/actions/workflows/ci.yml/badge.svg)](https://github.com/oetherington/smetana/actions/workflows/ci.yml)
[![Coverage Status](https://coveralls.io/repos/github/oetherington/smetana/badge.svg?branch=main)](https://coveralls.io/github/oetherington/smetana?branch=main)
[![https://goreportcard.com/report/github.com/oetherington/smetana](https://goreportcard.com/badge/github.com/oetherington/smetana)](https://goreportcard.com/report/github.com/oetherington/smetana)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

Smetana is an HTML and CSS generator for Go, designed for server-side webpage
rendering.

## Features
 - Simple component-like [API](#usage)
 - Built-in support for all HTML5 tags, and easily extensible for web
   components
 - Strongly-typed CSS [units](#adding-units) and [colors](#using-colors)
 - Robust with 100% test coverage
 - Zero dependencies outside the Go standard library

## Installation

```bash
$ go get github.com/oetherington/smetana
```

## Usage

Import the library with:

```go
import (
	s "github.com/oetherington/smetana"
)
```

Aliasing to `s` (or even `.`) is optional but advised.

### Typical usage

Here is an example of typical usage with separate styles for light and dark
mode, making use of the `Smetana` context:

```go
smetana := NewSmetanaWithPalettes(Palettes{
	"light": {
		"bg": Hex("#eee"),
		"fg": Hex("#222"),
	},
	"dark": {
		"bg": Hex("#363636"),
		"fg": Hex("#ddd"),
	},
})

font := smetana.Styles.AddFont("OpenSans", "/OpenSans.woff2")
container := smetana.Styles.AddAnonClass(CssProps{
	{"font-family", font},
	{"padding",     EM(2)},
	{"background",  PaletteValue("bg")},
	{"color",       PaletteValue("fg")},
})

css := smetana.RenderStyles()
lightCss = css["light"]
darkCss = css["dark"]

node := Html(
	Head(
		Title("My HTML Document"),
		LinkHref("stylesheet", "/styles/index.css"),
	),
	Body(
		container,
		H1("Hello world")
		P("foobar"),
	),
)

html := RenderHtml(node)
```

### Building HTML

To build an HTML tag we simply need to call the function with the name of that
tag. For instance, this:
```go
html := Html(
	Head(
		Title("My HTML Document"),
		Charset("UTF-8"),
		LinkHref("stylesheet", "/styles/index.css"),
	),
	Body(
		Div(
			ClassName("container"),
			H1("Hello world"),
		),
	),
)
```
can be rendered with `RenderHtml(html)` to produce the following HTML string:
```html
<!DOCTYPE html>
<html>
	<head>
		<title>My HTML Document</title>
		<meta charset="UTF-8">
		<link rel="stylesheet" href="/styles/index.css">
	</head>
	<body>
		<div class="container">
			<h1>Hello world</h1>
		</div>
	</body>
</html>
```
Note that the actual output will be automatically minified.

You may notice in the example above that we used `Charset` and `LinkHref` which
aren't the names of HTML tags. We could have instead directly used created a
`<meta>` DOM node (or the [MetaNode helper](#meta-tags)), but a number of
[special helpers](#special-case-helpers) are also included to avoid boilerplate
code for the most common use cases.

We also used `ClassName` to [apply a CSS class](#css-and-stylesheets).

By default, all of the DOM nodes take a variadic list of arguments which are
passed onto the `NewDomNode` function. This function accepts a variety of
different types of arguments to make generating your HTML as ergonomic as
possible. See the `NewDomNode` documentation for the full list.

#### Special-case helpers

Several frequently used tags have extra helper functions for their most common
use-cases:

 - `func AHref(href string, args ...any) DomNode` builds an `<a>` tag with
   the given URL.
 - `func BaseHref(href string) DomNode` builds a `<base>` tag with the given
   URL.
 - `func Charset(value string) DomNode` builds a charset `<meta>` node with the
   given value. If the empty string is passed in then it defaults to "UTF-8".
 - `func H(level int, args ...any) DomNode` builds a header tag from `<h1>` to
   `<h6>` with the given level.
 - `func LinkHref(rel string, href string)` builds a `<link>` tag with the
   given rel and URL attributes.
 - `func LinkStylesheet(href string)` builds a `<link>` tag with
   `rel="stylesheet"` and the given URL attribute.
 - `func LinkStylesheetMedia(href string, media string)` builds a `<link>` tag
   with `rel="stylesheet"` and the given URL and media attributes.
 - `func ScriptSrc(src string) DomNode` builds a `<script>` tag with the given
   src.

#### Meta tags

It is simple to create a `<meta>` tag using `NewDomNode`, but in most cases we
only want to set the "name" and "content" attributes, so there's a helper
function to do just that: `func Meta(name string, content string) MetaNode`.

There are several higher-level helpers that automatically set the "name"
property and so only take a single "value" string:
 - `Keywords`
 - `Description`
 - `Author`
 - `Viewport` (pass the empty string for the default value of
   "width=device-width, initial-scale=1.0")

as well as the `Charset` function mentioned above.

Smetana also supports natively "http-equiv" meta tags:
 - `func Equiv(equiv string, content string) EquivNode` builds a `<meta>` tag
    with "http-equiv" set to `equiv` and "content" set to `content`.
 - `func Refresh(value uint) DomNode` builds a `<meta>` tag with
   "http-equiv" set to "refresh"` and "content" set to the given integer value.
 - `func XUaCompatible(value string) EquivNode` builds a `<meta>` tag with
   "http-equiv" set to "x-ua-compatible" and "content" set to `value`.

#### Fragment nodes

Sometimes we want to combine multiple nodes at the same level of a document to
treat them as a single unit. In some cases it may be acceptable to wrap them in
another node such as a `div` or `span` but this is often undesirable as it
alters the generated markup.

In these cases, the children nodes can instead be wrapped in a `FragmentNode`
to treat them as a single entity but without adding an extra layer to the
generated markup.

```go
node := Fragment(
	Div(
		H(1, "Foo"),
		P("Bar"),
	),
	Span("Hello world"),
);
```

#### Transforming arrays

It's common to need to apply some operation to an array of data to turn it into
an array of DOM nodes. For this purpose, Smetana has a utility function called
`Xform` that functions similarly to `map` in Haskell or Javascript.

```go
titles := []string{"Foo", "Bar", "Baz"}
node := Div(
	Xform(titles, func (title string) Node {
		return H1(title)
	}),
)
```

#### Text nodes

Raw text inside of a tag is implemented by the `TextNode` struct. You should
rarely need to use `TextNode` explicitly as `Span("Hello world")` is
automatically converted into `Span(Text("Hello world"))`, but it's mentioned
here for completeness.

#### Creating custom nodes

To create your own `Node` types, you simply need to implement the interface:
```go
type Node interface {
	ToHtml(b *Builder)
}
```

`Builder` contains a `strings.Builder` called `Buf` which you can write your
HTML into. For instance:

```go
type CustomNode struct {
	Value string
}

func (node CustomNode) ToHtml(b *Builder) {
	b.Buf.WriteString(node.Value)
}
```

### CSS and StyleSheets

Smetana also supports generating CSS stylesheets along with your HTML.

Create a stylesheet with `styles := NewStyleSheet()`. This can later be
compiled into a CSS string with `RenderCss(styles, Palette{})`.

You can then add classes to the stylesheet with
```go
container := styles.AddClass("container", CssProps{
	{"cursor", "pointer"},
})
```

`container` is now a class name that can be passed directly into a DOM node:
```go
Div(
	container,
	"Hello world",
)
```
which will render to
```html
<div class="container">Hello world</div>
```
and
```css
.container{cursor:pointer}
```

If you don't require class names to be stable between builds then you can
generate a random class name with `addAnonClass`:
```go
container := styles.addAnonClass(CssProps{
	{"cursor", "pointer"},
})
```

`CssProps` is an array of items each of type `CssProp`, which is a struct
containing 2 fields: `Key` which is the name of the CSS property as a string,
and `Value` which can be any CSS value (see the documentation and source for
`WriteCssValue` for details). Note that the style shown in the documentation
without field names (ie; `{"cursor", "pointer"}` instead of
`{Key: "cursor", Value: "pointer"}`) will cause a lint error from `go vet`, but
is often still preferable when writing large amounts of styles. This error can
be silenced by instead using `go vet -composites=false`, but note that this is
a compromise and, if possible, should be limited to as little code as possible
rather than to your entire code base. Alternatively, you can use
`golangci-lint` with `// nolint` comments (see `examples/main.go`).

To use arbitrary CSS selectors you can instead use `AddBlock`:
```go
styles.AddBlock("body", CssProps{{"background", "red"}})
styles.AddBlock(".container > div", CssProps{{"display", "flex"}})
```

`NewStyleSheet` is also a variadic function which can take an arbitrary
number of `StyleSheetElement`s (the building blocks that make up a stylesheet).
This can be useful for cleanly adding global styles without using `AddBlock`:
```go
styles := NewStyleSheet(
	StylesCss(`
		body {
			padding: 3em;
		}
		p {
			font-family: sans-serif;
		}
	`),
	StylesBlock("div", CssProps{
		{"border-radius", PX(5)},
	}),
)
```

#### Using palettes

Stylesheets can be parameterized by using `Palette`s. This can be used, for
example, to use a single `StyleSheet` to generate separate CSS files for a
light mode and a dark mode. For instance:

```go
styles := NewStyleSheet(StylesBlock("body", CssProps{
	{"background", PaletteValue("bg")},
	{"color",      PaletteValue("fg")},
}))
darkStyles := RenderCss(styles, Palette{
	"bg": Hex("#000"),
	"fg": Hex("#fff"),
})
lightStyles := RenderCss(styles, Palette{
	"bg": Hex("#fff"),
	"fg": Hex("#000"),
})
```

The values in a `Palette` are not limited to `Color`s, but can actually be
any valid CSS value, such as `Unit`s, numbers, or strings.

#### Using colors

Instead of entering CSS color strings by hand, Smetana provides several helper
types and function to make color handling easier and more programmatic. For
instance, we can add an RGB background color property with:
```go
CssProps{{"background", Rgb(255, 255, 0)}}
```
which will compile to `background: #FFFF00` in CSS.

Aside from `RGB`, there are also helpers for `RGBA`, `HSL` and `HSLA`.

The `Hex` function will create an `RGB` color from a 4-digit or 7-digit CSS
hex color string, such as `#ab4` or `#FF00FF`.

For easier manipulation, all colors have `ToHsla()` and `ToRgba()` methods.

Colors can also be lightened or darkened by a certain amount with the `Lighten`
and `Darken` functions:
```go
red := Rgb(255, 0, 0)
darkerRed := Darken(red, 0.4)
lighterRed := Lighten(red, 0.4)
```

#### Adding units

Helpers are also provided to strongly type CSS units. For example,
```go
CssProps{{"margin", PX(10)}}
```
will compile to `margin: 10px`;

The following unit functions are provided: `PX`, `EM`, `REM`, `CM`, `MM`, `IN`,
`PT`, `PC`, `EX`, `CH`, `VW`, `VH`, `VMin`, `VMax` and `Perc` (for
percentages).

#### Applying classes to HTML

Class names can be passed directly to any DOM node by being typed as a
`ClassName` or `Classes` type. When passing multiple of these to the same node
they will be combined together with the `ClassNames` function:
```go
ClassNames("foo", "bar")
```
compiles to `"foo bar"`.

`ClassNames` takes a variadic list of arguments that may be strings (as above)
to combine together, or instances of the `Classes` type to conditionally
apply classnames:
```go
ClassNames("foo", "bar", Classes{"baz": false, "bop": true, "boz": 2 > 1})
```
compiles to `"foo bar bop boz"`;

#### Custom fonts

Smetana can also generate `@font-face` directives to load custom fonts like so:
```go
styles := NewStyleSheet()
font := styles.AddFont("OpenSans", "OpenSans.ttf", "OpenSans.woff2")
class := styles.AddClass(CssProps{
	{"font-family", font},
})
```

The `AddFont` function takes the name of the font family as the first argument
(which is also returned for convenience), followed by a variadic list of URLs
to the font sources. The type of each source is detected from its extension
which should be one of "ttf", "woff", "woff2" or "otf".

#### Inserting raw CSS

For more complex or obscure features it may be desirabe to add some manually
written CSS. This can be done with the `AddCss` function:
```go
styles := NewStyleSheet()
styles.AddCss("@media only screen and (max-width:600px) {body{width:100%;}}")
```

### Sitemaps

Smetana can also generate [XML sitemaps](https://www.sitemaps.org/protocol.html).

Simply construct an array of `SitemapLocation` structs using the provided
constructors then call `RenderSitemap` to get an XML string:
```go
sitemap := Sitemap{
	SitemapLocationUrl("https://duckduckgo.com"),
	SitemapLocationMod("https://lobste.rs", time.Now()),
	NewSitemapLocation(
		"https://news.ycombinator.com",
		time.Now(),
		ChangeFreqAlways,
		0.9,
	),
}
resultXml := RenderSitemap(sitemap)
```

The constructors are as follows:
 - `SitemapLocationUrl` takes only a URL.
 - `SitemapLocationMod` takes a URL and a last modified date.
 - `NewSitemapLocation` takes a URL, a last modified date, a change frequency
   and a priority.

The URL is a `string`, the modified date is a `time.Time`, the change frequency
is one of `ChangeFreqNone`, `ChangeFreqAlways`, `ChangeFreqHourly`,
`ChangeFreqDaily`, `ChangeFreqWeekly`, `ChangeFreqMonthly`, `ChangeFreqYearly`
or `ChangeFreqNever`, and the priority is a `float64` between 0 and 1 inclusive.

## License

Smetana is free software under the MIT license.

Copyright © 2023 Ollie Etherington.

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
