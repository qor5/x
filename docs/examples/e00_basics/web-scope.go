package e00_basics

import (
	"github.com/goplaid/web"
	. "github.com/goplaid/x/vuetify"
	. "github.com/theplant/htmlgo"
)

//@snippet_begin(WebScopeUseLocalsSample1)
func UseLocals(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = VCard(
		VBtn("Test Can not change other scope").Attr("@click", `locals.btnLabel = "YES"`),
		web.Scope(
			VCard(
				VBtn("").
					Attr("v-text", "locals.btnLabel").
					Attr("@click", `
if (locals.btnLabel == "Add") {
	locals.items.push({text: "B", icon: "done"}); 
	locals.btnLabel = "Remove";
} else {
	locals.items.pop({text: "B", icon: "done"}); 
	locals.btnLabel = "Add";
}`),

				VList(
					VSubheader(
						Text("REPORTS"),
					),
					VListItemGroup(
						VListItem(
							VListItemIcon(
								VIcon("").Attr("v-text", "item.icon"),
							),
							VListItemContent(
								VListItemTitle().Attr("v-text", "item.text"),
							),
						).Attr("v-for", "(item, i) in locals.items").
							Attr("x-bind:key", "i"),
					).Attr("v-model", "locals.selectedItem").
						Attr("color", "primary"),
				).Attr("dense", ""),
			).Class("mx-auto").
				Attr("max-width", "300").
				Attr("tile", ""),
		).Init(`{ selectedItem: 1, btnLabel:"Add", items: [{text: "A", icon: "clock"}]}`).
			VSlot("{ locals }"),
	)
	return
}

//@snippet_end

const WebScopeUseLocalsPagePath = "/samples/web-scope-use-locals"
