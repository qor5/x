package basics

import (
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e21_presents"
	"github.com/goplaid/x/docs/utils"
	"github.com/theplant/docgo/ch"

	. "github.com/theplant/docgo"
)

var Filter = Doc(
	Markdown(`

To add a basic filter to the list page

For example:
`),
	ch.Code(examples.FilterSample).Language("go"),
	utils.Demo("Basic filter", e21_presents.PresetsBasicFilterPath+"/customers", "e21_presents/filter.go"),
	Markdown(`
	Call ~FilterDataFunc~ on a ~ListingBuilder~
`),
).Title("Filters").
	Slug("basics/filter")
