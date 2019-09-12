package overview

import (
	"github.com/goplaid/web"
	ch "github.com/goplaid/x/codehighlight"
	samples2 "github.com/goplaid/x/docs/samples"
	utils2 "github.com/goplaid/x/docs/utils"
	"github.com/goplaid/x/md"
	. "github.com/theplant/htmlgo"
)

func Index(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Schema = Components(
		utils2.Anchor(H1(""), "The Type-safe HTML Builder in Go"),
		md.Markdown(`
Although Go can't do flexible builder syntax like [Kotlin](https://kotlinlang.org/docs/reference/type-safe-builders.html) does,
But it can also do quite well.

Consider the following code:
`),
		ch.Code(samples2.TypeSafeBuilderSample),
		md.Markdown(`
It's basically assembled what Kotlin can do, Also is legitimate Go code.
`),
		utils2.Demo("Check the rendered html of above code", samples2.TypeSafeBuilderSamplePath),
	)
	return
}
