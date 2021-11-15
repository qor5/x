package vuetify_components

import (
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e17_hello_lazy_portals_and_reload"
	"github.com/goplaid/x/docs/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var LazyPortalsAndReload = Doc(
	Markdown(`
Use ~web.Portal().EventFunc("menuItems").Name("menuContent")~ to put a portal place holder inside a part of html, and it will load specified event func's response body inside the place holder after the main page is rendered in a separate AJAX request. Later in an event func, you could also use ~r.ReloadPortals = []string{"menuContent"}~ to reload the portal.
`),
	ch.Code(examples.LazyPortalsAndReloadSample).Language("go"),
	utils.Demo("Lazy Portals", e17_hello_lazy_portals_and_reload.LazyPortalsAndReloadPath, "e17_hello_lazy_portals_and_reload/page.go"),
).Title("Lazy Portals").
	Slug("vuetify-components/lazy-portals")
