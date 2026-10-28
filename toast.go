package kumo

import (
	"time"

	"github.com/amigoer/kumo-mygo/icons/phosphor"
	"github.com/amigoer/kumo-mygo/icons/phosphor/fill"
	"github.com/egoist/mygo/ui"
)

// Toast kinds, the Type of a ui.Toast, which color its icon, title and
// ring as Kumo's toast variants do.
const (
	ToastDefault = ""
	ToastSuccess = "success"
	ToastInfo    = "info"
	ToastWarning = "warning"
	ToastError   = "error"
)

// Notify shows a toast, as Kumo's toast manager's notify: a title, a
// description and a kind (ToastSuccess, …, or ToastDefault). It returns
// its ID, for c.CloseToast. For an action or another timeout, call
// c.AddToast with a ui.Toast of your own; Toaster draws it the same way.
func Notify(c *ui.Context, title, description, kind string) string {
	return c.AddToast(ui.Toast{Title: title, Description: description, Type: kind})
}

var toastMotion = ui.ElementTransition{Duration: 250 * time.Millisecond,
	Enter: &ui.Motion{Opacity: 0, Y: 40}, Exit: &ui.Motion{Opacity: 0, Collapse: true}}

// Toaster draws the window's toasts in Kumo's look, in place of MyGo's:
// stacked in the bottom right corner, 340 DIPs wide, each a translucent
// card of the base fill, rounded 12 DIPs, ringed in the color of its kind.
// Call it once in the root view, wherever; add toasts with Notify or
// c.AddToast.
func Toaster(c *ui.Context) ui.Element {
	k := tok(c)
	return ui.ToastViewportBase(c, func(viewport ui.Element, toasts []ui.Toast) {
		w, _ := c.Size()
		edge := float32(32)
		if w < 640 {
			edge = 16
		}
		viewport.Padding(edge).AlignItems(ui.End).Justify(ui.End).Gap(12)
		for _, t := range toasts {
			part := ui.ToastBase(c, t)
			root := part.Root.Column().Width(min(340, w-2*edge)).Padding(16).Gap(4).Radius(12).
				TextColor(k.TextDefault).Transition(toastMotion)
			root.Shadow(0, 10, 15, -3, k.ShadowDrop).Shadow(0, 4, 6, -4, k.ShadowDrop)
			// Kumo layers a translucent tint of the kind over the card.
			accent, base, tint := k.TextDefault, k.Base, k.Base.Alpha(0.9)
			var ic IconFunc
			switch t.Type {
			case ToastSuccess:
				accent, ic, tint = k.Success, fill.CheckCircle, k.SuccessTint.Alpha(0.2)
			case ToastInfo:
				accent, ic, tint, base = k.Info, fill.Info, k.InfoTint.Alpha(0.5), k.Control
			case ToastWarning:
				accent, ic, tint = k.Warning, fill.Warning, k.WarningTint.Alpha(0.5)
			case ToastError:
				accent, ic, tint = k.Danger, fill.WarningOctagon, k.DangerTint.Alpha(0.5)
			}
			root.Background(tint.Over(base))
			if ic != nil {
				ring(root, 0.5, accent)
			} else {
				ring(root, 1, k.Line)
			}
			t := t
			root.Children(func() {
				ui.Row(c).Gap(8).AlignItems(ui.Start).Padding(0, 20, 0, 0).Children(func() {
					if ic != nil {
						icon(c, ic, 16).TextColor(accent).Margin(2, 0, 0, 0)
					}
					ui.Column(c).Gap(4).Grow(1).MinWidth(0).Children(func() {
						if t.Title != "" {
							title := ui.Text(c, t.Title).FontSize(15.6 * k.TextBase / 14).FontWeight(500).FixedLineHeight(20)
							if ic != nil {
								title.TextColor(accent)
							}
						}
						if t.Description != "" {
							ui.Text(c, t.Description).FontSize(14.8 * k.TextBase / 14).FixedLineHeight(20).TextColor(k.TextDefault.Alpha(0.7))
						}
					})
				})
				if t.Action != "" {
					ui.Row(c).Gap(8).Margin(8, 0, 0, 0).Children(func() {
						a := part.ActionButton().Label(t.Action)
						StyleButton(c, a, ButtonOptions{Size: SizeSM})
						a.Children(func() { ui.Text(c, t.Action).SingleLine() })
					})
				}
				xc := k.TextSubtle
				if ic != nil {
					xc = accent
				}
				x := part.CloseButton().Label(Strings.Close).Absolute().Top(8).Right(8).Size(20, 20).
					Radius(4).TextColor(xc).Cursor(ui.CursorPointer)
				if x.Hovered() {
					x.Background(xc.Alpha(0.15))
				}
				x.Children(func() { icon(c, phosphor.X, 12) })
			})
		}
	})
}
