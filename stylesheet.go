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

// [StyleSheet] aggregates the CSS styles for a page and compiles them
// from the in-code representation into a CSS string for the browser.
type StyleSheet struct {
	FontFaces map[string][]string
	Classes   map[ClassName]CssProps
}

// Create a new empty [StyleSheet].
func NewStyleSheet() StyleSheet {
	return StyleSheet{
		map[string][]string{},
		map[ClassName]CssProps{},
	}
}

// Add a new @font-face to the [StyleSheet]. `family` is the name to give
// to the CSS "font-family". `srcs` is an array of strings containing the
// URLs of the font files. The type of each src is automatically determined
// based on the file extension which should be one of "ttf", "woff", "woff2"
// or "otf".
func (styles *StyleSheet) AddFont(family string, srcs ...string) string {
	styles.FontFaces[family] = srcs
	return family
}

// Add a new class to a [StyleSheet].
func (styles *StyleSheet) AddClass(props CssProps) ClassName {
	name := ClassName(RandomString(8))
	styles.Classes[name] = props
	return name
}

// Compile a [StyleSheet] into a CSS String.
func (styles StyleSheet) ToCss(builder *Builder) {
	for family, srcs := range styles.FontFaces {
		builder.Buf.WriteString("@font-face{font-family:")
		builder.Buf.WriteString(family)
		builder.Buf.WriteString(";src:")
		for i, src := range srcs {
			if i > 0 {
				builder.Buf.WriteByte(',')
			}
			writeFontSrc(builder, src)
		}
		builder.Buf.WriteString(";}")
	}

	for name, props := range styles.Classes {
		builder.Buf.WriteByte('.')
		builder.Buf.WriteString(string(name))
		builder.Buf.WriteByte('{')
		writeProps(builder, props)
		builder.Buf.WriteByte('}')
	}
}

func isValidFontExtension(ext string) bool {
	return ext == ".ttf" || ext == ".woff" || ext == ".woff2" || ext == ".otf"
}

func writeFontSrc(builder *Builder, src string) {
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

func writeProps(builder *Builder, props CssProps) {
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
