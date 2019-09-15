package setup

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/goplaid/web"
	"github.com/goplaid/x/codehighlight"
	"github.com/goplaid/x/docs"
	"github.com/goplaid/x/docs/root/basics"
	getting_started "github.com/goplaid/x/docs/root/getting-started"
	"github.com/goplaid/x/docs/samples"
	"github.com/goplaid/x/docs/utils"
	. "github.com/theplant/htmlgo"
)

type section struct {
	title string
	slug  string
	items []*pageItem
}

type pageItem struct {
	section string
	slug    string
	title   string
	doc     HTMLComponent
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
		pr.PageTitle = cp.title + " - " + "GoPlaid"
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

		demo := innerPr.Body

		ctx.Injector.HeadHTML(`
		<style>
			[v-cloak] {
				display: none;
			}
		</style>
		`)

		pr.Body = Components(
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

// @snippet_begin(DemoLayoutSample)
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

		pr.Body = innerPr.Body

		return
	}
}

// @snippet_end

func rf(comp HTMLComponent, p *pageItem) web.PageFunc {
	return func(ctx *web.EventContext) (r web.PageResponse, err error) {
		r.Body = Components(
			utils.Anchor(H1(""), p.title),
			comp,
		)
		return
	}
}

var tbd = Text("TBD")

