package advanced_functions

import (
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e00_basics"
	"github.com/goplaid/x/docs/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var ReloadPageWithAFlash = Doc(
	Markdown(`
The results of an ~web.EventFunc~ could be:

- Go to a new page
- Reload the whole current page
- Refresh part of the current page

Let's demonstrate reload the whole current page:
`),
	ch.Code(examples.ReloadWithFlashSample).Language("go"),
	utils.Demo("Reload Page With a Flash", e00_basics.ReloadWithFlashPath, "e00_basics/reload-with-a-flash.go"),
	Markdown(`
~ctx.Flash~ Object is used to pass data between ~web.EventFunc~ to ~web.PageFunc~ just after the event func is executed. quite similar to [Rails's Flash](https://api.rubyonrails.org/classes/ActionDispatch/Flash.html).
Different is here you can pass in any complicated struct. as long as the page func to use that flash properly.

~er.Reload = true~ tells it will reload the whole page by running page func again, and with the result's body to replace the browser's html content. the event func and page func are executed in one AJAX request in the server.
`),
).Title("Reload Page with a Flash").
	Slug("basics/reload-page-with-a-flash")
