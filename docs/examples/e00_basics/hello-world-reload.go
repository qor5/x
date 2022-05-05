package e00_basics

// @snippet_begin(HelloWorldReloadSample)
import (
	"time"

	"github.com/goplaid/web"
	. "github.com/theplant/htmlgo"
)

func HelloWorldReload(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = Div(
		H1("Hello World"),
		Text(time.Now().Format(time.RFC3339Nano)),
		Button("Reload Page").Attr("@click", web.GET().
			EventFunc(reloadEvent).
			Go()),
	)
	return
}

func update(ctx *web.EventContext) (er web.EventResponse, err error) {
	er.Reload = true
	return
}

const reloadEvent = "reload"

var HelloWorldReloadPB = web.Page(HelloWorldReload).
	EventFunc(reloadEvent, update)

const HelloWorldReloadPath = "/samples/hello_world_reload"

// @snippet_end
