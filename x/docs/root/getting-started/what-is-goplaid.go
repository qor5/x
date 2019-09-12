package getting_started

import (
	ch "github.com/goplaid/x/codehighlight"
	"github.com/goplaid/x/docs/samples"
	"github.com/goplaid/x/docs/utils"
	"github.com/goplaid/x/md"
	. "github.com/theplant/htmlgo"
)

var WhatIsGoPlaid = Components(
	utils.Anchor(H1(""), "What is GoPlaid?"),
	md.Markdown(`
GoPlaid is yet another Go library to build web applications. 
different from other MVC frameworks. the concepts in GoPlaid is **Page**, **Event**, **Component**.

A Page composite different kinds of Components, and Components trigger Events. 
A Page contains many event handlers, and renders one view, and event handlers reload the whole page,
Or update certain part of the page, Or go to a different Page. 

GoPlaid is opinionated in several ways:

- It prefers writing HTML in static typing Go language, rather than a certain type of template language, Not even go template.
- It try to minify the needs to write any JavaScript/Typescript for building interactive web applications
- It maximize the reusability of Components.
- It prefers chain methods to set optional parameters of Component
- It uses [Vue](https://vuejs.org/) js under the hood. and only Vue Component can be integrated

`),
	utils.Anchor(H2(""), "Hello World"),
	md.Markdown(`
Here is the most sample hello world, that show the header with Hello World.
`),
	ch.Code(samples.HelloWorldSample),
	md.Markdown(`
The above is the code you mostly writing. the following is the boilerplate code that needs to write one time.
`),
	ch.Code(samples.HelloWorldMuxSample1),
	ch.Code(samples.HelloWorldMuxSample2),
	ch.Code(samples.HelloWorldMainSample),
	utils.Demo("Check the demo", samples.HelloWorldPath),

	md.Markdown("If you wondering why `H1(\"Hello World\")` and how this worked, Please go ahead and checkout next page"),

)
