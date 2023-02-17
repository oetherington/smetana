package smetana

import (
	"fmt"
	"strconv"
)

// A color value, suitable for use in CSS.
type Color interface {
	ToCssColor() string
	ToHsla() HSLA
	ToRgba() RGBA
}

// Structure representing an [RGB] [Color]. All values are unsigned from 0-255
// inclusive. Also see [Rgb].
type RGB struct {
	R uint8
	G uint8
	B uint8
}

// Convert an [RGB] color to a CSS string.
func (c RGB) ToCssColor() string {
	return fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
}

// Create an [RGB] color.
func Rgb(r uint8, g uint8, b uint8) RGB {
	return RGB{r, g, b}
}

// Convert an [RGB] into an [HSLA].
func (c RGB) ToHsla() HSLA {
	return Rgba(c.R, c.G, c.B, 255).ToHsla()
}

// Convert an [RGB] into an [RGBA].
func (c RGB) ToRgba() RGBA {
	return Rgba(c.R, c.G, c.B, 255)
}

// Convert a 7-digit hex string to an unsigned integer. It is the _callers_
// responsibility to check the input is the correct length.
func longHexToInt(s string) uint32 {
	if s[0] != '#' {
		return 0
	}
	result, err := strconv.ParseInt(s[1:], 16, 32)
	if err != nil {
		return 0
	}
	return uint32(result)
}

// Convert a 4-digit hex string to an unsigned integer. It is the _callers_
// responsibility to check the input is the correct length.
func shortHexToInt(s string) uint32 {
	long := []byte{s[0], s[1], s[1], s[2], s[2], s[3], s[3]}
	return longHexToInt(string(long))
}

// Create an [RGB] color from a hex string (ie; "#FFFFFF").
func Hex(s string) RGB {
	var result uint32
	if len(s) == 7 {
		result = longHexToInt(s)
	} else if len(s) == 4 {
		result = shortHexToInt(s)
	} else {
		return RGB{0, 0, 0}
	}
	// TODO: Check endianness
	return RGB{
		uint8((result >> 16) & 0xff),
		uint8((result >> 8) & 0xff),
		uint8(result & 0xff),
	}
}

// Structure representing an [RGB] [Color] plus an alpha channel. All values
// are unsigned from 0-255 inclusive. Also see [Rgba].
type RGBA struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

// Convert an [RGBA] color to a CSS string.
func (c RGBA) ToCssColor() string {
	alpha := float32(c.A) / 255.0
	return fmt.Sprintf("rgba(%d, %d, %d, %.2f)", c.R, c.G, c.B, alpha)
}

// Create an [RGBA] color.
func Rgba(r uint8, g uint8, b uint8, a uint8) RGBA {
	return RGBA{r, g, b, a}
}

// Convert an [RGBA] into an [HSLA].
func (c RGBA) ToHsla() HSLA {
	r := float32(c.R) / 255.0
	g := float32(c.G) / 255.0
	b := float32(c.B) / 255.0
	a := float32(c.A) / 255.0

	l := max(max(r, g), b)
	s := l - min(min(r, g), b)

	var h float32
	if s > 0 {
		if l == r {
			h = (g - b) / s
		} else if l == g {
			h = 2 + (b-r)/s
		} else {
			h = 4 + (r-g)/s
		}
	}

	var H uint16
	if 60*h < 0 {
		H = uint16(60*h + 360)
	} else {
		H = uint16(60 * h)
	}

	var S float32
	if s > 0 {
		if l <= 0.5 {
			S = s / (2*l - s)
		} else {
			S = s / (2 - (2*l - s))
		}
	}

	L := l - s/2

	return HSLA{H, S, L, a}
}

// Convert an [RGBA] into an [RGBA].
func (c RGBA) ToRgba() RGBA {
	return c
}

