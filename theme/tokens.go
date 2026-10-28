package theme

import "github.com/egoist/mygo/ui"

// Tokens are Kumo's semantic design tokens for one appearance, the native
// counterpart of the --color-kumo-* and --text-color-kumo-* custom
// properties of @cloudflare/kumo. Components never use raw colors: they read
// these by role, so a whole app restyles by changing one Tokens value.
//
// Field names drop the kumo- prefix: bg-kumo-base is Base, ring-kumo-line
// is Line and text-kumo-subtle is TextSubtle.
type Tokens struct {
	// Dark is set for the dark appearance.
	Dark bool

	// Surfaces, from the window outward.

	// Canvas is the page behind everything.
	Canvas ui.Color
	// Elevated lifts a region off the canvas, as a sidebar or the header
	// strip of a LayerCard.
	Elevated ui.Color
	// Recessed sinks a region below the canvas, as a well or a code block.
	Recessed ui.Color
	// Base is the face of cards, inputs and secondary buttons.
	Base ui.Color
	// Tint is a quiet fill, as a hovered ghost button or row.
	Tint ui.Color
	// Contrast is the strongest fill, nearly the text color.
	Contrast ui.Color
	// Overlay fills popovers, menus and dialogs.
	Overlay ui.Color
	// Control fills form controls: checkboxes, radios, the switch's knob.
	Control ui.Color
	// Interact is the fill of controls that are off but can be turned on,
	// as an unchecked switch's track.
	Interact ui.Color
	// Fill and FillHover are neutral fills, as the track of a meter.
	Fill      ui.Color
	FillHover ui.Color

	// Brand is the accent: primary buttons, checked controls, the focus
	// ring. BrandHover is it hovered.
	Brand      ui.Color
	BrandHover ui.Color

	// Line outlines controls and cards; Hairline divides within them.
	Line     ui.Color
	Hairline ui.Color
	// Focus is the neutral ring of a focused field.
	Focus ui.Color
	// ShadowEdge and ShadowDrop make Kumo's two-part shadow: a crisp 1px
	// edge and a soft drop.
	ShadowEdge ui.Color
	ShadowDrop ui.Color

	// Status colors, each with a translucent Tint to fill behind it.
	Info        ui.Color
	InfoTint    ui.Color
	Warning     ui.Color
	WarningTint ui.Color
	Danger      ui.Color
	DangerTint  ui.Color
	Success     ui.Color
	SuccessTint ui.Color
	// BannerInfo and BannerWarning fill banners of those kinds.
	BannerInfo    ui.Color
	BannerWarning ui.Color

	// Badge fills, with white text unless noted on Badge's variants.
	BadgeRed      ui.Color
	BadgeGreen    ui.Color
	BadgeOrange   ui.Color
	BadgePurple   ui.Color
	BadgeTeal     ui.Color
	BadgeBlue     ui.Color
	BadgeNeutral  ui.Color
	BadgeInverted ui.Color

	// Text colors.
	TextDefault     ui.Color
	TextInverse     ui.Color
	TextStrong      ui.Color
	TextSubtle      ui.Color
	TextInactive    ui.Color
	TextPlaceholder ui.Color
	// TextBrand is Cloudflare orange, for the logo and brand marks; Brand,
	// not TextBrand, is the interface accent.
	TextBrand   ui.Color
	TextLink    ui.Color
	TextInfo    ui.Color
	TextSuccess ui.Color
	TextDanger  ui.Color
	TextWarning ui.Color
	// Text on the subtle badge variants and on the inverted badge.
	TextBadgeOrangeSubtle  ui.Color
	TextBadgeTealSubtle    ui.Color
	TextBadgeNeutralSubtle ui.Color
	TextBadgeInverted      ui.Color

	// Switch colors, which Kumo takes from its raw palette rather than the
	// brand: the track on and off, each with its 1px ring, and the knob on
	// and off.
	SwitchOn, SwitchOnRing, SwitchOff, SwitchOffRing ui.Color
	SwitchThumbOn, SwitchThumbOff                    ui.Color

	// Type scale in DIPs: Kumo's text-xs, text-sm, text-base and text-lg.
	TextXS, TextSM, TextBase, TextLG float32
	// Font is the family of text, "" for the system's; MonoFont that of code.
	Font, MonoFont string
}

