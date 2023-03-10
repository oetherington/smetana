package smetana

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestCanCreateStylesCss(t *testing.T) {
	str := "body{padding:5px;}"
	css := StylesCss(str)
	assertEqual(t, str, string(css))
}

func TestCanCreateStylesFontFace(t *testing.T) {
	font := StylesFontFace("OpenSans", "OpenSans.ttf")
	expected := StyleSheetFontFace{"OpenSans", []string{"OpenSans.ttf"}}
	assertEqual(t, expected, font)
}

func TestCanCreateStylesBlock(t *testing.T) {
	block := StylesBlock("body", CssProps{{"background", "red"}})
	expected := StyleSheetBlock{"body", CssProps{{"background", "red"}}}
	assertEqual(t, expected, block)
}

func TestCanRenderEmptyStyleSheet(t *testing.T) {
	styles := NewStyleSheet()
	assertEqual(t, "", RenderCss(styles, Palette{}))
}

func TestCanCreateStyleSheetWithInitialElements(t *testing.T) {
	css := `
		body {
			padding: 3em;
			background: #ddd;
		}
	`
	styles := NewStyleSheet(StylesCss(css))
	assertEqual(t, css, RenderCss(styles, Palette{}))
}

func TestCanConvertFontNameToExtension(t *testing.T) {
	fmt, err := fontUrlToFormat("a.ttf")
	assertEqual(t, err, nil)
	assertEqual(t, fmt, "truetype")

	fmt, err = fontUrlToFormat("a.otf")
	assertEqual(t, err, nil)
	assertEqual(t, fmt, "opentype")

	fmt, err = fontUrlToFormat("a.woff")
	assertEqual(t, err, nil)
	assertEqual(t, fmt, "woff")

	fmt, err = fontUrlToFormat("a.woff2")
	assertEqual(t, err, nil)
	assertEqual(t, fmt, "woff2")

	fmt, err = fontUrlToFormat("a.png")
	assertEqual(t, fmt, "")
	assertEqual(t, err, errors.New("Invalid font URL: a.png"))
}

func TestCanAddARawCssString(t *testing.T) {
	styles := NewStyleSheet()
	styles.AddCss(".hello{background:red;}")
	assertEqual(t, ".hello{background:red;}", RenderCss(styles, Palette{}))
}

func TestCanAddFontFace(t *testing.T) {
	styles := NewStyleSheet()
	font := styles.AddFont("OpenSans", "OpenSans.ttf", "OpenSans.woff2")
	assertEqual(t, "OpenSans", font)
	css := RenderCss(styles, Palette{})
	expected := "@font-face{font-family:OpenSans;src:url(OpenSans.ttf)format('truetype'),url(OpenSans.woff2)format('woff2');}"
	assertEqual(t, expected, css)
}

func TestReportsErrorRenderingInvalidFontFormat(t *testing.T) {
	styles := NewStyleSheet()
	font := styles.AddFont("OpenSans", "OpenSans.png")
	assertEqual(t, "OpenSans", font)

	var buf strings.Builder
	logger := log.New(&buf, "", 0)
	css := RenderCssOpts(styles, Palette{}, logger)
	expected := "@font-face{font-family:OpenSans;src:url(OpenSans.png)format('');}"
	assertEqual(t, expected, css)
	assertEqual(t, "Invalid font URL: OpenSans.png\n", buf.String())
}

func TestCanAddClassWithStringProp(t *testing.T) {
	styles := NewStyleSheet()
	class := styles.AddClass("container", CssProps{
		{"cursor", "pointer"},
	})
	assertEqual(t, "container", class)
	css := RenderCss(styles, Palette{})
	assertEqual(t, fmt.Sprintf(".%s{cursor:pointer;}", class), css)
}

type CustomProp struct {
	Value int
}

func (prop CustomProp) String() string {
	return fmt.Sprintf("%d", prop.Value*2)
}

