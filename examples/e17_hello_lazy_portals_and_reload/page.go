package e17_hello_lazy_portals_and_reload

import (
	"github.com/sunfmin/bran/ui"
	. "github.com/sunfmin/bran/vuetify"
	h "github.com/theplant/htmlgo"
)

type mystate struct {
	Company string
}

var listItems = []string{"Apple", "Microsoft", "Google"}

func HelloLazyPortalsAndReload(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("addItem", addItem)
	ctx.Hub.RegisterEventFunc("menuItems", menuItems)

	ctx.StateOrInit(&mystate{})

	pr.Schema = VApp(
		VContent(
			VContainer(
				VMenu(
					ui.Slot(
						VBtn("Select").Color("primary").On("on"),
					).Name("activator").Scope("{ on }"),
					ui.LazyPortal("menuItems").Name("menuContent").Visible("true"),
				).OffsetY(true),
			),
		),
	)
	return
}

func menuItems(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	var items []h.HTMLComponent
	for _, item := range listItems {
		items = append(items, VListTile(
			VListTileTitle(h.Text(item)),
		))
	}

	items = append(items, VDivider())

	items = append(items,
		VDialog(
			ui.Slot(
				VListTile(
					VBtn("Create New").Flat(true).On("on"),
				),
			).Name("activator").Scope("{ on }"),

			VCard(
				VCardText(
					VTextField().FieldName("Company"),
				),
				VCardActions(
					VBtn("Create").Color("primary").OnClick("addItem"),
				),
			),
		).Width("500"),
	)

	r.Schema = VList(items...)
	return
}

func addItem(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	s := ctx.State.(*mystate)
	listItems = append(listItems, s.Company)
	s.Company = ""
	r.ReloadPortal = "menuContent"
	return
}
