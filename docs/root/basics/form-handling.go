package basics

import (
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e00_basics"
	"github.com/goplaid/x/docs/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
	. "github.com/theplant/htmlgo"
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
Use ~web.Bind(...).FieldName("Abc")~ to set the field name, make the name matches your data struct field name.
So that you can ~ctx.UnmarshalForm(&fv)~ to set the values to data object. value of input must be set manually to set the initial value of form field.

The fields which are bind with ~web.Bind(...).FieldName("Abc")~ are always submitted with every event func. A browser refresh, new page load will clear the form value.
`),
	utils.Anchor(H2(""), "Select Element"),
	Markdown(`
Use ~VSelect()~ to create a drop-down list. This components are used for collecting user provided information from a list of options.

The ~ItemText()~ method accept one string to set property of items’s text value and ~ItemValue()~ method accept one string to set property of items’s value.

Both of strings need to correspond to the object of the ~Items()~ parameter. Vuetify will marshal the object as JSON and get the corresponding value for each item as text and value.

For example:
`),
	ch.Code(`ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
	users := []user.User{}

	if err := db.Where("role = ?", "Admin").Find(&users).Error; err != nil {
		return h.Text(err.Error())
	}

	return VSelect().FieldName(field.Name).
		Label(field.Label).Value(field.Value(obj)).
		Items(users).ItemText("Name").ItemValue("ID")
})`),
	ch.Code(`type User struct {
	gorm.Model

	Name          string
	Email         string
	Role          string 
}`),
).Title("Form Handling").
	Slug("basics/form-handling")
