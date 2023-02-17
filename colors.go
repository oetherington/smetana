package smetana

import (
	"fmt"
	"strconv"
)

// A color value, suitable for use in CSS.
type Color interface {
	ToCssColor() string
}

// Structure representing an RGB color. All values are unsigned from 0-255
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
	return RGB{
		uint8((result >> 16) & 0xff),
		uint8((result >> 8) & 0xff),
		uint8(result & 0xff),
	}
}

// Structure representing an [RGB] color plus an alpha channel. All values are
// unsigned from 0-255 inclusive. Also see [Rgba].
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

// Structure representing an HSL color. "H" is an unsigned value between 0-360
// inclusive representing a position on the color wheel. 0 is red, 120 is
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

// Structure representing an HSL color plus as alpha channel. See [HSL] for
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
