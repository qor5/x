package basics

import (
	ch "github.com/goplaid/x/codehighlight"
	"github.com/goplaid/x/docs/samples"
	"github.com/goplaid/x/docs/utils"
	"github.com/goplaid/x/md"
	. "github.com/theplant/htmlgo"
)

var PartialRefreshWithPortal = Components(
	md.Markdown(`
As said before, The results of an ~web.EventFunc~ could be:

- Go to a new page
- Reload the whole current page
- Refresh part of the current page

We have covered two. Now let's demonstrate refresh part of the current page: 
`),
	ch.Code(samples.PartialUpdateSample),
	utils.Demo("", samples.PartialUpdatePagePath),
	md.Markdown(`
~web.Portal().Name("part1")~ Place a placeholder inside you page, and append ~web.PortalUpdate~ to ~er.UpdatePortals~ to update the portal with that name.
Multiple portal can be updated at the same time.
`),
	utils.Anchor(H2(""), "Load Portal in separate AJAX request"),
	md.Markdown(`
With ~web.Portal~, We can also load the portal with a separate AJAX request after page load.
It is useful for the type of the content is not that important to the page, But load them are
quite heavy. Like related products of a product detail page of a ECommerce site.
`),
	ch.Code(samples.PartialReloadSample),
	utils.Demo("", samples.PartialReloadPagePath),
	md.Markdown(`
It is not only load the portal in separate AJAX request, Also you can reload it with ease ~er.ReloadPortals = []string{"related_products"}~ in an event func.
`),
)
