package basics

import (
	ch "github.com/goplaid/x/codehighlight"
	"github.com/goplaid/x/docs/samples"
	"github.com/goplaid/x/md"
	. "github.com/theplant/htmlgo"
)

var LayoutFunctionAndPageInjector = Components(
	md.Markdown("Read this code first, Guess what it does."),
	ch.Code(samples.DemoLayoutSample),
)
