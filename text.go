package kumo

import "github.com/egoist/mygo/ui"

// TextVariant is the role of a text, as Kumo's Text variant prop.
type TextVariant int

const (
	// TextBody, the default, is text in the default color.
	TextBody TextVariant = iota
	// TextSecondary is quieter text in the subtle color.
	TextSecondary
	// TextSuccess and TextError color an outcome.
	TextSuccess
	TextError
	// TextMono and TextMonoSecondary set code and identifiers.
	TextMono
	TextMonoSecondary
	// TextHeading1, TextHeading2 and TextHeading3 are headings of 30, 24
	// and 16 DIPs in semibold; assistive technology sees their level.
	TextHeading1
	TextHeading2
	TextHeading3
)

// TextOptions configures Text.
type TextOptions struct {
	Variant TextVariant
	// Size sets the size of a non-heading text: SizeXS 12, SizeSM 13,
	// SizeBase 14 (the default) or SizeLG 16 DIPs.
	Size Size
	// Bold sets the weight to semibold.
	Bold bool
	// Truncate keeps the text on one line, ending in an ellipsis.
	Truncate bool
}

// Text creates a text styled as Kumo's Text. It wraps at the width it
// gets; it returns the element for more styling.
func Text(c *ui.Context, s string, opts ...TextOptions) ui.Element {
	o := first(opts)
	k := tok(c)
	e := ui.Text(c, s)
	size := k.TextBase
	switch o.Size {
	case SizeXS:
		size = k.TextXS
	case SizeSM:
		size = k.TextSM
	case SizeLG:
		size = k.TextLG
	}
	color := k.TextDefault
	switch o.Variant {
	case TextSecondary:
		color = k.TextSubtle
	case TextSuccess:
		color = k.TextSuccess
	case TextError:
		color = k.TextDanger
	case TextMono:
		e.Font(k.MonoFont)
	case TextMonoSecondary:
		e.Font(k.MonoFont)
		color = k.TextSubtle
	case TextHeading1:
		size = 30 * k.TextBase / 14
		e.FontWeight(600).Level(1).LineHeight(1.2)
	case TextHeading2:
		size = 24 * k.TextBase / 14
		e.FontWeight(600).Level(2).LineHeight(1.25)
	case TextHeading3:
		size = k.TextLG
		e.FontWeight(600).Level(3)
	}
	if o.Bold {
		e.FontWeight(600)
	}
	if o.Truncate {
		e.SingleLine()
	}
	return e.FontSize(size).TextColor(color)
}
