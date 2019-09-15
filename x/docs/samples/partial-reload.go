package samples

//@snippet_begin(PartialReloadSample)
import (
	"time"

	"github.com/goplaid/web"
	. "github.com/theplant/htmlgo"
)

func PartialReloadPage(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("related", related)
	ctx.Hub.RegisterEventFunc("reload3", reload3)
	ctx.Injector.HeadHTML(`
<style>
.rp {
	float: left;
	width: 200px;
	height: 200px;
	margin-right: 20px;
	background-color: orange;
}
</style>
`,
	)
	pr.Body = Div(
		Style(`
`,
		),
		H1("Partial Load and Reload"),
		Div(
			H2("Product 1"),
		).Style("height: 200px; background-color: grey;"),
		H2("Related Products"),
		web.Portal().Name("related_products").EventFunc("related"),
		web.Bind(
			A().Href("javascript:;").Text("Reload Related Products"),
		).OnClick("reload3"),
	)
	return
}

func related(ctx *web.EventContext) (er web.EventResponse, err error) {
	er.Body = Div(

		Div(
			H3("Product A"),
			Div().Text(time.Now().Format(time.RFC3339Nano)),
		).Class("rp"),
		Div(
			H3("Product B"),
			Div().Text(time.Now().Format(time.RFC3339Nano)),
		).Class("rp"),
		Div(
			H3("Product C"),
			Div().Text(time.Now().Format(time.RFC3339Nano)),
		).Class("rp"),
	)
	return
}

func reload3(ctx *web.EventContext) (er web.EventResponse, err error) {
	er.ReloadPortals = []string{"related_products"}
	return
}

//@snippet_end

const PartialReloadPagePath = "/samples/partial_reload"
