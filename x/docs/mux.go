package docs

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/goplaid/x/docs/examples/e21_present_hello_world"

	"github.com/goplaid/x/docs/examples/e15_vuetify_navigation_drawer"

	"github.com/goplaid/x/docs/examples/e10_vuetify_autocomplete"

	"github.com/goplaid/web"
	"github.com/goplaid/x/codehighlight"
	"github.com/goplaid/x/docs/examples/e00_basics"
	"github.com/goplaid/x/docs/examples/e11_vuetify_basic_inputs"
	"github.com/goplaid/x/docs/examples/e13_vuetify_list"
	"github.com/goplaid/x/docs/examples/e14_vuetify_menu"
	"github.com/goplaid/x/docs/root/basics"
	components_guide "github.com/goplaid/x/docs/root/components-guide"
	getting_started "github.com/goplaid/x/docs/root/getting-started"
	"github.com/goplaid/x/docs/root/presets"
	vuetify_components "github.com/goplaid/x/docs/root/vuetify-components"
	"github.com/goplaid/x/docs/utils"
	"github.com/goplaid/x/tiptap"
	v "github.com/goplaid/x/vuetify"
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
					A(
						Span("").Class("marker"),
						Span(p.title).Class("text"),
					).Class("tree-item-title tree-leaf-title js-item-title js-leaf-title").
						Href(fmt.Sprintf("%s/%s/%s", prefix, sec.slug, p.slug)),
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

