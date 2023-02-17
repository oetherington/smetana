package smetana

import (
	"fmt"
	"path/filepath"
)

// A map from CSS property names to their values. For instance,
//
//	{"background": "red"}.
//
// For properties that take size values it is recommended to use a
// unit helper rather than setting the value with a string. For
// example, instead of
//
//	{"padding": "4px"}
//
// you can use
//
//	{"padding": PX(4)}
type CssProps map[string]any

// The name of a CSS class.
type ClassName string

// Interface representing an abstract element to be inserted into a CSS.
// [StyleSheet]
type StyleSheetElement interface {
	ToCss(builder *Builder)
}

// Raw CSS type implementing [StyleSheetElement].
type StyleSheetCss string

// Convert [StyleSheetCSS] into a CSS string.
func (css StyleSheetCss) ToCss(builder *Builder) {
	builder.Buf.WriteString(string(css))
}

// @font-face type implementing [StyleSheetElement].
type StyleSheetFontFace struct {
	Family string
	Srcs   []string
}

func isValidFontExtension(ext string) bool {
	return ext == ".ttf" || ext == ".woff" || ext == ".woff2" || ext == ".otf"
}

// Convert a [StyleSheetFontFace] into a CSS string.
func (font StyleSheetFontFace) ToCss(builder *Builder) {
	builder.Buf.WriteString("@font-face{font-family:")
	builder.Buf.WriteString(font.Family)
	builder.Buf.WriteString(";src:")
	for i, src := range font.Srcs {
		if i > 0 {
			builder.Buf.WriteByte(',')
		}
		builder.Buf.WriteString("url(")
		builder.Buf.WriteString(src)
		builder.Buf.WriteString(")format('")
		ext := filepath.Ext(src)
		if isValidFontExtension(ext) {
			builder.Buf.WriteString(ext[1:])
		} else {
			err := fmt.Errorf("Invalid extension for font '%s'", src)
			builder.Logger.Panicln(err)
		}
		builder.Buf.WriteString("')")
	}
	builder.Buf.WriteString(";}")
}

// CSS block type implementing [StyleSheetElement].
type StyleSheetBlock struct {
	Selector string
	Props    CssProps
}

// Convert a [StyleSheetClass] into a CSS string.
func (block StyleSheetBlock) ToCss(builder *Builder) {
	builder.Buf.WriteString(block.Selector)
	builder.Buf.WriteByte('{')
	writeClassProps(builder, block.Props)
	builder.Buf.WriteByte('}')
}

func writeClassProps(builder *Builder, props CssProps) {
	for key, value := range props {
		builder.Buf.WriteString(key)
		builder.Buf.WriteByte(':')

		switch item := value.(type) {
		case string:
			builder.Buf.WriteString(item)
		case fmt.Stringer:
			builder.Buf.WriteString(item.String())
		case Color:
			builder.Buf.WriteString(item.ToCssColor())
		case int:
			builder.Buf.WriteString(fmt.Sprintf("%dpx", item))
		default:
			err := fmt.Errorf("Invalid CSS value: %v", item)
			builder.Logger.Println(err)
		}

		builder.Buf.WriteByte(';')
	}
}

// [StyleSheet] aggregates the CSS styles for a page and compiles them
// from the in-code representation into a CSS string for the browser.
type StyleSheet struct {
	Elements []StyleSheetElement
}

// Create a new empty [StyleSheet].
func NewStyleSheet() StyleSheet {
	return StyleSheet{[]StyleSheetElement{}}
}

// Add a raw CSS string to the [StyleSheet]
func (styles *StyleSheet) AddCss(css StyleSheetCss) {
	styles.Elements = append(styles.Elements, css)
}

// Add a new @font-face to the [StyleSheet]. `family` is the name to give
// to the CSS "font-family". `srcs` is an array of strings containing the
// URLs of the font files. The type of each src is automatically determined
// based on the file extension which should be one of "ttf", "woff", "woff2"
// or "otf".
func (styles *StyleSheet) AddFont(family string, srcs ...string) string {
	styles.Elements = append(styles.Elements, StyleSheetFontFace{
		family,
		srcs,
	})
	return family
}

// Add a new class to a [StyleSheet].
func (styles *StyleSheet) AddClass(props CssProps) ClassName {
	name := ClassName(RandomString(8))
	styles.Elements = append(styles.Elements, StyleSheetBlock{
		fmt.Sprintf(".%s", name),
		props,
	})
	return name
}

// Add a new block to a [StyleSheet].
func (styles *StyleSheet) AddBlock(selector string, props CssProps) {
	styles.Elements = append(styles.Elements, StyleSheetBlock{
		selector,
		props,
	})
}

// Compile a [StyleSheet] into a CSS String.
func (styles StyleSheet) Compile(builder *Builder) {
	for _, element := range styles.Elements {
		element.ToCss(builder)
	}
}
