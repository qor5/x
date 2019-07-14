package e15_hello_vuetify_navigation_drawer

import (
	"github.com/sunfmin/bran/ui"
	h "github.com/theplant/htmlgo"

	. "github.com/sunfmin/bran/vuetify"
)

func HelloVuetifyNavigationDrawer(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	wrapper := func(children ...h.HTMLComponent) h.HTMLComponent {
		return VApp(
			VContent(
				VContainer(
					children...,
				),
			),
		).Id("mainapp")
	}

	pr.Schema = wrapper(
		VNavigationDrawer(
			VToolbar(
				VList(
					VListTile(h.Text("Application")).Class("title"),
				),
			).Flat(true),
			VDivider(),
			VList(
				VListTile(
					VListTileAction(
						VIcon("dashboard"),
					),
					VListTileContent(
						VListTileTitle(h.Text("Home")),
					),
				).On("click", ""),
				VListTile(
					VListTileAction(
						VIcon("question_answer"),
					),
					VListTileContent(
						VListTileTitle(h.Text("About")),
					),
				).On("click", ""),
			).Dense(true).Class("pt-0"),
		).Permanent(true),
	)

	return
}
