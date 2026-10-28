package kumo

import (
	"math"
	"time"

	"github.com/egoist/mygo/ui"
)

// Loader sizes, as Kumo's sm, base and lg.
const (
	LoaderSM   float32 = 16
	LoaderBase float32 = 24
	LoaderLG   float32 = 32
)

// Loader creates Kumo's spinner, size DIPs square (LoaderBase when 0): an
// arc that grows, travels and shrinks as it turns over a faint ring, drawn
// in color, or in the subtle text color when color is zero. Assistive
// technology sees a progress indicator named Strings.Loading; rename it with
// Label.
func Loader(c *ui.Context, size float32, color ...ui.Color) ui.Element {
	if size <= 0 {
		size = LoaderBase
	}
	k := tok(c)
	col := k.TextSubtle
	if len(color) > 0 && color[0].A != 0 {
		col = color[0]
	}
	e := ui.Box(c).Size(size, size).Shrink(0).Role(ui.RoleProgress).Label(Strings.Loading)
	if c.Preferences().ReduceMotion {
		// A still quarter arc says "busy" without moving.
		e.Draw(func(p *ui.Painter, r ui.Rect) { paintLoader(p, r, col, 0.125, 0.25, 0) })
		return e
	}
	turn := e.Loop("turn", 2*time.Second, ui.Linear)
	phase := e.Loop("dash", 1500*time.Millisecond, ui.Linear)
	e.Draw(func(p *ui.Painter, r ui.Rect) {
		// SVG's animate: the dash grows 0→42 of 150 over the first half,
		// and its offset goes 0→-16→-59 (of a 59.7 circumference).
		var dash, offset float32
		if phase < 0.5 {
			dash = 42 * phase * 2
			offset = -16 * phase * 2
		} else {
			dash = 42
			offset = -16 - 43*(phase-0.5)*2
		}
		paintLoader(p, r, col, turn, dash/59.69, -offset/59.69)
	})
	return e
}

// paintLoader draws the faint ring and an arc of length frac (of the
// circle) starting start of the way around, turned by turn.
func paintLoader(p *ui.Painter, r ui.Rect, col ui.Color, turn, frac, start float32) {
	s := min(r.W, r.H) / 24
	cx, cy := r.X+r.W/2, r.Y+r.H/2
	rad := 9.5 * s
	w := 2 * s
	var ring ui.Path
	ring.Circle(cx, cy, rad)
	p.StrokePath(&ring, w, col.Alpha(0.1))
	if frac <= 0.001 {
		frac = 0.001
	}
	const segs = 48
	n := max(2, int(float32(segs)*frac)+1)
	a0 := (turn + start) * 2 * math.Pi
	span := frac * 2 * math.Pi
	var arc ui.Path
	for i := 0; i <= n; i++ {
		a := float64(a0 + span*float32(i)/float32(n) - math.Pi/2)
		x, y := cx+rad*float32(math.Cos(a)), cy+rad*float32(math.Sin(a))
		if i == 0 {
			arc.MoveTo(x, y)
		} else {
			arc.LineTo(x, y)
		}
	}
	arc.Cap(ui.RoundCap)
	p.StrokePath(&arc, w, col)
}
