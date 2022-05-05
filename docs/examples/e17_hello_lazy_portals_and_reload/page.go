package e17_hello_lazy_portals_and_reload

// @snippet_begin(LazyPortalsAndReloadSample)

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

func LazyPortalsAndReload(ctx *web.EventContext) (pr web.PageResponse, err error) {

	pr.Body = VApp(
		VMain(
			VContainer(
				VDialog(
					web.Slot(
						VBtn("Select").Color("primary").Attr("v-on", "on"),
					).Name("activator").Scope("{ on }"),
					web.Portal().Loader(web.POST().EventFunc("menuItems")).Name("menuContent"),
				),

				h.Div(
					h.H1("Portal A"),
					web.Portal().Loader(web.POST().EventFunc("portal1")).Name("portalA"),
				).Style("border: 2px solid blue;"),

				h.Div(
					h.H1("Portal B"),
					web.Portal().Loader(web.POST().EventFunc("portal1")).Name("portalB"),
				).Style("border: 2px solid red;"),

				VBtn("Reload Portal A and B").OnClick("reloadAB").Color("orange").Dark(true),

				h.Div(
					h.H1("Portal C"),
					web.Portal().Name("portalC"),
				).Style("border: 2px solid blue;"),

				h.Div(
					h.H1("Portal D"),
					web.Portal().Name("portalD"),
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
			web.Portal().Loader(web.POST().EventFunc("addItemForm")).Name("addItemForm").Visible("true"),
		).Width("500"),
	)

	r.Body = VList(items...)
	return
}

func addItemForm(ctx *web.EventContext) (r web.EventResponse, err error) {
	var s = &mystate{}
	ctx.MustUnmarshalForm(s)

	textField := VTextField().FieldName("Company")

	if len(s.Error) > 0 {
		textField.Error(true).ErrorMessages(s.Error)
	}

	r.Body = VCard(
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
	r.Body = h.Text(fmt.Sprint(time.Now().UnixNano()))
	return
}

func reloadAB(ctx *web.EventContext) (r web.EventResponse, err error) {
	r.ReloadPortals = []string{"portalA", "portalB"}
	return
}

func updateCD(ctx *web.EventContext) (r web.EventResponse, err error) {
	r.UpdatePortals = append(r.UpdatePortals,
		&web.PortalUpdate{
			Name: "portalC",
			Body: h.Text(fmt.Sprint(time.Now().UnixNano())),
		},
		&web.PortalUpdate{
			Name: "portalD",
			Body: h.Text(fmt.Sprint(time.Now().UnixNano())),
		},
	)
	return
}

var LazyPortalsAndReloadPB = web.Page(LazyPortalsAndReload).
	EventFunc("addItem", addItem).
	EventFunc("menuItems", menuItems).
	EventFunc("addItemForm", addItemForm).
	EventFunc("portal1", portal1).
	EventFunc("reloadAB", reloadAB).
	EventFunc("updateCD", updateCD)

const LazyPortalsAndReloadPath = "/samples/lazy-portals-and-reload"

// @snippet_end
