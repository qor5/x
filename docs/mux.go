package docs

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/middleware"
	"github.com/goplaid/web"
	"github.com/goplaid/x/docs/examples/e00_basics"
	"github.com/goplaid/x/docs/examples/e10_vuetify_autocomplete"
	"github.com/goplaid/x/docs/examples/e11_vuetify_basic_inputs"
	"github.com/goplaid/x/docs/examples/e13_vuetify_list"
	"github.com/goplaid/x/docs/examples/e14_vuetify_menu"
	"github.com/goplaid/x/docs/examples/e15_vuetify_navigation_drawer"
	"github.com/goplaid/x/docs/examples/e17_hello_lazy_portals_and_reload"
	"github.com/goplaid/x/docs/examples/e21_presents"
	"github.com/goplaid/x/docs/examples/e22_vuetify_variant_sub_form"
	"github.com/goplaid/x/docs/examples/e23_vuetify_components_kitchen"
	"github.com/goplaid/x/docs/root"
	"github.com/goplaid/x/docs/utils"
	"github.com/goplaid/x/presets"
	"github.com/goplaid/x/tiptap"
	v "github.com/goplaid/x/vuetify"
	"github.com/goplaid/x/vuetifyx"
	"github.com/theplant/docgo"
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

func addGA(ctx *web.EventContext) {
	if strings.Index(ctx.R.Host, "localhost") >= 0 {
		return
	}
	ctx.Injector.HeadHTML(`
<!-- Global site tag (gtag.js) - Google Analytics -->
<script async src="https://www.googletagmanager.com/gtag/js?id=UA-149605708-1"></script>
<script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());

  gtag('config', 'UA-149605708-1');
</script>
`)
}

