package kumo

import (
	"github.com/amigoer/kumo-mygo/icons/phosphor"
	"github.com/egoist/mygo/ui"
)

// Collapsible creates a disclosure, as Kumo's Collapsible: a medium-weight
// title with a caret that turns as it opens, showing what fn builds below
// it, indented behind a 2 DIP rule, while *open. The content grows into
// view. The trigger's Changed reports a click.
func Collapsible(c *ui.Context, title string, open *bool, fn func()) ui.Element {
	k := tok(c)
	col := ui.Column(c)
	col.Children(func() {
		parts := ui.CollapsibleBase(c, open)
		t := parts.Trigger
		t.Row().Gap(4).AlignItems(ui.Center).AlignSelf(ui.Start).FontSize(k.TextBase).FontWeight(500).
			TextColor(k.TextDefault).Cursor(ui.CursorPointer).FocusRing(false).Radius(4)
		focusRing(t, k, 4)
		turn := parts.Progress()
		t.Label(title).Children(func() {
			ui.Text(c, title).Unselectable()
			icon(c, phosphor.CaretDown, 12).Rotate(180 * turn)
		})
		parts.Panel(func() {
			ui.Column(c).Margin(8, 0).Gap(16).Padding(4, 4, 4, 16).
				BorderWidth(0, 0, 0, 2).BorderColor(k.Fill).Children(fn)
		})
	})
	return col
}
