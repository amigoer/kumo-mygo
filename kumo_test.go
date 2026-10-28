package kumo

import (
	"testing"

	"github.com/amigoer/kumo-mygo/icons/phosphor"
	"github.com/amigoer/kumo-mygo/theme"
	"github.com/egoist/mygo/ui"
)

// The views run headless, in both appearances, as a user clicks them.

func each(t *testing.T, fn func(t *testing.T, dark bool)) {
	for _, dark := range []bool{false, true} {
		name := "light"
		if dark {
			name = "dark"
		}
		t.Run(name, func(t *testing.T) { fn(t, dark) })
	}
}

func TestButtonClicks(t *testing.T) {
	each(t, func(t *testing.T, dark bool) {
		n := 0
		tt := ui.NewTester(func(c *ui.Context) {
			theme.Apply(c)
			for _, v := range []ButtonVariant{ButtonSecondary, ButtonPrimary, ButtonGhost, ButtonDestructive, ButtonSecondaryDestructive, ButtonOutline} {
				if Button(c.Key(v), "Go", ButtonOptions{Variant: v}).Clicked() && v == ButtonPrimary {
					n++
				}
			}
			if Button(c, "Add", ButtonOptions{Shape: ShapeSquare, Icon: phosphor.Plus}).Clicked() {
				n += 10
			}
			Button(c, "Off", ButtonOptions{Disabled: true})
		}, 600, 300)
		tt.SetDark(dark)
		if err := tt.Click("Add"); err != nil {
			t.Fatal(err)
		}
		if n != 10 {
			t.Errorf("n = %d after clicking the icon button, want 10", n)
		}
		// The accessible name is a label; drawn text would add a second one.
		shown := 0
		for _, s := range tt.Texts() {
			if s == "Add" {
				shown++
			}
		}
		if shown != 1 {
			t.Errorf("an icon-only button has %d labels, want 1", shown)
		}
	})
}

func TestChoices(t *testing.T) {
	var agree, on bool
	plan := "free"
	tab := 0
	tt := ui.NewTester(func(c *ui.Context) {
		theme.Apply(c)
		Checkbox(c, &agree, "Agree")
		Switch(c, &on, "Wi-Fi")
		Radio(c, &plan, "free", "Free")
		Radio(c, &plan, "pro", "Pro")
		Tabs(c, &tab, []string{"One", "Two"})
	}, 600, 400)
	for _, l := range []string{"Agree", "Wi-Fi", "Pro", "Two"} {
		if err := tt.Click(l); err != nil {
			t.Fatal(err)
		}
	}
	if !agree || !on || plan != "pro" || tab != 1 {
		t.Errorf("agree %v, on %v, plan %q, tab %d", agree, on, plan, tab)
	}
}

func TestDialogCloses(t *testing.T) {
	open := true
	tt := ui.NewTester(func(c *ui.Context) {
		theme.Apply(c)
		Dialog(c, &open, DialogOptions{Title: "Rename"}, func() { Text(c, "Body") })
	}, 800, 600)
	if !tt.HasText("Body") {
		t.Fatalf("dialog content missing: %q", tt.Texts())
	}
	if err := tt.Click(Strings.Close); err != nil {
		t.Fatal(err)
	}
	if open {
		t.Error("the close button left the dialog open")
	}
}

func TestPagination(t *testing.T) {
	page := 1
	tt := ui.NewTester(func(c *ui.Context) {
		theme.Apply(c)
		Pagination(c, &page, PaginationOptions{PerPage: 10, Total: 35})
	}, 800, 100)
	if err := tt.Click(Strings.Next); err != nil {
		t.Fatal(err)
	}
	if err := tt.Click(Strings.Last); err != nil {
		t.Fatal(err)
	}
	if page != 4 {
		t.Errorf("page %d, want 4", page)
	}
}

func TestThemeUse(t *testing.T) {
	l, d := theme.Light(), theme.Dark()
	l.Brand = ui.Hex("#7c3aed")
	theme.Use(l, d)
	defer theme.Use(theme.Light(), theme.Dark())
	var got ui.Color
	tt := ui.NewTester(func(c *ui.Context) {
		theme.Apply(c)
		got = theme.Of(c).Brand
	}, 100, 100)
	tt.Frame()
	if got != l.Brand {
		t.Errorf("brand %v, want %v", got, l.Brand)
	}
}
