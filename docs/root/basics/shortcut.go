package basics

import (
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e00_basics"
	"github.com/goplaid/x/docs/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var ShortCut = Doc(
	Markdown(`
To add keyboard shortcut to a button:

Trigger the event by [GlobalEvents](https://www.npmjs.com/package/vue-global-events). 
You can configure your own keyboard event like ~@keyup.ctrl.enter~ to trigger the event.

Also you can setup the ~filter~ function to limit when this event can be triggered by shortcut. 
In the example, the event would only be triggered when ~locals.shortCutEnabled~ is opened.
`),

	ch.Code(examples.ShortCutSample).Language("go"),
	utils.Demo("Shortcut", e00_basics.ShortCutSamplePath, "e00_basics/shortcut.go"),
).Slug("basics/shortcut").Title("Keyboard Shortcut")
