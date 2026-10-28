package kumo

import "github.com/egoist/mygo/ui"

// TableColumn describes a column of a Table.
type TableColumn struct {
	// Header is the column's heading.
	Header string
	// Width fixes the column's width in DIPs; columns without one share
	// the rest by Grow (1 when 0).
	Width, Grow float32
	// Align places the column's cells: ui.Start (the default), ui.Center
	// or ui.End, as numbers are.
	Align ui.Align
}

// TableOptions configures Table.
type TableOptions struct {
	// Striped fills every other row with the elevated color.
	Striped bool
	// CompactHeader draws Kumo's compact header: small strong text on
	// the elevated fill.
	CompactHeader bool
	// Selected, when set, adds a checkbox column: it returns the flag of
	// row i, which the checkbox toggles; the header's checks them all.
	Selected func(row int) *bool
	// OnRowClick, when set, makes rows clickable.
	OnRowClick func(row int)
	// Label names the table for assistive technology.
	Label string
}

// Table creates a table of rows by columns, as Kumo's Table: a header in
// semibold over a rule of the fill color, cells padded 12 DIPs, rows tinted
// as the pointer rests on them or as they are selected. cell builds the
// content of a row's column, as a Text or a Badge.
//
//	kumo.Table(c, cols, len(app.workers), func(r, col int) {
//		w := app.workers[r]
//		switch col {
//		case 0: kumo.Text(c, w.Name)
//		case 1: kumo.Badge(c, w.Status, kumo.BadgeOptions{Variant: kumo.BadgeSuccess})
//		}
//	}, kumo.TableOptions{Striped: true})
func Table(c *ui.Context, columns []TableColumn, rows int, cell func(row, col int), opts ...TableOptions) ui.Element {
	o := first(opts)
	k := tok(c)
	t := ui.Column(c).FillWidth().Background(k.Base).Role(ui.RoleTable).FontSize(k.TextBase).TextColor(k.TextDefault)
	if o.Label != "" {
		t.Label(o.Label)
	}
	place := func(e ui.Element, col TableColumn) ui.Element {
		if col.Width > 0 {
			e.Width(col.Width).Shrink(0)
		} else {
			g := col.Grow
			if g <= 0 {
				g = 1
			}
			e.Grow(g).Basis(0).MinWidth(0)
		}
		return e.Justify(col.Align)
	}
	t.Children(func() {
		head := ui.Row(c).AlignItems(ui.Center).Role(ui.RoleRow).BorderWidth(0, 0, 1, 0).BorderColor(k.Fill)
		py := float32(12)
		if o.CompactHeader {
			head.Background(k.Elevated).FontSize(k.TextXS).TextColor(k.TextStrong)
			py = 8
		} else {
			head.Background(k.Base).FontWeight(600)
		}
		head.Children(func() {
			if o.Selected != nil {
				all := rows > 0
				some := false
				for i := 0; i < rows; i++ {
					if *o.Selected(i) {
						some = true
					} else {
						all = false
					}
				}
				ui.Row(c).Padding(py, 12).Shrink(0).Role(ui.RoleColumnHeader).Children(func() {
					v := all
					if Checkbox(c, &v, "", CheckboxOptions{Indeterminate: some && !all}).Label("Select all").Changed() {
						for i := 0; i < rows; i++ {
							*o.Selected(i) = v
						}
					}
				})
			}
			for _, col := range columns {
				place(ui.Row(c).Padding(py, 12).Role(ui.RoleColumnHeader), col).Children(func() {
					ui.Text(c, col.Header).SingleLine()
				})
			}
		})
		for r := 0; r < rows; r++ {
			row := ui.Row(c.Key(r)).AlignItems(ui.Center).Role(ui.RoleRow).BorderWidth(0, 0, 1, 0).BorderColor(k.Hairline)
			selected := o.Selected != nil && *o.Selected(r)
			switch {
			case selected:
				row.Background(k.Tint)
			case o.OnRowClick != nil && row.Hovered():
				row.Background(k.Tint)
			case o.Striped && r%2 == 1:
				row.Background(k.Elevated)
			}
			if o.OnRowClick != nil {
				row.Cursor(ui.CursorPointer)
				if row.Clicked() {
					o.OnRowClick(r)
				}
			}
			row.Children(func() {
				if o.Selected != nil {
					ui.Row(c).Padding(12).Shrink(0).Role(ui.RoleCell).Children(func() {
						Checkbox(c, o.Selected(r), "").Label("Select row")
					})
				}
				for i, col := range columns {
					place(ui.Row(c).Padding(12).AlignItems(ui.Center).Gap(8).Role(ui.RoleCell), col).Children(func() {
						cell(r, i)
					})
				}
			})
		}
	})
	return t
}
