package kumo

import (
	"github.com/amigoer/kumo-mygo/internal/oklch"
	"github.com/amigoer/kumo-mygo/theme"
	"github.com/egoist/mygo/ui"
)

// ButtonVariant is a button's visual weight, as Kumo's variant prop.
type ButtonVariant int

const (
	// ButtonSecondary, the default, is the bordered button for most
	// actions.
	ButtonSecondary ButtonVariant = iota
	// ButtonPrimary is the brand-filled main action of a view.
	ButtonPrimary
	// ButtonGhost has no background until hovered, for tertiary actions.
	ButtonGhost
	// ButtonDestructive is the danger-filled button for deleting.
	ButtonDestructive
	// ButtonSecondaryDestructive is a secondary button with danger text,
	// for a dangerous action that is not the main one.
	ButtonSecondaryDestructive
	// ButtonOutline is bordered over a transparent background.
	ButtonOutline
)

// ButtonShape is a button's shape.
type ButtonShape int

const (
	// ShapeBase is the rectangular button with a label.
	ShapeBase ButtonShape = iota
	// ShapeSquare and ShapeCircle are icon-only buttons; their label is
	// the accessible name and the tooltip.
	ShapeSquare
	ShapeCircle
)

// ButtonOptions configures Button. The zero value is a base-size
// secondary button.
type ButtonOptions struct {
	Variant ButtonVariant
	Size    Size
	Shape   ButtonShape
	// Icon is drawn before the label, as phosphor.Plus.
	Icon IconFunc
	// Loading shows a spinner in place of the icon and disables the button.
	Loading bool
	// Disabled disables the button.
	Disabled bool
	// Tooltip, when set, shows on hover; icon-only buttons show their label
	// when it is empty.
	Tooltip string
	// FullWidth stretches the button across its parent and centers its
	// content.
	FullWidth bool
}

// buttonMetrics are the height, horizontal padding, gap, radius, font size
// and icon size of a button size: Kumo's h-*, px-*, gap-*, rounded-* and
// text-* classes.
type buttonMetrics struct {
	h, px, gap, radius, font, icon float32
}

func buttonSize(k *theme.Tokens, s Size) buttonMetrics {
	switch s {
	case SizeXS:
		return buttonMetrics{h: 20, px: 6, gap: 4, radius: 4, font: k.TextXS, icon: 12}
	case SizeSM:
		return buttonMetrics{h: 26, px: 8, gap: 4, radius: 6, font: k.TextXS, icon: 14}
	case SizeLG:
		return buttonMetrics{h: 40, px: 16, gap: 8, radius: 8, font: k.TextBase, icon: 18}
	}
	return buttonMetrics{h: 36, px: 12, gap: 6, radius: 8, font: k.TextBase, icon: 16}
}

// Button creates a button showing label, styled as Kumo's Button. Ask
// Clicked whether it was clicked:
//
//	if kumo.Button(c, "Deploy", kumo.ButtonOptions{Variant: kumo.ButtonPrimary, Icon: phosphor.Rocket}).Clicked() {
//		app.deploy()
//	}
//
// The label is also the accessible name. An icon-only button (ShapeSquare
// or ShapeCircle) still takes a label, which names it and becomes its
// tooltip.
func Button(c *ui.Context, label string, opts ...ButtonOptions) ui.Element {
	o := first(opts)
	b := ui.ButtonBase(c)
	if o.Disabled || o.Loading {
		b.Disabled(true)
	}
	if label != "" {
		b.Label(label)
	}
	StyleButton(c, b, o)
	k := tok(c)
	m := buttonSize(k, o.Size)
	iconOnly := o.Shape != ShapeBase
	b.Children(func() {
		switch {
		case o.Loading:
			Loader(c, m.icon, buttonForeground(k, o.Variant))
		case o.Icon != nil:
			icon(c, o.Icon, m.icon)
		}
		if !iconOnly && label != "" {
			ui.Text(c, label).SingleLine()
		}
	})
	tip := o.Tooltip
	if tip == "" && iconOnly {
		tip = label
	}
	if tip != "" {
		Tooltip(c, b, tip)
	}
	return b
}

// StyleButton gives an element that takes clicks, a ui.ButtonBase or the
// trigger of another widget, the look of a Kumo button: its size, colors,
// ring and the states of the pointer. Button calls it; use it to give
// MyGo's parts, as a SelectBase's Trigger, a button's face.
func StyleButton(c *ui.Context, b ui.Element, o ButtonOptions) ui.Element {
	k := tok(c)
	m := buttonSize(k, o.Size)
	b.Height(m.h).Gap(m.gap).Radius(m.radius).FontSize(m.font).FontWeight(500).
		AlignItems(ui.Center).Shrink(0).FocusRing(false).Cursor(ui.CursorPointer)
	switch o.Shape {
	case ShapeSquare:
		b.Width(m.h).Justify(ui.Center)
	case ShapeCircle:
		b.Width(m.h).Justify(ui.Center).Radius(m.h / 2)
	default:
		b.PaddingX(m.px)
		if o.Icon != nil || o.Loading {
			// Icons carry whitespace in their view box: Kumo trims the
			// leading padding to match the trailing one optically.
			b.PaddingStart(m.px - 2)
		}
	}
	if o.FullWidth {
		b.FillWidth().Justify(ui.Center)
	}
	hover := b.Hovered() && !b.IsDisabled()
	pressed := b.Pressed() && !b.IsDisabled()
	switch o.Variant {
	case ButtonPrimary, ButtonDestructive:
		token := k.Brand
		if o.Variant == ButtonDestructive {
			token = k.Danger
		}
		// Kumo's emphasis face: a top-lit gradient from the token mixed 15%
		// toward white down to the token, a 1px inner highlight of it mixed
		// 30% toward white, and a ring of it mixed 10% toward black. Hover
		// brightens the top of the gradient.
		bg := oklch.Lighten(token, 0.30)
		from := oklch.Lighten(token, 0.15)
		if hover {
			from = bg
		}
		to := token
		if pressed {
			from, to = token, oklch.Darken(token, 0.08)
		}
		b.LinearGradient(ui.LinearGradient{From: from, To: to, Angle: 180, Oklab: true}).
			InsetShadow(0, 1, 0, 0, bg).TextColor(ui.Hex("#ffffff"))
		ring(b, 1, oklch.Darken(token, 0.10))
		shadowXS(b)
	case ButtonSecondary:
		bg := k.Base
		if hover || pressed {
			bg = k.Tint
		}
		b.Background(bg).TextColor(k.TextDefault)
		ring(b, 1, k.Line)
		shadowXS(b)
	case ButtonSecondaryDestructive:
		line := k.Line
		if hover {
			line = k.Danger.Alpha(0.3)
		}
		b.Background(k.Base).TextColor(k.TextDanger)
		ring(b, 1, line)
		shadowXS(b)
	case ButtonGhost:
		if hover || pressed {
			b.Background(k.Tint)
		}
		b.TextColor(k.TextDefault)
	case ButtonOutline:
		line, fg := k.Line, k.TextDefault
		if hover {
			line, fg = k.Focus.Alpha(0.25), k.TextStrong
		}
		b.TextColor(fg)
		ring(b, 1, line)
	}
	focusRing(b, k, m.radius)
	return dimmed(b)
}

// buttonForeground is the text color of a variant, which its icons and
// spinner take.
func buttonForeground(k *theme.Tokens, v ButtonVariant) ui.Color {
	switch v {
	case ButtonPrimary, ButtonDestructive:
		return ui.Hex("#ffffff")
	case ButtonSecondaryDestructive:
		return k.TextDanger
	}
	return k.TextDefault
}
