package kumo

import (
	"github.com/amigoer/kumo-mygo/theme"
	"github.com/egoist/mygo/ui"
)

// Size is a control's size, as Kumo's size prop: SizeBase, the zero value,
// is the default; SizeXS and SizeSM are compact, SizeLG roomy. Components
// that have fewer sizes take the nearest.
type Size int

const (
	SizeBase Size = iota
	SizeXS
	SizeSM
	SizeLG
)

// IconFunc is an icon accessor, as phosphor.Plus (passed without calling
// it). Options take accessors so that an icon is parsed only once drawn.
type IconFunc = func() *ui.SVG

// tok returns the Kumo tokens of the frame.
func tok(c *ui.Context) *theme.Tokens { return theme.Of(c) }

// dimmed gives a disabled control Kumo's disabled look. The bases own the
// disabled semantics: no clicks, no focus, and assistive technology told.
func dimmed(e ui.Element) ui.Element {
	if e.IsDisabled() {
		e.Opacity(0.5).Cursor(ui.CursorNotAllowed)
	}
	return e
}

// shadowXS is Tailwind's shadow-xs, which Kumo's controls cast:
// 0 1px 2px rgb(0 0 0 / 0.05).
func shadowXS(e ui.Element) ui.Element {
	return e.Shadow(0, 1, 2, 0, ui.RGBA(0, 0, 0, 0.05))
}

// shadowFloat is the shadow of Kumo's floating layers, popovers, menus and
// dialogs: a crisp edge and a soft drop (shadow-lg in the drop color).
func shadowFloat(e ui.Element, k *theme.Tokens) ui.Element {
	return e.Shadow(0, 10, 15, -3, k.ShadowDrop).Shadow(0, 4, 6, -4, k.ShadowDrop)
}

// ring draws Kumo's 1px ring, an outline that takes no room, as CSS's
// box-shadow rings do, outside the border box.
func ring(e ui.Element, width float32, c ui.Color) ui.Element {
	if c.A == 0 || width <= 0 {
		return e
	}
	return e.Shadow(0, 0, 0, width, c)
}

// focusRing draws Kumo's keyboard focus ring, 2px of the brand color, on an
// element that turned MyGo's own off with FocusRing(false).
func focusRing(e ui.Element, k *theme.Tokens, radius float32) ui.Element {
	if e.FocusVisible() {
		e.Shadow(0, 0, 0, 2, k.Brand)
	}
	return e
}

// icon builds an icon from an accessor, sized s, when there is one.
func icon(c *ui.Context, f IconFunc, s float32) ui.Element {
	if f == nil {
		return ui.Element{}
	}
	return ui.Icon(c, f()).Size(s, s).Shrink(0)
}

// first returns the first of opts, or the zero value.
func first[T any](opts []T) T {
	var o T
	if len(opts) > 0 {
		o = opts[0]
	}
	return o
}
