package smetana

import (
	"fmt"
	"path/filepath"
)

// A single CSS property. For instance,
//
//	CssProp{Key: "cursor", Value: "pointer"}
//
// The value may be any type supported by [CssValueToString].
//
// For properties that take size values it is recommended to use a unit helper
// rather than setting the value with a string. For example, instead of
//
//	CssProp{"padding", "4px"}
//
// you can use
//
//	CssProp{"padding", PX(4)}
type CssProp struct {
	Key   string
	Value any
}

// An array of CSS values of type [CssProp].
//
// Note that strict ordering is preserved, which is important in cases such as:
//
//	CssProps{
//		{"margin", "none"},
//		{"margin-top", PX(5)},
//	}
type CssProps []CssProp

// The name of a CSS class.
type ClassName string

// A palette for rendering a [Stylesheet] multiple times with different values.
// This can be used, for instance, to create separate styles for light-mode
// and dark-mode.
type Palette map[string]fmt.Stringer

// Use [PaletteValue] when creating a [Stylesheet] to mark a value as needing
// to be fetched from a [Palette].
//
//	styles := NewStyleSheet()
//	styles.AddClass("container", CssProps{
//		"background": PaletteValue("background-color"),
//	})
//	palette := Palette{"background-color", Hex("#f0f")}
//	css := RenderCss(styles, palette)
type PaletteValue string

// Convert a [PaletteValue] into a string.
func (value PaletteValue) String() string {
	return string(value)
}

// Interface representing an abstract element to be inserted into a CSS
// [StyleSheet].
type StyleSheetElement interface {
	ToCss(builder *Builder, palette Palette)
}

// Raw CSS type implementing [StyleSheetElement].
type StyleSheetCss string

// Create a [StyleSheetCss] [StyleSheetElement].
func StylesCss(css string) StyleSheetCss {
	return StyleSheetCss(css)
}

// Convert [StyleSheetCSS] into a CSS string.
func (css StyleSheetCss) ToCss(builder *Builder, palette Palette) {
	builder.Buf.WriteString(string(css))
}

// Instead of using a raw CSS string with [StyleSheetCss] you can instead use
// [StyleSheetPaletteCss] to provide a function that takes a palette which is
// used to build a CSS string.
type StyleSheetPaletteCss func(palette Palette) string

// Convert [StyleSheetPaletteCSS] into a CSS string.
func (css StyleSheetPaletteCss) ToCss(builder *Builder, palette Palette) {
	builder.Buf.WriteString(css(palette))
}

// @font-face type implementing [StyleSheetElement].
type StyleSheetFontFace struct {
	Family string
	Srcs   []string
}

// Create a [StyleSheetFontFace] [StyleSheetElement].
func StylesFontFace(family string, srcs ...string) StyleSheetFontFace {
	return StyleSheetFontFace{family, srcs}
}

func fontUrlToFormat(url string) (string, error) {
	ext := filepath.Ext(url)
	switch ext {
	case ".ttf":
		return "truetype", nil
	case ".otf":
		return "opentype", nil
	case ".woff":
		return "woff", nil
	case ".woff2":
		return "woff2", nil
	}
	return "", fmt.Errorf("Invalid font URL: %s", url)
}

