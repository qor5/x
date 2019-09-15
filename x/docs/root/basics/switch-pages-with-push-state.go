package basics

import (
	ch "github.com/goplaid/x/codehighlight"
	"github.com/goplaid/x/docs/samples"
	"github.com/goplaid/x/docs/utils"
	"github.com/goplaid/x/md"
	. "github.com/theplant/htmlgo"
)

var SwitchPagesWithPushState = Components(
	md.Markdown(`Ways that page transition (between ~web.PageFunc~) in GoPlaid web app:

- Use a link to a new page by url
- Use a button etc to trigger post to an ~web.EventFunc~ that do some logic, then go to a new page

These transitions can use traditional link, Or more modern ajax with [push state](https://developer.mozilla.org/en-US/docs/Web/API/History_API#Examples)

Let's check this example:
`),
	ch.Code(samples.PageTransitionSample),
	utils.Demo("", samples.Page1Path),
	md.Markdown(`
When running the above demo, If you check Chrome Developer Tools about Network requests, 
You will see that the PushState link and the Button is actually doing an AJAX request to the other page.

Look like this:
~~~
POST /samples/page_2?__execute_event__=__reload__ HTTP/1.1
~~~

The result is an JSON object with page's html inside. 
~__reload__~ is another ~web.EventFunc~ that is the same as ~doAction2~, 
But it is default added to every ~web.PageFunc~. So that the web page can
both respond to normal HTTP request from Browser, Search Engine, Or from
other pages in the same web app that can do push state link.
`),
	utils.Anchor(H2(""), "Summary"),
	md.Markdown(`
- Write once with PageFunc, you get both normal html page render, and AJAX JSON page render
- EventFunc is always called with AJAX request, and you can return to a different page, or rerender the current page, 
- EventFunc is not wrapped with layout function.
- EventFunc is used to do data operations, triggered by page's html element. and it's result can be:
	1. Reload the whole current page with new updated data
	2. Update partial of the current page
	3. Go to a different page after data updated

Next we will talk about how to update partial of the current page.
`),
)
