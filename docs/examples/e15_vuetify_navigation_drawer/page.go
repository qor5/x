package e15_vuetify_navigation_drawer

// @snippet_begin(VuetifyNavigationDrawerSample)
import (
	"fmt"
	"time"

	"github.com/goplaid/web"
	. "github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
)

func VuetifyNavigationDrawer(ctx *web.EventContext) (pr web.PageResponse, err error) {

	pr.Body = VContainer(
		h.H2("A drawer that has close button"),

		VBtn("show").On("click", "vars.drawer1 = !vars.drawer1"),

		VNavigationDrawer(
			h.Text("Hi"),
			VBtn("Close").On("click", "vars.drawer1 = false"),
		).Temporary(true).
			Attr("v-model", "vars.drawer1").
			Right(true).
			Bottom(true).
			Absolute(true).
			Width(600),

		h.H2("Load a drawer from remote and show it").Class("pt-8"),

		VBtn("Show Drawer 2").OnClick("showDrawer"),

		web.Portal().Name("drawer2UpdateContent"),

		web.Portal().Name("drawer2"),
	).Attr(web.InitContextVars, `{drawer1: false, drawer2: false}`)

	return
}

func showDrawer(ctx *web.EventContext) (er web.EventResponse, err error) {
	er.UpdatePortals = append(er.UpdatePortals,
		&web.PortalUpdate{
			Name: "drawer2",
			Body: VNavigationDrawer(
				h.Text("Drawer 2"),
				web.Portal(
					textField(""),
				).Name("InputPortal"),
				VBtn("Update parent and close").
					OnClick("updateParentAndClose"),
			).Right(true).
				Attr("v-model", "vars.drawer2").
				Bottom(true).
				Temporary(true).
				Absolute(true).
				Value(true).
				Width(800),
		},
	)

	er.VarsScript = `setTimeout(function(){ vars.drawer2 = true }, 100)`
	return
}

func textField(value string, fieldErrors ...string) h.HTMLComponent {
	return VTextField().
		FieldName("Drawer2Input").
		ErrorMessages(fieldErrors...).
		Value(value)
}

func updateParentAndClose(ctx *web.EventContext) (er web.EventResponse, err error) {
	if len(ctx.R.FormValue("Drawer2Input")) < 10 {
		er.UpdatePortals = append(er.UpdatePortals, &web.PortalUpdate{
			Name: "InputPortal",
			Body: textField(ctx.R.FormValue("Drawer2Input"), "input more then 10 characters"),
		})
		return
	}

	er.UpdatePortals = append(er.UpdatePortals, &web.PortalUpdate{
		Name: "drawer2UpdateContent",
		Body: h.Text(fmt.Sprintf("Updated content at %s", time.Now())),
	})

	er.VarsScript = "vars.drawer2 = false"
	return
}

var VuetifyNavigationDrawerPB = web.Page(VuetifyNavigationDrawer).
	EventFunc("showDrawer", showDrawer).
	EventFunc("updateParentAndClose", updateParentAndClose)

const VuetifyNavigationDrawerPath = "/samples/vuetify-navigation-drawer"

// @snippet_end
