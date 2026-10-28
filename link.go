package kumo

import (
	"github.com/amigoer/kumo-mygo/icons/phosphor"
	"github.com/amigoer/kumo-mygo/theme"
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
	// A text element cannot hold the arrow, so the text and the arrow are
	// a row.
	if o.External && url != "" {
		r := ui.Row(c).Gap(3).AlignItems(ui.Center)
		r.Children(func() {
			styleLink(ui.Link(c, label, url), o, k)
			icon(c, phosphor.ArrowSquareOut, 14)
		})
		return r
	}
	// An empty url has no link element to name itself: the button takes the
	// label for assistive technology and draws it.
	custom := url == ""
	var l ui.Element
	if custom {
		l = ui.ButtonBase(c).Role(ui.RoleLink).Label(label)
	} else {
		l = ui.Link(c, label, url)
	}
	styleLink(l, o, k)
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

func styleLink(l ui.Element, o LinkOptions, k *theme.Tokens) {
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
}
