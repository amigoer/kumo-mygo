package kumo

import "github.com/egoist/mygo/ui"

// RadioOptions configures Radio.
type RadioOptions struct {
	// Description shows below the label, in subtle text.
	Description string
	// Error rings the circle in the danger color.
	Error bool
	// Disabled disables the radio button.
	Disabled bool
	// Card draws the radio button as Kumo's choice card.
	Card bool
}

// Radio creates a radio button choosing value into *selected, with a
// label, as an item of Kumo's Radio.Group: a 16 DIP circle, filled with the
// contrast color around a dot of the base color when chosen. Build the
// buttons of one choice in a RadioGroup. Changed reports a click.
func Radio[T comparable](c *ui.Context, selected *T, value T, label string, opts ...RadioOptions) ui.Element {
	o := first(opts)
	k := tok(c)
	row := ui.RadioBase(c, selected, value).Gap(8).AlignItems(ui.Start).FocusRing(false).Cursor(ui.CursorPointer)
	if o.Disabled {
		row.Disabled(true)
	}
	if label != "" {
		row.Label(label)
	}
	if o.Description != "" {
		row.Description(o.Description)
	}
	on := *selected == value
	if o.Card {
		styleChoiceCard(row, k, on, o.Error)
	}
	focus := row.FocusVisible()
	row.Children(func() {
		dot := ui.Box(c).Size(16, 16).Radius(8).Shrink(0).PassThrough().Center()
		if label != "" {
			dot.Margin(2, 0, 0, 0)
		}
		line := k.Line
		if o.Error {
			line = k.Danger
		}
		if on {
			dot.Background(k.Contrast)
		} else {
			dot.Background(k.Base)
			if row.Hovered() && !o.Error {
				line = k.Hairline
			}
		}
		ring(dot, 1, line)
		if focus {
			ring(dot, 2, k.Brand)
		}
		if on {
			dot.Children(func() { ui.Box(c).Size(8, 8).Radius(4).Background(k.Base) })
		}
		if label != "" || o.Description != "" {
			ui.Column(c).Gap(2).Shrink(1).Children(func() {
				if label != "" {
					ui.Text(c, label).FontSize(k.TextBase).FontWeight(500).TextColor(k.TextDefault)
				}
				if o.Description != "" {
					ui.Text(c, o.Description).FontSize(k.TextSM).TextColor(k.TextSubtle).LineHeight(1.375)
				}
			})
		}
	})
	return dimmed(row)
}

// RadioGroup lays out the radio buttons fn builds under a legend, as Kumo's
// Radio.Group, and tells assistive technology they make one choice.
func RadioGroup(c *ui.Context, o CheckboxGroupOptions, fn func()) ui.Element {
	return choiceGroup(c, o, ui.RoleRadioGroup, fn)
}

// RadioSegmented creates Kumo's segmented radio group: the options side by
// side in a bordered row, the chosen one filled with the contrast color.
// It suits a choice among a few short labels.
func RadioSegmented[T comparable](c *ui.Context, selected *T, values []T, labels []string) ui.Element {
	k := tok(c)
	row := ui.Row(c).Role(ui.RoleRadioGroup).Radius(9).Clip().Background(k.Control)
	shadowXS(row)
	ring(row, 1, k.Line)
	row.Children(func() {
		for i, v := range values {
			name := ""
			if i < len(labels) {
				name = labels[i]
			}
			seg := ui.RadioBase(c.Key(i), selected, v).Height(34).PaddingX(12).AlignItems(ui.Center).
				FontSize(k.TextSM).FontWeight(500).FocusRing(false).Cursor(ui.CursorPointer).Label(name)
			on := *selected == v
			switch {
			case on:
				seg.Background(k.Contrast).TextColor(k.TextInverse)
			case seg.Hovered() && !seg.IsDisabled():
				seg.Background(k.Contrast.Alpha(0.07)).TextColor(k.TextDefault)
			default:
				seg.TextColor(k.TextDefault)
			}
			if i < len(values)-1 {
				seg.BorderWidth(0, 1, 0, 0).BorderColor(k.Line.Alpha(0.6))
			}
			if seg.FocusVisible() {
				seg.InsetShadow(0, 0, 0, 2, k.Brand)
			}
			seg.Children(func() { ui.Text(c, name).SingleLine() })
			dimmed(seg)
		}
	})
	return row
}
