package e00_basics

import (
	"fmt"
	"time"

	"github.com/goplaid/web"
	. "github.com/goplaid/x/vuetify"
	. "github.com/theplant/htmlgo"
)

func EventHandlingPage(ctx *web.EventContext) (pr web.PageResponse, err error) {
	api := ctx.R.URL.Query().Get("api")
	switch api {
	case "url":
		return EventHandlingURL(ctx)
	case "pushstate":
		return EventHandlingPushState(ctx)
	case "eventfunc":
		return EventHandlingEventFunc(ctx)
	case "reload":
		return EventHandlingReload(ctx)
	default:
		pr.Body = Div()
		return
	}
}

// @snippet_begin(EventHandlingURLSample)
func EventHandlingURL(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = Div(
		VBtn("Go").Attr("@click", web.GET().URL(EventExamplePagePath).Go()),
	)
	return
}

// @snippet_end

// @snippet_begin(EventHandlingPushStateSample)
func EventHandlingPushState(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = Div(
		VBtn("Go").Attr("@click", web.GET().URL(EventExamplePagePath).PushState(true).Go()),
	)
	return
}

// @snippet_end

// @snippet_begin(EventHandlingEventFuncSample)
func EventHandlingEventFunc(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("hello", func(ctx *web.EventContext) (r web.EventResponse, err error) {
		r.VarsScript = `alert("Hello World")`
		return
	})

	pr.Body = Div(
		VBtn("Go").Attr("@click", web.Plaid().EventFunc("hello").Go()),
	)
	return
}

// @snippet_end

// @snippet_begin(EventHandlingReloadSample)
func EventHandlingReload(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = Div(
		Text(fmt.Sprintf("Now: %s", time.Now().Format(time.RFC3339Nano))),
		VBtn("Reload").Attr("@click", web.Plaid().Reload().Go()),
	)
	return
}

// @snippet_end

func ExamplePage(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = Div(
		H1("ExamplePage"),
	)
	return
}

const EventHandlingPagePath = "/samples/event_handling"
const EventExamplePagePath = "/samples/event_handling/example"
