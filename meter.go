package kumo

import (
	"fmt"
	"time"

	"github.com/egoist/mygo/ui"
)

// MeterOptions configures Meter.
type MeterOptions struct {
	// Label names the meter, above its bar on the left.
	Label string
	// Max is the value of a full bar, 100 when 0.
	Max float64
	// HideValue hides the value above the bar on the right.
	HideValue bool
	// Value, when set, is shown in place of the percentage, as "4.2 GB of
	// 10 GB".
	Value string
	// Color fills the bar in place of the brand color, as Danger near a
	// limit.
	Color ui.Color
}

// Meter shows value out of Max as a bar, as Kumo's Meter: a label and the
// value above an 8 DIP rounded track of the fill color, filled from the
// left in the brand color, easing to new values.
func Meter(c *ui.Context, value float64, opts ...MeterOptions) ui.Element {
	o := first(opts)
	k := tok(c)
	hi := o.Max
	if hi <= 0 {
		hi = 100
	}
	frac := float32(max(0, min(value/hi, 1)))
	col := ui.Column(c).FillWidth().Gap(8).Role(ui.RoleMeter).Range(0, hi, value)
	if o.Label != "" {
		col.Label(o.Label)
	}
	shown := o.Value
	if shown == "" {
		shown = fmt.Sprintf("%.0f%%", float64(frac)*100)
	}
	col.Value(shown)
	fill := k.Brand
	if o.Color.A != 0 {
		fill = o.Color
	}
	w := col.AnimateTo("value", frac, ui.Animation{Duration: 300 * time.Millisecond, Ease: ui.EaseOut})
	col.Children(func() {
		if o.Label != "" || !o.HideValue {
			ui.Row(c).Justify(ui.SpaceBetween).AlignItems(ui.Center).Gap(16).Children(func() {
				ui.Text(c, o.Label).FontSize(k.TextXS).TextColor(k.TextSubtle)
				if !o.HideValue {
					ui.Text(c, shown).FontSize(k.TextSM).FontWeight(500).TextColor(k.TextDefault).FontFeatures("tnum")
				}
			})
		}
		ui.Box(c).FillWidth().Height(8).Radius(4).Background(k.Fill).Draw(func(p *ui.Painter, r ui.Rect) {
			if w > 0 {
				p.Fill(ui.Rect{X: r.X, Y: r.Y, W: max(r.H, r.W*w), H: r.H}, fill, r.H/2)
			}
		})
	})
	return col
}
