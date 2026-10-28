package kumo

import (
	"fmt"

	"github.com/egoist/mygo/ui"
)

// SliderOptions configures Slider.
type SliderOptions struct {
	// Label names the slider, above it.
	Label string
	// Min and Max bound the value, 0 and 100 when both are 0.
	Min, Max float64
	// Step is the increment of the arrows and of dragging, 1 when 0.
	Step float64
	// Size is SizeBase (32 DIPs) or SizeSM (24 DIPs).
	Size Size
	// Format writes a value for the badge under the thumb and the bounds
	// below the track, "%g" when nil.
	Format func(v float64) string
	// HideBounds hides Min and Max below the track.
	HideBounds bool
	// Disabled disables the slider.
	Disabled bool
}

// Slider creates Kumo's slider setting *value: a recessed track whose
// filled part, a raised panel, ends in a grip the pointer drags, with the
// value in a badge under it while it is dragged or focused. The arrows,
// Home and End move it too. It returns the slider, whose Changed reports a
// new value.
func Slider(c *ui.Context, value *float64, opts ...SliderOptions) ui.Element {
	o := first(opts)
	k := tok(c)
	lo, hi := o.Min, o.Max
	if lo == 0 && hi == 0 {
		hi = 100
	}
	step := o.Step
	if step <= 0 {
		step = 1
	}
	format := o.Format
	if format == nil {
		format = func(v float64) string { return fmt.Sprintf("%g", v) }
	}
	h, outer, inner, font := float32(32), float32(8), float32(5), k.TextSM
	if o.Size == SizeSM || o.Size == SizeXS {
		h, outer, inner, font = 24, 6, 4, k.TextXS
	}
	const pad, thumb = 3, 16
	var s ui.Element
	col := ui.Column(c).FillWidth().Gap(8).FontSize(font)
	col.Children(func() {
		if o.Label != "" {
			ui.Text(c, o.Label).FontWeight(500).TextColor(k.TextDefault)
		}
		s = ui.SliderBase(c, value, lo, hi).Step(step).FillWidth().Height(h).
			PaddingX(pad + thumb/2).FocusRing(false).Cursor(ui.CursorGrab)
		if o.Disabled {
			s.Disabled(true)
		}
		if o.Label != "" {
			s.Label(o.Label)
		}
		s.Value(format(*value))
		dragging := s.Dragging()
		if dragging {
			s.Cursor(ui.CursorGrabbing)
		}
		frac := float32(0)
		if hi > lo {
			frac = float32((max(lo, min(*value, hi)) - lo) / (hi - lo))
		}
		rtl := c.Direction() == ui.RightToLeft
		focus := s.FocusVisible()
		recessed, line, overlay, grip, focusCol, drop := k.Recessed, k.Line, k.Overlay, k.Contrast.Alpha(0.25), k.Focus, k.ShadowDrop
		s.Draw(func(p *ui.Painter, r ui.Rect) {
			p.Shadow(r, outer, 0, 0, 0, 1, line)
			p.Fill(r, recessed, outer)
			travel := r.W - 2*pad - thumb
			x := r.X + pad + travel*frac
			if rtl {
				x = r.X + pad + travel*(1-frac)
			}
			fill := ui.Rect{X: r.X + pad, Y: r.Y + pad, W: x + thumb - (r.X + pad), H: r.H - 2*pad}
			if rtl {
				fill = ui.Rect{X: x, Y: r.Y + pad, W: r.X + r.W - pad - x, H: r.H - 2*pad}
			}
			p.Shadow(fill, inner, 0, 1, 3, 0, drop)
			p.Shadow(fill, inner, 0, 0, 0, 1, line)
			p.Fill(fill, overlay, inner)
			t := ui.Rect{X: x, Y: r.Y + pad, W: thumb, H: r.H - 2*pad}
			if focus {
				p.Shadow(t, inner, 0, 0, 0, 2, focusCol)
			}
			gh := t.H / 2
			p.Fill(ui.Rect{X: t.X + t.W/2 - 1, Y: t.Y + (t.H-gh)/2, W: 2, H: gh}, grip, 1)
		})
		if (dragging || focus) && !o.Disabled {
			badge := format(*value)
			inv, contrast := k.TextInverse, k.Contrast
			s.DrawOver(func(p *ui.Painter, r ui.Rect) {
				travel := r.W - 2*pad - thumb
				x := r.X + pad + travel*frac + thumb/2
				if rtl {
					x = r.X + pad + travel*(1-frac) + thumb/2
				}
				span := ui.Span{Text: badge, Size: font, Weight: 500, Color: inv}
				w, th := p.MeasureText(0, span)
				b := ui.Rect{X: x - w/2 - 6, Y: r.Y + r.H + 11, W: w + 12, H: th + 2}
				p.Fill(b, contrast, 4)
				p.RichText(b.X+6, b.Y+1, 0, span)
			})
		}
		if !o.HideBounds {
			ui.Row(c).Justify(ui.SpaceBetween).TextColor(k.TextSubtle).FontFeatures("tnum").Children(func() {
				ui.Text(c, format(lo))
				ui.Text(c, format(hi))
			})
		}
	})
	if o.Disabled {
		col.Opacity(0.5)
	}
	return s
}
