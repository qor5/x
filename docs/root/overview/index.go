package overview

import (
	ch "github.com/sunfmin/bran/codehighlight"
	"github.com/sunfmin/bran/docs/samples"
	"github.com/sunfmin/bran/docs/utils"
	"github.com/sunfmin/bran/md"
	"github.com/sunfmin/bran/ui"
	. "github.com/theplant/htmlgo"
)

func Index(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	pr.Schema = Components(
		utils.Anchor(H1(""), "The Type-safe HTML Builder in Go"),
		md.Markdown(`
Although Go can't do flexible builder syntax like [Kotlin](https://kotlinlang.org/docs/reference/type-safe-builders.html) does, 
But it can also do quite well.

Consider the following code:
`),
		ch.Code(samples.TypeSafeBuilderSample),
		md.Markdown(`
It's basically assembled what Kotlin can do, Also is legitimate Go code.
`),
		utils.Demo("Check the rendered html of above code", samples.TypeSafeBuilderSamplePath),
	)
	return
}
