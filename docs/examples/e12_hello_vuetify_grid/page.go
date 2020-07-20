package e12_hello_vuetify_grid

import (
	"fmt"

	"github.com/goplaid/web"
	. "github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
)

func HelloVuetifyGrid(ctx *web.EventContext) (pr web.PageResponse, err error) {

	row := func(col int, count int, color string) (r []h.HTMLComponent) {
		for i := 0; i < count; i++ {
			r = append(r, VFlex(
				VCard(
					VCardText(h.Text(fmt.Sprint(col))),
				).Dark(true).Color(color),
			).Col(Xs, col))
		}
		return
	}

	var lc []h.HTMLComponent
	lc = append(lc, row(12, 1, "primary")...)
	lc = append(lc, row(6, 2, "secondary")...)
	lc = append(lc, row(4, 3, "primary")...)
	lc = append(lc, row(3, 4, "secondary")...)
	lc = append(lc, row(2, 6, "primary")...)
	lc = append(lc, row(1, 12, "secondary")...)

	pr.Body = VApp(
		VMain(
			VContainer(
				VLayout(
					lc...,
				).Row(true).Wrap(true),
			).GridList(Md).TextAlign(Xs, Center),
		),
	).Id("mainapp")
	return
}
