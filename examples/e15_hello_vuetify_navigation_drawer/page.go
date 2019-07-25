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
					VListItem(h.Text("Application")).Class("title"),
				),
			).Flat(true),
			VDivider(),
			VList(
				VListItem(
					VListItemAction(
						VIcon("dashboard"),
					),
					VListItemContent(
						VListItemTitle(h.Text("Home")),
					),
				).On("click", ""),
				VListItem(
					VListItemAction(
						VIcon("question_answer"),
					),
					VListItemContent(
						VListItemTitle(h.Text("About")),
					),
				).On("click", ""),
			).Dense(true).Class("pt-0"),
		).Permanent(true),
	)

	return
}
