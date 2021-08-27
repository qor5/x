package e00_basics

//@snippet_begin(HelloWorldReloadSample)
import (
	"time"

	"github.com/goplaid/web"
	. "github.com/theplant/htmlgo"
)

func HelloWorldReload(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("reload", update)
	pr.Body = Div(
		H1("Hello World"),
		Text(time.Now().Format(time.RFC3339Nano)),
		Button("Reload Page").Attr("@click", web.Plaid().EventFunc("reload").Go()),
	)
	return
}

func update(ctx *web.EventContext) (er web.EventResponse, err error) {
	er.Reload = true
	return
}

//@snippet_end

const HelloWorldReloadPath = "/samples/hello_world_reload"
