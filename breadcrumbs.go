package kumo

import (
	"github.com/amigoer/kumo-mygo/icons/phosphor"
	"github.com/egoist/mygo/ui"
)

// BreadcrumbsOptions configures Breadcrumbs.
type BreadcrumbsOptions struct {
	// Size is SizeBase (48 DIPs high) or SizeSM (40 DIPs).
	Size Size
	// Icons, when set, are drawn before the items of the same index.
	Icons []IconFunc
	// Label names the trail for assistive technology.
	Label string
}

// Breadcrumbs creates a path of items, as Kumo's Breadcrumbs: links in the
// subtle color separated by carets, ending in the current item. A click on
// a link, or Enter on it, sets *chosen to its index, which the returned
// row's Changed reports.
func Breadcrumbs(c *ui.Context, items []string, chosen *int, opts ...BreadcrumbsOptions) ui.Element {
	o := first(opts)
	k := tok(c)
	h, gap, font := float32(48), float32(4), k.TextBase
	if o.Size == SizeSM || o.Size == SizeXS {
		h, gap, font = 40, 2, k.TextSM
	}
	row := ui.Row(c).Height(h).Gap(gap).AlignItems(ui.Center).MinWidth(0).FontSize(font).FontWeight(500).
		Role(ui.RoleGroup).FocusGroup(ui.Horizontal)
	if o.Label != "" {
		row.Label(o.Label)
	}
	changed := false
	row.Children(func() {
		for i, name := range items {
			var ic IconFunc
			if i < len(o.Icons) {
				ic = o.Icons[i]
			}
			if i == len(items)-1 {
				ui.Row(c).Gap(4).AlignItems(ui.Center).MinWidth(0).TextColor(k.TextDefault).Children(func() {
					if ic != nil {
						icon(c, ic, 16)
					}
					ui.Text(c, name).SingleLine()
				})
				continue
			}
			l := ui.ButtonBase(c.Key(i)).Role(ui.RoleLink).Label(name).Gap(4).Shrink(0).Radius(4).
				FocusRing(false).Cursor(ui.CursorPointer)
			fg := k.TextSubtle
			if l.Hovered() {
				fg = k.TextDefault
			}
			l.TextColor(fg)
			if l.FocusVisible() {
				ring(l, 2, k.Brand)
			}
			l.Children(func() {
				if ic != nil {
					icon(c, ic, 16)
				}
				ui.Text(c, name).SingleLine()
			})
			if l.Clicked() {
				*chosen = i
				changed = true
			}
			icon(c, phosphor.CaretRight, 12).TextColor(k.TextInactive)
		}
	})
	if changed {
		row.SetChanged()
	}
	return row
}
