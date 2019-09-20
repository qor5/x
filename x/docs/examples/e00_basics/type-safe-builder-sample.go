package e00_basics

// @snippet_begin(TypeSafeBuilderSample)
import (
	"github.com/goplaid/web"
	. "github.com/theplant/htmlgo"
)

func result(args ...HTMLComponent) HTMLComponent {

	var converted []HTMLComponent
	for _, arg := range args {
		converted = append(converted, Div(arg).Class("wrapped"))
	}

	return HTML(
		Head(
			Title("XML encoding with Go"),
		),
		Body(
			H1("XML encoding with Go"),
			P().Text("this format can be used as an alternative markup to XML"),
			A().Href("http://golang.org").Text("Go"),
			P(
				Text("this is some"),
				B("mixed"),
				Text("text. For more see the"),
				A().Href("http://golang.org").Text("Go"),
				Text("project"),
			),
			P().Text("some text"),

			P(converted...),
		),
	)
}

func TypeSafeBuilderSamplePF(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = result(H5("1"), B("2"), Strong("3"))
	return
}

// @snippet_end

const TypeSafeBuilderSamplePath = "/samples/type_safe_builder_sample"
