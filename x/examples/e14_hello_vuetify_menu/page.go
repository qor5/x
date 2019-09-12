package e14_hello_vuetify_menu

import (
	"github.com/sunfmin/bran/web"
	. "github.com/sunfmin/bran/x/vuetify"
	h "github.com/theplant/htmlgo"
)

func HelloVuetifyMenu(ctx *web.EventContext) (pr web.PageResponse, err error) {
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
			web.Slot(
				VBtn("Menu as Popover").
					Attr("v-on", "on").
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
								VIcon("favorite"),
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
