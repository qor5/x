package presets

import (
	ch "github.com/goplaid/x/codehighlight"
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e21_presents"
	"github.com/goplaid/x/docs/utils"
	"github.com/goplaid/x/md"
	. "github.com/theplant/htmlgo"
)

var ListingCustomizations = Components(
	md.Markdown(`
We get a default listing page with default columns, But default columns from database
columns rarely fit the needs for any real application.

How do we change the columns of the table? and how to we change the content display of a columns?
`),
	ch.Code(examples.PresetListingCustomization01Sample).Language("go"),
	utils.Demo("", e21_presents.PresetsListingCustomization01PATH+"/customers"),
	md.Markdown(`
What we did with above code:

- We added a new field to listing table that not exists on the struct ~Customer~
- We define the listing display for the listing table by using the ~Td()~ and fetch the company data from a different table with associated column value
- We link the company name in the listing to link the edit drawer of company
- We limit the edit drawer field to only have ~Name~ and ~CompanyID~
- We made the ~CompanyID~ field a vuetify ~VSelect~ component
- We add the companies as a new section, that you can manage
`),
)
