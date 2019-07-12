package e13_hello_vuetify_list

import (
	"github.com/sunfmin/bran/ui"
	h "github.com/theplant/htmlgo"

	. "github.com/sunfmin/bran/vuetify"
)

func HelloVuetifyList(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	wrapper := func(children ...h.HTMLComponent) h.HTMLComponent {
		return VApp(
			VContent(
				VContainer(
					VLayout(
						VFlex(
							VCard(children...),
						).Col(Xs, 6).Offset(Sm, 3),
					).Row(true),
				).GridList(Md).TextAlign(Xs, Center),
			),
		).Id("mainapp")
	}

	pr.Schema = wrapper(
		VToolbar(
			VToolbarSideIcon(),
			VToolbarTitle("Inbox"),
			VSpacer(),
			VBtn("").Icon(true).Children(
				VIcon("search"),
			),
		).Color("cyan").Dark(true),
		VList(
			VSubheader(h.Text("Today")),
			VListTile(
				VListTileAvatar(
					h.Img("https://cdn.vuetifyjs.com/images/lists/1.jpg"),
				),
				VListTileContent(
					VListTileTitle(h.Text("Brunch this weekend?")),
					VListTileSubTitle(
						h.Span("Ali Connors").Class("text--primary"),
						h.Text("&mdash; I'll be in your neighborhood doing errands this weekend. Do you want to hang out?"),
					),
				),
			).Avatar(true),
			VDivider().Inset(true),
			VListTile(
				VListTileAvatar(
					h.Img("https://cdn.vuetifyjs.com/images/lists/2.jpg"),
				),
				VListTileContent(
					VListTileTitle(h.RawHTML(`Summer BBQ <span class="grey--text text--lighten-1">4</span>`)),
					VListTileSubTitle(h.RawHTML(`<span class='text--primary'>to Alex, Scott, Jennifer</span> &mdash; Wish I could come, but I'm out of town this weekend.`)),
				),
			).Avatar(true),
			VDivider().Inset(true),
			VListTile(
				VListTileAvatar(
					h.Img("https://cdn.vuetifyjs.com/images/lists/3.jpg"),
				),
				VListTileContent(
					VListTileTitle(h.Text(`Oui oui`)),
					VListTileSubTitle(h.RawHTML(`<span class='text--primary'>Sandra Adams</span> &mdash; Do you have Paris recommendations? Have you ever been?`)),
				),
			).Avatar(true),
		).TwoLine(true),
	)

	return
}
