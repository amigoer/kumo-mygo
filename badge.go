package kumo

import "github.com/egoist/mygo/ui"

// BadgeVariant is a badge's color, as Kumo's Badge variant prop.
type BadgeVariant int

const (
	// BadgePrimary, the default, is inverted: dark in the light
	// appearance, light in the dark.
	BadgePrimary BadgeVariant = iota
	BadgeSecondary
	BadgeError
	BadgeWarning
	BadgeSuccess
	BadgeInfo
	BadgeDestructive
	// BadgeBeta is a dashed brand outline, for experimental features.
	BadgeBeta
	BadgeOutline
	BadgeRed
	BadgeGreen
	BadgeNeutral
	BadgeOrange
	BadgePurple
	BadgeTeal
	BadgeTealSubtle
	BadgeBlue
)

// BadgeOptions configures Badge.
type BadgeOptions struct {
	Variant BadgeVariant
	// Dot draws the status appearance: an outlined pill with a colored dot
	// for BadgeSuccess, BadgeWarning, BadgeError or BadgeNeutral.
	Dot bool
	// Icon is drawn before the text, 12 DIPs.
	Icon IconFunc
}

// Badge creates a small pill of text, as Kumo's Badge: 12 DIP medium text,
// 8 DIPs of padding across, fully rounded.
func Badge(c *ui.Context, text string, opts ...BadgeOptions) ui.Element {
	o := first(opts)
	k := tok(c)
	b := ui.Row(c).AlignItems(ui.Center).Gap(4).Padding(2, 8).Radius(999).Shrink(0).
		FontSize(k.TextXS).FontWeight(500).AlignSelf(ui.Start)
	if o.Icon != nil {
		b.PaddingStart(6)
	}
	var dot ui.Color
	white, black := ui.Hex("#ffffff"), ui.Hex("#000000")
	if o.Dot {
		b.Gap(6).TextColor(k.TextDefault)
		ring(b, 1, k.Hairline)
		switch o.Variant {
		case BadgeSuccess:
			dot = k.Success
		case BadgeWarning:
			dot = k.BadgeOrange
		case BadgeError:
			dot = k.BadgeRed
		case BadgeNeutral:
			dot = k.BadgeNeutral
		}
	} else {
		var bg, fg ui.Color
		switch o.Variant {
		case BadgePrimary:
			bg, fg = k.BadgeInverted, k.TextBadgeInverted
		case BadgeSecondary:
			bg, fg = k.Fill, k.TextBadgeNeutralSubtle
		case BadgeError:
			bg, fg = k.DangerTint, k.Danger
		case BadgeWarning:
			bg, fg = k.WarningTint, k.Warning
		case BadgeSuccess:
			bg, fg = k.SuccessTint, k.Success
		case BadgeInfo:
			bg, fg = k.InfoTint, k.Info
		case BadgeDestructive, BadgeRed:
			bg, fg = k.BadgeRed, white
		case BadgeBeta:
			fg = k.TextLink
			b.Border(1, k.Brand).BorderStyle(ui.BorderDashed)
		case BadgeOutline:
			bg, fg = k.Base, k.TextDefault
			b.Border(1, k.Fill)
		case BadgeGreen:
			bg, fg = k.BadgeGreen, white
		case BadgeNeutral:
			bg, fg = k.BadgeNeutral, white
		case BadgeOrange:
			bg, fg = k.BadgeOrange, black
		case BadgePurple:
			bg, fg = k.BadgePurple, white
		case BadgeTeal:
			bg, fg = k.BadgeTeal, white
		case BadgeTealSubtle:
			bg, fg = k.BadgeTeal.Alpha(0.15), k.TextBadgeTealSubtle
		case BadgeBlue:
			bg, fg = k.BadgeBlue, white
		}
		b.Background(bg).TextColor(fg)
	}
	b.Children(func() {
		if dot.A != 0 {
			ui.Box(c).Size(7, 7).Radius(4).Background(dot).Shrink(0)
		}
		if o.Icon != nil {
			icon(c, o.Icon, 12)
		}
		ui.Text(c, text).SingleLine()
	})
	return b
}
