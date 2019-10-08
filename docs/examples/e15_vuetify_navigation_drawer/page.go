package e15_vuetify_navigation_drawer

// @snippet_begin(VuetifyNavigationDrawerSample)
import (
	"github.com/goplaid/web"
	. "github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
)

func VuetifyNavigationDrawer(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("showDrawer", showDrawer)

	pr.Body = VContainer(

		VBtn("show").On("click", "vars.drawer1 = !vars.drawer1"),

		VNavigationDrawer(
			h.Text("Hi"),
			VBtn("Close").On("click", "vars.drawer1 = false"),
		).Temporary(true).
			Attr("v-model", "vars.drawer1").
			Right(true).
			Bottom(true).
			Absolute(true).
			Width(600).Attr(web.InitContextVars, `{drawer1: false}`),

		VBtn("Show Drawer 2").OnClick("showDrawer"),

		web.Portal().EventFunc("").Name("drawer2"),
	)

	return
}

func showDrawer(ctx *web.EventContext) (er web.EventResponse, err error) {
	er.UpdatePortals = append(er.UpdatePortals,
		&web.PortalUpdate{
			Name: "drawer2",
			Body: VNavigationDrawer(
				h.Text("Drawer 2"),
			).Right(true).
				Attr("v-model", "vars.drawer2").
				Bottom(true).
				Temporary(true).
				Absolute(true).
				Value(true).
				Width(800).
				Attr(web.InitContextVars, `{drawer2: false}`),
			AfterLoaded: `setTimeout(function(){ comp.vars.drawer2 = true }, 100)`,
		},
	)
	return
}

// @snippet_end
const VuetifyNavigationDrawerPath = "/samples/vuetify-navigation-drawer"