// @snippet_begin(TipTapLayoutSample)
func tiptapLayout(in web.PageFunc) (out web.PageFunc) {
	return func(ctx *web.EventContext) (pr web.PageResponse, err error) {

		ctx.Injector.HeadHTML(`
			<link rel="stylesheet" href="/assets/tiptap.css">
			<script src='/assets/vue.js'></script>
		`)

		ctx.Injector.TailHTML(`
<script src='/assets/tiptap.js'></script>
<script src='/assets/main.js'></script>
`)
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

// @snippet_begin(DemoBootstrapLayoutSample)
func demoBootstrapLayout(in web.PageFunc) (out web.PageFunc) {
	return func(ctx *web.EventContext) (pr web.PageResponse, err error) {

		ctx.Injector.HeadHTML(`
<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
<script src='/assets/vue.js'></script>
		`)

		ctx.Injector.TailHTML(`
<script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js" integrity="sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1" crossorigin="anonymous"></script>
<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js" integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM" crossorigin="anonymous"></script>
<script src='/assets/main.js'></script>

`)
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

// @snippet_begin(DemoVuetifyLayoutSample)
func demoVuetifyLayout(in web.PageFunc) (out web.PageFunc) {
	return func(ctx *web.EventContext) (pr web.PageResponse, err error) {

		ctx.Injector.HeadHTML(`
			<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto+Mono" async>
			<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:300,400,500" async>
			<link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons" async>
			<link rel="stylesheet" href="/assets/vuetify.css">
			<script src='/assets/vue.js'></script>
		`)

		ctx.Injector.TailHTML(`
<script src='/assets/vuetify.js'></script>
<script src='/assets/main.js'></script>

`)
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

		pr.Body = v.VApp(
			v.VContent(
				innerPr.Body,
			),
		)

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

func Mux(prefix string) http.Handler {

	// @snippet_begin(HelloWorldMuxSample1)
	mux := http.NewServeMux()
	// @snippet_end

	// @snippet_begin(ComponentsPackSample)
	ub := web.New()
	mux.Handle("/assets/main.js",
		ub.PacksHandler("text/javascript",
			JSComponentsPack(),
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
			CSSComponentsPack(),
		),
	)
	// @snippet_end

	// @snippet_begin(TipTapComponentsPackSample)
	mux.Handle("/assets/tiptap.js",
		ub.PacksHandler("text/javascript",
			tiptap.JSComponentsPack(),
		),
	)

	mux.Handle("/assets/tiptap.css",
		ub.PacksHandler("text/css",
			tiptap.CSSComponentsPack(),
		),
	)
	// @snippet_end

	// @snippet_begin(VuetifyComponentsPackSample)
	mux.Handle("/assets/vuetify.js",
		ub.PacksHandler("text/javascript",
			v.JSComponentsPack(),
		),
	)

	mux.Handle("/assets/vuetify.css",
		ub.PacksHandler("text/css",
			v.CSSComponentsPack(),
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
					title: "Manipulate Page URL in Event Func",
					slug:  "manipulate-page-url-in-event-func.html",
					doc:   basics.ManipulatePageURLInEventFunc,
				},
				{
					title: "Form Handling",
					slug:  "form-handling.html",
					doc:   basics.FormHandling,
				},
				{
					title: "Summary of Event Response",
					slug:  "summary-of-event-response.html",
					doc:   basics.SummaryOfEventResponse,
				},
			},
		},
		{
			title: "Components Guide",
			slug:  "components-guide",
			items: []*pageItem{
				{
					title: "Composite new Component With Go",
					slug:  "composite-new-component-with-go.html",
					doc:   components_guide.CompositeNewComponentWithGo,
				},
				{
					title: "Integrate a heavy Vue Component",
					slug:  "integrate-a-heavy-vue-component.html",
					doc:   components_guide.IntegrateAHeavyVueComponent,
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
					doc:   vuetify_components.ATasteOfUsingVuetifyInGo,
				},
				{
					title: "Basic Inputs",
					slug:  "basic-inputs.html",
					doc:   vuetify_components.BasicInputs,
				},
				{
					title: "Auto Complete",
					slug:  "auto-complete.html",
					doc:   vuetify_components.AutoComplete,
				},
				{
					title: "Navigation Drawer",
					slug:  "navigation-drawer.html",
					doc:   vuetify_components.NavigationDrawer,
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
					doc:   presets.ItsTheWholeHouse,
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

	mux.Handle(e00_basics.TypeSafeBuilderSamplePath, emptyUb.Page(e00_basics.TypeSafeBuilderSamplePF))

	// @snippet_begin(HelloWorldMuxSample2)
	wb := web.New()
	mux.Handle(e00_basics.HelloWorldPath, wb.Page(e00_basics.HelloWorld))
	// @snippet_end

	// @snippet_begin(HelloWorldReloadMuxSample1)
	mux.Handle(
		e00_basics.HelloWorldReloadPath,
		wb.Page(
			demoLayout(
				e00_basics.HelloWorldReload,
			),
		),
	)
	// @snippet_end

	mux.Handle(
		e00_basics.Page1Path,
		wb.Page(
			demoLayout(
				e00_basics.Page1,
			),
		),
	)
	mux.Handle(
		e00_basics.Page2Path,
		wb.Page(
			demoLayout(
				e00_basics.Page2,
			),
		),
	)

	mux.Handle(
		e00_basics.ReloadWithFlashPath,
		wb.Page(
			demoLayout(
				e00_basics.ReloadWithFlash,
			),
		),
	)

	mux.Handle(
		e00_basics.PartialUpdatePagePath,
		wb.Page(
			demoLayout(
				e00_basics.PartialUpdatePage,
			),
		),
	)

	mux.Handle(
		e00_basics.PartialReloadPagePath,
		wb.Page(
			demoLayout(
				e00_basics.PartialReloadPage,
			),
		),
	)

	mux.Handle(
		e00_basics.MultiStatePagePath,
		wb.Page(
			demoLayout(
				e00_basics.MultiStatePage,
			),
		),
	)

	mux.Handle(
		e00_basics.FormHandlingPagePath,
		wb.Page(
			demoLayout(
				e00_basics.FormHandlingPage,
			),
		),
	)

	mux.Handle(
		e00_basics.CompositeComponentSample1PagePath,
		wb.Page(
			demoBootstrapLayout(
				e00_basics.CompositeComponentSample1Page,
			),
		),
	)

	mux.Handle(
		e00_basics.HelloWorldTipTapPath,
		wb.Page(
			tiptapLayout(
				e00_basics.HelloWorldTipTap,
			),
		),
	)

	mux.Handle(
		e13_vuetify_list.HelloVuetifyListPath,
		wb.Page(
			demoVuetifyLayout(
				e13_vuetify_list.HelloVuetifyList,
			),
		),
	)

	mux.Handle(
		e14_vuetify_menu.HelloVuetifyMenuPath,
		wb.Page(
			demoVuetifyLayout(
				e14_vuetify_menu.HelloVuetifyMenu,
			),
		),
	)

	mux.Handle(
		e11_vuetify_basic_inputs.VuetifyBasicInputsPath,
		wb.Page(
			demoVuetifyLayout(
				e11_vuetify_basic_inputs.VuetifyBasicInputs,
			),
		),
	)

	mux.Handle(
		e10_vuetify_autocomplete.VuetifyAutoCompletePath,
		wb.Page(
			demoVuetifyLayout(
				e10_vuetify_autocomplete.VuetifyAutocomplete,
			),
		),
	)

	mux.Handle(
		e15_vuetify_navigation_drawer.VuetifyNavigationDrawerPath,
		wb.Page(
			demoVuetifyLayout(
				e15_vuetify_navigation_drawer.VuetifyNavigationDrawer,
			),
		),
	)

	// @snippet_begin(MountPresetHelloWorldSample)
	mux.Handle(
		e21_present_hello_world.PresetHelloWorldPath+"/",
		e21_present_hello_world.PresetHelloWorld01(),
	)
	// @snippet_end

	home := secs[0].items[0]
	mux.Handle("/", ub.Page(layout(rf(home.doc, home), secs, prefix, home)))
	return mux
}
