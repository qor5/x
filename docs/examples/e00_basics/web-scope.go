package e00_basics

import (
	"github.com/goplaid/web"
	"github.com/goplaid/x/docs/utils"
	. "github.com/goplaid/x/vuetify"
	. "github.com/theplant/htmlgo"
)

// @snippet_begin(WebScopeUseLocalsSample1)
func UseLocals(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = VCard(
		VBtn("Test Can Not Change Other Scope").Attr("@click", `locals.btnLabel = "YES"`),
		web.Scope(
			VCard(
				VBtn("").
					Attr("v-text", "locals.btnLabel").
					Attr("@click", `
if (locals.btnLabel == "Add") {
	locals.items.push({text: "B", icon: "done"}); 
	locals.btnLabel = "Remove";
} else {
	locals.items.pop(); 
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

var UseLocalsPB = web.Page(UseLocals)

// @snippet_end

// @snippet_begin(WebScopeUsePlaidFormSample1)
var materialID, materialName, rawMaterialID, rawMaterialName, countryID, countryName, productName string

func UsePlaidForm(ctx *web.EventContext) (pr web.PageResponse, err error) {

	pr.Body = Div(
		H3("Form Content"),
		utils.PrettyFormAsJSON(ctx),

		Div(
			Div(
				Fieldset(
					Legend("Product Form"),
					Div(
						Label("Product Name"),
						Input("").Value(productName).Type("text").Attr(web.VFieldName("ProductName")...),
					),
					Div(
						Label("Material ID"),
						Input("").Value(materialID).Type("text").Disabled(true).Attr(web.VFieldName("MaterialID")...),
					),

					web.Scope(
						Fieldset(
							Legend("Material Form"),

							Div(
								Label("Material Name"),
								Input("").Value(materialName).Type("text").Attr(web.VFieldName("MaterialName")...),
							),
							Div(
								Label("Raw Material ID"),
								Input("").Value(rawMaterialID).Type("text").Disabled(true).Attr(web.VFieldName("RawMaterialID")...),
							),
							web.Scope(
								Fieldset(
									Legend("Raw Material Form"),

									Div(
										Label("Raw Material Name"),
										Input("").Value(rawMaterialName).Type("text").Attr(web.VFieldName("RawMaterialName")...),
									),

									Button("Send").Style(`background: orange;`).Attr("@click", web.POST().EventFunc("updateValue").Go()),
								).Style(`background: orange;`),
							).VSlot("{ plaidForm }"),

							Button("Send").Style(`background: brown;`).Attr("@click", web.POST().EventFunc("updateValue").Go()),
						).Style(`background: brown;`),
					).VSlot("{ plaidForm }"),

					Div(
						Label("Country ID"),
						Input("").Value(countryID).Type("text").Disabled(true).Attr(web.VFieldName("CountryID")...),
					),

					web.Scope(
						Fieldset(
							Legend("Country Of Origin Form"),

							Div(
								Label("Country Name"),
								Input("").Value(countryName).Type("text").Attr(web.VFieldName("CountryName")...),
							),

							Button("Send").Style(`background: red;`).Attr("@click", web.POST().EventFunc("updateValue").Go()),
						).Style(`background: red;`),
					).VSlot("{ plaidForm }"),

					Div(
						Button("Send").Style(`background: grey;`).Attr("@click", web.POST().EventFunc("updateValue").Go())),
				).Style(`background: grey;`)),
		).Style(`width:600px;`),
	)

	return
}

func updateValue(ctx *web.EventContext) (er web.EventResponse, err error) {
	ctx.R.ParseForm()
	if v := ctx.R.Form.Get("ProductName"); v != "" {
		productName = v
	}
	if v := ctx.R.Form.Get("MaterialName"); v != "" {
		materialName = v
		materialID = "66"
	}
	if v := ctx.R.Form.Get("RawMaterialName"); v != "" {
		rawMaterialName = v
		rawMaterialID = "88"
	}
	if v := ctx.R.Form.Get("CountryName"); v != "" {
		countryName = v
		countryID = "99"
	}
	er.Reload = true
	return
}

var UsePlaidFormPB = web.Page(UsePlaidForm).
	EventFunc("updateValue", updateValue)

// @snippet_end

const WebScopeUseLocalsPagePath = "/samples/web-scope-use-locals"
const WebScopeUsePlaidFormPagePath = "/samples/web-scope-use-plaid-form"
