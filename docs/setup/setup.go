package setup

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/sunfmin/bran/docs/samples"

	"github.com/sunfmin/bran"
	"github.com/sunfmin/bran/codehighlight"
	"github.com/sunfmin/bran/core"
	"github.com/sunfmin/bran/docs"
	"github.com/sunfmin/bran/docs/root/overview"
	"github.com/sunfmin/bran/ui"
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
	renderFunc ui.PageFunc
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
					A(
						Span("").Class("marker"),
						Span(p.title).Class("text"),
					).Href(fmt.Sprintf("%s/%s/%s", prefix, sec.slug, p.slug)).Class("tree-item-title tree-leaf-title js-item-title js-leaf-title"),
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
						A().Href("/").Text("Learn").Class("nav-item"),
						A().Href("/").Text("Github").Class("nav-item"),
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

func layout(in ui.PageFunc, secs []*section, prefix string, cp *pageItem) (out ui.PageFunc) {
	return func(ctx *ui.EventContext) (pr ui.PageResponse, err error) {

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

		var innerPr ui.PageResponse
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

func home(prefix string, pages []pageItem) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/e01_hello_button", 302)
		return
	}
}

func Setup(prefix string) http.Handler {
	ub := bran.New()

	mux := http.NewServeMux()

	mux.Handle("/assets/main.js",
		ub.PacksHandler("text/javascript",
			docs.JSComponentsPack(),
			core.JSComponentsPack(),
		),
	)

	mux.Handle("/assets/vue.js",
		ub.PacksHandler("text/javascript",
			core.JSVueComponentsPack(),
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
					renderFunc: overview.Index,
				},
				{
					title:      "Hello World",
					slug:       "hello-world.html",
					renderFunc: overview.Index,
				},
				{
					title:      "The Go HTML builder",
					slug:       "the-go-html-builder.html",
					renderFunc: overview.Index,
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
					renderFunc: overview.Index,
				},
				{
					title:      "Switch Pages with Push State",
					slug:       "switch-pages-with-push-state.html",
					renderFunc: overview.Index,
				},
				{
					title:      "Form Handling",
					slug:       "form-handling.html",
					renderFunc: overview.Index,
				},
				{
					title:      "File Uploads",
					slug:       "file-uploads.html",
					renderFunc: overview.Index,
				},
				{
					title:      "Partial Refresh with Portal",
					slug:       "partial-refresh-with-portal.html",
					renderFunc: overview.Index,
				},
				{
					title:      "Event Flash Object",
					slug:       "event-flash-object.html",
					renderFunc: overview.Index,
				},
				{
					title:      "Page Injector",
					slug:       "page-injector.html",
					renderFunc: overview.Index,
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
					renderFunc: overview.Index,
				},
				{
					title:      "Integrate My First Vue Component",
					slug:       "integrate-my-first-vue-component.html",
					renderFunc: overview.Index,
				},
				{
					title:      "Update Form Values",
					slug:       "update-form-values.html",
					renderFunc: overview.Index,
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
					renderFunc: overview.Index,
				},
				{
					title:      "Basic Inputs",
					slug:       "basic-inputs.html",
					renderFunc: overview.Index,
				},
				{
					title:      "Auto Complete",
					slug:       "auto-complete.html",
					renderFunc: overview.Index,
				},
				{
					title:      "Navigation Drawer",
					slug:       "navigation-drawer.html",
					renderFunc: overview.Index,
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
					renderFunc: overview.Index,
				},
				{
					title:      "Listing fields and their Component Func",
					slug:       "listing-fields-and-their-component-func.html",
					renderFunc: overview.Index,
				},
				{
					title:      "Listing Filters",
					slug:       "listing-filters.html",
					renderFunc: overview.Index,
				},
				{
					title:      "Listing Filter Tabs",
					slug:       "listing-filter-tabs.html",
					renderFunc: overview.Index,
				},
				{
					title:      "Bulk Actions",
					slug:       "bulk-actions.html",
					renderFunc: overview.Index,
				},
				{
					title:      "Global Search Box",
					slug:       "global-search-tabs.html",
					renderFunc: overview.Index,
				},
				{
					title:      "Edit simple object side by side",
					slug:       "edit-simple-object-side-by-side.html",
					renderFunc: overview.Index,
				},
				{
					title:      "Editing Field Component Func",
					slug:       "editing-field-component-func.html",
					renderFunc: overview.Index,
				},
				{
					title:      "Validations",
					slug:       "validations.html",
					renderFunc: overview.Index,
				},
				{
					title:      "Complex Object with a detail page",
					slug:       "complex-object-with-detail-page.html",
					renderFunc: overview.Index,
				},
				{
					title:      "Card and Data Table Component",
					slug:       "card-and-data-table-component.html",
					renderFunc: overview.Index,
				},
				{
					title:      "Key Info and Detail Info Component",
					slug:       "key-info-and-detail-info-component.html",
					renderFunc: overview.Index,
				},
				{
					title:      "Files and Images",
					slug:       "files-and-images.html",
					renderFunc: overview.Index,
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

	emptyUb := bran.New().LayoutFunc(bran.NoopLayoutFunc)

	mux.Handle(samples.TypeSafeBuilderSamplePath, mw(emptyUb.Page(samples.TypeSafeBuilderSamplePF)))
	mux.Handle("/", mw(ub.Page(layout(overview.Index, secs, prefix, secs[0].items[0]))))
	return mux
}
