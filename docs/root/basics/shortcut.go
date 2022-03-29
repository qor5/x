package basics

import (
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var ShortCut = Doc(
	Markdown(`
To add keyboard shortcut to a button:

Trigger the event by [GlobalEvents](https://www.npmjs.com/package/vue-global-events). 
You can configure your own keyboard event like ~~~@keyup.ctrl.enter~~~ to trigger the event.

Also you can setup the ~~~filter~~~ function to limit when this event can be triggered by shortcut. 
In the example, the event would only be triggered when Drawer is opened.
`),

	ch.Code(`
clickEvent := web.Plaid().
				EventFunc(eventAskForReview).
				Query("id", fmt.Sprint(dv.ID)).
				Go()

// Add shortcut for this button. only available when drawer is opened
web.GlobalEvents().Attr(":filter", "(event, handler, eventName) => 
	vars.presetsRightDrawer == true").Attr("@keyup.ctrl.enter", clickEvent),
`),
).Slug("basics/shortcut").Title("Keyboard Shortcut")
