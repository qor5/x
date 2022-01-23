package basics

import (
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e00_basics"
	"github.com/goplaid/x/docs/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
	. "github.com/theplant/htmlgo"
)

var EventHandling = Doc(
	Markdown(`Event Handling provides multiple ways to handle interaction events.`),

	utils.Anchor(H2(""), "API"),
	Markdown(`Open another page.`),
	H3("Example"),
	ch.Code(examples.EventHandlingURLSample).Language("go"),
	utils.Demo("Event Handling", e00_basics.EventHandlingPagePath+"?api=url", "e00_basics/event-handling.go#L23-L28"),

	utils.Anchor(H2(""), "PushState"),
	Markdown(`Open another page and also changing the window location.`),
	H3("Example"),
	ch.Code(examples.EventHandlingPushStateSample).Language("go"),
	utils.Demo("Event Handling", e00_basics.EventHandlingPagePath+"?api=pushstate", "e00_basics/event-handling.go#33-L38"),

	utils.Anchor(H2(""), "EventFunc"),
	Markdown(`Register an event func and call it when the event is triggered.`),
	H3("Example"),
	ch.Code(examples.EventHandlingEventFuncSample).Language("go"),
	utils.Demo("Event Handling", e00_basics.EventHandlingPagePath+"?api=eventfunc", "e00_basics/event-handling.go#45-L54"),

	utils.Anchor(H2(""), "Reload"),
	Markdown(`Refresh the page.`),
	H3("Example"),
	ch.Code(examples.EventHandlingReloadSample).Language("go"),
	utils.Demo("Event Handling", e00_basics.EventHandlingPagePath+"?api=reload", "e00_basics/event-handling.go#58-L62"),
).Title("Event Handling").Slug("basics/event-handling")
