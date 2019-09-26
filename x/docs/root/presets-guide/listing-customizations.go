package presets_guide

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

How do we change the columns of the list? and how to we change the content display of a columns?
`),
	ch.Code(examples.PresetsListingCustomizationFieldsSample).Language("go"),
	utils.Demo("", e21_presents.PresetsListingCustomizationFieldsPath+"/customers"),
	md.Markdown(`
What we did with above code:

- Added a new field to listing table that not exists on the struct ~Customer~
- Define the listing display for the listing table by using the ~Td()~ and fetch the company data from a different table with associated column value
- Link the company name in the listing to link the edit drawer of company
- Limit the edit drawer field to only have ~Name~ and ~CompanyID~
- Made the ~CompanyID~ field a vuetify ~VSelect~ component
- Add companies as a new navigation item, that you can manage companies data
`),

	utils.Anchor(H2(""), "Add Filters to the list"),
	md.Markdown(`
Here we continue to add filters for the list
`),
	ch.Code(examples.PresetsListingCustomizationFiltersSample).Language("go"),
	utils.Demo("", e21_presents.PresetsListingCustomizationFiltersPath+"/customers"),
	md.Markdown(`
~FilterDataFunc~ of ~presets.ListingBuilder~ setup to have the filter menu or not.
And how it will combine the sql conditions when doing query. the filter menu will 
change the url query strings with the filter values, and for date type in url query
string it uses unix epoch int value. So the sql condition has to convert the database
column data to unix epoch in order to compare with the value in url query string.

Current we support these types

- ~ItemTypeDate~: set it as a date filter item, which have many switches to support date and date range
- ~ItemTypeNumber~: set it to a number filter item, which have switches to support number and number range
- ~ItemTypeString~: set it to a string filter item, which have contains, and match exactly
- ~ItemTypeSelect~: set it to a select filter item, which have a options of values for selection
`),
)
