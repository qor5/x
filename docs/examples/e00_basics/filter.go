package e00_basics

//@snippet_begin(FilterSample)
import (
	"github.com/goplaid/web"
	"github.com/goplaid/x/presets"
	"github.com/goplaid/x/vuetifyx"
)

type Post struct {
	Name   string
	Status string
}

func PresetsBasicFilter(b *presets.Builder) {
	// create a ModelBuilder
	videoBuilder := b.Model(&Post{})

	// get its ListingBuilder
	listing := videoBuilder.Listing()

	// Call FilterDataFunc
	listing.FilterDataFunc(func(ctx *web.EventContext) vuetifyx.FilterData {
		// Prepare filter options, it is a two dimension array: [][]string{"text", "value"}
		options := []*vuetifyx.SelectItem{{
			Text:  "Draft",
			Value: "draft",
		}}

		return []*vuetifyx.FilterItem{
			{
				Key:      "status",
				Label:    "Status",
				ItemType: vuetifyx.ItemTypeString,
				// %s is the condition. e.g. >, >=, =, <, <=, like，
				// ？ is the value of of selected option
				SQLCondition: `status %s ?`,
				Options:      options,
			},
		}
	})
}

//@snippet_end

const PresetsBasicFilterPath = "/samples/basic_filter"
