package smetana

import "testing"

func TestCanConvertHexToRgb(t *testing.T) {
	assertEqual(t, RGB{0, 0, 0}, Hex("#000000"))
	assertEqual(t, RGB{255, 255, 255}, Hex("#FFFFFF"))
	assertEqual(t, RGB{255, 255, 255}, Hex("#ffffff"))
	assertEqual(t, RGB{255, 0, 0}, Hex("#ff0000"))
	assertEqual(t, RGB{0, 255, 0}, Hex("#00ff00"))
	assertEqual(t, RGB{0, 0, 255}, Hex("#0000ff"))
	assertEqual(t, RGB{136, 136, 136}, Hex("#888888"))
	assertEqual(t, RGB{0, 0, 0}, Hex("#000"))
	assertEqual(t, RGB{255, 255, 255}, Hex("#fff"))
	assertEqual(t, RGB{255, 0, 0}, Hex("#f00"))
	assertEqual(t, RGB{0, 255, 0}, Hex("#0f0"))
	assertEqual(t, RGB{0, 0, 255}, Hex("#00f"))
	assertEqual(t, RGB{136, 136, 136}, Hex("#888"))
	assertEqual(t, RGB{0, 0, 0}, Hex("invalid-color"))
	assertEqual(t, RGB{0, 0, 0}, Hex("#xxxxxx"))
	assertEqual(t, RGB{0, 0, 0}, Hex("#xxx"))
	assertEqual(t, RGB{0, 0, 0}, Hex("xxxxxxx"))
	assertEqual(t, RGB{0, 0, 0}, Hex("xxxx"))
}

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

func TestRgbToHsla(t *testing.T) {
	// TODO
}

func TestRgbaToHsla(t *testing.T) {
	black := Rgba(0, 0, 0, 255)
	assertEqual(t, Hsla(0, 0, 0, 1), black.ToHsla())
	white := Rgba(255, 255, 255, 255)
	assertEqual(t, Hsla(0, 0, 1, 1), white.ToHsla())
	red := Rgba(255, 0, 0, 255)
	assertEqual(t, Hsla(0, 1, 0.5, 1), red.ToHsla())
	green := Rgba(0, 255, 0, 255)
	assertEqual(t, Hsla(120, 1, 0.5, 1), green.ToHsla())
	blue := Rgba(0, 0, 255, 255)
	assertEqual(t, Hsla(240, 1, 0.5, 1), blue.ToHsla())
}

func TestHslToHsla(t *testing.T) {
	hsla := Hsl(120, 0.3, 0.6)
	assertEqual(t, Hsla(120, 0.3, 0.6, 1.0), hsla.ToHsla())
}

func TestHslaToHsla(t *testing.T) {
	hsla := Hsla(120, 0.3, 0.6, 0.8)
	assertEqual(t, hsla, hsla.ToHsla())
}
