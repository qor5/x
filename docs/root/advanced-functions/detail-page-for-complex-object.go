package advanced_functions

import (
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e21_presents"
	"github.com/goplaid/x/docs/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
	. "github.com/theplant/htmlgo"
)

var DetailPageForComplexObject = Doc(
	Markdown(`
By default, presets will only generate the listing page, editing page for a model,
It's for simple objects. But for a complicated object with a lots of relationships and connections,
and as the main data model of your system, It's better to have detail page for them. In there
You can add all kinds of operations conveniently.
`),
	ch.Code(examples.PresetsDetailPageTopNotesSample).Language("go"),
	utils.Demo("Presets Detail Page Top Notes", e21_presents.PresetsDetailPageTopNotesPath+"/customers", "e21_presents/detailing.go"),
	Markdown(`
- The name of detailing fields are just a place holder for decide ordering
- ~CellComponentFunc~ customize how the cell display
- ~stripeui~ package create basic components that similar to [Stripe Dashboard](https://dashboard.stripe.com)
- ~stripeui.DataTable~ create a data table, Which the Listing page uses the same component
- ~LoadMoreAt~ will only show for example 2 rows of data, and you can click load more to display all
- ~stripeui.Card~ display a card with toolbar you can setup action buttons
- We reference the new form drawer that ~b.Model(&Note{})~ creates, but hide notes in the menu
`),
	utils.Anchor(H2(""), "Details Info components and actions"),
	Markdown(`
A ~stripeui.DetailInfo~ component is used for display main detail field of the model.
And you can add any actions to the detail page with ease:
`),
	ch.Code(examples.PresetsDetailPageDetailsSample).Language("go"),
	utils.Demo("Presets Detail Page Details", e21_presents.PresetsDetailPageDetailsPath+"/customers", "e21_presents/detailing.go"),
	Markdown(`
- The ~stripui.Card~ Actions links to two event functions: Agree Terms, and Update Details
- Agree Terms show a drawer popup that edit the ~term_agreed_at~ field
- Update Details reuse the edit customer form
`),

	utils.Anchor(H2(""), "More Usage for Data Table"),
	Markdown(`
A ~stripeui.DataTable~ component is very featured rich, Here check out the row expandable example:
`),
	ch.Code(examples.PresetsDetailPageCardsSample).Language("go"),
	utils.Demo("Presets Detail Page Credit Cards", e21_presents.PresetsDetailPageCardsPath+"/customers", "e21_presents/detailing.go"),
	Markdown(`
- ~RowExpandFunc~ config the content when data table row expand
- ~cc.Editing~ setup the fields when edit
- ~cc.Creating~ setup the fields when create
`),
).Title("Detail page for complex object").
	Slug("presets-guide/detail-page-for-complex-object")
