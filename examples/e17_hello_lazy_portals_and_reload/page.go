package e17_hello_lazy_portals_and_reload

import (
	"github.com/sunfmin/bran/ui"
	. "github.com/sunfmin/bran/vuetify"
	h "github.com/theplant/htmlgo"
)

type mystate struct {
	Company string
	Error   string
}

var listItems = []string{"Apple", "Microsoft", "Google"}

func HelloLazyPortalsAndReload(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("addItem", addItem)
	ctx.Hub.RegisterEventFunc("menuItems", menuItems)
	ctx.Hub.RegisterEventFunc("addItemForm", addItemForm)

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
			ui.LazyPortal("addItemForm").Name("addItemForm").Visible("true"),
		).Width("500"),
	)

	r.Schema = VList(items...)
	return
}

func addItemForm(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	s := ctx.State.(*mystate)
	textField := VTextField().FieldName("Company")

	if len(s.Error) > 0 {
		textField.Error(true).ErrorMessages(s.Error)
	}

	r.Schema = VCard(
		VCardText(
			textField,
		),
		VCardActions(
			VBtn("Create").Color("primary").OnClick("addItem"),
		),
	)
	return
}

func addItem(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	s := ctx.State.(*mystate)
	if len(s.Company) < 5 {
		s.Error = "too short"
		r.ReloadPortal = "addItemForm"
		return
	}

	listItems = append(listItems, s.Company)
	s.Company = ""
	s.Error = ""
	r.ReloadPortal = "menuContent"
	return
}
