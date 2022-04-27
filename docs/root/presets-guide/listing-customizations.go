package presets_guide

import (
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e21_presents"
	"github.com/goplaid/x/docs/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
	. "github.com/theplant/htmlgo"
)

var ListingCustomizations = Doc(
	Markdown(`
We get a default listing page with default columns, But default columns from database
columns rarely fit the needs for any real application.

`),
	utils.Anchor(H2(""), "Change List Columns and Component of Field"),
	Markdown(`
Here is how do we change the columns of the list and how to we change the content display of a columns.
`),
	ch.Code(examples.PresetsListingCustomizationFieldsSample).Language("go"),
	utils.Demo("Presets Listing Customization Fields", e21_presents.PresetsListingCustomizationFieldsPath+"/customers", "e21_presents/listing.go"),
	Markdown(`
What we did with above code:

- Added a new field to listing table that not exists on the struct ~Customer~
- Define the listing display for the listing table by using the ~Td()~ and fetch the company data from a different table with associated column value
- Link the company name in the listing to link the edit drawer of company
- Limit the edit drawer field to only have ~Name~ and ~CompanyID~
- Made the ~CompanyID~ field a vuetify ~VSelect~ component
- Add companies as a new navigation item, that you can manage companies data
- ~.SearchColumns("name", "email")~ configure the top navigation search box searches which columns with sql like operation
`),

	utils.Anchor(H2(""), "Filters Panel"),
	Markdown(`
Here we continue to add filters for the list
`),
	ch.Code(examples.PresetsListingCustomizationFiltersSample).Language("go"),
	utils.Demo("Presets Listing Filters", e21_presents.PresetsListingCustomizationFiltersPath+"/customers", "e21_presents/listing.go"),
	Markdown(`
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

	utils.Anchor(H2(""), "Filter Tabs"),
	Markdown(`
Filter tabs is based on Filters configuration. But display as tabs above the list,
You can think it as a short cut that used very frequently to filter something instead of 
use the pop up panel of filter.
`),
	ch.Code(examples.PresetsListingCustomizationTabsSample).Language("go"),
	utils.Demo("Presets Listing Filter Tabs", e21_presents.PresetsListingCustomizationTabsPath+"/customers", "e21_presents/listing.go"),
	Markdown(`
~Query~ string name must be from the Filter's item configuration key field.
`),

	utils.Anchor(H2(""), "Bulk Actions"),
	Markdown(`
Bulk actions makes the list row show checkboxes, and you can select one or many rows,
Later do an bulk update data for all of them.

Here is how to use it:
`),
	ch.Code(examples.PresetsListingCustomizationBulkActionsSample).Language("go"),
	utils.Demo("Presets Listing Bulk Actions", e21_presents.PresetsListingCustomizationBulkActionsPath+"/customers", "e21_presents/listing.go"),
	Markdown(`
- ~ComponentFunc~ of the bulk action configure the component that will show to user to input after user clicked the bulk action button
- ~UpdateFunc~ configure the logic that the bulk action execute
`),

	utils.Anchor(H2(""), "Actions"),
	Markdown(`
Action allows you to update database or do anything else with a button.

Here is how to use it:
`),
	ch.Code(examples.PresetsListingCustomizationActionsSample).Language("go"),
	utils.Demo("Presets Listing Actions", e21_presents.PresetsListingCustomizationActionsPath+"/customers", "e21_presents/listing.go"),
	Markdown(`
- ~ComponentFunc~ and ~UpdateFunc~ does the same thing in Actions as in Bulk Actions.
`),
).Title("Listing Customizations").
	Slug("presets-guide/listing-customizations")
