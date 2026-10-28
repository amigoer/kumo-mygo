package theme

import (
	"sync"

	"github.com/egoist/mygo/ui"
)

var (
	mu            sync.Mutex
	light, dark   = Light(), Dark()
	cache         = map[cacheKey]*Tokens{}
	uiThemeCached = map[cacheKey]*ui.Theme{}
)

type cacheKey struct {
	dark, contrast bool
	scale          float32
}

// Use replaces the tokens of the light and the dark appearance for the whole
// app, as a brand of your own over Kumo's: start from Light and Dark and
// change what differs. A nil argument keeps that appearance as it is. Call
// it before the first window, or call Context.Invalidate after it.
//
//	l, d := theme.Light(), theme.Dark()
//	l.Brand, d.Brand = ui.Hex("#7c3aed"), ui.Hex("#8b5cf6")
//	theme.Use(l, d)
func Use(l, d *Tokens) {
	mu.Lock()
	defer mu.Unlock()
	if l != nil {
		cp := *l
		cp.Dark = false
		light = &cp
	}
	if d != nil {
		cp := *d
		cp.Dark = true
		dark = &cp
	}
	clear(cache)
	clear(uiThemeCached)
}

// Of returns the tokens for the frame's appearance: the dark ones when the
// frame's theme is dark (c.Theme().Dark), the light ones otherwise. Like
// MyGo's own theme, they follow the desktop's higher contrast, darkening
// lines and secondary text, and its text scale. Components call it every
// frame; the result is shared, so treat it as read-only.
func Of(c *ui.Context) *Tokens {
	p := c.Preferences()
	key := cacheKey{dark: c.Theme().Dark, contrast: p.HighContrast, scale: p.TextScale}
	if key.scale <= 0 {
		key.scale = 1
	}
	mu.Lock()
	defer mu.Unlock()
	if t, ok := cache[key]; ok {
		return t
	}
	t := derive(key)
	cache[key] = t
	return t
}

func derive(key cacheKey) *Tokens {
	src := light
	if key.dark {
		src = dark
	}
	t := *src
	if key.contrast {
		t.Line = t.Line.Mix(t.TextDefault, 0.45)
		t.Line.A = 255
		t.Hairline = t.Hairline.Mix(t.TextDefault, 0.3)
		t.TextSubtle = t.TextSubtle.Mix(t.TextDefault, 0.4)
		t.TextPlaceholder = t.TextPlaceholder.Mix(t.TextDefault, 0.3)
	}
	if key.scale != 1 {
		t.TextXS *= key.scale
		t.TextSM *= key.scale
		t.TextBase *= key.scale
		t.TextLG *= key.scale
	}
	return &t
}

// Kumo returns a MyGo theme made from the tokens of the frame's appearance,
// so that MyGo's own widgets, scroll bars, menus, focus rings and toasts
// match Kumo's components. Its signature fits a shell's Theme option.
func Kumo(c *ui.Context) *ui.Theme {
	p := c.Preferences()
	key := cacheKey{dark: c.Theme().Dark, contrast: p.HighContrast, scale: p.TextScale}
	if key.scale <= 0 {
		key.scale = 1
	}
	tok := Of(c)
	mu.Lock()
	defer mu.Unlock()
	if t, ok := uiThemeCached[key]; ok {
		return t
	}
	t := UITheme(tok)
	uiThemeCached[key] = t
	return t
}

// UITheme maps tokens onto MyGo's theme, for MyGo's widgets.
func UITheme(k *Tokens) *ui.Theme {
	var t *ui.Theme
	if k.Dark {
		t = ui.DarkTheme()
	} else {
		t = ui.LightTheme()
	}
	t.Dark = k.Dark
	t.Background = k.Canvas
	t.Surface = k.Base
	t.SurfaceHover = k.Tint
	t.SurfacePressed = k.Fill
	t.Border = k.Line.Over(k.Base)
	t.Text = k.TextDefault
	t.TextMuted = k.TextSubtle
	t.Accent = k.Brand
	t.AccentHover = k.BrandHover
	t.AccentPressed = k.BrandHover.Mix(black, 0.1)
	t.AccentText = white
	t.Danger = k.Danger
	t.Warning = k.Warning
	t.Success = k.Success
	t.Selection = k.Brand.Alpha(0.25)
	t.Focus = k.Brand.Alpha(0.7)
	t.Inverse = k.Contrast
	t.InverseText = k.TextInverse
	t.Scrollbar = k.TextDefault.Alpha(0.3)
	t.Radius = 8
	t.Spacing = 4
	t.FontSize = k.TextBase
	t.Font = k.Font
	return t
}

// Apply installs Kumo's MyGo theme for the current frame. Call it first in
// the root view, before creating any widget: SetTheme affects only what is
// built after it. MyGo reruns the view as the appearance changes.
func Apply(c *ui.Context) { c.SetTheme(Kumo(c)) }
