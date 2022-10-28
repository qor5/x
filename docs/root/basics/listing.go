package basics

import (
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/example_basics"
	"github.com/goplaid/x/docs/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var Listing = Doc(
	Markdown(`
By the [1 Minute Quick Start](/getting-started/one-minute-quick-start.html), We get a default listing page with default columns, But default columns from database columns rarely fit the needs for any real application. Here we will introduce common customizations on the list page.

- Configure fields that displayed on the page
- Modify the display value
- Display a virtual field
- Default scope
- Extend the dot menu

There would be a runable example at the last.

## Configure fields that displayed on the page
Suppose we added a new model called ~Category~, the ~Post~ belongs to ~Category~. Then we want to display ~CategoryID~ on the list page.
`),

	ch.Code(`
type Post struct {
	ID    uint
	Title string
	Body  string

	CategoryID uint

	UpdatedAt time.Time
	CreatedAt time.Time
}

type Category struct {
	ID   uint
	Name string

	UpdatedAt time.Time
	CreatedAt time.Time
}

postModelBuilder.Listing("ID", "Title", "Body", "CategoryID")
`),

	Markdown(`
## Modify the display value
To display the category name rather than category id in the post listing page. The ~ComponentFunc~ would do the work.
The ~obj~ is the ~Post~ record, and ~field~ is the ~CategoryID~ field of this ~Post~ record. You can get the value by ~field.Value(obj)~ function.
`),

	ch.Code(`postModelBuilder.Listing().Field("CategoryID").Label("Category").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
	c := models.Category{}
	cid, _ := field.Value(obj).(uint)
	if err := db.Where("id = ?", cid).Find(&c).Error; err != nil {
		// ignore err in the example
	}
	return h.Td(h.Text(c.Name))
})
`).Language("go"),

	Markdown(`
## Display virtual fields
`),
	ch.Code(`postModelBuilder.Listing("ID", "Title", "Body", "CategoryID", "VirtualValue")
postModelBuilder.Listing().Field("VirtualField").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
	return h.Td(h.Text("virtual field"))
})
`),

	Markdown(`
## DefaultScope
If we want to display ~Post~ with ~category_id~ only. Use the ~Listing().Searcher~ to apply SQL conditions.
`),
	ch.Code(`postModelBuilder.Listing().Searcher = func(model interface{}, params *presets.SearchParams, ctx *web.EventContext) (r interface{}, totalCount int, err error){
	qdb := db.Where("category_id != 0")
	return gorm2op.DataOperator(qdb).Search(model, params, ctx)
}
`),

	Markdown(`
## Extend the dot menu
You can extend the dot menu by calling the ~RowMenuItem~ function. If you want to overwrite the default ~Edit~ and ~Delete~ link, you can pass the items you wanted to ~Listing().RowMenu()~
`),
	ch.Code(`rmn := postModelBuilder.Listing().RowMenu()
rmn.RowMenuItem("Show").ComponentFunc(func(obj interface{}, id string, ctx *web.EventContext) h.HTMLComponent {
	return h.Text("Fake Show")
})
`),

	Markdown(`
## Full Example
`),
	ch.Code(examples.PresetsListingSample).Language("go"),
	utils.Demo("Presets Listing Customization Fields", example_basics.ListingSamplePath+"/posts", "example_basics/listing.go"),
).Title("Listing Page").
	Slug("basics/listing")
