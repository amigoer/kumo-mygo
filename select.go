package kumo

import (
	"github.com/amigoer/kumo-mygo/icons/phosphor"
	"github.com/amigoer/kumo-mygo/theme"
	"github.com/egoist/mygo/ui"
)

// SelectOption is one choice of a Select.
type SelectOption[T comparable] struct {
	Value T
	// Label is what the option shows; the trigger shows the chosen one's.
	Label string
	// Description shows under the label in the popup, in subtle text.
	Description string
	// Icon is drawn before the label.
	Icon IconFunc
	// Group starts a new group of options under this heading, with a rule
	// above it unless it is the first.
	Group string
}

// SelectOptions configures Select.
type SelectOptions struct {
	FieldOptions
	// Size is the trigger's size, as Input's.
	Size Size
	// Placeholder shows while *selected matches no option
	// (Strings.SelectEmpty when empty).
	Placeholder string
	// Disabled disables the select.
	Disabled bool
}

// Select creates a drop-down choosing *selected among options, as Kumo's
// Select: a trigger with the look of an Input and an up-and-down caret,
// opening a popup of the base fill whose highlighted option is tinted and
// whose chosen one is checked. The arrows open it and move among the
// options; Enter chooses. It returns the trigger, whose Changed reports a
// new choice.
//
//	regions := []kumo.SelectOption[string]{{Value: "weur", Label: "Western Europe"}, {Value: "enam", Label: "Eastern North America"}}
//	kumo.Select(c, &app.region, regions, kumo.SelectOptions{FieldOptions: kumo.FieldOptions{Label: "Region"}})
func Select[T comparable](c *ui.Context, selected *T, options []SelectOption[T], opts ...SelectOptions) ui.Element {
	o := first(opts)
	var trigger ui.Element
	build := func() { trigger = selectBox(c, selected, options, o) }
	if o.Label != "" || o.Description != "" || o.Error != "" {
		Field(c, o.FieldOptions, build)
	} else {
		build()
	}
	return trigger
}

func selectBox[T comparable](c *ui.Context, selected *T, options []SelectOption[T], o SelectOptions) ui.Element {
	k := tok(c)
	m := inputSize(k, o.Size)
	sel := ui.SelectBase(c, selected)
	t := sel.Trigger
	if o.Disabled {
		t.Disabled(true)
	}
	if o.Label != "" {
		t.Label(o.Label)
	}
	if o.Error != "" {
		t.Error(o.Error)
	}
	t.Height(m.h).PaddingX(m.px).Gap(m.gap).Radius(m.radius).FontSize(m.font).
		Justify(ui.SpaceBetween).AlignItems(ui.Center).Background(k.Control).
		TextColor(k.TextDefault).FocusRing(false).Cursor(ui.CursorPointer).MinWidth(160)
	switch {
	case o.Error != "":
		ring(t, 1, k.Danger)
		if t.Focused() {
			t.Shadow(0, 0, 0, 1.5, k.Danger.Alpha(0.5))
		}
	case t.FocusVisible():
		ring(t, 2, k.Brand)
	case t.Focused() || sel.Open():
		ring(t, 1.5, k.Focus.Alpha(0.5))
	default:
		ring(t, 1, k.Line)
	}
	t.Expanded(sel.Open())
	var chosen *SelectOption[T]
	for i := range options {
		if options[i].Value == *selected {
			chosen = &options[i]
			break
		}
	}
	t.Children(func() {
		ui.Row(c).Gap(m.gap).AlignItems(ui.Center).MinWidth(0).Shrink(1).Children(func() {
			if chosen != nil {
				if chosen.Icon != nil {
					icon(c, chosen.Icon, m.icon)
				}
				ui.Text(c, chosen.Label).SingleLine()
			} else {
				ph := o.Placeholder
				if ph == "" {
					ph = Strings.SelectEmpty
				}
				ui.Text(c, ph).SingleLine().TextColor(k.TextPlaceholder)
			}
		})
		icon(c, phosphor.CaretUpDown, m.icon).TextColor(k.TextSubtle)
	})
	sel.Popup(func(panel ui.Element) {
		stylePopup(panel, k)
		panel.Margin(4, 0, 0, 0).PaddingY(6).MinWidth(160).MaxHeight(320)
		panel.Children(func() {
			ui.Scroll(c).FillWidth().Children(func() {
				for i, opt := range options {
					if opt.Group != "" {
						if i > 0 {
							ui.Box(c).Height(1).Margin(4, 0).Background(k.Hairline)
						}
						ui.Text(c, opt.Group).Padding(6, 14).FontSize(k.TextSM).FontWeight(600).TextColor(k.TextSubtle).Unselectable()
					}
					item := sel.Item(opt.Value).Row().Gap(8).AlignItems(ui.Center).Justify(ui.SpaceBetween).
						Margin(0, 6).Padding(6, 8).Radius(4).FontSize(k.TextBase).Cursor(ui.CursorPointer)
					if opt.Label != "" {
						item.Label(opt.Label)
					}
					if item.Highlighted() {
						item.Background(k.Tint)
					}
					on := opt.Value == *selected
					item.Children(func() {
						ui.Row(c).Gap(8).AlignItems(ui.Center).MinWidth(0).Children(func() {
							if opt.Icon != nil {
								icon(c, opt.Icon, 16)
							}
							ui.Column(c).MinWidth(0).Children(func() {
								ui.Text(c, opt.Label).SingleLine()
								if opt.Description != "" {
									ui.Text(c, opt.Description).FontSize(k.TextSM).TextColor(k.TextSubtle)
								}
							})
						})
						if on {
							icon(c, phosphor.Check, 16)
						} else {
							ui.Box(c).Size(16, 16)
						}
					})
				}
			})
		})
	})
	return dimmed(t)
}

// stylePopup gives a floating panel, a select's popup, a menu or a
// popover, Kumo's look: the base fill, rounded 8 DIPs in a ring of the
// line color, over a large soft shadow.
func stylePopup(panel ui.Element, k *theme.Tokens) {
	panel.Column().Radius(8).Background(k.Base).TextColor(k.TextDefault).Clip()
	ring(panel, 1, k.Line)
	shadowFloat(panel, k)
}
