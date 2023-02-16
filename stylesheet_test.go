package smetana

import (
	"fmt"
	"testing"
)

func TestCanRenderEmptyStyleSheet(t *testing.T) {
	styles := NewStyleSheet()
	assertEqual(t, "", RenderCss(styles))
}

func TestCanValidateFontExtensions(t *testing.T) {
	assertEqual(t, true, isValidFontExtension(".ttf"))
	assertEqual(t, true, isValidFontExtension(".woff"))
	assertEqual(t, true, isValidFontExtension(".woff2"))
	assertEqual(t, true, isValidFontExtension(".otf"))
	assertEqual(t, false, isValidFontExtension(".txt"))
	assertEqual(t, false, isValidFontExtension(".png"))
	assertEqual(t, false, isValidFontExtension(".html"))
	assertEqual(t, false, isValidFontExtension(".go"))
}

func TestCanAddFontFace(t *testing.T) {
	styles := NewStyleSheet()
	font := styles.AddFont("OpenSans", "OpenSans.ttf", "OpenSans.woff2")
	assertEqual(t, "OpenSans", font)
	css := RenderCss(styles)
	expected := "@font-face{font-family:OpenSans;src:url(OpenSans.ttf)format('ttf'),url(OpenSans.woff2)format('woff2');}"
	assertEqual(t, expected, css)
}

func TestCanAddClassWithStringProp(t *testing.T) {
	styles := NewStyleSheet()
	class := styles.AddClass(CssProps{
		"cursor": "pointer",
	})
	assertEqual(t, 8, len(class))
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
	class := styles.AddClass(CssProps{
		"padding": CustomProp{4},
	})
	assertEqual(t, 8, len(class))
	assertEqual(t, fmt.Sprintf(".%s{padding:8;}", class), RenderCss(styles))
}

func TestCanAddClassWithIntProp(t *testing.T) {
	styles := NewStyleSheet()
	class := styles.AddClass(CssProps{
		"margin": 10,
	})
	assertEqual(t, 8, len(class))
	assertEqual(t, fmt.Sprintf(".%s{margin:10px;}", class), RenderCss(styles))
}
