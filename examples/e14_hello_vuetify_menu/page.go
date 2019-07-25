package e14_hello_vuetify_menu

import (
	"github.com/sunfmin/bran/ui"
	h "github.com/theplant/htmlgo"

	. "github.com/sunfmin/bran/vuetify"
)

func HelloVuetifyMenu(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
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
		VMenu(
			ui.Slot(
				VBtn("Menu as Popover").
					On("on").
					Dark(true).
					Color("indigo"),
			).Name("activator").
				Scope("{ on }"),
			VCard(
				VList(
					VListItem(
						VListItemAvatar(
							h.Img("https://cdn.vuetifyjs.com/images/john.jpg").Alt("John"),
						),
						VListItemContent(
							VListItemTitle(h.Text("John Leider")),
							VListItemSubtitle(h.Text("Founder of Vuetify.js")),
						),
						VListItemAction(
							VBtn("").Icon(true).Children(
								VIcon("mdi-heart"),
							),
						),
					),
				),
				VDivider(),
				VList(
					VListItem(
						VListItemAction(
							VSwitch().Color("purple"),
						),
						VListItemTitle(h.Text("Enable messages")),
					),
					VListItem(
						VListItemAction(
							VSwitch().Color("purple"),
						),
						VListItemTitle(h.Text("Enable hints")),
					),
				),

				VCardActions(
					VSpacer(),
					VBtn("Cancel").Text(true),
					VBtn("Save").Color("primary").Text(true),
				),
			),
		).CloseOnContentClick(false).
			NudgeWidth(200).
			OffsetX(true),
	)

	return
}