// Convert a [StyleSheetFontFace] into a CSS string.
func (font StyleSheetFontFace) ToCss(builder *Builder, palette Palette) {
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
		format, err := fontUrlToFormat(src)
		if err == nil {
			builder.Buf.WriteString(format)
		} else {
			builder.Logger.Println(err)
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

// Create a [StyleSheetBlock] [StyleSheetElement].
func StylesBlock(selector string, props CssProps) StyleSheetBlock {
	return StyleSheetBlock{selector, props}
}

// Convert a [StyleSheetClass] into a CSS string.
func (block StyleSheetBlock) ToCss(builder *Builder, palette Palette) {
	builder.Buf.WriteString(block.Selector)
	builder.Buf.WriteByte('{')
	for _, prop := range block.Props {
		builder.Buf.WriteString(prop.Key)
		builder.Buf.WriteByte(':')
		WriteCssValue(builder, palette, prop.Value)
		builder.Buf.WriteByte(';')
	}
	builder.Buf.WriteByte('}')
}

// A helper that allows you to format text including values from a [Palette].
// This can be used fo cases such as:
//
//	CssProps{"border", PalettePrintf(
//		"1px solid %s",
//		PaletteValue("border-color"),
//	)}
type PalettePrintfData struct {
	Format string
	Args   []any
}

// Create a [PalettePrintfData].
func PalettePrintf(
	format string,
	args ...any,
) PalettePrintfData {
	return PalettePrintfData{format, args}
}

// Convert a [PalettePrintfData] into a CSS string.
func (data PalettePrintfData) Render(palette Palette) string {
	args := make([]any, len(data.Args))
	for i, value := range data.Args {
		switch item := value.(type) {
		case PaletteValue:
			args[i] = palette[string(item)]
		default:
			args[i] = item
		}
	}
	return fmt.Sprintf(data.Format, args...)
}

// Convert the given CSS value into a string using the given [Palette] if
// applicable.
//
// The value argument may be any of the following types: PaletteValue, string,
// fmt.Stringer (which includes all of the Smetana unit types), or an int
// (which will be interpreted as a quantity in pixels).
func CssValueToString(palette Palette, value any) (string, error) {
	switch item := value.(type) {
	case PaletteValue:
		insertion := palette[string(item)]
		if insertion == nil {
			return "inherit", fmt.Errorf("Missing palette value: %s", item)
		}
		return insertion.String(), nil
	case PalettePrintfData:
		return item.Render(palette), nil
	case string:
		return item, nil
	case fmt.Stringer:
		return item.String(), nil
	case int:
		return fmt.Sprintf("%dpx", item), nil
	default:
		return "inherit", fmt.Errorf("Invalid CSS value: %v", item)
	}
}

// Write the given value as a string to the [Builder], using the given
// [Palette] is applicable. This is a low-level function that should rarely
// be needed to be called directly by library consumers, but it's included in
// the public API for flexibility.
//
// The value argument may be any type supported by [CssValueToString].
func WriteCssValue(builder *Builder, palette Palette, value any) {
	str, err := CssValueToString(palette, value)
	if err != nil {
		builder.Logger.Println(err)
	}
	builder.Buf.WriteString(str)
}

// [StyleSheet] aggregates the CSS styles for a page and compiles them
// from the in-code representation into a CSS string for the browser. Note that
// [StyleSheet] is itself a [StyleSheetElement], so they can be nested.
type StyleSheet struct {
	Elements []StyleSheetElement
}

// Create a new empty [StyleSheet].
func NewStyleSheet(elements ...StyleSheetElement) StyleSheet {
	return StyleSheet{elements}
}

// Add a raw CSS string to the [StyleSheet].
func (styles *StyleSheet) AddCss(css StyleSheetCss) {
	styles.Elements = append(styles.Elements, css)
}

// Add [StyleSheetPaletteCss] generator function to the [StyleSheet].
func (styles *StyleSheet) AddPaletteCss(css StyleSheetPaletteCss) {
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
func (styles *StyleSheet) AddClass(name ClassName, props CssProps) ClassName {
	styles.Elements = append(styles.Elements, StyleSheetBlock{
		fmt.Sprintf(".%s", name),
		props,
	})
	return name
}

// Add a new class to a [StyleSheet] with a random name.
func (styles *StyleSheet) AddAnonClass(props CssProps) ClassName {
	name := ClassName(RandomString(8))
	return styles.AddClass(name, props)
}

// Add a new block to a [StyleSheet].
func (styles *StyleSheet) AddBlock(selector string, props CssProps) {
	styles.Elements = append(styles.Elements, StyleSheetBlock{
		selector,
		props,
	})
}

// Compile a [StyleSheet] into a CSS String.
func (styles StyleSheet) ToCss(builder *Builder, palette Palette) {
	for _, element := range styles.Elements {
		element.ToCss(builder, palette)
	}
}
