package e15_hello_vuetify_navigation_drawer

import (
	"github.com/goplaid/web"
	. "github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
)

func HelloVuetifyNavigationDrawer(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("showDrawer", showDrawer)
	wrapper := func(children ...h.HTMLComponent) h.HTMLComponent {
		return VApp(
			VContent(
				VContainer(
					children...,
				),
			),
		).Id("mainapp").Class("overflow-hidden")
	}

	pr.Body = wrapper(
		VNavigationDrawer(
			VToolbar(
				VList(
					VListItem(h.Text("Application")).Class("title"),
				),
			).Flat(true),
			VDivider(),
			VList(
				VListItem(
					VListItemAction(
						VIcon("dashboard"),
					),
					VListItemContent(
						VListItemTitle(h.Text("Home")),
					),
				).On("click", ""),
				VListItem(
					VListItemAction(
						VIcon("question_answer"),
					),
					VListItemContent(
						VListItemTitle(h.Text("About")),
					),
				).On("click", ""),
			).Dense(true).Class("pt-0"),
		).Permanent(true),

		VBtn("show").On("click", "vars.drawer1 = !vars.drawer1"),

		VNavigationDrawer(
			h.Text("Hi"),
		).Temporary(true).
			Attr("v-model", "vars.drawer1").
			Right(true).
			Bottom(true).
			Absolute(true).
			Width(600).Attr("v-init-context-vars", `{drawer1: true}`),

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
				Attr("v-init-context-vars", `{drawer2: false}`),
			AfterLoaded: `setTimeout(function(){ comp.vars.drawer2 = true }, 100)`,
		},
	)
	return
}
