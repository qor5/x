package e16_hello_vuetify_simple_components

import (
	"github.com/goplaid/web"
	. "github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
)

func HelloVuetifySimpleComponents(ctx *web.EventContext) (pr web.PageResponse, err error) {
	wrapper := func(children ...h.HTMLComponent) h.HTMLComponent {
		return VApp(
			VContent(
				VContainer(
					children...,
				),
			),
		).Id("mainapp")
	}

	pr.Body = wrapper(
		h.Div(
			VAvatar(
				h.Img("https://vuetifyjs.com/apple-touch-icon-180x180.png"),
			).Color("grey lighten-3").Size(32),
			VAvatar(
				h.Img("https://vuetifyjs.com/apple-touch-icon-180x180.png"),
			).Tile(true).Color("grey lighten-3").Size(32),
		).Style("margin-bottom: 40px"),

		h.Div(
			VBadge(
				web.Slot(
					h.Span("6"),
				).Name("badge"),
				VIcon("shopping_cart").
					Large(true).
					Color("grey lighten-1"),
			).Left(true),
		).Style("margin-bottom: 40px"),

		h.Div(
			VChip(h.Text("Example Chip")),
			VChip(h.Text("Example Chip")).Close(true),
			VChip(
				VAvatar(h.Img("https://randomuser.me/api/portraits/men/35.jpg")),
				h.Text("Trevor Hansen"),
			).Close(true),
			VChip(
				VAvatar(h.Text("A")).Class("teal"),
				h.Text("ANZ Bank"),
			),
		).Style("margin-bottom: 40px"),

		h.Div(
			VAlert(h.Text("This is a success alert.")).Type("success").Value(true),
		).Style("margin-bottom: 40px"),
	)

	return
}
