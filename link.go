package kumo

import (
	"github.com/amigoer/kumo-mygo/icons/phosphor"
	"github.com/egoist/mygo/ui"
)

// LinkVariant is a link's look, as Kumo's Link variant prop.
type LinkVariant int

const (
	// LinkInline, the default, is underlined text in the link color.
	LinkInline LinkVariant = iota
	// LinkCurrent is underlined in the color of the text around it.
	LinkCurrent
	// LinkPlain is the link color without an underline, fading on hover.
	LinkPlain
)

// LinkOptions configures Link.
type LinkOptions struct {
	Variant LinkVariant
	// External adds an arrow out after the text, for a link leaving the
	// app.
	External bool
}

// Link creates a text that opens url when clicked, styled as Kumo's Link.
// A url without a scheme goes to a page of the window's Router, as MyGo's
// Link does. With an empty url, it only reports Clicked.
func Link(c *ui.Context, label, url string, opts ...LinkOptions) ui.Element {
	o := first(opts)
	k := tok(c)
	// With children of its own, the link takes an empty label and names
	// itself for assistive technology.
	custom := url == "" || o.External
	var l ui.Element
	switch {
	case url == "":
		l = ui.ButtonBase(c).Role(ui.RoleLink).Label(label)
	case custom:
		l = ui.Link(c, "", url).Label(label)
	default:
		l = ui.Link(c, label, url)
	}
	hover := l.Hovered()
	switch o.Variant {
	case LinkCurrent:
		l.Underline().DecorationThickness(1)
		if hover {
			l.DecorationColor(k.TextDefault)
		} else {
			l.DecorationColor(k.TextDefault.Alpha(0.35))
		}
	case LinkPlain:
		col := k.TextLink
		if hover {
			col = col.Alpha(0.7)
		}
		l.TextColor(col)
	default:
		l.TextColor(k.TextLink).Underline().DecorationThickness(1)
		if hover {
			l.DecorationColor(k.TextLink)
		} else {
			l.DecorationColor(k.TextLink.Alpha(0.35))
		}
	}
	l.Cursor(ui.CursorPointer)
	if custom {
		l.Gap(3).AlignItems(ui.Center).Children(func() {
			ui.Text(c, label)
			if o.External {
				icon(c, phosphor.ArrowSquareOut, 14)
			}
		})
	}
	return l
}
