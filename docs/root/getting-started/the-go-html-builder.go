package getting_started

import (
	ch "github.com/goplaid/x/codehighlight"
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e00_basics"
	"github.com/goplaid/x/docs/utils"
	"github.com/goplaid/x/md"
	. "github.com/theplant/htmlgo"
)

var TheGoHTMLBuilder = Components(
	md.Markdown(`
Like at the beginning we said, That we don't use interpreted template language (eg go html/template)
to generate html page. We think they are:

- error prone without static type enforcing
- hard to refactor
- difficult to abstract out to component
- yet another tedious syntax to learn
- not flexible to use helper functions

We like to use standard Go code. the library [htmlgo](https://github.com/theplant/htmlgo) is just for that.

Although Go can't do flexible builder syntax like [Kotlin](https://kotlinlang.org/docs/reference/type-safe-builders.html) does,
But it can also do quite well.

Consider the following code:
`),
	ch.Code(examples.TypeSafeBuilderSample).Language("go"),
	md.Markdown(`
It's basically assembled what Kotlin can do, Also is legitimate Go code.
`),
	utils.Demo("", e00_basics.TypeSafeBuilderSamplePath),
)
