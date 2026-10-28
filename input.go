package kumo

import (
	"github.com/amigoer/kumo-mygo/icons/phosphor"
	"github.com/amigoer/kumo-mygo/theme"
	"github.com/egoist/mygo/ui"
)

// InputOptions configures Input, InputArea and SensitiveInput. The
// embedded FieldOptions give it a label, description and error, as Kumo's
// Input props; with none, the bare field is built.
type InputOptions struct {
	FieldOptions
	// Size is the height: 20, 26, 36 (the default) or 40 DIPs.
	Size Size
	// Placeholder shows while the input is empty.
	Placeholder string
	// Disabled and ReadOnly stop editing; a read-only input can still be
	// focused, selected and copied.
	Disabled, ReadOnly bool
	// Icon is drawn at the start of the field, as a search glass.
	Icon IconFunc
	// Prefix and Suffix are fixed text inside the field around what is
	// typed, as "https://" or ".workers.dev": Kumo's InputGroup addons.
	Prefix, Suffix string
	// Password hides what is typed.
	Password bool
	// AutoFocus gives the input the focus as it first shows.
	AutoFocus bool
	// Lines is the number of lines InputArea shows (3 when 0).
	Lines int
}

type inputMetrics struct{ h, px, gap, radius, font, icon float32 }

func inputSize(k *theme.Tokens, s Size) inputMetrics {
	m := buttonSize(k, s)
	return inputMetrics{h: m.h, px: m.px, gap: m.gap, radius: m.radius, font: m.font, icon: m.icon}
}

// Input creates a single-line text input editing *value, as Kumo's Input:
// the control fill with a 1px ring of the line color, which turns to the
// focus color as it takes the keyboard focus, or the danger color with an
// Error. It returns the text input itself, whose Changed reports an edit
// and Submitted the Enter key.
//
//	kumo.Input(c, &app.name, kumo.InputOptions{
//		FieldOptions: kumo.FieldOptions{Label: "Worker name", Description: "Lowercase letters and dashes."},
//		Placeholder:  "my-worker",
//	})
func Input(c *ui.Context, value *string, opts ...InputOptions) ui.Element {
	o := first(opts)
	var input ui.Element
	build := func() { input = inputBox(c, value, o, false) }
	if o.Label != "" || o.Description != "" || o.Error != "" {
		Field(c, o.FieldOptions, build)
	} else {
		build()
	}
	return input
}

// InputArea creates a multi-line text input editing *value, as Kumo's
// InputArea, Lines lines high.
func InputArea(c *ui.Context, value *string, opts ...InputOptions) ui.Element {
	o := first(opts)
	var input ui.Element
	build := func() { input = inputBox(c, value, o, true) }
	if o.Label != "" || o.Description != "" || o.Error != "" {
		Field(c, o.FieldOptions, build)
	} else {
		build()
	}
	return input
}

// SensitiveInput creates a password-like input whose value hides behind
// dots until its eye button shows it, as Kumo's SensitiveInput, for API
// tokens and secrets.
func SensitiveInput(c *ui.Context, value *string, opts ...InputOptions) ui.Element {
	o := first(opts)
	var input ui.Element
	build := func() {
		box := ui.Box(c)
		shown := ui.Local(box, "shown", func() bool { return false })
		box.Children(func() {
			oo := o
			oo.Password = !*shown
			input = inputBox(c, value, oo, false, func(m inputMetrics) {
				label, ic := Strings.Show, phosphor.Eye
				if *shown {
					label, ic = Strings.Hide, phosphor.EyeSlash
				}
				if Button(c, label, ButtonOptions{Variant: ButtonGhost, Size: SizeXS, Shape: ShapeSquare, Icon: ic}).Clicked() {
					*shown = !*shown
				}
			})
		})
	}
	if o.Label != "" || o.Description != "" || o.Error != "" {
		Field(c, o.FieldOptions, build)
	} else {
		build()
	}
	return input
}

// inputBox builds the field: a row with the control's fill and ring that
// holds the icon, prefix, the unstyled text input and suffix, and trailing
// controls that end builds. It returns the text input.
func inputBox(c *ui.Context, value *string, o InputOptions, multiline bool, end ...func(m inputMetrics)) ui.Element {
	k := tok(c)
	m := inputSize(k, o.Size)
	box := ui.Row(c).Radius(m.radius).Background(k.Control).TextColor(k.TextDefault).
		FontSize(m.font).Gap(m.gap).PaddingX(m.px).Cursor(ui.CursorText)
	if multiline {
		box.AlignItems(ui.Start).PaddingY(8)
	} else {
		box.Height(m.h).AlignItems(ui.Center)
	}
	var input ui.Element
	box.Children(func() {
		if o.Icon != nil {
			icon(c, o.Icon, m.icon).TextColor(k.TextSubtle)
		}
		if o.Prefix != "" {
			ui.Text(c, o.Prefix).SingleLine().TextColor(k.TextSubtle).Unselectable()
		}
		if multiline {
			lines := o.Lines
			if lines <= 0 {
				lines = 3
			}
			input = ui.TextAreaBase(c, value).Lines(lines, lines+6)
		} else {
			input = ui.TextInputBase(c, value)
			if o.Password {
				input.Password()
			}
		}
		input.Grow(1).MinWidth(0).FocusRing(false).ReadOnly(o.ReadOnly)
		if o.Disabled {
			input.Disabled(true)
		}
		if o.Placeholder != "" {
			input.Placeholder(o.Placeholder)
		}
		if o.Label != "" {
			input.Label(o.Label)
		}
		if o.Description != "" {
			input.Description(o.Description)
		}
		if o.Error != "" {
			input.Error(o.Error)
		}
		if o.AutoFocus {
			input.AutoFocus()
		}
		if o.Suffix != "" {
			ui.Text(c, o.Suffix).SingleLine().TextColor(k.TextSubtle).Unselectable()
		}
		for _, f := range end {
			f(m)
		}
	})
	// A click on the padding, the icon or the addons edits too.
	if box.Clicked() && !o.Disabled {
		input.Focus()
	}
	focused := input.Focused() && !o.ReadOnly
	switch {
	case o.Error != "":
		ring(box, 1, k.Danger)
		if focused {
			box.Shadow(0, 0, 0, 1.5, k.Danger.Alpha(0.5))
		}
	case focused:
		ring(box, 1.5, k.Focus.Alpha(0.5))
	default:
		ring(box, 1, k.Line)
	}
	if o.ReadOnly {
		box.Background(k.Recessed)
	}
	if o.Disabled {
		box.Opacity(0.5).Cursor(ui.CursorNotAllowed).TextColor(k.TextInactive)
	}
	return input
}