func Setup(prefix string) http.Handler {

	// @snippet_begin(HelloWorldMuxSample1)
	mux := http.NewServeMux()
	// @snippet_end

	// @snippet_begin(ComponentsPackSample)
	ub := web.New()
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
	// @snippet_end

	mux.Handle("/favicon.ico", http.NotFoundHandler())

	var secs = []*section{
		{
			title: "Getting Started",
			slug:  "getting-started",
			items: []*pageItem{
				{
					title: "What is GoPlaid?",
					slug:  "what-is-goplaid.html",
					doc:   getting_started.WhatIsGoPlaid,
				},
				{
					title: "The Go HTML builder",
					slug:  "the-go-html-builder.html",
					doc:   getting_started.TheGoHTMLBuilder,
				},
			},
		},

		{
			title: "Basics",
			slug:  "basics",
			items: []*pageItem{
				{
					title: "Page Func and Event Func",
					slug:  "page-func-and-event-func.html",
					doc:   basics.PageFuncAndEventFunc,
				},
				{
					title: "Layout Function and Page Injector",
					slug:  "layout-function-and-page-injector.html",
					doc:   basics.LayoutFunctionAndPageInjector,
				},
				{
					title: "Switch Pages with Push State",
					slug:  "switch-pages-with-push-state.html",
					doc:   basics.SwitchPagesWithPushState,
				},
				{
					title: "Reload Page with a Flash",
					slug:  "reload-page-with-a-flash.html",
					doc:   basics.ReloadPageWithAFlash,
				},
				{
					title: "Partial Refresh with Portal",
					slug:  "partial-refresh-with-portal.html",
					doc:   basics.PartialRefreshWithPortal,
				},
				{
					title: "Form Handling",
					slug:  "form-handling.html",
					doc:   tbd,
				},
				{
					title: "File Uploads",
					slug:  "file-uploads.html",
					doc:   tbd,
				},
			},
		},
		{
			title: "Components Guide",
			slug:  "components-guide",
			items: []*pageItem{
				{
					title: "Composite With Go",
					slug:  "composite-with-go.html",
					doc:   tbd,
				},
				{
					title: "Integrate My First Vue Component",
					slug:  "integrate-my-first-vue-component.html",
					doc:   tbd,
				},
				{
					title: "Update Form Values",
					slug:  "update-form-values.html",
					doc:   tbd,
				},
			},
		},
		{
			title: "Vuetify Components",
			slug:  "vuetify-components",
			items: []*pageItem{
				{
					title: "A Taste of using Vuetify in Go",
					slug:  "a-taste-of-using-vuetify-in-go.html",
					doc:   tbd,
				},
				{
					title: "Basic Inputs",
					slug:  "basic-inputs.html",
					doc:   tbd,
				},
				{
					title: "Auto Complete",
					slug:  "auto-complete.html",
					doc:   tbd,
				},
				{
					title: "Navigation Drawer",
					slug:  "navigation-drawer.html",
					doc:   tbd,
				},
			},
		},
		{
			title: "Presets",
			slug:  "presets",
			items: []*pageItem{
				{
					title: "Not just scaffolding, it's the whole house",
					slug:  "its-the-whole-house.html",
					doc:   tbd,
				},
				{
					title: "Listing fields and their Component Func",
					slug:  "listing-fields-and-their-component-func.html",
					doc:   tbd,
				},
				{
					title: "Listing Filters",
					slug:  "listing-filters.html",
					doc:   tbd,
				},
				{
					title: "Listing Filter Tabs",
					slug:  "listing-filter-tabs.html",
					doc:   tbd,
				},
				{
					title: "Bulk Actions",
					slug:  "bulk-actions.html",
					doc:   tbd,
				},
				{
					title: "Global Search Box",
					slug:  "global-search-tabs.html",
					doc:   tbd,
				},
				{
					title: "Edit simple object side by side",
					slug:  "edit-simple-object-side-by-side.html",
					doc:   tbd,
				},
				{
					title: "Editing Field Component Func",
					slug:  "editing-field-component-func.html",
					doc:   tbd,
				},
				{
					title: "Validations",
					slug:  "validations.html",
					doc:   tbd,
				},
				{
					title: "Complex Object with a detail page",
					slug:  "complex-object-with-detail-page.html",
					doc:   tbd,
				},
				{
					title: "Card and Data Table Component",
					slug:  "card-and-data-table-component.html",
					doc:   tbd,
				},
				{
					title: "Key Info and Detail Info Component",
					slug:  "key-info-and-detail-info-component.html",
					doc:   tbd,
				},
				{
					title: "Files and Images",
					slug:  "files-and-images.html",
					doc:   tbd,
				},
			},
		},
	}

	for _, sec := range secs {
		for _, p := range sec.items {
			url := fmt.Sprintf("/%s/%s", sec.slug, p.slug)
			log.Println(url)
			mux.Handle(
				url,
				ub.Page(layout(rf(p.doc, p), secs, prefix, p)),
			)
		}
	}

	emptyUb := web.New().LayoutFunc(web.NoopLayoutFunc)

	mux.Handle(samples.TypeSafeBuilderSamplePath, emptyUb.Page(samples.TypeSafeBuilderSamplePF))

	// @snippet_begin(HelloWorldMuxSample2)
	wb := web.New()
	mux.Handle(samples.HelloWorldPath, wb.Page(samples.HelloWorld))
	// @snippet_end

	// @snippet_begin(HelloWorldReloadMuxSample1)
	mux.Handle(
		samples.HelloWorldReloadPath,
		wb.Page(
			demoLayout(
				samples.HelloWorldReload,
			),
		),
	)
	// @snippet_end

	mux.Handle(
		samples.Page1Path,
		wb.Page(
			demoLayout(
				samples.Page1,
			),
		),
	)
	mux.Handle(
		samples.Page2Path,
		wb.Page(
			demoLayout(
				samples.Page2,
			),
		),
	)

	mux.Handle(
		samples.ReloadWithFlashPath,
		wb.Page(
			demoLayout(
				samples.ReloadWithFlash,
			),
		),
	)

	mux.Handle(
		samples.PartialUpdatePagePath,
		wb.Page(
			demoLayout(
				samples.PartialUpdatePage,
			),
		),
	)

	mux.Handle(
		samples.PartialReloadPagePath,
		wb.Page(
			demoLayout(
				samples.PartialReloadPage,
			),
		),
	)

	home := secs[0].items[0]
	mux.Handle("/", ub.Page(layout(rf(home.doc, home), secs, prefix, home)))
	return mux
}
