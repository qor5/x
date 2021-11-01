package e00_basics

//@snippet_begin(MultiStatePageSample)
import (
	"net/url"

	"github.com/goplaid/web"
	. "github.com/theplant/htmlgo"
)

func MultiStatePage(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("openPanel", openPanel)
	ctx.Hub.RegisterEventFunc("update5", update5)

	title := "Multi State Page"
	if len(ctx.R.URL.Query().Get("title")) > 0 {
		title = ctx.R.URL.Query().Get("title")
	}
	var panel HTMLComponent
	if len(ctx.R.URL.Query().Get("panel")) > 0 {
		panel = Div(
			Fieldset(
				Div(
					Label("Name"),
					Input("").Type("text"),
				),
				Div(
					Label("Date"),
					Input("").Type("date"),
				),
			),
			Button("Update").Attr("@click", web.Plaid().EventFunc("update5").Go()),
		).Style("border: 5px solid orange; height: 200px;")
	}

	pr.Body = Div(
		H1(title),
		Ol(
			Li(
				A().Text("change page title").Href("javascript:;").
					Attr("@click", web.Plaid().Queries(url.Values{"title": []string{"Hello"}}).Go()),
			),
			Li(
				A().Text("show panel").Href("javascript:;").Attr("@click", web.Plaid().EventFunc("openPanel").Go()),
			),
		),
		panel,

		Table(
			Thead(
				Th("Name"),
				Th("Date"),
			),
			Tbody(
				Tr(
					Td(Text("Felix")),
					Td(Text("2019-01-02")),
				),
			),
		),
	)
	return
}

func openPanel(ctx *web.EventContext) (er web.EventResponse, err error) {
	er.PushState = web.Location(url.Values{"panel": []string{"1"}}).MergeQuery(true)
	return
}

func update5(ctx *web.EventContext) (er web.EventResponse, err error) {
	er.PushState = web.Location(url.Values{"panel": []string{""}}).MergeQuery(true)
	return
}

//@snippet_end

const MultiStatePagePath = "/samples/multi_state_page"