func TestCanAddClassWithCustomStringerProp(t *testing.T) {
	styles := NewStyleSheet()
	class := styles.AddClass("container", CssProps{
		{"padding", CustomProp{4}},
	})
	assertEqual(t, "container", class)
	css := RenderCss(styles, Palette{})
	assertEqual(t, fmt.Sprintf(".%s{padding:8;}", class), css)
}

func TestCanAddClassWithIntProp(t *testing.T) {
	styles := NewStyleSheet()
	class := styles.AddClass("container", CssProps{
		{"margin", 10},
	})
	assertEqual(t, "container", class)
	css := RenderCss(styles, Palette{})
	assertEqual(t, fmt.Sprintf(".%s{margin:10px;}", class), css)
}

func TestCanAddClassWithColorProp(t *testing.T) {
	styles := NewStyleSheet()
	class := styles.AddClass("container", CssProps{
		{"color", Rgb(255, 0, 0)},
	})
	assertEqual(t, "container", class)
	css := RenderCss(styles, Palette{})
	assertEqual(t, fmt.Sprintf(".%s{color:#FF0000;}", class), css)
}

func TestCanAddAnonClass(t *testing.T) {
	styles := NewStyleSheet()
	class := styles.AddAnonClass(CssProps{
		{"cursor", "pointer"},
	})
	assertEqual(t, 8, len(class))
	css := RenderCss(styles, Palette{})
	assertEqual(t, fmt.Sprintf(".%s{cursor:pointer;}", class), css)
}

func TestCanAddBlock(t *testing.T) {
	styles := NewStyleSheet()
	styles.AddBlock("body", CssProps{
		{"background", "red"},
	})
	css := RenderCss(styles, Palette{})
	assertEqual(t, "body{background:red;}", css)
}

func TestCanAddBlockWithPaletteValues(t *testing.T) {
	styles := NewStyleSheet()
	styles.AddBlock("body", CssProps{
		{"background", PaletteValue("background-color")},
	})
	palette := Palette{
		"background-color": Hex("#FF00FF"),
	}
	css := RenderCss(styles, palette)
	assertEqual(t, "body{background:#FF00FF;}", css)
}

func TestCanAddBlockWithFormattedPaletteValues(t *testing.T) {
	styles := NewStyleSheet()
	styles.AddBlock("div", CssProps{
		{"border", PalettePrintf(
			"%s solid %s",
			PX(2),
			PaletteValue("border-color"),
		)},
	})
	palette := Palette{
		"border-color": Hex("#FF00FF"),
	}
	css := RenderCss(styles, palette)
	assertEqual(t, "div{border:2px solid #FF00FF;}", css)
}

func TestCanAddBlockWithMissingPaletteValues(t *testing.T) {
	styles := NewStyleSheet()
	styles.AddBlock("body", CssProps{
		{"background", PaletteValue("background-color")},
	})
	var buf strings.Builder
	logger := log.New(&buf, "", 0)
	css := RenderCssOpts(styles, Palette{}, logger)
	assertEqual(t, "body{background:inherit;}", css)
	assertEqual(t, "Missing palette value: background-color\n", buf.String())
}

func TestCanAddBlockWithInvalidCssValue(t *testing.T) {
	styles := NewStyleSheet()
	styles.AddBlock("body", CssProps{
		{"background", NewStyleSheet()},
	})
	var buf strings.Builder
	logger := log.New(&buf, "", 0)
	css := RenderCssOpts(styles, Palette{}, logger)
	assertEqual(t, "body{background:inherit;}", css)
	assertEqual(t, "Invalid CSS value: {[]}\n", buf.String())
}

func TestCanAddPaletteCss(t *testing.T) {
	styles := NewStyleSheet()
	styles.AddPaletteCss(func(palette Palette) string {
		color, err := CssValueToString(
			palette,
			PaletteValue("background-color"),
		)
		assertEqual(t, nil, err)
		return fmt.Sprintf("body{background:%s;}", color)
	})
	palette := Palette{
		"background-color": Hex("#FF00FF"),
	}
	css := RenderCss(styles, palette)
	assertEqual(t, "body{background:#FF00FF;}", css)
}
