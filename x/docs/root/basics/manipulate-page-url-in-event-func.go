package basics

import (
	ch "github.com/goplaid/x/codehighlight"
	"github.com/goplaid/x/docs/samples"
	"github.com/goplaid/x/docs/utils"
	"github.com/goplaid/x/md"
	. "github.com/theplant/htmlgo"
)

var ManipulatePageURLInEventFunc = Components(
	md.Markdown(`
Encode page state into query strings in url is useful. because user can paste the link to another person,
That can open the page to the exact state of the page being sent, Not the initial state of the page.

For example:
`),
	ch.Code(samples.MultiStatePageSample),
	utils.Demo("", samples.MultiStatePagePath),
	md.Markdown(`
This page have several state that encoded in the url:

- Page title have a default value, but if provided with a ~title~ query string, it will use that value
- The edit panel can be open, or closed based on having the ~panel~ query string or not

~web.PushState(url.Values{"panel": []string{"1"}}).MergeQuery(true)~ means it will do a push state request to current page, with panel query string panel=1. 
~MergeQuery~ means that it will not touch other query strings like ~title=1~ we mentioned above.

In ~update5~ event func, which is when you click the update button after open the panel, ~web.PushState(url.Values{"panel": []string{""}}).MergeQuery(true)~ basically removes the query string panel=1, and won't touch any other query strings.

Don't have to be in event func to use push state query, can use a simple ~web.Bind~ to directly change the query string like:

~~~go
web.Bind(
	A().Text("change page title").Href("javascript:;"),
).PushStateQuery(url.Values{"title": []string{"Hello"}}),
~~~

This don't have ~.MergeQuery(true)~, So it will replace the whole query string to only ~title=Hello~ 

`),
)
