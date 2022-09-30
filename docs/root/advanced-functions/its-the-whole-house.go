package advanced_functions

import (
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e21_presents"
	"github.com/goplaid/x/docs/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var ItsTheWholeHouse = Doc(
	Markdown(`
Presets let you config generalized data management UI interface for database.
It's not a scaffolding to generate source code. But provide more abstract and
flexible API to enrich features along the way.

`),
	ch.Code(examples.PresetHelloWorldSample).Language("go"),
	Markdown(`
And this ~*presets.Builder~ instance is actually also a ~http.Handler~, So that we can mount it
to the http serve mux directly like this:
`),
	ch.Code(examples.MountPresetHelloWorldSample).Language("go"),
	utils.Demo("Presets Hello World", e21_presents.PresetsHelloWorldPath+"/customers", "e21_presents/listing.go"),
	Markdown(`
With ~r.Model(&Customer{})~:

- It setup the global layout with the left navigation menu
- It setup the listing page with a data table
- It add the new button to create a new record
- It setup the editing and creating form as a right side drawer
- It setup each row of data have a operation menu that you have edit and delete operations
- It setup the global search box, can search the model's all string columns
`),
).Title("Not just scaffolding, it's the whole house").
	Slug("presets-guide/its-the-whole-house")
