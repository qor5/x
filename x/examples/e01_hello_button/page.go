package e01_hello_button

import (
	"github.com/goplaid/web"
	. "github.com/theplant/htmlgo"
)

type mystate struct {
	Message string
}

func HelloButton(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("reload", reload)

	var s = &mystate{}
	if ctx.Flash != nil {
		s = ctx.Flash.(*mystate)
	}

	pr.Body = Div(
		web.Bind(Button("Hello")).
			OnClick("reload"),
		web.Bind(Tag("input").
			Attr("type", "text").
			Attr("value", s.Message)).
			OnInput("reload").
			FieldName("Message"),
		Div().
			Style("font-family: monospace;").
			Text(s.Message),
	)
	return
}

func reload(ctx *web.EventContext) (r web.EventResponse, err error) {
	var s = &mystate{}
	ctx.MustUnmarshalForm(s)
	ctx.Flash = s

	r.Reload = true
	return
}
