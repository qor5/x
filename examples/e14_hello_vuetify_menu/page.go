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
					VListTile(
						VListTileAvatar(
							h.Img("https://cdn.vuetifyjs.com/images/john.jpg").Alt("John"),
						),
						VListTileContent(
							VListTileTitle(h.Text("John Leider")),
							VListTileSubTitle(h.Text("Founder of Vuetify.js")),
						),
						VListTileAction(
							VBtn("").Icon(true).Children(
								VIcon("favorite"),
							),
						),
					).Avatar(true),
				),
				VDivider(),
				VList(
					VListTile(
						VListTileAction(
							VSwitch().Color("purple"),
						),
						VListTileTitle(h.Text("Enable messages")),
					),
					VListTile(
						VListTileAction(
							VSwitch().Color("purple"),
						),
						VListTileTitle(h.Text("Enable hints")),
					),
				),

				VCardActions(
					VSpacer(),
					VBtn("Cancel").Flat(true),
					VBtn("Save").Color("primary").Flat(true),
				),
			),
		).CloseOnContentClick(false).
			NudgeWidth(200).
			OffsetX(true),
	)

	return
}
