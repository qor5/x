package components_guide

import (
	ch "github.com/goplaid/x/codehighlight"
	"github.com/goplaid/x/docs/samples"
	"github.com/goplaid/x/docs/utils"
	"github.com/goplaid/x/md"
	. "github.com/theplant/htmlgo"
)

var CompositeNewComponentWithGo = Components(
	md.Markdown(`
Any Go function that returns an ~htmlgo.HTMLComponent~ is a component,
Any Go struct that implements ~MarshalHTML(ctx context.Context) ([]byte, error)~ function is an component.
They can be composite into a new component very easy. 

This example is ported from [Bootstrap4 Navbar](https://getbootstrap.com/docs/4.3/components/navbar/):
`),
	ch.Code(samples.CompositeComponentSample1),
	utils.Demo("", samples.CompositeComponentSample1PagePath),
	md.Markdown(`
You can see from the example, It is easy to pass in components as parameter, and wrap components.
By utilizing the power of flexible Go language, Any component can be abstracted and reused with enough parameters. 

It is a responsive navigation header, Resizing your window, the nav bar will react to device window size and change to nav bar popup and hide search form.

For this ~Navbar~ component to work, I have to import Bootstrap assets in this new layout function:
`),
	ch.Code(samples.DemoBootstrapLayoutSample),

	md.Markdown(`
You can utilize the command line tool [html2go](https://github.com/sunfmin/html2go) to convert existing html code to htmlgo code.
By writing html in Go you get:

- The static type checking
- Abstract out easily to different functions
- Easier refactor with IDE like GoLand
- Loop and variable replacing is just like in Go
- Invoke helper functions is just like in Go
- Almost as readable as normal HTML
- Not possible to have html tag not closed, Or not matched.

Once you have these, Why generate html in any interpreted template language!

`),
)
