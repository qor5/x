package setup

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/goplaid/x/docs"
	"github.com/goplaid/x/docs/root/getting-started"
	"github.com/goplaid/x/docs/samples"

	"github.com/goplaid/web"
	"github.com/goplaid/x/codehighlight"
	"github.com/theplant/appkit/contexts"
	"github.com/theplant/appkit/server"
	. "github.com/theplant/htmlgo"
)

type section struct {
	title string
	slug  string
	items []*pageItem
}

type pageItem struct {
	section    string
	slug       string
	title      string
	renderFunc web.PageFunc
}

func menuLinks(prefix string, secs []*section) (comp HTMLComponent) {
	var nav = Nav().Class("side-tree-nav")
	for _, sec := range secs {
		secdiv := Div(
			Div(
				Div().Class("marker"),
				Div().Class("text").Text(sec.title),
			).Class("tree-item-title tree-branch-title js-item-title js-branch-title is_active"),
		).Class("tree-item tree-branch js-item js-branch _opened")
		for _, p := range sec.items {
			secdiv.AppendChildren(
				Div(
					web.Bind(A(
						Span("").Class("marker"),
						Span(p.title).Class("text"),
					).Class("tree-item-title tree-leaf-title js-item-title js-leaf-title").Href("javascript:;")).
						PushStateURL(fmt.Sprintf("%s/%s/%s", prefix, sec.slug, p.slug)),
				).Class("tree-item tree-leaf js-item js-leaf"),
			)
		}
		nav.AppendChildren(secdiv)
	}
	comp = Aside(
		Div(nav).Class("js-side-tree-nav"),
	).Class("g-3")

	return
}

func header() HTMLComponent {
	return Header(
		Div(
			Div(
				A().Href("/").Class("global-header-logo").Text("GoPlaid"),
				Nav(
					Div(
						A().Href("https://github.com/goplaid").Text("Github").Class("nav-item"),
					).Class("nav-links"),
				).Class("global-nav"),
			).Class("g-layout"),
		).Class("global-header-panel"),
	).Class("global-header")
}

func footer() HTMLComponent {
	return Footer(
		Div(
			Div(
				Div(
					Div(
						Div().Class("terms-copyright").Text("Licensed under the MIT license"),
					).Class("global-footer-row"),
				).Class("global-footer-container"),
			).Class("g-layout"),
		).Class("global-footer-terms"),
	).Role("contentinfo").Class("global-footer")
}

func layout(in web.PageFunc, secs []*section, prefix string, cp *pageItem) (out web.PageFunc) {
	return func(ctx *web.EventContext) (pr web.PageResponse, err error) {

		tailScript := `<script src='/assets/main.js'></script>`
		if len(os.Getenv("DEV")) > 0 {
			fmt.Println("Using Dev environment, make sure you did: yarn start")
			tailScript = `
				<script src='http://localhost:3300/app.js'></script>
				<script src='http://localhost:3100/app.js'></script>
			`
		} else {
			ctx.Injector.HeadHTML(`
				<link rel="stylesheet" href="/assets/main.css">
			`)
		}

		ctx.Injector.Title(cp.title)
		ctx.Injector.HeadHTML(`
			<script src='/assets/vue.js'></script>
			<script src='/assets/codehighlight.js'></script>
		`)

		ctx.Injector.TailHTML(tailScript)

		var innerPr web.PageResponse
		innerPr, err = in(ctx)
		if err != nil {
			panic(err)
		}

		demo := innerPr.Schema

		ctx.Injector.HeadHTML(`
		<style>
			[v-cloak] {
				display: none;
			}
		</style>
		`)

		pr.Schema = Components(
			Div(
				header(),
				Div(
					Div(
						menuLinks(prefix, secs),
						Article(demo.(HTMLComponent)).Class("page-content g-9").Role("main"),
					).Class("g-grid"),
				).Class("g-layout global-content"),
			).Class("global-layout"),
			footer(),
		)

		return
	}
}

func demoLayout(in web.PageFunc) (out web.PageFunc) {
	return func(ctx *web.EventContext) (pr web.PageResponse, err error) {

		ctx.Injector.HeadHTML(`
			<script src='/assets/vue.js'></script>
		`)

		ctx.Injector.TailHTML(`<script src='/assets/main.js'></script>`)
		ctx.Injector.HeadHTML(`
		<style>
			[v-cloak] {
				display: none;
			}
		</style>
		`)

		var innerPr web.PageResponse
		innerPr, err = in(ctx)
		if err != nil {
			panic(err)
		}

		pr.Schema = innerPr.Schema

		return
	}
}

func rf(comp HTMLComponent) web.PageFunc {
	return func(ctx *web.EventContext) (r web.PageResponse, err error) {
		r.Schema = comp
		return
	}
}