// Structure representing an HSL [Color]. "H" is an unsigned value between
// 0-360 inclusive representing a position on the color wheel. 0 is red, 120 is
// green, 240 is blue, and other colors are interpolated between. S is
// saturation and must be a float between 0.0-1.0 inclusive. L is the
// lightness and must also be a float between 0.0-1.0 inclusive. Also see
// [Hsl].
type HSL struct {
	H uint16
	S float32
	L float32
}

// Convert an [HSL] color to a CSS string.
func (c HSL) ToCssColor() string {
	s := c.S * 100.0
	l := c.L * 100.0
	return fmt.Sprintf("hsl(%d, %.1f%%, %.1f%%)", c.H, s, l)
}

// Create an [HSL] color.
func Hsl(h uint16, s float32, l float32) HSL {
	return HSL{h, s, l}
}

// Convert an [HSL] into an [HSLA].
func (c HSL) ToHsla() HSLA {
	return HSLA{c.H, c.S, c.L, 1.0}
}

// Convert an [HSL] into an [RGBA].
func (c HSL) ToRgba() RGBA {
	return c.ToHsla().ToRgba()
}

// Structure representing an HSL [Color] plus as alpha channel. See [HSL] for
// more info. The alpha is stored as a float between 0.0-1.0 inclusive. Also
// see [Hsla].
type HSLA struct {
	H uint16
	S float32
	L float32
	A float32
}

// Convert an [HSLA] color to a CSS string.
func (c HSLA) ToCssColor() string {
	s := c.S * 100.0
	l := c.L * 100.0
	return fmt.Sprintf("hsla(%d, %.1f%%, %.1f%%, %.2f)", c.H, s, l, c.A)
}

// Create an [HSL] color.
func Hsla(h uint16, s float32, l float32, a float32) HSLA {
	return HSLA{h, s, l, a}
}

// Convert an [HSLA] into an [HSLA].
func (c HSLA) ToHsla() HSLA {
	return c
}

const oneThird = 1.0 / 3.0
const twoThirds = 2.0 / 3.0

func hueToRgb(v1 float32, v2 float32, vH float32) float32 {
	if vH < 0 {
		vH += 1
	} else if vH > 1 {
		vH -= 1
	}

	if (6 * vH) < 1 {
		return v1 + (v2-v1)*6*vH
	}

	if (2 * vH) < 1 {
		return v2
	}

	if (3 * vH) < 2 {
		return v1 + (v2-v1)*(twoThirds-vH)*6
	}

	return v1
}

// Convert an [HSLA] into an [RGBA].
func (c HSLA) ToRgba() RGBA {
	a := uint8(c.A * 255)

	if c.S <= 0 {
		value := uint8(c.L * 255)
		return RGBA{value, value, value, a}
	}

	hue := float32(c.H) / 360

	var v2 float32
	if c.L < 0.5 {
		v2 = c.L * (1 + c.S)
	} else {
		v2 = (c.L + c.S) - (c.L * c.S)
	}

	v1 := 2*c.L - v2

	return RGBA{
		uint8(255 * hueToRgb(v1, v2, hue+oneThird)),
		uint8(255 * hueToRgb(v1, v2, hue)),
		uint8(255 * hueToRgb(v1, v2, hue-oneThird)),
		a,
	}
}

// Darken a [Color] by the given amount, which should be a float32 between 0.0
// and 1.0, inclusive. Passing a value between 0.0 and -1.0 is equivalent to
// calling [Lighten] with a positive value.
func Darken(c Color, amount float32) HSLA {
	hsla := c.ToHsla()
	hsla.L = clamp(hsla.L-hsla.L*amount, 0.0, 1.0)
	return hsla
}

// Lighten a [Color] by the given amount, which should be a float32 between 0.0
// and 1.0, inclusive. Passing a value between 0.0 and -1.0 is equivalent to
// calling [Darken] with a positive value.
func Lighten(c Color, amount float32) HSLA {
	hsla := c.ToHsla()
	hsla.L = clamp(hsla.L+hsla.L*amount, 0.0, 1.0)
	return hsla
}
