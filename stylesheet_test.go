package smetana

import (
	"errors"
	"fmt"
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
	block := StylesBlock("body", CssProps{"background": "red"})
	expected := StyleSheetBlock{"body", CssProps{"background": "red"}}
	assertEqual(t, expected, block)
}

func TestCanRenderEmptyStyleSheet(t *testing.T) {
	styles := NewStyleSheet()
	assertEqual(t, "", RenderCss(styles))
}

func TestCanCreateStyleSheetWithInitialElements(t *testing.T) {
	css := `
		body {
			padding: 3em;
			background: #ddd;
		}
	`
	styles := NewStyleSheet(StylesCss(css))
	assertEqual(t, css, RenderCss(styles))
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
	assertEqual(t, ".hello{background:red;}", RenderCss(styles))
}

func TestCanAddFontFace(t *testing.T) {
	styles := NewStyleSheet()
	font := styles.AddFont("OpenSans", "OpenSans.ttf", "OpenSans.woff2")
	assertEqual(t, "OpenSans", font)
	css := RenderCss(styles)
	expected := "@font-face{font-family:OpenSans;src:url(OpenSans.ttf)format('truetype'),url(OpenSans.woff2)format('woff2');}"
	assertEqual(t, expected, css)
}

func TestCanAddClassWithStringProp(t *testing.T) {
	styles := NewStyleSheet()
	class := styles.AddClass("container", CssProps{
		"cursor": "pointer",
	})
	assertEqual(t, "container", class)
	assertEqual(t, fmt.Sprintf(".%s{cursor:pointer;}", class), RenderCss(styles))
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
		"padding": CustomProp{4},
	})
	assertEqual(t, "container", class)
	assertEqual(t, fmt.Sprintf(".%s{padding:8;}", class), RenderCss(styles))
}

func TestCanAddClassWithIntProp(t *testing.T) {
	styles := NewStyleSheet()
	class := styles.AddClass("container", CssProps{
		"margin": 10,
	})
	assertEqual(t, "container", class)
	assertEqual(t, fmt.Sprintf(".%s{margin:10px;}", class), RenderCss(styles))
}

func TestCanAddClassWithColorProp(t *testing.T) {
	styles := NewStyleSheet()
	class := styles.AddClass("container", CssProps{
		"color": Rgb(255, 0, 0),
	})
	assertEqual(t, "container", class)
	assertEqual(t, fmt.Sprintf(".%s{color:#FF0000;}", class), RenderCss(styles))
}

func TestCanAddAnonClass(t *testing.T) {
	styles := NewStyleSheet()
	class := styles.AddAnonClass(CssProps{
		"cursor": "pointer",
	})
	assertEqual(t, 8, len(class))
	assertEqual(t, fmt.Sprintf(".%s{cursor:pointer;}", class), RenderCss(styles))
}

func TestCanAddBlock(t *testing.T) {
	styles := NewStyleSheet()
	styles.AddBlock("body", CssProps{
		"background": "red",
	})
	assertEqual(t, "body{background:red;}", RenderCss(styles))
}