func Setup(prefix string) http.Handler {
	ub := web.New()

	// @snippet_begin(HelloWorldMuxSample1)
	mux := http.NewServeMux()
	// @snippet_end

	mux.Handle("/assets/main.js",
		ub.PacksHandler("text/javascript",
			docs.JSComponentsPack(),
			web.JSComponentsPack(),
		),
	)

	mux.Handle("/assets/vue.js",
		ub.PacksHandler("text/javascript",
			web.JSVueComponentsPack(),
		),
	)

	mux.Handle("/assets/codehighlight.js",
		ub.PacksHandler("text/javascript",
			codehighlight.JSComponentsPack(),
		),
	)

	mux.Handle("/assets/main.css",
		ub.PacksHandler("text/css",
			codehighlight.CSSComponentsPack(),
			docs.CSSComponentsPack(),
		),
	)

	mux.Handle("/favicon.ico", http.NotFoundHandler())

	var secs = []*section{
		{
			title: "Getting Started",
			slug:  "getting-started",
			items: []*pageItem{
				{
					title:      "What is GoPlaid?",
					slug:       "what-is-goplaid.html",
					renderFunc: rf(getting_started.WhatIsGoPlaid),
				},
				{
					title:      "The Go HTML builder",
					slug:       "the-go-html-builder.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
			},
		},

		{
			title: "Basics",
			slug:  "basics",
			items: []*pageItem{
				{
					title:      "Page Func and Event Func",
					slug:       "page-func-and-event-func.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
				{
					title:      "Layout Function to setup common things",
					slug:       "layout-function-to-setup-common-things.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
				{
					title:      "Switch Pages with Push State",
					slug:       "switch-pages-with-push-state.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
				{
					title:      "Form Handling",
					slug:       "form-handling.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
				{
					title:      "File Uploads",
					slug:       "file-uploads.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
				{
					title:      "Partial Refresh with Portal",
					slug:       "partial-refresh-with-portal.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
				{
					title:      "Event Flash Object",
					slug:       "event-flash-object.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
				{
					title:      "Page Injector",
					slug:       "page-injector.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
			},
		},
		{
			title: "Components Guide",
			slug:  "components-guide",
			items: []*pageItem{
				{
					title:      "Composite With Go",
					slug:       "composite-with-go.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
				{
					title:      "Integrate My First Vue Component",
					slug:       "integrate-my-first-vue-component.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
				{
					title:      "Update Form Values",
					slug:       "update-form-values.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
			},
		},
		{
			title: "Vuetify Components",
			slug:  "vuetify-components",
			items: []*pageItem{
				{
					title:      "A Taste of using Vuetify in Go",
					slug:       "a-taste-of-using-vuetify-in-go.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
				{
					title:      "Basic Inputs",
					slug:       "basic-inputs.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
				{
					title:      "Auto Complete",
					slug:       "auto-complete.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
				{
					title:      "Navigation Drawer",
					slug:       "navigation-drawer.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
			},
		},
		{
			title: "Presets",
			slug:  "presets",
			items: []*pageItem{
				{
					title:      "Not just scaffolding, it's the whole house",
					slug:       "its-the-whole-house.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
				{
					title:      "Listing fields and their Component Func",
					slug:       "listing-fields-and-their-component-func.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
				{
					title:      "Listing Filters",
					slug:       "listing-filters.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
				{
					title:      "Listing Filter Tabs",
					slug:       "listing-filter-tabs.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
				{
					title:      "Bulk Actions",
					slug:       "bulk-actions.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
				{
					title:      "Global Search Box",
					slug:       "global-search-tabs.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
				{
					title:      "Edit simple object side by side",
					slug:       "edit-simple-object-side-by-side.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
				{
					title:      "Editing Field Component Func",
					slug:       "editing-field-component-func.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
				{
					title:      "Validations",
					slug:       "validations.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
				{
					title:      "Complex Object with a detail page",
					slug:       "complex-object-with-detail-page.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
				{
					title:      "Card and Data Table Component",
					slug:       "card-and-data-table-component.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
				{
					title:      "Key Info and Detail Info Component",
					slug:       "key-info-and-detail-info-component.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
				{
					title:      "Files and Images",
					slug:       "files-and-images.html",
					renderFunc: rf(getting_started.TheGoHTMLBuilder),
				},
			},
		},
	}

	mw := server.Compose(
		// server.LogRequest,
		// log.WithLogger(l),
		contexts.WithHTTPStatus,
	)

	for _, sec := range secs {
		for _, p := range sec.items {
			url := fmt.Sprintf("/%s/%s", sec.slug, p.slug)
			log.Println(url)
			mux.Handle(
				url,
				mw(ub.Page(layout(p.renderFunc, secs, prefix, p))),
			)
		}
	}

	emptyUb := web.New().LayoutFunc(web.NoopLayoutFunc)

	mux.Handle(samples.TypeSafeBuilderSamplePath, mw(emptyUb.Page(samples.TypeSafeBuilderSamplePF)))

	// @snippet_begin(HelloWorldMuxSample2)
	wb := web.New()
	mux.Handle(samples.HelloWorldPath, wb.Page(samples.HelloWorld))
	// @snippet_end

	mux.Handle("/", mw(ub.Page(layout(rf(getting_started.TheGoHTMLBuilder), secs, prefix, secs[0].items[0]))))
	return mux
}
