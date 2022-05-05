package e00_basics

import (
	"fmt"
	"net/url"
	"time"

	"github.com/goplaid/web"
	. "github.com/goplaid/x/vuetify"
	. "github.com/theplant/htmlgo"
)

// @snippet_begin(EventHandlingURLSample)
func EventHandlingURL(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = Div(
		VCard(
			VCardTitle(Text("URL")),
			VCardActions(VBtn("Go").Attr("@click", web.GET().URL(EventExamplePagePath).Go())),
		),
	)
	return
}

// @snippet_end

// @snippet_begin(EventHandlingPushStateSample)
func EventHandlingPushState(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = Div(
		VCard(
			VCardTitle(Text("PushState")),
			VCardActions(VBtn("Go").Attr("@click", web.GET().URL(EventExamplePagePath).PushState(true).Go())),
		),
	)
	return
}

// @snippet_end

// @snippet_begin(EventHandlingReloadSample)
func EventHandlingReload(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = Div(
		VCard(
			VCardTitle(Text("Reload")),
			Text(fmt.Sprintf("Now: %s", time.Now().Format(time.RFC3339Nano))),
			VCardActions(VBtn("Reload").Attr("@click", web.POST().Reload().Go())),
		),
	)
	return
}

// @snippet_end

// @snippet_begin(EventHandlingQuerySample)
func EventHandlingQuery(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = Div(
		VCard(
			VCardTitle(Text("Query")),
			VCardActions(VBtn("Go").Attr("@click", web.GET().URL(EventExamplePagePath).PushState(true).Query("address", "tokyo").Go())),
		),
	)
	return
}

// @snippet_end

// @snippet_begin(EventHandlingMergeQuerySample)
func EventHandlingMergeQuery(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = Div(
		VCard(
			VCardTitle(Text("MergeQuery")),
			VCardActions(VBtn("Go").Attr("@click", web.GET().URL(EventExamplePagePath+"?address=beijing&name=qor5&email=qor5@theplant.jp").PushState(true).Query("address", "tokyo").MergeQuery(true).Go())),
		),
	)
	return
}

// @snippet_end

// @snippet_begin(EventHandlingClearMergeQuerySample)
func EventHandlingClearMergeQueryQuery(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = Div(
		VCard(
			VCardTitle(Text("ClearMergeQuery")),
			VCardActions(VBtn("Go").Attr("@click", web.GET().URL(EventExamplePagePath+"?address=beijing&name=qor5&email=qor5@theplant.jp").PushState(true).Query("address", "tokyo").ClearMergeQuery([]string{"name"}).Go())),
		),
	)
	return
}

// @snippet_end

// @snippet_begin(EventHandlingStringQuerySample)
func EventHandlingStringQuery(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = Div(
		VCard(
			VCardTitle(Text("StringQuery")),
			VCardActions(VBtn("Go").Attr("@click", web.GET().URL(EventExamplePagePath).PushState(true).StringQuery("address=tokyo").Go())),
		),
	)
	return
}

// @snippet_end

// @snippet_begin(EventHandlingQueriesSample)
func EventHandlingQueries(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = Div(
		VCard(
			VCardTitle(Text("Queries")),
			VCardActions(VBtn("Go").Attr("@click", web.GET().URL(EventExamplePagePath).PushState(true).Queries(url.Values{"address": []string{"tokyo"}}).Go())),
		),
	)
	return
}

// @snippet_end

// @snippet_begin(EventHandlingPushStateURLSample)
func EventHandlingPushStateURL(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = Div(
		VCard(
			VCardTitle(Text("PushStateURL")),
			VCardActions(VBtn("Go").Attr("@click", web.GET().PushStateURL(EventExamplePagePath).Go())),
		),
	)
	return
}

// @snippet_end

// @snippet_begin(EventHandlingLocationSample)
func EventHandlingLocation(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = Div(
		VCard(
			VCardTitle(Text("Location")),
			VCardActions(VBtn("Go").Attr("@click", web.POST().PushState(true).Location(&web.LocationBuilder{MyURL: EventExamplePagePath, MyStringQuery: "address=test"}).Go())),
		),
	)
	return
}

// @snippet_end

// @snippet_begin(EventHandlingFieldValueSample)
func EventHandlingFileValue(ctx *web.EventContext) (pr web.PageResponse, err error) {

	pr.Body = Div(
		VCard(
			VCardTitle(Text("FieldValue")),
			VCardActions(VBtn("Go").Attr("@click", web.POST().EventFunc("form").FieldValue("name", "qor5").Go())),
		),
	)
	return
}

// @snippet_end

// @snippet_begin(EventHandlingFormClearSample)
func EventHandlingFormClear(ctx *web.EventContext) (pr web.PageResponse, err error) {

	pr.Body = Div(
		VCard(
			VCardTitle(Text("FormClear")),
			VCardActions(VBtn("Go").Attr("@click", web.POST().EventFunc("form").FieldValue("name", "qor5").FormClear().Go())),
		),
	)
	return
}

// @snippet_end

// @snippet_begin(EventHandlingEventFuncSample)
func EventHandlingEventFunc(ctx *web.EventContext) (pr web.PageResponse, err error) {

	pr.Body = Div(
		VBtn("Go").Attr("@click", web.POST().EventFunc("hello").Go()),
	)
	return
}

// @snippet_end

// @snippet_begin(EventHandlingBeforeScriptSample)
func EventHandlingScript(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = Div(
		VCard(
			VCardTitle(Text("Script")),
			VCardActions(VBtn("Go").Attr("@click", web.POST().ThenScript(`alert("this is then script")`).AfterScript(`alert("this is after script")`).BeforeScript(`alert("this is before script")`).Go())),
		),
	)
	return
}

// @snippet_end

// @snippet_begin(EventHandlingRawSample)
func EventHandlingRaw(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = Div(
		VCard(
			VCardTitle(Text("Raw")),
			VCardActions(VBtn("Go").Attr("@click", web.POST().Raw(`pushStateURL("/samples/event_handling/example")`).Go())),
		),
	)
	return
}

// @snippet_end

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
	case "query":
		return EventHandlingQuery(ctx)
	case "merge_query":
		return EventHandlingMergeQuery(ctx)
	case "clear_merge_query":
		return EventHandlingClearMergeQueryQuery(ctx)
	case "string_query":
		return EventHandlingStringQuery(ctx)
	case "queries":
		return EventHandlingQueries(ctx)
	case "pushstateurl":
		return EventHandlingPushStateURL(ctx)
	case "fieldvalue":
		return EventHandlingFileValue(ctx)
	case "formclear":
		return EventHandlingFormClear(ctx)
	case "script":
		return EventHandlingScript(ctx)
	case "location":
		return EventHandlingLocation(ctx)
	case "raw":
		return EventHandlingRaw(ctx)
	default:
		pr.Body = Div()
		return
	}
}

func ExamplePage(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = Div(
		H1("ExamplePage"),
	)
	return
}

var ExamplePagePB = web.Page(ExamplePage).
	EventFunc("form", func(ctx *web.EventContext) (r web.EventResponse, err error) {
		r.VarsScript = fmt.Sprintf(`alert("form data is %s")`, ctx.R.FormValue("name"))
		return
	}).
	EventFunc("hello", func(ctx *web.EventContext) (r web.EventResponse, err error) {
		r.VarsScript = `alert("Hello World")`
		return
	})

var EventHandlingPagePB = web.Page(EventHandlingPage)

const EventHandlingPagePath = "/samples/event_handling"
const EventExamplePagePath = "/samples/event_handling/example"