// Kumo's raw palette: Tailwind v4's oklch values and Kumo's own neutral
// steps (kumo-binding.css), which its tokens resolve to.
// Lightness is 0 to 1, as ui.Oklch takes it.
var (
	white = ui.Hex("#ffffff")
	black = ui.Hex("#000000")

	neutral = func(l float32) ui.Color { return ui.Oklch(l, 0, 0) }

	neutral900 = neutral(0.205)
	neutral100 = neutral(0.97)

	blue400 = ui.Oklch(0.707, 0.165, 254.624)
	blue300 = ui.Oklch(0.809, 0.105, 251.813)
	blue500 = ui.Oklch(0.623, 0.214, 259.815)
	blue600 = ui.Oklch(0.546, 0.245, 262.881)
	blue700 = ui.Oklch(0.488, 0.243, 264.376)
	blue800 = ui.Oklch(0.424, 0.199, 265.638)

	red400 = ui.Oklch(0.704, 0.191, 22.216)
	red500 = ui.Oklch(0.637, 0.237, 25.331)
	red600 = ui.Oklch(0.577, 0.245, 27.325)
	red700 = ui.Oklch(0.505, 0.213, 27.518)

	emerald200 = ui.Oklch(0.905, 0.093, 164.15)
	emerald400 = ui.Oklch(0.765, 0.177, 163.223)
	emerald600 = ui.Oklch(0.596, 0.145, 163.225)
	emerald700 = ui.Oklch(0.508, 0.118, 165.612)
	emerald800 = ui.Oklch(0.432, 0.095, 166.913)

	orange200 = ui.Oklch(0.901, 0.076, 70.697)
	orange400 = ui.Oklch(0.75, 0.183, 55.934)
	orange650 = ui.Oklch(0.815, 0.197, 76)
	orange800 = ui.Oklch(0.47, 0.157, 37.304)

	teal200 = ui.Oklch(0.91, 0.096, 180.426)
	teal650 = ui.Oklch(0.549, 0.096, 184.565)
	teal700 = ui.Oklch(0.511, 0.096, 186.391)
	teal800 = ui.Oklch(0.437, 0.078, 188.216)

	purple600 = ui.Oklch(0.558, 0.288, 302.321)
	purple700 = ui.Oklch(0.496, 0.265, 301.924)

	yellow100 = ui.Oklch(0.973, 0.071, 103.193)

	// brand is Kumo's accent blue; the dark appearance mixes it 10% toward
	// black in oklch, which scales its lightness and chroma by 0.9.
	brand     = ui.Oklch(0.5772, 0.2324, 260)
	brandDark = ui.Oklch(0.5772*0.9, 0.2324*0.9, 260)

	cloudflareOrange = ui.Hex("#f6821f")
)

// Light returns Kumo's light tokens, a new value the caller may change.
func Light() *Tokens {
	return &Tokens{
		Canvas:    neutral(0.9875),
		Elevated:  neutral(0.98),
		Recessed:  neutral(0.965),
		Base:      white,
		Tint:      neutral100,
		Contrast:  neutral(0.12),
		Overlay:   neutral(0.9875),
		Control:   white,
		Interact:  neutral(0.87),
		Fill:      neutral(0.922),
		FillHover: neutral(0.965),

		Brand:      brand,
		BrandHover: blue700,

		Line:       neutral(0.145).Alpha(0.1),
		Hairline:   neutral(0.935),
		Focus:      neutral(0.15),
		ShadowEdge: black.Alpha(0.12),
		ShadowDrop: black.Alpha(0.08),

		Info:          blue500,
		InfoTint:      ui.Oklch(0.932, 0.032, 255.6).Alpha(0.45),
		Warning:       ui.Oklch(0.739, 0.177, 58.2),
		WarningTint:   ui.Oklch(0.931, 0.107, 94.6).Alpha(0.2),
		Danger:        red500,
		DangerTint:    ui.Oklch(0.936, 0.032, 17.7).Alpha(0.42),
		Success:       emerald600,
		SuccessTint:   ui.Oklch(0.962, 0.043, 156.7).Alpha(0.57),
		BannerInfo:    ui.Oklch(0.932, 0.032, 255.585).Alpha(0.7),
		BannerWarning: yellow100,

		BadgeRed:      red600,
		BadgeGreen:    emerald600,
		BadgeOrange:   orange650,
		BadgePurple:   purple600,
		BadgeTeal:     teal650,
		BadgeBlue:     blue600,
		BadgeNeutral:  neutral(0.556),
		BadgeInverted: neutral(0.145),

		TextDefault:     neutral900,
		TextInverse:     neutral100,
		TextStrong:      neutral(0.145),
		TextSubtle:      neutral(0.556),
		TextInactive:    neutral(0.87),
		TextPlaceholder: neutral(0.708),
		TextBrand:       cloudflareOrange,
		TextLink:        blue800,
		TextInfo:        blue800,
		TextSuccess:     emerald800,
		TextDanger:      red700,
		TextWarning:     ui.Oklch(0.597, 0.144, 57.5),

		TextBadgeOrangeSubtle:  orange800,
		TextBadgeTealSubtle:    teal800,
		TextBadgeNeutralSubtle: neutral(0.269),
		TextBadgeInverted:      white,

		SwitchOn:       blue500,
		SwitchOnRing:   blue600,
		SwitchOff:      neutral(0.922),
		SwitchOffRing:  neutral(0.87),
		SwitchThumbOn:  white,
		SwitchThumbOff: white,

		TextXS: 12, TextSM: 13, TextBase: 14, TextLG: 16,
		MonoFont: "monospace",
	}
}