func layout(in web.PageFunc, secs []*section, prefix string, cp *pageItem) (out web.PageFunc) {
	return func(ctx *web.EventContext) (pr web.PageResponse, err error) {
		addGA(ctx)
		pr.PageTitle = cp.title + " - " + "GoPlaid"

		ctx.Injector.HeadHTML(`
				<link rel="stylesheet" href="/assets/main.css">
			`)

		ctx.Injector.Title(cp.title)
		ctx.Injector.HeadHTML(`
			<script src='/assets/vue.js'></script>
			<script src='/assets/codehighlight.js'></script>
		`)

		ctx.Injector.TailHTML(coreJSTags)

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
		addGA(ctx)

		ctx.Injector.HeadHTML(`
			<script src='/assets/vue.js'></script>
		`)

		ctx.Injector.TailHTML(coreJSTags)
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
		addGA(ctx)

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
		addGA(ctx)

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

var coreJSTags = func() string {
	if len(os.Getenv("DEV_CORE_JS")) > 0 {
		return `
<script src='http://localhost:3100/js/chunk-vendors.js'></script>
<script src='http://localhost:3100/js/app.js'></script>
`
	}
	return `<script src='/assets/main.js'></script>`
}()

var vuetifyJSTags = func() string {
	if len(os.Getenv("DEV_VUETIFY_JS")) > 0 {
		return `
<script src='http://localhost:3080/js/chunk-vendors.js'></script>
<script src='http://localhost:3080/js/app.js'></script>
`
	}
	return `<script src='/assets/vuetify.js'></script>`
}()

// @snippet_begin(DemoVuetifyLayoutSample)
func demoVuetifyLayout(in web.PageFunc) (out web.PageFunc) {
	return func(ctx *web.EventContext) (pr web.PageResponse, err error) {
		addGA(ctx)

		ctx.Injector.HeadHTML(`
			<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto+Mono" async>
			<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:300,400,500" async>
			<link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons" async>
			<link rel="stylesheet" href="/assets/vuetify.css">
			<script src='/assets/vue.js'></script>
		`)

		ctx.Injector.TailHTML(fmt.Sprintf("%s %s", vuetifyJSTags, coreJSTags))
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
			v.VMain(
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

func Mux(prefix string) http.Handler {

	// @snippet_begin(HelloWorldMuxSample1)
	mux := http.NewServeMux()
	// @snippet_end

	// @snippet_begin(ComponentsPackSample)
	ub := web.New()
	mux.Handle("/assets/main.js",
		ub.PacksHandler("text/javascript",
			web.JSComponentsPack(),
		),
	)

	mux.Handle("/assets/vue.js",
		ub.PacksHandler("text/javascript",
			web.JSVueComponentsPack(),
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
			v.Vuetify(""),
			v.JSComponentsPack(),
			vuetifyx.JSComponentsPack(),
		),
	)

	mux.Handle("/assets/vuetify.css",
		ub.PacksHandler("text/css",
			v.CSSComponentsPack(),
		),
	)
	// @snippet_end

	mux.Handle("/favicon.ico", http.NotFoundHandler())

	samplesMux := SamplesHandler(prefix)
	mux.Handle("/samples/",
		middleware.Logger(
			middleware.RequestID(
				samplesMux,
			),
		),
	)

	mux.Handle("/",
		middleware.Logger(
			middleware.RequestID(
				docgo.New().
					SitePrefix(prefix).
					Assets("/assets/", root.Assets).
					Home(root.Home).
					Build(),
			),
		),
	)
	return mux
}

func SamplesHandler(prefix string) http.Handler {
	mux := http.NewServeMux()
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
		e00_basics.EventExamplePagePath,
		wb.Page(
			demoVuetifyLayout(
				e00_basics.ExamplePage,
			),
		),
	)

	mux.Handle(
		e00_basics.EventHandlingPagePath,
		wb.Page(
			demoVuetifyLayout(
				e00_basics.EventHandlingPage,
			),
		),
	)

	mux.Handle(
		e00_basics.WebScopeUseLocalsPagePath,
		wb.Page(
			demoVuetifyLayout(
				e00_basics.UseLocals,
			),
		),
	)

	mux.Handle(
		e00_basics.WebScopeUsePlaidFormPagePath,
		wb.Page(
			demoLayout(
				e00_basics.UsePlaidForm,
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
		e22_vuetify_variant_sub_form.VuetifyVariantSubFormPath,
		wb.Page(
			demoVuetifyLayout(
				e22_vuetify_variant_sub_form.VuetifyVariantSubForm,
			),
		),
	)

	mux.Handle(
		e23_vuetify_components_kitchen.VuetifyComponentsKitchenPath,
		wb.Page(
			demoVuetifyLayout(
				e23_vuetify_components_kitchen.VuetifyComponentsKitchen,
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

	mux.Handle(
		e17_hello_lazy_portals_and_reload.LazyPortalsAndReloadPath,
		wb.Page(
			demoVuetifyLayout(
				e17_hello_lazy_portals_and_reload.LazyPortalsAndReload,
			),
		),
	)

	// @snippet_begin(MountPresetHelloWorldSample)
	c00 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsHelloWorld(c00)
	mux.Handle(
		e21_presents.PresetsHelloWorldPath+"/",
		c00,
	)
	// @snippet_end

	c01 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsListingCustomizationFields(c01)
	mux.Handle(
		e21_presents.PresetsListingCustomizationFieldsPath+"/",
		c01,
	)

	c02 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsListingCustomizationFilters(c02)
	mux.Handle(
		e21_presents.PresetsListingCustomizationFiltersPath+"/",
		c02,
	)

	c03 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsListingCustomizationTabs(c03)
	mux.Handle(
		e21_presents.PresetsListingCustomizationTabsPath+"/",
		c03,
	)

	c04 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsListingCustomizationBulkActions(c04)
	mux.Handle(
		e21_presents.PresetsListingCustomizationBulkActionsPath+"/",
		c04,
	)

	c05 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsEditingCustomizationDescription(c05)
	mux.Handle(
		e21_presents.PresetsEditingCustomizationDescriptionPath+"/",
		c05,
	)

	c06 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsEditingCustomizationFileType(c06)
	mux.Handle(
		e21_presents.PresetsEditingCustomizationFileTypePath+"/",
		c06,
	)

	c07 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsEditingCustomizationValidation(c07)
	mux.Handle(
		e21_presents.PresetsEditingCustomizationValidationPath+"/",
		c07,
	)

	c08 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsDetailPageTopNotes(c08)
	mux.Handle(
		e21_presents.PresetsDetailPageTopNotesPath+"/",
		c08,
	)

	c09 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsDetailPageDetails(c09)
	mux.Handle(
		e21_presents.PresetsDetailPageDetailsPath+"/",
		c09,
	)

	c10 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsDetailPageCards(c10)
	mux.Handle(
		e21_presents.PresetsDetailPageCardsPath+"/",
		c10,
	)

	c11 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsPermissions(c11)
	mux.Handle(
		e21_presents.PresetsPermissionsPath+"/",
		c11,
	)

	c12 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsModelBuilderExtensions(c12)
	mux.Handle(
		e21_presents.PresetsModelBuilderExtensionsPath+"/",
		c12,
	)

	c13 := presets.New().AssetFunc(addGA)
	e21_presents.PresetsBasicFilter(c13)
	mux.Handle(
		e21_presents.PresetsBasicFilterPath+"/",
		c13,
	)

	mux.Handle(
		e00_basics.ShortCutSamplePath,
		wb.Page(
			demoVuetifyLayout(
				e00_basics.ShortCutSample,
			),
		),
	)

	c14 := presets.New().AssetFunc(addGA)
	e00_basics.PresetsNotificationCenterSample(c14)
	mux.Handle(
		e00_basics.NotificationCenterSamplePath+"/",
		c14,
	)

	return mux
}
