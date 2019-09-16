package samples

//@snippet_begin(CompositeComponentSample1)
import (
	"github.com/goplaid/web"
	. "github.com/theplant/htmlgo"
)

func Navbar(title string, activeIndex int, items ...HTMLComponent) HTMLComponent {
	ul := Ul().Class("navbar-nav mr-auto")

	for i, item := range items {
		ul.AppendChildren(
			Li(
				item,
			).Class("nav-item").ClassIf("active", activeIndex == i),
		)
	}

	return Nav(
		A(Text(title)).Class("navbar-brand").
			Href("#"),

		Button("").Class("navbar-toggler").
			Type("button").
			Attr("data-toggle", "collapse").
			Attr("data-target", "#navbarNav").
			Attr("aria-controls", "navbarNav").
			Attr("aria-expanded", "false").
			Attr("aria-label", "Toggle navigation").
			Children(
				Span("").Class("navbar-toggler-icon"),
			),

		Div(
			ul,
			Form(
				Input("").Class("form-control mr-sm-2").
					Type("search").
					Placeholder("Search").
					Attr("aria-label", "Search"),
				Button("Search").Class("btn btn-outline-light my-2 my-sm-0").
					Type("submit"),
			).Class("form-inline my-2 my-lg-0"),
		).Class("collapse navbar-collapse").
			Id("navbarNav"),
	).Class("navbar navbar-expand-lg navbar-dark bg-primary")
}

func CompositeComponentSample1Page(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = Div(
		Navbar(
			"Hello",
			1,

			A(
				Text("Home"),
			).Class("nav-link").
				Href("#"),

			A(
				Text("Features"),
			).Class("nav-link").
				Href("#"),

			A(
				Text("Pricing"),
			).Class("nav-link").
				Href("#"),

			A(
				Text("Disabled"),
			).Class("nav-link disabled").
				Href("#").
				TabIndex(-1).
				Attr("aria-disabled", "true"),
		),
	)
	return
}

//@snippet_end

const CompositeComponentSample1PagePath = "/samples/composite-component-sample1"
