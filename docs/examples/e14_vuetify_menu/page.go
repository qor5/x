package e14_vuetify_menu

// @snippet_begin(VuetifyMenuSample)

import (
	"github.com/goplaid/web"
	"github.com/goplaid/x/docs/utils"
	. "github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
)

type formData struct {
	EnableMessages bool
	EnableHints    bool
}

var globalFavored bool

const favoredIconPortalName = "favoredIcon"

func HelloVuetifyMenu(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("submit", submit)
	ctx.Hub.RegisterEventFunc("toggleFavored", toggleFavored)

	var fv formData
	err = ctx.UnmarshalForm(&fv)
	if err != nil {
		return
	}

	pr.Body = VContainer(
		utils.PrettyFormAsJSON(ctx),

		VMenu(
			web.Slot(
				VBtn("Menu as Popover").
					On("click", "vars.myMenuShow = true").
					Dark(true).
					Color("indigo"),
			).Name("activator"),

			VCard(
				VList(
					VListItem(
						VListItemAvatar(
							h.Img("https://cdn.vuetifyxjs.com/images/john.jpg").Alt("John"),
						),
						VListItemContent(
							VListItemTitle(h.Text("John Leider")),
							VListItemSubtitle(h.Text("Founder of Vuetify.js")),
						),
						VListItemAction(
							web.Portal(
								favoredIcon(),
							).Name(favoredIconPortalName),
						),
					),
				),
				VDivider(),
				VList(
					VListItem(
						VListItemAction(
							VSwitch().Color("purple").
								FieldName("EnableMessages").
								InputValue(fv.EnableMessages),
						),
						VListItemTitle(h.Text("Enable messages")),
					),
					VListItem(
						VListItemAction(
							VSwitch().Color("purple").
								FieldName("EnableHints").
								InputValue(fv.EnableHints),
						),
						VListItemTitle(h.Text("Enable hints")),
					),
				),

				VCardActions(
					VSpacer(),
					VBtn("Cancel").Text(true).
						On("click", "vars.myMenuShow = false"),
					VBtn("Save").Color("primary").
						Text(true).OnClick("submit"),
				),
			),
		).CloseOnContentClick(false).
			NudgeWidth(200).
			OffsetY(true).
			Attr("v-model", "vars.myMenuShow"),
	).Attr(web.InitContextVars, `{myMenuShow: false}`)

	return
}

func favoredIcon() h.HTMLComponent {
	color := ""
	if globalFavored {
		color = "red"
	}

	return VBtn("").Icon(true).Children(
		VIcon("favorite").Color(color),
	).OnClick("toggleFavored")
}

func toggleFavored(ctx *web.EventContext) (er web.EventResponse, err error) {
	globalFavored = !globalFavored
	er.UpdatePortals = append(er.UpdatePortals, &web.PortalUpdate{
		Name: favoredIconPortalName,
		Body: favoredIcon(),
	})
	return
}

func submit(ctx *web.EventContext) (er web.EventResponse, err error) {
	er.Reload = true
	er.VarsScript = "vars.myMenuShow = false"
	return
}

// @snippet_end
const HelloVuetifyMenuPath = "/samples/hello-vuetify-menu"
