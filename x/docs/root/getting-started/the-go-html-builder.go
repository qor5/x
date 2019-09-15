package getting_started

import (
	ch "github.com/goplaid/x/codehighlight"
	"github.com/goplaid/x/docs/samples"
	"github.com/goplaid/x/docs/utils"
	"github.com/goplaid/x/md"
	. "github.com/theplant/htmlgo"
)

var TheGoHTMLBuilder = Components(
	md.Markdown(`
Like at the beginning we said, That we don't use interpretation template language to generate html page.
We think they are error prone, hard to refactor, and difficult to abstract out to component.
We like to use standard Go code. the library [htmlgo](https://github.com/theplant/htmlgo) is just for that.

Although Go can't do flexible builder syntax like [Kotlin](https://kotlinlang.org/docs/reference/type-safe-builders.html) does,
But it can also do quite well.

Consider the following code:
`),
	ch.Code(samples.TypeSafeBuilderSample),
	md.Markdown(`
It's basically assembled what Kotlin can do, Also is legitimate Go code.
`),
	utils.Demo("", samples.TypeSafeBuilderSamplePath),
)
