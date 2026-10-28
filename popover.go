package kumo

import (
	"time"

	"github.com/egoist/mygo/ui"
)

// PopoverSide is where a popover or menu opens beside its anchor.
type PopoverSide int

const (
	// SideBottom, the default, opens below the anchor, or above it where
	// there is no room below.
	SideBottom PopoverSide = iota
	SideTop
	SideLeft
	SideRight
)

// PopoverOptions configures Popover.
type PopoverOptions struct {
	// Title and Description head the popover.
	Title, Description string
	Side               PopoverSide
	// Width fixes the popover's width; it fits its content otherwise, up
	// to 360 DIPs.
	Width float32
}

var popTransition = ui.ElementTransition{Duration: 150 * time.Millisecond,
	Enter: &ui.Motion{Opacity: 0, Scale: 0.96}, Exit: &ui.Motion{Opacity: 0, Scale: 0.96}}

// Popover shows a panel beside anchor while *open, as Kumo's Popover: the
// base fill, rounded 8 DIPs in a ring of the line color, padded 16 by 12.
// The app opens it, as from anchor's Clicked; a press outside or Escape
// closes it. Call it every frame after building anchor. It returns the
// panel, invalid while closed.
//
//	b := kumo.Button(c, "Share")
//	if b.Clicked() { app.share = !app.share }
//	kumo.Popover(c, b, &app.share, kumo.PopoverOptions{Title: "Share"}, func() { … })
func Popover(c *ui.Context, anchor ui.Element, open *bool, o PopoverOptions, body func()) ui.Element {
	k := tok(c)
	anchor.Expanded(*open)
	return ui.PopoverBase(c, anchor, open, func(panel ui.Element) {
		placePopup(panel, anchor, o.Side, 8)
		stylePopup(panel, k)
		panel.Padding(12, 16).Gap(8).FontSize(k.TextSM).MaxWidth(360).Transition(popTransition)
		if o.Width > 0 {
			panel.Width(o.Width)
		}
		if o.Title != "" {
			panel.Label(o.Title)
		}
		panel.Children(func() {
			if o.Title != "" || o.Description != "" {
				ui.Column(c).Children(func() {
					if o.Title != "" {
						ui.Text(c, o.Title).FontSize(k.TextBase).FontWeight(500).FixedLineHeight(24)
					}
					if o.Description != "" {
						ui.Text(c, o.Description).FontSize(k.TextBase).TextColor(k.TextSubtle).FixedLineHeight(24)
					}
				})
			}
			if body != nil {
				body()
			}
		})
	})
}

// placePopup puts a panel on a side of its anchor, gap DIPs away.
func placePopup(panel, anchor ui.Element, side PopoverSide, gap float32) {
	switch side {
	case SideTop:
		panel.AttachTo(anchor, ui.AnchorTopStart, ui.AnchorBottomStart).Margin(0, 0, gap, 0)
	case SideLeft:
		panel.AttachTo(anchor, ui.AnchorTopStart, ui.AnchorTopEnd).MarginEnd(gap)
	case SideRight:
		panel.AttachTo(anchor, ui.AnchorTopEnd, ui.AnchorTopStart).MarginStart(gap)
	default:
		panel.Margin(gap, 0, 0, 0)
	}
}
