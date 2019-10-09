package vuetify_components

import (
	ch "github.com/goplaid/x/codehighlight"
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e13_vuetify_list"
	"github.com/goplaid/x/docs/examples/e14_vuetify_menu"
	"github.com/goplaid/x/docs/utils"
	"github.com/goplaid/x/md"
	. "github.com/theplant/htmlgo"
)

var ATasteOfUsingVuetifyInGo = Components(
	md.Markdown(`
[Vuetify](https://vuetifyjs.com/en/) is a really mature Vue components library for
[Material Design](https://material.io/design/). We have made the efforts to 
integrate most all of it as a go package. You can use it with ease just like any
other go package.
`),
	utils.Anchor(H2(""), "Use container, toolbar, list, list item etc"),
	md.Markdown(`
This example is purely render, we didn't integrate any interaction (event func) to it.
`),
	ch.Code(examples.VuetifyListSample).Language("go"),
	utils.Demo("Vuetify List", e13_vuetify_list.HelloVuetifyListPath, "e13_vuetify_list/page.go"),

	utils.Anchor(H2(""), "Use menu, card, list, etc"),
	md.Markdown(`
This example uses the menu popup, card, list component. and some interactions of clicking 
buttons on the menu popup.
`),
	ch.Code(examples.VuetifyMenuSample).Language("go"),
	md.Markdown(`
~.Attr(web.INIT_CONTEXT_VARS, "{myMenuShow: false}")~ is a special vue directive that
we created to initialize vue context component data variables. It will initialize
~vars.myMenuShow~ to ~false~. So that you don't need to modify javascript code to do 
the initialization. It's often useful to control dialog, popups. At this example, 
We add it, So that the cancel button on the menu, could actually close the menu without
requesting server backend.

~toggleFavored~ event func did an partial update only to the favorite icon button. So that it won't close the 
menu popup, but updated the button to toggle the favorite icon.
`),
	utils.Demo("Vuetify Menu", e14_vuetify_menu.HelloVuetifyMenuPath, "e14_vuetify_menu/page.go"),
)
