package kumo

import (
	"time"

	"github.com/egoist/mygo/ui"
)

// SwitchOptions configures Switch.
type SwitchOptions struct {
	// Size is SizeSM, SizeBase or SizeLG (SizeXS is SizeSM).
	Size Size
	// Neutral draws the monochrome variant, for subtle toggles.
	Neutral bool
	// Disabled disables the switch.
	Disabled bool
	// LabelFirst puts the label before the switch.
	LabelFirst bool
}

// Switch creates a switch turning *on on and off, with a label, as Kumo's
// Switch: a squircle track whose square knob slides across, in blue when
// on. The label is the accessible name; a click on it toggles the switch
// too. Changed reports a click.
func Switch(c *ui.Context, on *bool, label string, opts ...SwitchOptions) ui.Element {
	o := first(opts)
	k := tok(c)
	// Track height and width: Kumo's h-4 w-8, h-4.5 w-9 and h-5 w-10.
	h, w := float32(18), float32(36)
	switch o.Size {
	case SizeXS, SizeSM:
		h, w = 16, 32
	case SizeLG:
		h, w = 20, 40
	}
	row := ui.SwitchBase(c, on).Gap(8).AlignItems(ui.Center).FocusRing(false).Cursor(ui.CursorPointer)
	if o.Disabled {
		row.Disabled(true)
	}
	if label != "" {
		row.Label(label)
	}
	if o.LabelFirst {
		row.Reverse().Justify(ui.End)
	}
	row.Checked(*on)
	focus := row.FocusVisible()
	pos := row.AnimateTo("knob", boolToFloat(*on), ui.Animation{Duration: 150 * time.Millisecond, Ease: ui.EaseOut})
	row.Children(func() {
		trackOn, ringOn, thumbOn := k.SwitchOn, k.SwitchOnRing, k.SwitchThumbOn
		trackOff, ringOff, thumbOff := k.SwitchOff, k.SwitchOffRing, k.SwitchThumbOff
		if o.Neutral {
			if k.Dark {
				trackOn, ringOn, thumbOn = k.Base, ui.Oklch(0.371, 0, 0), ui.Oklch(0.708, 0, 0)
				trackOff, ringOff, thumbOff = k.Base, k.Hairline, ui.Oklch(0.24, 0, 0)
			} else {
				trackOn, ringOn, thumbOn = ui.Oklch(0.556, 0, 0), ui.Oklch(0.439, 0, 0), k.Base
				trackOff, ringOff, thumbOff = ui.Oklch(0.935, 0, 0), k.Hairline, k.Base
			}
		}
		track := ui.Box(c).Size(w, h).Radius(5).Shrink(0).PassThrough().
			Background(trackOff.Mix(trackOn, pos))
		ring(track, 1, ringOff.Mix(ringOn, pos))
		if focus {
			ring(track, 2, k.Brand)
		}
		thumb := thumbOff.Mix(thumbOn, pos)
		edge, drop := k.ShadowEdge, k.ShadowDrop
		track.Draw(func(p *ui.Painter, r ui.Rect) {
			x := r.X + (r.W-r.H)*pos
			knob := ui.Rect{X: x, Y: r.Y, W: r.H, H: r.H}
			p.Shadow(knob, 5, 0, 1, 2, 0, drop)
			p.Shadow(knob, 5, 0, 0, 1, 0.5, edge)
			p.Fill(knob, thumb, 5)
		})
		if label != "" {
			ui.Text(c, label).FontSize(k.TextBase).FontWeight(500).TextColor(k.TextDefault)
		}
	})
	return dimmed(row)
}

func boolToFloat(b bool) float32 {
	if b {
		return 1
	}
	return 0
}
