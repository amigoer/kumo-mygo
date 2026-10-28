package kumo

import "github.com/egoist/mygo/ui"

// Code creates code within text, in the monospace font at the small size.
func Code(c *ui.Context, code string) ui.Element {
	k := tok(c)
	return ui.Text(c, code).Font(k.MonoFont).FontSize(k.TextSM).TextColor(k.TextDefault).
		TextBackground(k.Tint)
}

// CodeBlockOptions configures CodeBlock.
type CodeBlockOptions struct {
	// Copy adds a button copying the code at the top right.
	Copy bool
	// MaxHeight scrolls the code past this height, when set.
	MaxHeight float32
}

// CodeBlock creates a block of code, as Kumo's CodeBlock: monospace lines
// 20 DIPs apart in a rounded, bordered panel of the base fill, selectable
// and kept as written, scrolling sideways when long.
func CodeBlock(c *ui.Context, code string, opts ...CodeBlockOptions) ui.Element {
	o := first(opts)
	k := tok(c)
	panel := ui.Box(c).FillWidth().MinWidth(0).Radius(6).Border(1, k.Fill).Background(k.Base)
	panel.Children(func() {
		scroll := ui.Scroll(c).FillWidth()
		if o.MaxHeight > 0 {
			scroll.MaxHeight(o.MaxHeight)
		}
		scroll.Padding(10).Children(func() {
			ui.Text(c, code).Font(k.MonoFont).FontSize(k.TextSM).FixedLineHeight(20).
				TextColor(k.TextDefault).NoWrap().Selectable()
		})
		if o.Copy {
			ui.Box(c).Absolute().Top(4).Right(4).Children(func() { copyButton(c, code, SizeSM) })
		}
	})
	return panel
}

// ClipboardTextOptions configures ClipboardText.
type ClipboardTextOptions struct {
	// Size is the field's size, as Input's.
	Size Size
	// Label names the value for assistive technology.
	Label string
}

// ClipboardText shows a value to copy, as an API token or an ID, in a
// read-only monospace field with a copy button at its end, as Kumo's
// ClipboardText.
func ClipboardText(c *ui.Context, value string, opts ...ClipboardTextOptions) ui.Element {
	o := first(opts)
	k := tok(c)
	m := inputSize(k, o.Size)
	box := ui.Row(c).Height(m.h).AlignItems(ui.Stretch).Radius(m.radius).Background(k.Base).Clip()
	ring(box, 1, k.Line)
	if o.Label != "" {
		box.Role(ui.RoleGroup).Label(o.Label)
	}
	box.Children(func() {
		ui.Row(c).Grow(1).MinWidth(0).AlignItems(ui.Center).PaddingX(m.px).Children(func() {
			ui.Text(c, value).Font(k.MonoFont).FontSize(m.font - 1).SingleLine().Selectable().TextColor(k.TextDefault)
		})
		ui.Row(c).AlignItems(ui.Center).PaddingX(4).BorderWidth(0, 0, 0, 1).BorderColor(k.Line).Children(func() {
			copyButton(c, value, SizeSM)
		})
	})
	return box
}
