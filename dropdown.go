package kumo

import (
	"github.com/amigoer/kumo-mygo/icons/phosphor"
	"github.com/egoist/mygo/ui"
)

// MenuItem is an entry of a DropdownMenu. Set one of Label, Header or
// Separator.
type MenuItem struct {
	// Label is the item's text; OnSelect runs as it is chosen, by a
	// click, Enter or Space, and the menu closes.
	Label    string
	OnSelect func()
	// Icon is drawn before the label; items of a menu with icons line up.
	Icon IconFunc
	// Shortcut shows a key combination at the end, as "⌘K".
	Shortcut string
	// Danger colors a destructive item, as Delete.
	Danger bool
	// Disabled grays the item out and keeps it from being chosen.
	Disabled bool
	// Checked, when not nil, makes it a checkbox item toggling *Checked,
	// marked with a check while true; the menu stays open.
	Checked *bool
	// Header makes it a group heading in semibold, not chosen.
	Header bool
	// Separator makes it a hairline between groups.
	Separator bool
}

// DropdownMenu shows a menu of items beside anchor while *open, as Kumo's
// DropdownMenu: a panel of the control fill, rounded 8 DIPs in a ring of
// the line color, whose items tint as the pointer or the arrows reach
// them. The app opens it, as from anchor's Clicked; a press outside,
// Escape or choosing an item closes it. Call it every frame after building
// anchor. It returns the panel, invalid while closed.
//
//	b := kumo.Button(c, "Actions", kumo.ButtonOptions{Icon: phosphor.DotsThree})
//	if b.Clicked() { app.menu = !app.menu }
//	kumo.DropdownMenu(c, b, &app.menu, []kumo.MenuItem{
//		{Label: "Rename", Icon: phosphor.PencilSimple, OnSelect: app.rename},
//		{Separator: true},
//		{Label: "Delete", Icon: phosphor.Trash, Danger: true, OnSelect: app.delete},
//	})
func DropdownMenu(c *ui.Context, anchor ui.Element, open *bool, items []MenuItem, side ...PopoverSide) ui.Element {
	k := tok(c)
	anchor.Expanded(*open)
	// The arrows open the menu from a focused anchor, as Enter does.
	if !*open && (anchor.Shortcut(0, ui.KeyDown) || anchor.Shortcut(0, ui.KeyUp)) {
		*open = true
	}
	hasIcons := false
	for _, it := range items {
		if it.Icon != nil || it.Checked != nil {
			hasIcons = true
		}
	}
	var s PopoverSide
	if len(side) > 0 {
		s = side[0]
	}
	return ui.PopoverBase(c, anchor, open, func(panel ui.Element) {
		placePopup(panel, anchor, s, 4)
		stylePopup(panel, k)
		panel.Background(k.Control).MinWidth(144).MaxWidth(320).Padding(6).
			Role(ui.RoleMenu).FocusGroup(ui.Vertical).Transition(popTransition)
		panel.Children(func() {
			first := true
			for i, it := range items {
				switch {
				case it.Separator:
					ui.Box(c).Height(1).Margin(4, -4).Background(k.Hairline)
					continue
				case it.Header:
					h := ui.Text(c, it.Label).Padding(6, 8).FontSize(k.TextBase).FontWeight(600).Unselectable()
					if hasIcons {
						h.PaddingStart(32)
					}
					continue
				}
				item := ui.ButtonBase(c.Key(i)).Role(ui.RoleMenuItem).Label(it.Label).
					Justify(ui.Start).Gap(8).Padding(6, 8).Radius(6).FontSize(k.TextBase).
					FocusRing(false).Cursor(ui.CursorPointer)
				if it.Checked != nil {
					item.Role(ui.RoleMenuItemCheckBox).Checked(*it.Checked)
				}
				if it.Disabled {
					item.Disabled(true)
				}
				if first && !it.Disabled {
					item.AutoFocus()
					first = false
				}
				lit := (item.Hovered() || item.Focused()) && !it.Disabled
				fg := k.TextDefault
				if it.Danger {
					fg = k.TextDanger
				}
				item.TextColor(fg)
				if lit {
					if it.Danger {
						item.Background(k.Danger.Alpha(0.05))
					} else {
						item.Background(k.Tint)
					}
				}
				if item.FocusVisible() {
					item.InsetShadow(0, 0, 0, 2, k.Brand)
				}
				it := it
				item.Children(func() {
					switch {
					case it.Checked != nil:
						if *it.Checked {
							icon(c, phosphor.Check, 16)
						} else {
							ui.Box(c).Size(16, 16).Shrink(0)
						}
					case it.Icon != nil:
						icon(c, it.Icon, 16)
					case hasIcons:
						ui.Box(c).Size(16, 16).Shrink(0)
					}
					ui.Text(c, it.Label).SingleLine().Grow(1)
					if it.Shortcut != "" {
						ui.Text(c, it.Shortcut).FontSize(k.TextXS).LetterSpacing(1.2).Opacity(0.6).MarginStart(16)
					}
				})
				dimmed(item)
				if item.Clicked() {
					if it.Checked != nil {
						*it.Checked = !*it.Checked
					} else {
						*open = false
					}
					if it.OnSelect != nil {
						it.OnSelect()
					}
				}
			}
		})
	})
}
