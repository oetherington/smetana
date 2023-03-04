package smetana

import (
	"reflect"
	"runtime/debug"
	"testing"
)

func assert[T any](t *testing.T, exp T, got T, equal bool) {
	if reflect.DeepEqual(exp, got) != equal {
		debug.PrintStack()
		t.Fatalf("Expecting '%v' got '%v'\n", exp, got)
	}
}

func assertEqual[T any](t *testing.T, exp T, got T) {
	assert(t, exp, got, true)
}

func assertNotEqual[T any](t *testing.T, exp T, got T) {
	assert(t, exp, got, false)
}

func assertOneOf[T any](t *testing.T, exp []T, got T) {
	for _, option := range exp {
		if reflect.DeepEqual(option, got) {
			return
		}
	}

	debug.PrintStack()
	t.Fatalf("Expecting one of '%v' got '%v'\n", exp, got)
}

func TestCanCreateSmetanaContextWithDefaults(t *testing.T) {
	smetana := NewSmetana()
	assertEqual(t, 0, len(smetana.Styles.Elements))
	assertEqual(t, 0, len(smetana.Palettes))
}

func TestCanCreateSmetanaContextWithPalettes(t *testing.T) {
	smetana := NewSmetanaWithPalettes(Palettes{
		"default": {
			"color": Hex("#FFFFFF"),
		},
	})
	assertEqual(t, 0, len(smetana.Styles.Elements))
	assertEqual(t, 1, len(smetana.Palettes))
	assertNotEqual(t, nil, smetana.Palettes["default"])
	assertNotEqual(t, nil, smetana.Palettes["default"]["color"])
}

func TestCanAddAPaletteToASmetanaContext(t *testing.T) {
	smetana := NewSmetana()
	smetana.AddPalette("default", Palette{
		"color": Hex("#FFFFFF"),
	})
	assertEqual(t, 1, len(smetana.Palettes))
	assertNotEqual(t, nil, smetana.Palettes["default"])
	assertNotEqual(t, nil, smetana.Palettes["default"]["color"])
}

func TestCanRenderStylesFromASmetanaContext(t *testing.T) {
	smetana := NewSmetanaWithPalettes(Palettes{
		"light": {
			"bg": Hex("#FFFFFF"),
		},
		"dark": {
			"bg": Hex("#000000"),
		},
	})
	smetana.Styles.AddBlock("body", CssProps{
		"background": PaletteValue("bg"),
	})
	css := smetana.RenderStyles()
	assertEqual(t, "body{background:#FFFFFF;}", css["light"])
	assertEqual(t, "body{background:#000000;}", css["dark"])
}
