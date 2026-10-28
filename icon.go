package kumo

import "github.com/egoist/mygo/ui"

// Icon creates an icon from a parsed SVG or an accessor such as
// phosphor.MagnifyingGlass (passed without calling it), 16 DIPs square as
// Kumo draws Phosphor icons in text. It takes the color of the text around
// it; set another with TextColor, another size with Size. Icons are
// decorations that assistive technology skips, unless Label names them.
func Icon[T *ui.SVG | func() *ui.SVG](c *ui.Context, source T) ui.Element {
	var svg *ui.SVG
	switch v := any(source).(type) {
	case *ui.SVG:
		svg = v
	case func() *ui.SVG:
		if v != nil {
			svg = v()
		}
	}
	return ui.Icon(c, svg).Size(16, 16).Shrink(0)
}
