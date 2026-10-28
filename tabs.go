package kumo

import (
	"time"

	"github.com/egoist/mygo/ui"
)

// TabsVariant is the look of a tab list.
type TabsVariant int

const (
	// TabsSegmented, the default, raises the chosen tab on a base pill
	// over a recessed track.
	TabsSegmented TabsVariant = iota
	// TabsUnderline marks the chosen tab with a brand line under it.
	TabsUnderline
)

// TabsOptions configures Tabs.
type TabsOptions struct {
	Variant TabsVariant
	// Size is SizeBase (36 DIPs) or SizeSM (26 DIPs, as a small Input).
	Size Size
	// Label names the tab list for assistive technology.
	Label string
}

const tabsSlide = 200 * time.Millisecond

// Tabs creates a row of tabs showing labels, of which *selected is the
// index of the one chosen, as Kumo's Tabs: the indicator slides to the tab
// chosen. A click chooses a tab, as do the arrows, Home and End; the
// List's Changed reports a new choice. Build the chosen page below it from
// *selected.
func Tabs(c *ui.Context, selected *int, labels []string, opts ...TabsOptions) ui.TabsParts {
	o := first(opts)
	k := tok(c)
	sm := o.Size == SizeSM || o.Size == SizeXS
	tabs := ui.TabsBase(c, selected, len(labels))
	list := tabs.List.FontWeight(500).AlignSelf(ui.Start).MaxWidthPercent(100)
	if o.Label != "" {
		list.Label(o.Label)
	}
	font := k.TextBase
	if sm {
		font = k.TextXS
	}
	list.FontSize(font)
	segmented := o.Variant == TabsSegmented
	if segmented {
		h, r := float32(36), float32(8)
		if sm {
			h, r = 26, 6
		}
		list.Height(h).Radius(r).Background(k.Recessed).PaddingX(2).AlignItems(ui.Stretch)
		ring(list, 1, k.Hairline.Alpha(0.7))
	} else {
		h := float32(30)
		if sm {
			h = 26
		}
		list.Height(h+8).Gap(16).PaddingY(0).AlignItems(ui.Stretch).
			BorderWidth(0, 0, 1, 0).BorderColor(k.Hairline)
	}
	var active ui.Element
	list.Children(func() {
		for i, label := range labels {
			tab := tabs.Tab(i).AlignItems(ui.Center).FocusRing(false).Cursor(ui.CursorPointer)
			on := i == *selected
			hover := tab.Hovered() && !tab.IsDisabled()
			fg := k.TextSubtle
			if on || hover {
				fg = k.TextDefault
			}
			tab.TextColor(fg)
			if segmented {
				px, r := float32(10), float32(6)
				if sm {
					px, r = 8, 4
				}
				tab.PaddingX(px).MarginY(2).Radius(r)
			} else {
				tab.Padding(0, 8).MarginY(0).Radius(4).Margin(0, 0, 8, 0)
				if sm {
					tab.PaddingX(6)
				}
				if hover {
					tab.Background(k.Tint)
				}
				if !on {
					tab.FontWeight(400)
				}
			}
			if tab.FocusVisible() {
				tab.InsetShadow(0, 0, 0, 2, k.Brand)
			}
			if on {
				active = tab
			}
			tab.Children(func() { ui.Text(c, label).SingleLine() })
			dimmed(tab)
		}
	})
	if len(labels) == 0 || !active.Valid() {
		return tabs
	}
	reduce := c.Preferences().ReduceMotion
	slide := func(r ui.Rect) ui.Rect {
		b := active.LayoutBounds()
		x := list.Animate("ind-x", b.X-r.X, tabsSlide)
		w := list.Animate("ind-w", b.W, tabsSlide)
		if reduce {
			x, w = b.X-r.X, b.W
		}
		return ui.Rect{X: r.X + x, Y: b.Y, W: w, H: b.H}
	}
	if segmented {
		rad := float32(6)
		if sm {
			rad = 4
		}
		base, line, drop := k.Base, k.Line, k.ShadowDrop
		list.Draw(func(p *ui.Painter, r ui.Rect) {
			pill := slide(r)
			p.Shadow(pill, rad, 0, 1, 3, 0, drop)
			p.Shadow(pill, rad, 0, 0, 0, 1, line)
			p.Fill(pill, base, rad)
		})
	} else {
		brand := k.Brand
		list.DrawOver(func(p *ui.Painter, r ui.Rect) {
			s := slide(r)
			p.Fill(ui.Rect{X: s.X, Y: r.Y + r.H - 2, W: s.W, H: 2}, brand, 0)
		})
	}
	return tabs
}
