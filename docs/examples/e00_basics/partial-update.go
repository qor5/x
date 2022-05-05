package e00_basics

// @snippet_begin(PartialUpdateSample)
import (
	"time"

	"github.com/goplaid/web"
	. "github.com/theplant/htmlgo"
)

func PartialUpdatePage(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = Div(
		H1("Partial Update"),
		A().Text("Edit").Href("javascript:;").
			Attr("@click", web.POST().EventFunc("edit1").Go()),
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
			Button("Update").
				Attr("@click", web.POST().EventFunc("reload2").Go()),
		),
	})
	return
}

func reload2(ctx *web.EventContext) (er web.EventResponse, err error) {
	er.Reload = true
	return
}

var PartialUpdatePagePB = web.Page(PartialUpdatePage).
	EventFunc("edit1", edit1).
	EventFunc("reload2", reload2)

const PartialUpdatePagePath = "/samples/partial_update"

// @snippet_end
