package e23_vuetify_components_kitchen

// @snippet_begin(VuetifyComponentsKitchen)

import (
	"github.com/goplaid/web"
	"github.com/goplaid/x/docs/utils"
	. "github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
)

var globalCities = []string{"Tokyo", "Hangzhou", "Shanghai"}

type formVals struct {
	Cities1 []string
	Cities2 []string
}

var fv = formVals{
	Cities1: []string{
		"TK",
		"LD",
	},

	Cities2: []string{
		"Hangzhou",
		"Shanghai",
	},
}

func VuetifyComponentsKitchen(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("removeCity", removeCity)
	ctx.Hub.RegisterEventFunc("submit", submit)

	var chips h.HTMLComponents
	for _, city := range globalCities {
		chips = append(chips,
			web.Bind(
				VChip(h.Text(city)).
					Close(true),
			).
				On("click:close").
				EventFunc("removeCity", city),
		)
	}

	pr.Body = VContainer(
		h.H1("Chips delete"),
		chips,

		h.H1("Chips group"),
		utils.PrettyFormAsJSON(ctx),
		VChipGroup(
			VChip(h.Text("Hangzhou")).Value("HZ"),
			VChip(h.Text("Shanghai")).Value("SH").Filter(true),
			VChip(h.Text("Tokyo")).Value("TK").Filter(true),
			VChip(h.Text("New York")).Value("NY"),
			VChip(h.Text("London")).Value("LD"),
		).ActiveClass("indigo darken-3 white--text").
			//Mandatory(true).
			FieldName("Cities1").
			Value(fv.Cities1).
			Multiple(true),
		VBtn("Submit").
			OnClick("submit"),

		VAutocomplete().
			Items(globalCities).
			FieldName("Cities2").
			Value(fv.Cities2),
	)
	return
}

func submit(ctx *web.EventContext) (r web.EventResponse, err error) {
	fv = formVals{}
	ctx.MustUnmarshalForm(&fv)

	r.Reload = true
	return
}

func removeCity(ctx *web.EventContext) (r web.EventResponse, err error) {
	city := ctx.Event.Params[0]
	var newCities []string
	for _, c := range globalCities {
		if c != city {
			newCities = append(newCities, c)
		}
	}
	globalCities = newCities
	r.Reload = true
	return
}

// @snippet_end

const VuetifyComponentsKitchenPath = "/samples/vuetify-components-kitchen"
