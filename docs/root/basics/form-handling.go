package basics

import (
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e00_basics"
	"github.com/goplaid/x/docs/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var FormHandling = Doc(
	Markdown(`
Form handling is an important part of web development. to make handling form easy, 
we have a global form that always be submitted with any event func. What you need to do
is just to give an input a name.

For example:
`),
	ch.Code(examples.FormHandlingSample).Language("go"),
	utils.Demo("Form Handling", e00_basics.FormHandlingPagePath, "e00_basics/form-handling.go"),
	Markdown(`
Use ~.Attr(web.VFieldName("Abc")...)~ to set the field name, make the name matches your data struct field name.
So that you can ~ctx.UnmarshalForm(&fv)~ to set the values to data object. value of input must be set manually to set the initial value of form field.

The fields which are bind with ~.Attr(web.VFieldName("Abc")...)~ are always submitted with every event func. A browser refresh, new page load will clear the form value.

~web.Scope(...).VSlot("{ plaidForm }")~ to nest a new form inside outside form, EventFunc inside will only post form values inside the scope.
`),
).Title("Form Handling").
	Slug("basics/form-handling")
