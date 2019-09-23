package vuetify_components

import (
	ch "github.com/goplaid/x/codehighlight"
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e15_vuetify_navigation_drawer"
	"github.com/goplaid/x/docs/utils"
	"github.com/goplaid/x/md"
	. "github.com/theplant/htmlgo"
)

var NavigationDrawer = Components(
	md.Markdown(`
Vuetify navigation drawer provide a popup layer that show on the side of the window.

Here is one example:
`),
	ch.Code(examples.VuetifyNavigationDrawerSample).Language("go"),
	utils.Demo("", e15_vuetify_navigation_drawer.VuetifyNavigationDrawerPath),
)
