package kumo

import "github.com/egoist/mygo/ui"

// ButtonGroupItem is one button of a ButtonGroup.
type ButtonGroupItem struct {
	// Label is the button's text, or its accessible name and tooltip when
	// it shows only Icon.
	Label string
	Icon  IconFunc
	// IconOnly hides the label, showing a square button of Icon.
	IconOnly bool
	Disabled bool
}

// ButtonGroup creates buttons joined side by side, as Kumo's ButtonGroup:
// their inner corners square and their rings shared. Every button takes
// the variant and size of o. It returns the index of the button clicked
// in this frame, or -1.
func ButtonGroup(c *ui.Context, items []ButtonGroupItem, o ButtonOptions) int {
	clicked := -1
	k := tok(c)
	m := buttonSize(k, o.Size)
	row := ui.Row(c).AlignItems(ui.Center).Role(ui.RoleGroup)
	shadowXS(row.Radius(m.radius))
	row.Children(func() {
		for i, it := range items {
			bo := o
			bo.Icon, bo.Disabled = it.Icon, it.Disabled || o.Disabled
			if it.IconOnly {
				bo.Shape = ShapeSquare
			} else {
				bo.Shape = ShapeBase
			}
			b := ui.Box(c.Key(i))
			b.Children(func() {
				btn := Button(c, it.Label, bo)
				// Square the corners between neighbours: top-left, top-right,
				// bottom-right and bottom-left, as CSS orders them.
				l, r := float32(0), float32(0)
				if i == 0 {
					l = m.radius
				}
				if i == len(items)-1 {
					r = m.radius
				}
				btn.Radius(l, r, r, l)
				if btn.Clicked() {
					clicked = i
				}
			})
			if i > 0 {
				b.MarginStart(-1)
			}
		}
	})
	return clicked
}
