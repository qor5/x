package basics

import (
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e00_basics"
	"github.com/goplaid/x/docs/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var ManipulatePageURLInEventFunc = Doc(
	Markdown(`
Encode page state into query strings in url is useful. because user can paste the link to another person,
That can open the page to the exact state of the page being sent, Not the initial state of the page.

For example:
`),
	ch.Code(examples.MultiStatePageSample).Language("go"),
	utils.Demo("Manipulate Page URL In Event Func", e00_basics.MultiStatePagePath, "e00_basics/manipulate-page-url.go"),
	Markdown(`
This page have several state that encoded in the url:

- Page title have a default value, but if provided with a ~title~ query string, it will use that value
- The edit panel can be open, or closed based on having the ~panel~ query string or not

~web.Location(url.Values{"panel": []string{"1"}}).MergeQuery(true)~ means it will do a push state request to current page, with panel query string panel=1. 
~MergeQuery~ means that it will not touch other query strings like ~title=1~ we mentioned above.

In ~update5~ event func, which is when you click the update button after open the panel, ~web.Location(url.Values{"panel": []string{""}}).MergeQuery(true)~ basically removes the query string panel=1, and won't touch any other query strings.

Don't have to be in event func to use push state query, can use a simple ~web.Bind~ to directly change the query string like:

~~~go
web.Bind(
	A().Text("change page title").Href("javascript:;"),
).Queries(url.Values{"title": []string{"Hello"}}),
~~~

This don't have ~.MergeQuery(true)~, So it will replace the whole query string to only ~title=Hello~ 

`),
).Title("Manipulate Page URL in Event Func").
	Slug("basics/manipulate-page-url-in-event-func")
