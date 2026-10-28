// Package oklch converts MyGo colors to and from the OKLCH space, to mix
// them as CSS's color-mix(in oklch, …) does: Kumo derives the faces of its
// emphasis buttons that way from the brand and danger tokens.
package oklch

import (
	"math"

	"github.com/egoist/mygo/ui"
)

// LCH is a color in OKLCH: lightness 0 to 1, chroma about 0 to 0.4, hue in
// degrees, and alpha 0 to 1.
type LCH struct{ L, C, H, A float64 }

func lin(u uint8) float64 {
	c := float64(u) / 255
	if c <= 0.04045 {
		return c / 12.92
	}
	return math.Pow((c+0.055)/1.055, 2.4)
}

// From converts c's sRGB value to OKLCH.
func From(c ui.Color) LCH {
	r, g, b := lin(c.R), lin(c.G), lin(c.B)
	l := math.Cbrt(0.4122214708*r + 0.5363325363*g + 0.0514459929*b)
	m := math.Cbrt(0.2119034982*r + 0.6806995451*g + 0.1073969566*b)
	s := math.Cbrt(0.0883024619*r + 0.2817188376*g + 0.6299787005*b)
	L := 0.2104542553*l + 0.7936177850*m - 0.0040720468*s
	A := 1.9779984951*l - 2.4285922050*m + 0.4505937099*s
	B := 0.0259040371*l + 0.7827717662*m - 0.8086757660*s
	h := math.Atan2(B, A) * 180 / math.Pi
	if h < 0 {
		h += 360
	}
	return LCH{L: L, C: math.Hypot(A, B), H: h, A: float64(c.A) / 255}
}

// Color returns the color, through ui.Oklch so that one outside sRGB keeps
// its wide value where the screen shows it.
func (o LCH) Color() ui.Color {
	return ui.Oklch(float32(o.L), float32(max(o.C, 0)), float32(o.H)).Alpha(float32(o.A))
}

// Toward returns c mixed p of the way toward an achromatic color of
// lightness l, as color-mix(in oklch, c, white p) does for l = 1 and black
// for l = 0: the gray has no hue, so c's hue stays and its chroma fades.
func Toward(c ui.Color, l float64, p float64) ui.Color {
	o := From(c)
	o.L = o.L*(1-p) + l*p
	o.C *= 1 - p
	return o.Color()
}

// Lighten mixes c p of the way toward white in OKLCH.
func Lighten(c ui.Color, p float64) ui.Color { return Toward(c, 1, p) }

// Darken mixes c p of the way toward black in OKLCH.
func Darken(c ui.Color, p float64) ui.Color { return Toward(c, 0, p) }
