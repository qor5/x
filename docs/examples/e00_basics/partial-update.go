package e00_basics

//@snippet_begin(PartialUpdateSample)
import (
	"time"

	"github.com/goplaid/web"
	. "github.com/theplant/htmlgo"
)

func PartialUpdatePage(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("edit1", edit1)
	ctx.Hub.RegisterEventFunc("reload2", reload2)

	pr.Body = Div(
		H1("Partial Update"),
		web.Bind(
			A().Text("Edit").Href("javascript:;"),
		).OnClick("edit1"),
		web.Portal(
			Text("original portal content here"),
		).Name("part1"),
		Div().Text(time.Now().Format(time.RFC3339Nano)),
	)
	return
}

func edit1(ctx *web.EventContext) (er web.EventResponse, err error) {
	er.UpdatePortals = append(er.UpdatePortals, &web.PortalUpdate{
		Name: "part1",
		Body: Div(
			Fieldset(
				Legend("Input value"),
				Div(
					Label("Title"),
					Input("").Type("text"),
				),

				Div(
					Label("Date"),
					Input("").Type("date"),
				),
			),
			web.Bind(
				Button("Update"),
			).OnClick("reload2"),
		),
	})
	return
}

func reload2(ctx *web.EventContext) (er web.EventResponse, err error) {
	er.Reload = true
	return
}

//@snippet_end

const PartialUpdatePagePath = "/samples/partial_update"
