package samples

import (
	"fmt"
	"net/url"

	"github.com/goplaid/web"
	. "github.com/theplant/htmlgo"
)

var page1Title = "Page 1"

//@snippet_begin(PageTransitionSample)

const Page1Path = "/samples/page_1"
const Page2Path = "/samples/page_2"

func Page1(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = Div(
		H1(page1Title),
		Ul(
			Li(
				A().Href(Page2Path).
					Text("To Page 2 With Normal Link"),
			),
			Li(
				web.Bind(
					A().Href("javascript:;").
						Text("To Page 2 With Push State Link"),
				).PushStateURL(Page2Path),
			),
		),
		fromParam(ctx),
	).Style("color: green; font-size: 24px;")
	return
}

func Page2(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("doAction1", doAction1)
	ctx.Hub.RegisterEventFunc("doAction2", doAction2)

	pr.Body = Div(
		H1("Page 2"),
		Ul(
			Li(
				web.Bind(
					A().Href("javascript:;").
						Text("To Page 1 With Normal Link"),
				).PushStateURL(Page1Path).
					PushStateQuery(
						url.Values{"from": []string{"page 2 link 1"}},
					),
			),
			Li(
				web.Bind(
					Button("Do an action then go to Page 1 with push state and parameters"),
				).OnClick("doAction2", "42"),
			),
			Li(
				web.Bind(
					Button("Do an action then go to Page 1 with redirect url"),
				).OnClick("doAction1", "41"),
			),
		),
	).Style("color: orange; font-size: 24px;")
	return
}

func fromParam(ctx *web.EventContext) HTMLComponent {
	var from HTMLComponent
	val := ctx.R.FormValue("from")
	if len(val) > 0 {
		from = Components(
			B("from:"),
			Text(val),
		)
	}
	return from
}

func doAction1(ctx *web.EventContext) (er web.EventResponse, err error) {
	updateDatabase(ctx.Event.ParamAsInt(0))
	er.RedirectURL = Page1Path + "?" + url.Values{"from": []string{"page2 with redirect"}}.Encode()
	return
}

func doAction2(ctx *web.EventContext) (er web.EventResponse, err error) {
	updateDatabase(ctx.Event.ParamAsInt(0))
	er.PushState = web.PushState(url.Values{"from": []string{"page2"}}).
		URL(Page1Path)
	return
}

//@snippet_end

func updateDatabase(val int) {
	page1Title = fmt.Sprintf("Page 1 (Updated by Page2 to %d)", val)
}
