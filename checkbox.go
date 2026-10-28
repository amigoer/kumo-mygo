package kumo

import (
	"github.com/amigoer/kumo-mygo/theme"
	"github.com/egoist/mygo/ui"
)

// CheckboxOptions configures Checkbox.
type CheckboxOptions struct {
	// Indeterminate draws a dash, for a box over a group of which some are
	// checked; assistive technology hears "mixed". A click still toggles
	// *checked.
	Indeterminate bool
	// Error rings the box in the danger color.
	Error bool
	// Disabled disables the checkbox.
	Disabled bool
	// Description shows below the label, in subtle text.
	Description string
	// LabelFirst puts the label before the box.
	LabelFirst bool
	// Card draws the checkbox as Kumo's choice card: bordered and padded,
	// tinted while checked.
	Card bool
}

// Checkbox creates a check box toggling *checked, with a label, as Kumo's
// Checkbox: a 16 DIP square, filled with the contrast color and checked in
// the inverse text color. The label is the accessible name; a click on it
// toggles the box too. Changed reports a click.
func Checkbox(c *ui.Context, checked *bool, label string, opts ...CheckboxOptions) ui.Element {
	o := first(opts)
	k := tok(c)
	row := ui.CheckboxBase(c, checked).Gap(8).AlignItems(ui.Start).FocusRing(false).Cursor(ui.CursorPointer)
	if o.Disabled {
		row.Disabled(true)
	}
	if label != "" {
		row.Label(label)
	}
	if o.Indeterminate {
		row.Mixed()
	}
	if o.Description != "" {
		row.Description(o.Description)
	}
	if o.LabelFirst {
		row.Reverse().Justify(ui.End)
	}
	on := *checked || o.Indeterminate
	if o.Card {
		styleChoiceCard(row, k, *checked, o.Error)
	}
	focus := row.FocusVisible()
	row.Children(func() {
		box := ui.Box(c).Size(16, 16).Radius(4).Shrink(0).PassThrough()
		if label != "" {
			box.Margin(2, 0, 0, 0)
		}
		line := k.Hairline
		if o.Error {
			line = k.Danger
		}
		if on {
			box.Background(k.Contrast)
			line = k.Contrast
		} else {
			box.Background(k.Base)
		}
		ring(box, 1, line)
		if focus {
			ring(box, 2, k.Brand)
		}
		mixed := o.Indeterminate
		ink := k.TextInverse
		box.DrawOver(func(p *ui.Painter, r ui.Rect) {
			if !on {
				return
			}
			// Phosphor's bold check and minus at 12 DIPs, centered.
			var path ui.Path
			cx, cy := r.X+r.W/2, r.Y+r.H/2
			if mixed {
				path.MoveTo(cx-4, cy).LineTo(cx+4, cy)
			} else {
				path.MoveTo(cx-4.3, cy+0.2).LineTo(cx-1.4, cy+3).LineTo(cx+4.4, cy-2.9)
			}
			path.Cap(ui.RoundCap).Join(ui.RoundJoin)
			p.StrokePath(&path, 1.75, ink)
		})
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

// styleChoiceCard gives a checkbox or radio row Kumo's card appearance.
func styleChoiceCard(row ui.Element, k *theme.Tokens, on, invalid bool) {
	bg, border := k.Base, k.Hairline
	switch {
	case on:
		bg, border = k.Tint, k.Interact
	case row.Hovered() && !row.IsDisabled():
		bg = k.Elevated
	}
	if invalid {
		border = k.Danger
	}
	row.Gap(12).Padding(12).Radius(8).Background(bg).Border(1, border)
}

// CheckboxGroupOptions configures CheckboxGroup and RadioGroup.
type CheckboxGroupOptions struct {
	// Legend names the group, above it.
	Legend string
	// Description and Error show below the group, as Field shows them.
	Description string
	Error       string
	// Horizontal lays the items in a row that wraps.
	Horizontal bool
}

// CheckboxGroup lays out the checkboxes fn builds under a legend, as Kumo's
// Checkbox.Group: a column of items 8 DIPs apart, 16 below the legend.
func CheckboxGroup(c *ui.Context, o CheckboxGroupOptions, fn func()) ui.Element {
	return choiceGroup(c, o, ui.RoleGroup, fn)
}

func choiceGroup(c *ui.Context, o CheckboxGroupOptions, role ui.Role, fn func()) ui.Element {
	k := tok(c)
	col := ui.Column(c).Gap(16).Role(role)
	if o.Legend != "" {
		col.Label(o.Legend)
	}
	col.Children(func() {
		if o.Legend != "" {
			ui.Text(c, o.Legend).FontSize(k.TextBase).FontWeight(500).TextColor(k.TextDefault).Unselectable()
		}
		items := ui.Column(c).Gap(8)
		if o.Horizontal {
			items.Row().Wrap().Gap(16)
		}
		items.Children(fn)
		switch {
		case o.Error != "":
			ui.Text(c, o.Error).FontSize(k.TextSM).TextColor(k.TextDanger).Role(ui.RoleStatus)
		case o.Description != "":
			ui.Text(c, o.Description).FontSize(k.TextSM).TextColor(k.TextSubtle)
		}
	})
	return col
}
