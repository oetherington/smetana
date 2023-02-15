package smetana

import "testing"

func TestRgbToCssColor(t *testing.T) {
	assertEqual(t, "#000000", Rgb(0, 0, 0).ToCssColor())
	assertEqual(t, "#FF0000", Rgb(255, 0, 0).ToCssColor())
	assertEqual(t, "#00FF00", Rgb(0, 255, 0).ToCssColor())
	assertEqual(t, "#0000FF", Rgb(0, 0, 255).ToCssColor())
	assertEqual(t, "#FFFFFF", Rgb(255, 255, 255).ToCssColor())
	assertEqual(t, "#828282", Rgb(130, 130, 130).ToCssColor())
}

func TestRgbaToCssColor(t *testing.T) {
	assertEqual(t, "rgba(0, 0, 0, 1.00)", Rgba(0, 0, 0, 255).ToCssColor())
	assertEqual(t, "rgba(255, 0, 0, 0.00)", Rgba(255, 0, 0, 0).ToCssColor())
	assertEqual(t, "rgba(0, 200, 0, 0.51)", Rgba(0, 200, 0, 130).ToCssColor())
}

func TestHslToCssColor(t *testing.T) {
	assertEqual(t, "hsl(0, 40.0%, 80.0%)", Hsl(0, 0.4, 0.8).ToCssColor())
	assertEqual(t, "hsl(120, 0.0%, 100.0%)", Hsl(120, 0.0, 1.0).ToCssColor())
	assertEqual(t, "hsl(240, 90.0%, 10.0%)", Hsl(240, 0.9, 0.1).ToCssColor())
}

func TestHslaToCssColor(t *testing.T) {
	assertEqual(t, "hsla(0, 40.0%, 80.0%, 0.00)", Hsla(0, 0.4, 0.8, 0.0).ToCssColor())
	assertEqual(t, "hsla(120, 0.0%, 100.0%, 1.00)", Hsla(120, 0.0, 1.0, 1.0).ToCssColor())
	assertEqual(t, "hsla(240, 90.0%, 10.0%, 0.60)", Hsla(240, 0.9, 0.1, 0.6).ToCssColor())
}
