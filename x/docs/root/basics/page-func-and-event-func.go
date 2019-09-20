package basics

import (
	ch "github.com/goplaid/x/codehighlight"
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e00_basics"
	"github.com/goplaid/x/docs/utils"
	"github.com/goplaid/x/md"
	. "github.com/theplant/htmlgo"
)

var PageFuncAndEventFunc = Components(
	md.Markdown(`
~PageFunc~ is used to build a web page, ~EventFunc~ is called when user interact with the page, For example button or link clicks.
`),
	ch.Code(examples.PageFuncAndEventFuncDefinition).Language("go"),
	md.Markdown(`~ctx.Hub.RegisterEventFunc~ is used to connect an ~EventFunc~ to a ~PageFunc~ and also provide an event name, 
to be used by ~web.Bind~ to attach to an html element.`),
	ch.Code(examples.EventFuncHubDefinition).Language("go"),
	md.Markdown("Here is a hello world with more interactions. User click the button will reload the page with latest time"),
	ch.Code(examples.HelloWorldReloadSample).Language("go"),
	utils.Demo("", e00_basics.HelloWorldReloadPath),
	md.Markdown("Note that you have to mount each `PageFunc` to http.ServeMux with a path to be able to access the `PageFunc` in your browser"),
	ch.Code(examples.HelloWorldReloadMuxSample1).Language("go"),
	md.Markdown("`wb.Page` convert any `PageFunc` into `http.Handler`, outside you can wrap any middleware that can use on Go standard `http.Handler`."),
	md.Markdown(`In case you don't know what is a http.Handler middleware, 
It's a function that takes http.Handler as input, might also with other parameters, 
And also return a new http.Handler, 
[gziphandler](https://github.com/nytimes/gziphandler) is an example.`),

	md.Markdown(`But What the heck is ~demoLayout~ there?
Well it's a ~PageFunc~ middleware. That takes an ~PageFunc~ as input, 
wrap it's ~PageResponse~ with layout html and return a new ~PageFunc~. 
If you follow the code to write your own ~PageFunc~, 
The button click might not work without this. 
Since there is no layout to import needed javascript to make this work. 
continue to next page to checkout how to add necessary javascript, css etc to make the demo work.`),
)
