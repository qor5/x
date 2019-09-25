package presets

import (
	ch "github.com/goplaid/x/codehighlight"
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e21_presents"
	"github.com/goplaid/x/docs/utils"
	"github.com/goplaid/x/md"
	. "github.com/theplant/htmlgo"
)

var ItsTheWholeHouse = Components(
	md.Markdown(`
Presets let you config generalized data management UI interface for database. 
It's not a scaffolding to generate source code. But provide more abstract and 
flexible API to enrich features along the way.

`),
	ch.Code(examples.PresetHelloWorldSample).Language("go"),
	md.Markdown(`
And this ~*presets.Builder~ instance is actually also a ~http.Handler~, So that we can mount it
to the http serve mux directly like this:
`),
	ch.Code(examples.MountPresetHelloWorldSample).Language("go"),
	utils.Demo("", e21_presents.PresetsHelloWorldPath+"/customers"),
	md.Markdown(`
With ~r.Model(&Customer{})~:

- It setup the global layout with the left navigation menu
- It setup the listing page with a data table
- It add the new button to create a new record
- It setup the editing and creating form as a right side drawer
- It setup each row of data have a operation menu that you have edit and delete operations
- It setup the global search box, can search the model's all string columns
`),
)
