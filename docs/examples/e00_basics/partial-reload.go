package e00_basics

// @snippet_begin(PartialReloadSample)
import (
	"time"

	"github.com/goplaid/web"
	. "github.com/theplant/htmlgo"
)

func PartialReloadPage(ctx *web.EventContext) (pr web.PageResponse, err error) {
	reloadCount = 0
	ctx.Hub.RegisterEventFunc("related", related)
	ctx.Hub.RegisterEventFunc("reload3", reload3)
	ctx.Hub.RegisterEventFunc("autoReload", autoReload)
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
		H1("Portal Reload Automatically"),

		web.Scope(
			web.Portal().Loader(web.Plaid().EventFunc("autoReload")).AutoReloadInterval("locals.interval"),
			Button("stop").Attr("@click", "locals.interval = 0"),
		).Init(`{interval: 2000}`).VSlot("{ locals }"),

		H1("Partial Load and Reload"),
		Div(
			H2("Product 1"),
		).Style("height: 200px; background-color: grey;"),
		H2("Related Products"),
		web.Portal().Name("related_products").Loader(web.Plaid().EventFunc("related").Query("productCode", "AH123")),
		A().Href("javascript:;").Text("Reload Related Products").
			Attr("@click", web.Plaid().EventFunc("reload3").Go()),
	)
	return
}

func related(ctx *web.EventContext) (er web.EventResponse, err error) {
	code := ctx.R.FormValue("productCode")
	er.Body = Div(

		Div(
			H3("Product A (related products of "+code+")"),
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

var reloadCount = 1

func autoReload(ctx *web.EventContext) (er web.EventResponse, err error) {
	er.Body = Span(time.Now().String())
	reloadCount++

	if reloadCount > 5 {
		er.VarsScript = `vars.interval = 0;`
	}
	return
}

// @snippet_end

const PartialReloadPagePath = "/samples/partial_reload"
