package kumo

import (
	"time"

	"github.com/egoist/mygo/ui"
)

// TooltipSide is where a tooltip shows beside its anchor.
type TooltipSide int

const (
	// TooltipTop, the default, shows above the anchor, or below it where
	// there is no room above.
	TooltipTop TooltipSide = iota
	TooltipBottom
	TooltipLeft
	TooltipRight
)

// TooltipOptions configures Tooltip.
type TooltipOptions struct {
	Side TooltipSide
}

var tooltipTransition = ui.ElementTransition{
	Duration: 150 * time.Millisecond,
	Enter:    &ui.Motion{Opacity: 0, Scale: 0.9},
	Exit:     &ui.Motion{Opacity: 0, Scale: 0.9},
}

// Tooltip shows text beside anchor while the pointer rests on it or the
// keyboard focus comes to it, in Kumo's look: a small card of the base
// surface with a hairline outline. Call it after building anchor, every
// frame. The text also becomes anchor's accessible description. It
// returns the tip, invalid while hidden.
func Tooltip(c *ui.Context, anchor ui.Element, text string, opts ...TooltipOptions) ui.Element {
	if text == "" || !anchor.Valid() {
		return ui.Element{}
	}
	o := first(opts)
	k := tok(c)
	anchor.Description(text)
	return ui.TooltipBase(c, anchor, func(tip ui.Element) {
		switch o.Side {
		case TooltipBottom:
			tip.AttachTo(anchor, ui.AnchorBottom, ui.AnchorTop).MarginY(10)
		case TooltipLeft:
			tip.AttachTo(anchor, ui.AnchorStart, ui.AnchorEnd).MarginX(10)
		case TooltipRight:
			tip.AttachTo(anchor, ui.AnchorEnd, ui.AnchorStart).MarginX(10)
		}
		tip.Column().MaxWidth(320).Padding(6, 10).Radius(6).
			Background(k.Base).TextColor(k.TextDefault).FontSize(k.TextSM).
			Border(1, k.Line).Transition(tooltipTransition)
		tip.Shadow(0, 4, 6, -1, k.ShadowDrop).Shadow(0, 2, 4, -2, k.ShadowDrop)
		tip.Children(func() { ui.Text(c, text) })
	})
}
