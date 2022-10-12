package advanced_functions

import (
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e00_basics"
	"github.com/goplaid/x/docs/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var TheGoHTMLBuilder = Doc(
	Markdown(`
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
	Markdown(`
It's basically assembled what Kotlin can do, Also is legitimate Go code.
`),
	utils.Demo("The Go HTML Builder", e00_basics.TypeSafeBuilderSamplePath, "e00_basics/type-safe-builder-sample.go"),
).Title("The Go HTML builder").
	Slug("advanced-functions/the-go-html-builder")
