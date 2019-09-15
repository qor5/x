package basics

import (
	"github.com/goplaid/x/md"
	. "github.com/theplant/htmlgo"
)

var PartialRefreshWithPortal = Components(
	md.Markdown(`
The results of an ~web.EventFunc~ could be:

- Go to a new page
- Reload the whole current page
- Refresh part of the current page

Now let's demonstrate refresh part of the current page: 
`),
)