// Dark returns Kumo's dark tokens, a new value the caller may change.
func Dark() *Tokens {
	return &Tokens{
		Dark:      true,
		Canvas:    neutral(0.10),
		Elevated:  neutral(0.12),
		Recessed:  neutral(0.15),
		Base:      neutral(0.17),
		Tint:      neutral(0.269),
		Contrast:  neutral(0.985),
		Overlay:   neutral(0.269),
		Control:   neutral900,
		Interact:  neutral(0.371),
		Fill:      neutral(0.269),
		FillHover: neutral(0.269),

		Brand:      brandDark,
		BrandHover: blue700,

		Line:       neutral(0.32),
		Hairline:   neutral(0.269),
		Focus:      neutral(0.935),
		ShadowEdge: white.Alpha(0.1),
		ShadowDrop: black.Alpha(0.3),

		Info:          blue500,
		InfoTint:      ui.Oklch(0.38, 0.145, 265.5).Alpha(0.22),
		Warning:       ui.Oklch(0.645, 0.168, 50),
		WarningTint:   ui.Oklch(0.353, 0.079, 65).Alpha(0.37),
		Danger:        red600,
		DangerTint:    ui.Oklch(0.429, 0.176, 28.7).Alpha(0.17),
		Success:       emerald400,
		SuccessTint:   ui.Oklch(0.393, 0.096, 152.3).Alpha(0.2),
		BannerInfo:    ui.Oklch(0.379, 0.146, 265.522).Alpha(0.5),
		BannerWarning: ui.Oklch(0.554, 0.135, 66.442).Alpha(0.5),

		BadgeRed:      red700,
		BadgeGreen:    emerald700,
		BadgeOrange:   orange650,
		BadgePurple:   purple700,
		BadgeTeal:     teal700,
		BadgeBlue:     blue700,
		BadgeNeutral:  neutral(0.439),
		BadgeInverted: white,

		TextDefault:     neutral100,
		TextInverse:     neutral(0.205),
		TextStrong:      neutral(0.985),
		TextSubtle:      neutral(0.708),
		TextInactive:    neutral(0.439),
		TextPlaceholder: neutral(0.556),
		TextBrand:       cloudflareOrange,
		TextLink:        blue400,
		TextInfo:        blue400,
		TextSuccess:     emerald200,
		TextDanger:      red400,
		TextWarning:     orange400,

		TextBadgeOrangeSubtle:  orange200,
		TextBadgeTealSubtle:    teal200,
		TextBadgeNeutralSubtle: neutral(0.922),
		TextBadgeInverted:      black,

		SwitchOn:       blue600,
		SwitchOnRing:   blue500,
		SwitchOff:      neutral(0.371),
		SwitchOffRing:  neutral(0.439),
		SwitchThumbOn:  blue300,
		SwitchThumbOff: neutral(0.24),

		TextXS: 12, TextSM: 13, TextBase: 14, TextLG: 16,
		MonoFont: "monospace",
	}
}
