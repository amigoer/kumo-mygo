// Package kumo is Cloudflare's Kumo design system as native MyGo
// components: the same tokens, variants and sizes as @cloudflare/kumo, drawn
// by MyGo on the GPU without a webview, HTML or JavaScript.
//
// Components are functions of a *ui.Context that build MyGo elements and
// return them, so the usual element methods still apply:
//
//	theme.Apply(c) // first in the root view
//	if kumo.Button(c, "Deploy", kumo.ButtonOptions{Variant: kumo.ButtonPrimary}).Clicked() {
//		app.deploy()
//	}
//	kumo.Input(c, &app.name, kumo.InputOptions{Label: "Name", Placeholder: "my-worker"})
//
// The application owns state: components edit values through pointers, as
// MyGo's widgets do, and report what happened with Clicked and Changed.
// Behavior, keyboard, focus and accessibility come from MyGo's widgets
// without a look (ButtonBase, SelectBase, DialogBase, …); this package gives
// them Kumo's look from the tokens of package theme, which follow the
// light or dark appearance. Options structs mirror Kumo's props; their zero
// values are Kumo's defaults.
//
// Element values belong to one build pass: don't keep them across frames.
// Key a parent (c.Key) to keep a stateful component's identity in lists.
package kumo
