package kumo

import "github.com/egoist/mygo/ui"

// Surface creates a card's surface, as Kumo's Surface and a LayerCard
// without sections: the base fill, rounded 8 DIPs, in a 1px ring of the
// line color with a faint shadow. It lays its children in a column; give
// it padding and children:
//
//	kumo.Surface(c).Padding(16).Gap(8).Children(func() { … })
func Surface(c *ui.Context) ui.Element {
	k := tok(c)
	s := ui.Column(c).Radius(8).Background(k.Base).Clip()
	ring(s, 1, k.Line)
	shadowXS(s)
	return s
}

// LayerCardOptions configures LayerCard.
type LayerCardOptions struct {
	// Title heads the card, in the elevated strip above its body.
	Title string
	// Icon is drawn before the title.
	Icon IconFunc
	// Actions builds controls at the end of the strip, as a Link or a
	// ghost Button of size SizeSM.
	Actions func()
}

// LayerCard creates Kumo's layered card: an elevated strip holding the
// title and actions over a base panel that body fills, each in its own
// ring, as LayerCard with its Secondary and Primary sections. It returns
// the body's panel, a column of children 8 DIPs apart padded 16.
func LayerCard(c *ui.Context, o LayerCardOptions, body func()) ui.Element {
	k := tok(c)
	root := ui.Column(c).FillWidth().Radius(8).Background(k.Elevated).Clip().FontSize(k.TextBase)
	ring(root, 1, k.Hairline)
	if o.Title != "" {
		root.Role(ui.RoleGroup).Label(o.Title)
	}
	var panel ui.Element
	root.Children(func() {
		if o.Title != "" || o.Actions != nil {
			ui.Row(c).Gap(8).AlignItems(ui.Center).Padding(8, 16).MinHeight(36).
				FontWeight(500).TextColor(k.TextSubtle).Children(func() {
				if o.Icon != nil {
					icon(c, o.Icon, 16)
				}
				ui.Text(c, o.Title).SingleLine().Grow(1)
				if o.Actions != nil {
					ui.Row(c).Gap(4).AlignItems(ui.Center).FontWeight(400).Children(o.Actions)
				}
			})
		}
		panel = ui.Column(c).Gap(8).Padding(16, 12, 16, 16).Radius(8).Background(k.Base).TextColor(k.TextDefault)
		ring(panel, 1, k.Fill)
		if body != nil {
			panel.Children(body)
		}
	})
	return panel
}
