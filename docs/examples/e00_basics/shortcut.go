package e00_basics

// @snippet_begin(ShortCutSample)
import (
	"github.com/goplaid/web"
	. "github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
)

func ShortCutSample(ctx *web.EventContext) (pr web.PageResponse, err error) {
	clickEvent := "locals.count += 1"
	pr.Body = VContainer(
		web.Scope(
			VRow(
				VCol(
					VRow(
						VBtn("count+1").Attr("@click", clickEvent).Class("mr-4"),
						h.Text("Shortcut: enter"),
					).Class("mb-10"),
					VRow(
						VBtn("toggle shortcut").Attr("@click", "locals.shortCutEnabled = !locals.shortCutEnabled"),
					),
				),
				VCol(
					VCard(
						VCardTitle(h.Text("Shortcut Enabled")),
						VCardText().Attr("v-text", "locals.shortCutEnabled"),
					).Class("mb-10"),

					VCard(
						VCardTitle(h.Text("Count")),
						VCardText().Attr("v-text", "locals.count"),
					),
				),
			).Class("mt-10"),
			// Add shortcut for this button. only available when drawer is opened
			web.GlobalEvents().Attr(":filter", `(event, handler, eventName) => locals.shortCutEnabled == true`).Attr("@keydown.enter", clickEvent),
		).Init(`{ shortCutEnabled: true, count: 0 }`).
			VSlot("{ locals }"),
	)
	return
}

// @snippet_end
const ShortCutSamplePath = "/samples/shortcut-sample"
