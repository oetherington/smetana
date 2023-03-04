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

func TestRgbToString(t *testing.T) {
	assertEqual(t, "#000000", Rgb(0, 0, 0).String())
	assertEqual(t, "#FF0000", Rgb(255, 0, 0).String())
	assertEqual(t, "#00FF00", Rgb(0, 255, 0).String())
	assertEqual(t, "#0000FF", Rgb(0, 0, 255).String())
	assertEqual(t, "#FFFFFF", Rgb(255, 255, 255).String())
	assertEqual(t, "#828282", Rgb(130, 130, 130).String())
}

func TestRgbaToString(t *testing.T) {
	assertEqual(t, "rgba(0, 0, 0, 1.00)", Rgba(0, 0, 0, 255).String())
	assertEqual(t, "rgba(255, 0, 0, 0.00)", Rgba(255, 0, 0, 0).String())
	assertEqual(t, "rgba(0, 200, 0, 0.51)", Rgba(0, 200, 0, 130).String())
}

func TestHslToString(t *testing.T) {
	assertEqual(t, "hsl(0, 40.0%, 80.0%)", Hsl(0, 0.4, 0.8).String())
	assertEqual(t, "hsl(120, 0.0%, 100.0%)", Hsl(120, 0.0, 1.0).String())
	assertEqual(t, "hsl(240, 90.0%, 10.0%)", Hsl(240, 0.9, 0.1).String())
}

func TestHslaToString(t *testing.T) {
	assertEqual(t, "hsla(0, 40.0%, 80.0%, 0.00)", Hsla(0, 0.4, 0.8, 0.0).String())
	assertEqual(t, "hsla(120, 0.0%, 100.0%, 1.00)", Hsla(120, 0.0, 1.0, 1.0).String())
	assertEqual(t, "hsla(240, 90.0%, 10.0%, 0.60)", Hsla(240, 0.9, 0.1, 0.6).String())
}

func TestRgbToHsla(t *testing.T) {
	black := Rgb(0, 0, 0)
	assertEqual(t, Hsla(0, 0, 0, 1), black.ToHsla())
	white := Rgb(255, 255, 255)
	assertEqual(t, Hsla(0, 0, 1, 1), white.ToHsla())
	red := Rgb(255, 0, 0)
	assertEqual(t, Hsla(0, 1, 0.5, 1), red.ToHsla())
	green := Rgb(0, 255, 0)
	assertEqual(t, Hsla(120, 1, 0.5, 1), green.ToHsla())
	blue := Rgb(0, 0, 255)
	assertEqual(t, Hsla(240, 1, 0.5, 1), blue.ToHsla())
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
	alpha := Rgba(0, 0, 255, 127)
	assertEqual(t, Hsla(240, 1, 0.5, 0.49803922), alpha.ToHsla())
}

func TestHslToHsla(t *testing.T) {
	hsl := Hsl(120, 0.3, 0.6)
	assertEqual(t, Hsla(120, 0.3, 0.6, 1.0), hsl.ToHsla())
}

func TestHslaToHsla(t *testing.T) {
	hsla := Hsla(120, 0.3, 0.6, 0.8)
	assertEqual(t, hsla, hsla.ToHsla())
}

func TestRgbToRgba(t *testing.T) {
	rgb := Rgb(100, 40, 230)
	assertEqual(t, Rgba(100, 40, 230, 255), rgb.ToRgba())
}

func TestRgbaToRgba(t *testing.T) {
	rgba := Rgba(100, 40, 230, 120)
	assertEqual(t, Rgba(100, 40, 230, 120), rgba.ToRgba())
}

func TestHslToRgba(t *testing.T) {
	black := Hsl(0, 0, 0)
	assertEqual(t, Rgba(0, 0, 0, 255), black.ToRgba())
	white := Hsl(0, 0, 1)
	assertEqual(t, Rgba(255, 255, 255, 255), white.ToRgba())
	red := Hsl(0, 1, 0.5)
	assertEqual(t, Rgba(255, 0, 0, 255), red.ToRgba())
	green := Hsl(120, 1, 0.5)
	assertEqual(t, Rgba(0, 255, 0, 255), green.ToRgba())
	blue := Hsl(240, 1, 0.5)
	assertEqual(t, Rgba(0, 0, 255, 255), blue.ToRgba())
}

func TestHslaToRgba(t *testing.T) {
	black := Hsla(0, 0, 0, 1)
	assertEqual(t, Rgba(0, 0, 0, 255), black.ToRgba())
	white := Hsla(0, 0, 1, 1)
	assertEqual(t, Rgba(255, 255, 255, 255), white.ToRgba())
	red := Hsla(0, 1, 0.5, 1)
	assertEqual(t, Rgba(255, 0, 0, 255), red.ToRgba())
	green := Hsla(120, 1, 0.5, 1)
	assertEqual(t, Rgba(0, 255, 0, 255), green.ToRgba())
	blue := Hsla(240, 1, 0.5, 1)
	assertEqual(t, Rgba(0, 0, 255, 255), blue.ToRgba())
	alpha := Hsla(240, 1, 0.5, 0.49803922)
	assertEqual(t, Rgba(0, 0, 255, 127), alpha.ToRgba())
}

func TestDarken(t *testing.T) {
	value := Hsla(120, 0.5, 0.5, 1.0)
	darkened := Darken(value, 0.1)
	assertEqual(t, Hsla(120, 0.5, 0.45, 1.0), darkened)
}

func TestLighten(t *testing.T) {
	value := Hsla(120, 0.5, 0.5, 1.0)
	lightened := Lighten(value, 0.1)
	assertEqual(t, Hsla(120, 0.5, 0.55, 1.0), lightened)
}
