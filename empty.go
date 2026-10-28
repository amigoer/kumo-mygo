package kumo

import (
	"time"

	"github.com/amigoer/kumo-mygo/icons/phosphor"
	"github.com/egoist/mygo/ui"
)

// EmptyOptions configures Empty.
type EmptyOptions struct {
	// Size is SizeSM, SizeBase or SizeLG: how much room surrounds it.
	Size Size
	// Icon is drawn large above the title, in the subtle color.
	Icon IconFunc
	// Title and Description say what is missing and what to do.
	Title, Description string
	// CommandLine shows a command to copy, as "npm create cloudflare".
	CommandLine string
	// Contents builds what follows, as a primary Button.
	Contents func()
}

// Empty creates an empty state, as Kumo's Empty: a rounded, bordered
// panel centering an icon, a title, a description, a command to copy and
// actions.
func Empty(c *ui.Context, o EmptyOptions) ui.Element {
	k := tok(c)
	py, px, gap := float32(64), float32(40), float32(24)
	switch o.Size {
	case SizeXS, SizeSM:
		py, px, gap = 32, 24, 16
	case SizeLG:
		py, px, gap = 80, 48, 32
	}
	panel := ui.Column(c).FillWidth().AlignItems(ui.Center).Padding(py, px).Gap(gap).
		Radius(12).Border(1, k.Fill).Background(k.Control).TextColor(k.TextDefault)
	panel.Children(func() {
		if o.Icon != nil {
			icon(c, o.Icon, 48).TextColor(k.TextSubtle)
		}
		if o.Title != "" || o.Description != "" {
			ui.Column(c).AlignItems(ui.Center).Gap(10).MaxWidth(560).Children(func() {
				if o.Title != "" {
					ui.Text(c, o.Title).FontSize(k.TextLG).FontWeight(600).TextAlign(ui.Center).Level(2)
				}
				if o.Description != "" {
					ui.Text(c, o.Description).FontSize(k.TextBase).TextColor(k.TextSubtle).TextAlign(ui.Center).LineHeight(1.5)
				}
			})
		}
		if o.CommandLine != "" {
			commandLine(c, o.CommandLine)
		}
		if o.Contents != nil {
			ui.Row(c).Gap(8).AlignItems(ui.Center).Children(o.Contents)
		}
	})
	return panel
}

// commandLine shows a command in a monospace pill with a copy button.
func commandLine(c *ui.Context, cmd string) {
	k := tok(c)
	pill := ui.Row(c).Height(40).MaxWidthPercent(80).AlignItems(ui.Center).Gap(8).
		Padding(0, 8, 0, 12).Radius(8).Background(k.Overlay).Font(k.MonoFont)
	ring(pill, 1, k.Line)
	shadowXS(pill)
	pill.Children(func() {
		ui.Text(c, "$").TextColor(k.TextSubtle).Unselectable()
		ui.Text(c, cmd).SingleLine().Selectable().Shrink(1).MinWidth(0)
		copyButton(c, cmd, SizeSM)
	})
}

// copyButton is a ghost icon button copying text, which shows a check for
// a moment after.
func copyButton(c *ui.Context, text string, size Size) ui.Element {
	k := tok(c)
	holder := ui.Box(c)
	copied := ui.Local(holder, "copied", func() time.Time { return time.Time{} })
	recent := !copied.IsZero() && c.Now().Sub(*copied) < 1500*time.Millisecond
	var b ui.Element
	holder.Children(func() {
		label, ic := Strings.Copy, IconFunc(phosphor.Copy)
		if recent {
			label, ic = Strings.Copied, phosphor.Check
		}
		b = Button(c, label, ButtonOptions{Variant: ButtonGhost, Size: size, Shape: ShapeSquare, Icon: ic})
		if recent {
			b.TextColor(k.TextSuccess)
			c.After(1500*time.Millisecond - c.Now().Sub(*copied))
		}
		if b.Clicked() {
			c.WriteClipboard(text)
			*copied = c.Now()
			c.Announce(Strings.Copied)
		}
	})
	return b
}
