package kumo

import (
	"time"

	"github.com/egoist/mygo/ui"
)

// SidebarOptions configures Sidebar.
type SidebarOptions struct {
	// Width is the open sidebar's width, 260 DIPs when 0.
	Width float32
	// Collapsed, when set and true, narrows the sidebar to its items'
	// icons, 57 DIPs; SidebarTrigger toggles it.
	Collapsed *bool
	// Header builds the strip at the top, 58 DIPs high, as a logo and a
	// SidebarTrigger; Footer the strip at the bottom, 48 DIPs high.
	Header, Footer func()
	// Floating draws the sidebar as a rounded card with a shadow, away
	// from the window's edges, in place of a column with a rule.
	Floating bool
}

const sidebarSlide = 200 * time.Millisecond

// Sidebar creates an app's navigation column, as Kumo's Sidebar: the base
// fill with a rule on its end side, a header, the groups and items body
// builds, scrolling, and a footer. Put it first in a row with the page
// beside it Growing.
func Sidebar(c *ui.Context, o SidebarOptions, body func()) ui.Element {
	k := tok(c)
	w := o.Width
	if w <= 0 {
		w = 260
	}
	collapsed := o.Collapsed != nil && *o.Collapsed
	target := w
	if collapsed {
		target = 57
	}
	side := ui.Column(c).FillHeight().Shrink(0).Background(k.Base).TextColor(k.TextDefault).Clip()
	width := side.Animate("width", target, sidebarSlide)
	if c.Preferences().ReduceMotion {
		width = target
	}
	side.Width(width)
	if o.Floating {
		side.Margin(8).Radius(12).Border(1, k.Line)
		shadowFloat(side, k)
	} else {
		side.BorderWidth(0, 1, 0, 0).BorderColor(k.Line)
	}
	side.Children(func() {
		if o.Header != nil {
			ui.Row(c).Height(58).Shrink(0).Gap(4).PaddingX(12).AlignItems(ui.Center).
				BorderWidth(0, 0, 1, 0).BorderColor(k.Line).Children(o.Header)
		}
		ui.Scroll(c).Grow(1).MinHeight(0).FillWidth().Children(func() {
			ui.Column(c).Gap(16).Padding(12, 8).Children(body)
		})
		if o.Footer != nil {
			ui.Row(c).Height(48).Shrink(0).Gap(16).PaddingX(12).AlignItems(ui.Center).
				BorderWidth(1, 0, 0, 0).BorderColor(k.Line).Children(o.Footer)
		}
	})
	return side
}

// SidebarGroup lays out items under a label, as Kumo's Sidebar.Group with
// its GroupLabel: small medium-weight subtle text above a column of items.
// While the sidebar is collapsed, pass an empty label.
func SidebarGroup(c *ui.Context, label string, items func()) ui.Element {
	k := tok(c)
	g := ui.Column(c).Gap(2).Role(ui.RoleGroup)
	if label != "" {
		g.Label(label)
	}
	g.Children(func() {
		if label != "" {
			ui.Text(c, label).Margin(4, 0, 8, 0).PaddingX(12).FontSize(k.TextSM).FontWeight(500).
				TextColor(k.TextSubtle).SingleLine().Unselectable()
		}
		items()
	})
	return g
}

// SidebarItem creates an item of a sidebar, as Kumo's Sidebar.MenuButton:
// an icon and a label in a 34 DIP row, tinted while hovered and, with the
// strong text color, while active, the page shown. With collapsed set it
// shows only the icon, its label in a tooltip. Ask Clicked whether it was
// chosen.
func SidebarItem(c *ui.Context, label string, ic IconFunc, active, collapsed bool) ui.Element {
	k := tok(c)
	b := ui.ButtonBase(c).Label(label).MinHeight(34).FillWidth().Gap(8).Radius(8).
		Justify(ui.Start).AlignItems(ui.Center).FontSize(k.TextSM).FontWeight(500).
		FocusRing(false).Cursor(ui.CursorPointer)
	if collapsed {
		b.Justify(ui.Center).PaddingX(0)
	} else {
		b.PaddingX(12)
	}
	fg := k.TextSubtle
	switch {
	case active:
		b.Background(k.Tint)
		fg = k.TextStrong
	case b.Hovered():
		b.Background(k.Tint)
		fg = k.TextDefault
	}
	b.TextColor(fg)
	if b.FocusVisible() {
		b.InsetShadow(0, 0, 0, 2, k.Brand)
	}
	b.Children(func() {
		if ic != nil {
			icon(c, ic, 16)
		}
		if !collapsed {
			ui.Text(c, label).SingleLine()
		}
	})
	if collapsed {
		Tooltip(c, b, label, TooltipOptions{Side: TooltipRight})
	}
	return b
}

// SidebarTrigger creates the button that collapses and expands a sidebar,
// toggling *collapsed: a 34 DIP square in the subtle color.
func SidebarTrigger(c *ui.Context, collapsed *bool, ic IconFunc) ui.Element {
	k := tok(c)
	b := ui.ButtonBase(c).Size(34, 34).Radius(8).Center().FocusRing(false).Cursor(ui.CursorPointer).
		Label("Toggle sidebar").Expanded(!*collapsed)
	fg := k.TextSubtle
	if b.Hovered() {
		b.Background(k.Tint)
		fg = k.TextDefault
	}
	b.TextColor(fg)
	if b.FocusVisible() {
		b.InsetShadow(0, 0, 0, 2, k.Brand)
	}
	b.Children(func() { icon(c, ic, 18) })
	if b.Clicked() {
		*collapsed = !*collapsed
	}
	return b
}
