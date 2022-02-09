package root

import (
	"embed"

	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e00_basics"
	"github.com/goplaid/x/docs/root/basics"
	components_guide "github.com/goplaid/x/docs/root/components-guide"
	getting_started "github.com/goplaid/x/docs/root/getting-started"
	presets_guide "github.com/goplaid/x/docs/root/presets-guide"
	vuetify_components "github.com/goplaid/x/docs/root/vuetify-components"
	"github.com/goplaid/x/docs/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
	. "github.com/theplant/htmlgo"
)

var Home = Doc(
	Markdown(`
## What is GoPlaid

GoPlaid is yet another Go library to build web applications. 
different from other MVC frameworks. the concepts in GoPlaid is **Page**, **Event**, **Component**. 
and doesn't include Model.

A Page composite different kinds of Components, and Components trigger Events. 
A Page contains many event handlers, and renders one view, and event handlers reload the whole page,
Or update certain part of the page, Or go to a different Page. 

GoPlaid is opinionated in several ways:

- It prefers writing HTML in static typing Go language, rather than a certain type of template language, Not even go template.
- It try to minify the needs to write any JavaScript/Typescript for building interactive web applications
- It maximize the reusability of Components. since it uses Go to write components, You can abstract component very easy, and use component from a third party Go package is also like using normal Go packages.
- It prefers chain methods to set optional parameters of Component
- It uses [Vue](https://vuejs.org/) js under the hood. and only Vue Component can be integrated

`),
	utils.Anchor(H2(""), "Hello World"),
	Markdown(`
Here is the most sample hello world, that show the header with Hello World.
`),
	ch.Code(examples.HelloWorldSample).Language("go"),
	Markdown(`
~H1("Hello World")~ is actually a simple component. it renders h1 html tag. and been set to page body.

The above is the code you mostly writing. the following is the boilerplate code that needs to write one time.
`),
	ch.Code(examples.HelloWorldMuxSample1).Language("go"),
	ch.Code(examples.HelloWorldMuxSample2).Language("go"),
	ch.Code(examples.HelloWorldMainSample).Language("go"),
	utils.Demo("Hello World", e00_basics.HelloWorldPath, "e00_basics/hello-world.go"),

	Markdown(`
If you wondering why ~H1("Hello World")~ and how this worked, Please go ahead and checkout next page
`),
).Title("GoPlaid").
	Slug("/").
	Tables(
		ChildrenTable(
			ContentGroup(
				getting_started.OneMinuteQuickStart,
				getting_started.TheGoHTMLBuilder,
			).Title("Getting Started"),

			ContentGroup(
				basics.PageFuncAndEventFunc,
				basics.Filter,
				basics.LayoutFunctionAndPageInjector,
				basics.SwitchPagesWithPushState,
				basics.ReloadPageWithAFlash,
				basics.PartialRefreshWithPortal,
				basics.ManipulatePageURLInEventFunc,
				basics.FormHandling,
				basics.SummaryOfEventResponse,
				basics.WebScope,
				basics.EventHandling,
			).Title("Basics"),

			ContentGroup(
				components_guide.CompositeNewComponentWithGo,
				components_guide.IntegrateAHeavyVueComponent,
			).Title("Components Guide"),

			ContentGroup(
				vuetify_components.ATasteOfUsingVuetifyInGo,
				vuetify_components.BasicInputs,
				vuetify_components.AutoComplete,
				vuetify_components.VariantSubForm,
				vuetify_components.NavigationDrawer,
				vuetify_components.LazyPortalsAndReload,
			).Title("Vuetify Components"),

			ContentGroup(
				presets_guide.ItsTheWholeHouse,
				presets_guide.ListingCustomizations,
				presets_guide.EditingCustomizations,
				presets_guide.DetailPageForComplexObject,
				presets_guide.Permissions,
			).Title("Presets Guide"),

			ContentGroup(
				Doc(utils.ExamplesDoc()).
					Title("All Demo Examples").
					Slug("appendix/all-demo-examples"),
			).Title("Appendix"),
		),
	)

//go:embed assets/**.*
var Assets embed.FS
