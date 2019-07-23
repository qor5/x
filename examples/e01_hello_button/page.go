package e01_hello_button

import (
	"github.com/sunfmin/bran/ui"
	. "github.com/theplant/htmlgo"
)

type mystate struct {
	Message string
}

func HelloButton(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("reload", reload)

	s := ctx.StateOrInit(&mystate{}).(*mystate)

	pr.Schema = Div(
		ui.Bind(Button("Hello")).
			OnClick("reload"),
		ui.Bind(Tag("input").
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

func reload(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	r.Reload = true
	return
}
