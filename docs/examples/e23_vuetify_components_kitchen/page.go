package e23_vuetify_components_kitchen

// @snippet_begin(VuetifyComponentsKitchen)

import (
	"github.com/goplaid/web"
	. "github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
)

var globalCities = []string{"Tokyo", "Hangzhou", "Shanghai"}

func VuetifyComponentsKitchen(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("removeCity", removeCity)

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
		chips,
	)
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
