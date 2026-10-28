package kumo

import (
	"github.com/amigoer/kumo-mygo/icons/phosphor"
	"github.com/egoist/mygo/ui"
)

// BannerVariant is a banner's kind, as Kumo's Banner variant prop.
type BannerVariant int

const (
	// BannerInfo, the default, is for general messages.
	BannerInfo BannerVariant = iota
	// BannerAlert is a warning.
	BannerAlert
	// BannerError is a critical issue.
	BannerError
	// BannerSecondary is a neutral message.
	BannerSecondary
)

// BannerOptions configures Banner.
type BannerOptions struct {
	Variant BannerVariant
	// Size is SizeBase or SizeSM, the compact banner for dialogs, which
	// keeps the title and the description on one line.
	Size Size
	// Title is the message in medium weight; Description follows it in
	// the small size.
	Title, Description string
	// Icon is drawn at the start, in the variant's color.
	Icon IconFunc
	// Action builds controls at the end, as a Button of size SizeSM.
	Action func()
	// OnDismiss, when set, adds a close button that calls it.
	OnDismiss func()
}

// Banner creates a full-width message bar, as Kumo's Banner: a tinted,
// rounded strip in the color of its variant.
//
//	kumo.Banner(c, kumo.BannerOptions{Variant: kumo.BannerAlert, Icon: phosphor.Warning,
//		Title: "Usage limit near", Description: "You've used 92% of this month's requests."})
func Banner(c *ui.Context, o BannerOptions) ui.Element {
	k := tok(c)
	compact := o.Size == SizeSM || o.Size == SizeXS
	var bg, fg, ic ui.Color
	switch o.Variant {
	case BannerAlert:
		bg, fg, ic = k.WarningTint, k.TextWarning, k.Warning
	case BannerError:
		bg, fg, ic = k.DangerTint, k.TextDanger, k.Danger
	case BannerSecondary:
		bg, fg, ic = k.Contrast.Alpha(0.05), k.TextDefault.Alpha(0.7), k.Interact
	default:
		bg, fg, ic = k.InfoTint, k.TextInfo, k.Info
	}
	row := ui.Row(c).Background(bg).TextColor(fg).Role(ui.RoleStatus)
	font := k.TextBase
	if compact {
		row.AlignItems(ui.Center).Gap(8).Radius(6).Padding(8, 12)
		font = k.TextSM
	} else {
		row.AlignItems(ui.Start).Gap(12).Radius(8).Padding(12, 16)
	}
	row.FontSize(font)
	if o.Title != "" {
		row.Label(o.Title)
	}
	row.Children(func() {
		if o.Icon != nil {
			h := font * 1.375
			ui.Box(c).Height(h).Center().Shrink(0).TextColor(ic).Children(func() { icon(c, o.Icon, 16) })
		}
		ui.Row(c).Grow(1).MinWidth(0).AlignItems(ui.Center).Gap(12).Children(func() {
			text := ui.Column(c).Grow(1).MinWidth(0).Gap(2)
			if compact {
				text.Row().Wrap().GapX(6).AlignItems(ui.Baseline)
			}
			text.Children(func() {
				if o.Title != "" {
					ui.Text(c, o.Title).FontWeight(500).LineHeight(1.375)
				}
				if o.Description != "" {
					ui.Text(c, o.Description).FontSize(k.TextSM).LineHeight(1.375)
				}
			})
			if o.Action != nil {
				ui.Row(c).Gap(8).Shrink(0).AlignItems(ui.Center).Children(o.Action)
			}
			if o.OnDismiss != nil {
				b := Button(c, Strings.Close, ButtonOptions{Variant: ButtonGhost, Size: SizeSM, Shape: ShapeSquare, Icon: phosphor.X})
				b.TextColor(fg)
				if b.Clicked() {
					o.OnDismiss()
				}
			}
		})
	})
	return row
}
