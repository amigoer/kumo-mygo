// Gallery shows every Kumo component in a native MyGo window, in the light
// and the dark appearance.
//
//	go run ./examples/gallery
//	go run ./examples/gallery -theme dark
package main

import (
	"flag"
	"log"

	kumo "github.com/amigoer/kumo-mygo"
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
)

func main() {
	appearance := flag.String("theme", "system", "appearance: system, light or dark")
	lang := flag.String("lang", "zh", "components' own words: zh or en")
	flag.Parse()
	if *lang == "zh" {
		kumo.Strings = kumo.StringsZhHans
	}
	switch *appearance {
	case "light":
		mygo.Theme.SetSource(mygo.ThemeLight)
	case "dark":
		mygo.Theme.SetSource(mygo.ThemeDark)
	}
	g := newGallery()
	mygo.App.WhenReady(func() {
		w := mygo.NewWindow(mygo.WindowOptions{
			Title:     "kumo-mygo",
			Width:     1180,
			Height:    800,
			MinWidth:  720,
			MinHeight: 520,
			Content:   ui.View(g.view),
		})
		w.OnClosed(func() { mygo.App.Quit() })
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
