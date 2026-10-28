package kumo

import (
	"time"

	"github.com/amigoer/kumo-mygo/icons/phosphor"
	"github.com/egoist/mygo/ui"
)

// DialogOptions configures Dialog. Size picks Kumo's widths: SizeSM 288,
// SizeBase 384, SizeLG 512 DIPs; Width sets any other, as Kumo's xl 768.
type DialogOptions struct {
	Size  Size
	Width float32
	// Title heads the dialog and names it; Description follows it in
	// subtle text.
	Title, Description string
	// Footer builds the actions at the bottom, aligned to the end, as a
	// secondary Cancel and a primary confirm Button.
	Footer func()
	// HideClose hides the close button at the top right.
	HideClose bool
	// Alert makes it Kumo's alert dialog, for a choice the user must make:
	// neither a click on the backdrop nor Escape closes it, only its
	// buttons, and assistive technology announces it as an alert.
	Alert bool
}

var (
	dialogFade = ui.ElementTransition{Duration: 150 * time.Millisecond, Colors: true,
		Enter: &ui.Motion{Opacity: 0}, Exit: &ui.Motion{Opacity: 0}}
	dialogPop = ui.ElementTransition{Duration: 150 * time.Millisecond,
		Enter: &ui.Motion{Opacity: 0, Scale: 0.9}, Exit: &ui.Motion{Opacity: 0, Scale: 0.9}}
)

// Dialog shows a modal dialog while *open, as Kumo's Dialog: a panel of
// the base fill, rounded 12 DIPs in a ring of the line color, near the top
// of the window over the recessed color dimming it. A click on the
// backdrop or Escape closes it, setting *open to false; so does its close
// button. body builds the content between the header and the footer.
// Focus moves into the dialog and stays there while it shows. It returns
// the panel, invalid while closed.
//
//	kumo.Dialog(c, &app.confirm, kumo.DialogOptions{
//		Title: "Delete worker?", Description: "This cannot be undone.",
//		Footer: func() {
//			if kumo.Button(c, "Cancel").Clicked() { app.confirm = false }
//			if kumo.Button(c, "Delete", kumo.ButtonOptions{Variant: kumo.ButtonDestructive}).Clicked() { app.delete() }
//		},
//	}, nil)
func Dialog(c *ui.Context, open *bool, o DialogOptions, body func()) ui.Element {
	k := tok(c)
	width := o.Width
	if width <= 0 {
		switch o.Size {
		case SizeXS, SizeSM:
			width = 288
		case SizeLG:
			width = 512
		default:
			width = 384
		}
	}
	wasOpen := *open
	panel := ui.DialogBase(c, open, func(backdrop, panel ui.Element) {
		ww, wh := c.Size()
		top := float32(64)
		if ww < 640 {
			top = 32
		}
		backdrop.Background(k.Recessed.Alpha(0.8)).Transition(dialogFade).
			Justify(ui.Start).AlignItems(ui.Center).Padding(top, 16, 16, 16)
		panel.Column().Width(width).MaxWidth(max(0, ww-32)).MaxHeight(max(0, wh-top-16)).
			Radius(12).Background(k.Base).TextColor(k.TextDefault).FontSize(k.TextBase).
			Clip().Transition(dialogPop)
		ring(panel, 1, k.Line)
		panel.Shadow(0, 20, 25, -5, ui.RGBA(0, 0, 0, 0.03)).Shadow(0, 8, 10, -6, ui.RGBA(0, 0, 0, 0.03))
		if o.Alert {
			panel.Role(ui.RoleAlertDialog)
		}
		if o.Title != "" {
			panel.Label(o.Title)
		}
		if o.Description != "" {
			panel.Description(o.Description)
		}
		panel.Children(func() {
			ui.Scroll(c).FillWidth().Grow(1).Shrink(1).MinHeight(0).Children(func() {
				ui.Column(c).Padding(32).Gap(16).Children(func() {
					if o.Title != "" || !o.HideClose {
						ui.Row(c).Gap(16).AlignItems(ui.Start).Justify(ui.SpaceBetween).Children(func() {
							ui.Column(c).Gap(8).Grow(1).MinWidth(0).Children(func() {
								if o.Title != "" {
									ui.Text(c, o.Title).FontSize(24 * k.TextBase / 14).FontWeight(600).Level(2).LineHeight(1.25)
								}
								if o.Description != "" {
									ui.Text(c, o.Description).TextColor(k.TextSubtle).LineHeight(1.5)
								}
							})
							if !o.HideClose {
								b := Button(c, Strings.Close, ButtonOptions{Shape: ShapeSquare, Icon: phosphor.X})
								if b.Clicked() {
									*open = false
								}
							}
						})
					}
					if body != nil {
						body()
					}
				})
			})
			if o.Footer != nil {
				ui.Row(c).Justify(ui.End).AlignItems(ui.Center).Gap(8).Padding(0, 32, 32, 32).Wrap().Children(o.Footer)
			}
		})
	})
	// An alert dialog waits for a choice: neither the backdrop nor Escape
	// closes it, only its own buttons.
	if o.Alert && wasOpen && panel.Valid() && panel.Dismissed() {
		*open = true
	}
	return panel
}
