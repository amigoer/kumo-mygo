package kumo

import (
	"fmt"

	"github.com/amigoer/kumo-mygo/icons/phosphor"
	"github.com/egoist/mygo/ui"
)

// PaginationOptions configures Pagination.
type PaginationOptions struct {
	// PerPage is the number of items on a page (10 when 0); Total the
	// number of items in all.
	PerPage, Total int
	// Simple shows only the previous and next buttons, as Kumo's
	// controls="simple".
	Simple bool
	// HideSummary hides "Showing 1-10 of 42" on the left.
	HideSummary bool
}

// Pagination creates the controls of paged items, as Kumo's Pagination: a
// summary of what shows on the left, and joined buttons for the first,
// previous, next and last pages around the page number on the right.
// *page counts from 1; the returned row's Changed reports a new page.
func Pagination(c *ui.Context, page *int, o PaginationOptions) ui.Element {
	k := tok(c)
	per := o.PerPage
	if per <= 0 {
		per = 10
	}
	pages := max(1, (o.Total+per-1)/per)
	*page = max(1, min(*page, pages))
	row := ui.Row(c).FillWidth().AlignItems(ui.Center).Justify(ui.SpaceBetween).Gap(8).Role(ui.RoleGroup)
	changed := false
	row.Children(func() {
		if !o.HideSummary {
			from, to := (*page-1)*per+1, min(*page*per, o.Total)
			if o.Total == 0 {
				from = 0
			}
			ui.Text(c, fmt.Sprintf(Strings.ShowingOf, from, to, o.Total)).FontSize(k.TextSM).TextColor(k.TextSubtle).Grow(1)
		} else {
			ui.Spacer(c)
		}
		ui.Row(c).Gap(8).AlignItems(ui.Center).Children(func() {
			first, last := *page <= 1, *page >= pages
			items := []ButtonGroupItem{
				{Label: Strings.First, Icon: phosphor.CaretDoubleLeft, IconOnly: true, Disabled: first},
				{Label: Strings.Previous, Icon: phosphor.CaretLeft, IconOnly: true, Disabled: first},
			}
			if o.Simple {
				items = items[1:]
			}
			if i := ButtonGroup(c.Key("prev"), items, ButtonOptions{Size: SizeSM}); i >= 0 {
				if o.Simple || i == 1 {
					*page--
				} else {
					*page = 1
				}
				changed = true
			}
			if !o.Simple {
				ui.Text(c, fmt.Sprintf(Strings.PageOf, *page, pages)).FontSize(k.TextSM).TextColor(k.TextSubtle).
					MinWidth(72).TextAlign(ui.Center).FontFeatures("tnum")
			}
			items = []ButtonGroupItem{
				{Label: Strings.Next, Icon: phosphor.CaretRight, IconOnly: true, Disabled: last},
				{Label: Strings.Last, Icon: phosphor.CaretDoubleRight, IconOnly: true, Disabled: last},
			}
			if o.Simple {
				items = items[:1]
			}
			if i := ButtonGroup(c.Key("next"), items, ButtonOptions{Size: SizeSM}); i >= 0 {
				if i == 0 {
					*page++
				} else {
					*page = pages
				}
				changed = true
			}
		})
	})
	if changed {
		row.SetChanged()
	}
	return row
}
