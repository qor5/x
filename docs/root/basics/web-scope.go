package basics

import (
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e00_basics"
	"github.com/goplaid/x/docs/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var WebScope = Doc(
	Markdown(`

### Use Locals to init vue variables

There is a concept of reactive object in vue. Reactive object can trigger view updates, and [Vue cannot detect normal property additions (e.g. this.myObject.newProperty = 'hi')](https://vuejs.org/v2/api/#Vue-set).
We pre-set the "locals" object as a reactive object, and then we can initialize various types of values and slot it into "locals". And the valid scopes of these values are all inside web.Scope().

For example:
`),
	ch.Code(examples.WebScopeUseLocalsSample1).Language("go"),
	utils.Demo("Web Scope Use Locals", e00_basics.WebScopeUseLocalsPagePath, "e00_basics/web-scope.go"),
	Markdown(`
Use ~web.Scope()~ to determine the effective scope of the variable, then use ~.Init(...).VSlot("{ locals }")~ to initialize the variable and slot it into the ~locals~ object. 

In ~VBtn("")~, you can use the ~click~ event to change the variable value in ~locals~ to achieve the effect that the page changes with the click.

In ~VBtn("Test Can Not Change Other Scope")~, values in ~locals~ will not change with the click, because the button is not in ~web.Scope()~.

Video Tutorial (<https://www.youtube.com/watch?v=UPuBvVRhUr0>)
`),
).Title("Scope Component").
	Slug("basics/scope-component")
