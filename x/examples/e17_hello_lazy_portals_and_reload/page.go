package e17_hello_lazy_portals_and_reload

import (
	"fmt"
	"time"

	"github.com/goplaid/web"
	. "github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
)

type mystate struct {
	Company string
	Error   string
}

var listItems = []string{"Apple", "Microsoft", "Google"}

func HelloLazyPortalsAndReload(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("addItem", addItem)
	ctx.Hub.RegisterEventFunc("menuItems", menuItems)
	ctx.Hub.RegisterEventFunc("addItemForm", addItemForm)
	ctx.Hub.RegisterEventFunc("portal1", portal1)
	ctx.Hub.RegisterEventFunc("reloadAB", reloadAB)
	ctx.Hub.RegisterEventFunc("updateCD", updateCD)

	pr.Schema = VApp(
		VContent(
			VContainer(
				VDialog(
					web.Slot(
						VBtn("Select").Color("primary").Attr("v-on", "on"),
					).Name("activator").Scope("{ on }"),
					web.Portal().EventFunc("menuItems").Name("menuContent"),
				),

				h.Div(
					h.H1("Portal A"),
					web.Portal().EventFunc("portal1").Name("portalA"),
				).Style("border: 2px solid blue;"),

				h.Div(
					h.H1("Portal B"),
					web.Portal().EventFunc("portal1").Name("portalB"),
				).Style("border: 2px solid red;"),

				VBtn("Reload Portal A and B").OnClick("reloadAB").Color("orange").Dark(true),

				h.Div(
					h.H1("Portal C"),
					web.Portal().EventFunc("").Name("portalC"),
				).Style("border: 2px solid blue;"),

				h.Div(
					h.H1("Portal D"),
					web.Portal().EventFunc("").Name("portalD"),
				).Style("border: 2px solid red;"),

				VBtn("Update Portal C and D").OnClick("updateCD").Color("primary").Dark(true),
			),
		),
	)
	return
}

func menuItems(ctx *web.EventContext) (r web.EventResponse, err error) {

	var items []h.HTMLComponent
	for _, item := range listItems {
		items = append(items, VListItem(
			VListItemTitle(h.Text(item)),
		))
	}

	items = append(items, VDivider())

	items = append(items,
		VDialog(
			web.Slot(
				VListItemAction(
					VBtn("Create New").Text(true).Attr("v-on", "on"),
				),
			).Name("activator").Scope("{ on }"),
			web.Portal().EventFunc("addItemForm").Name("addItemForm").Visible("true"),
		).Width("500"),
	)

	r.Schema = VList(items...)
	return
}

func addItemForm(ctx *web.EventContext) (r web.EventResponse, err error) {
	var s = &mystate{}
	ctx.MustUnmarshalForm(s)

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

func addItem(ctx *web.EventContext) (r web.EventResponse, err error) {
	var s = &mystate{}
	ctx.MustUnmarshalForm(s)

	if len(s.Company) < 5 {
		s.Error = "too short"
		r.ReloadPortals = []string{"addItemForm"}
		return
	}

	listItems = append(listItems, s.Company)
	s.Company = ""
	s.Error = ""
	r.ReloadPortals = []string{"menuContent"}
	return
}

func portal1(ctx *web.EventContext) (r web.EventResponse, err error) {
	r.Schema = h.Text(fmt.Sprint(time.Now().UnixNano()))
	return
}

func reloadAB(ctx *web.EventContext) (r web.EventResponse, err error) {
	r.ReloadPortals = []string{"portalA", "portalB"}
	return
}

func updateCD(ctx *web.EventContext) (r web.EventResponse, err error) {
	r.UpdatePortals = append(r.UpdatePortals,
		&web.PortalUpdate{
			Name:   "portalC",
			Schema: h.Text(fmt.Sprint(time.Now().UnixNano())),
		},
		&web.PortalUpdate{
			Name:   "portalD",
			Schema: h.Text(fmt.Sprint(time.Now().UnixNano())),
		},
	)
	return
}
