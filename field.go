package kumo

import (
	"github.com/amigoer/kumo-mygo/icons/phosphor"
	"github.com/egoist/mygo/ui"
)

// FieldOptions are what Field shows around a control, as Kumo's Field
// props. Input, Select and the other form controls take the same fields.
type FieldOptions struct {
	// Label names the control, above it. It is also the control's
	// accessible name when the component builds the control.
	Label string
	// Description explains the control, below it, in subtle text.
	Description string
	// Error, when set, replaces the description with an error in the
	// danger color, and gives the control its error look.
	Error string
	// Optional shows Strings.Optional after the label, as Kumo's
	// required={false}.
	Optional bool
	// LabelTooltip adds an info button after the label showing it.
	LabelTooltip string
	// HideLabel keeps the label from showing; it still names the control.
	HideLabel bool
	// Horizontal puts the label beside the control, Kumo's controlFirst
	// layout turned around: the label first, the control at the end.
	Horizontal bool
}

// Field lays out a form control with its label, description and error, as
// Kumo's Field: a column 8 DIPs apart, the label in medium weight. control
// builds the control. Give a control you build yourself its accessible
// name with Label, as Field cannot reach it.
func Field(c *ui.Context, o FieldOptions, control func()) ui.Element {
	k := tok(c)
	col := ui.Column(c).Gap(8).Role(ui.RoleGroup)
	if o.Label != "" {
		col.Label(o.Label)
	}
	col.Children(func() {
		if o.Horizontal {
			ui.Row(c).Gap(12).AlignItems(ui.Center).Children(func() {
				if o.Label != "" && !o.HideLabel {
					Label(c, o.Label, LabelOptions{Optional: o.Optional, Tooltip: o.LabelTooltip}).Grow(1)
				}
				control()
			})
		} else {
			if o.Label != "" && !o.HideLabel {
				Label(c, o.Label, LabelOptions{Optional: o.Optional, Tooltip: o.LabelTooltip})
			}
			control()
		}
		switch {
		case o.Error != "":
			ui.Text(c, o.Error).FontSize(k.TextSM).TextColor(k.TextDanger).LineHeight(1.375).Role(ui.RoleStatus)
		case o.Description != "":
			ui.Text(c, o.Description).FontSize(k.TextSM).TextColor(k.TextSubtle).LineHeight(1.375)
		}
	})
	return col
}

// LabelOptions configures Label.
type LabelOptions struct {
	// Optional shows Strings.Optional after the text.
	Optional bool
	// Tooltip adds an info button after the text that shows it.
	Tooltip string
}

// Label creates a form label, as Kumo's Label: medium weight in the base
// size, with an optional marker and an info tooltip.
func Label(c *ui.Context, text string, opts ...LabelOptions) ui.Element {
	o := first(opts)
	k := tok(c)
	row := ui.Row(c).Gap(4).AlignItems(ui.Center)
	row.Children(func() {
		ui.Text(c, text).FontSize(k.TextBase).FontWeight(500).TextColor(k.TextDefault).Unselectable()
		if o.Optional {
			ui.Text(c, Strings.Optional).FontSize(k.TextBase).TextColor(k.TextSubtle).Unselectable()
		}
		if o.Tooltip != "" {
			b := Button(c, Strings.MoreInfo, ButtonOptions{Variant: ButtonGhost, Size: SizeXS, Shape: ShapeSquare, Icon: phosphor.Info, Tooltip: o.Tooltip})
			b.TextColor(k.TextSubtle)
		}
	})
	return row
}
