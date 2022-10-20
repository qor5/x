package basics

import (
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var Listing = Doc(
	Markdown(`
By the [1 Minute Quick Start](/getting-started/one-minute-quick-start.html), We get a default listing page with default columns, But default columns from database columns rarely fit the needs for any real application. Here we will introduce common customizations on the list page.

- Configure fields that displayed on the page
- Modify the display value of a cell
- Display a virtual field
- Add default query conditions to the dataset
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
## Modify the display value of a cell
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
	ch.Code(`
	postModelBuilder.Listing("ID", "Title", "Body", "CategoryID", "VirtualValue")
	postModelBuilder.Listing().Field("VirtualField").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		return h.Td(h.Text("virtual field"))
	})
`),
).Title("Listing").
	Slug("basics/listing-customizations")
