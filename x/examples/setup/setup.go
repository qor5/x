package setup

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"

	e01_hello_button2 "github.com/sunfmin/bran/x/examples/e01_hello_button"
	e02_hello_material_button2 "github.com/sunfmin/bran/x/examples/e02_hello_material_button"
	e04_hello_material_grid2 "github.com/sunfmin/bran/x/examples/e04_hello_material_grid"
	e06_hello_drawer2 "github.com/sunfmin/bran/x/examples/e06_hello_drawer"
	e07_hello_lazy_portal_in_drawer2 "github.com/sunfmin/bran/x/examples/e07_hello_lazy_portal_in_drawer"
	e08_hello_popover2 "github.com/sunfmin/bran/x/examples/e08_hello_popover"
	e09_hello_dialog2 "github.com/sunfmin/bran/x/examples/e09_hello_dialog"
	e10_hello_vuetify_autocomplete2 "github.com/sunfmin/bran/x/examples/e10_hello_vuetify_autocomplete"
	e11_hello_vuetify_text_field2 "github.com/sunfmin/bran/x/examples/e11_hello_vuetify_text_field"
	e12_hello_vuetify_grid2 "github.com/sunfmin/bran/x/examples/e12_hello_vuetify_grid"
	e13_hello_vuetify_list2 "github.com/sunfmin/bran/x/examples/e13_hello_vuetify_list"
	e14_hello_vuetify_menu2 "github.com/sunfmin/bran/x/examples/e14_hello_vuetify_menu"
	e15_hello_vuetify_navigation_drawer2 "github.com/sunfmin/bran/x/examples/e15_hello_vuetify_navigation_drawer"
	e16_hello_vuetify_simple_components2 "github.com/sunfmin/bran/x/examples/e16_hello_vuetify_simple_components"
	e17_hello_lazy_portals_and_reload2 "github.com/sunfmin/bran/x/examples/e17_hello_lazy_portals_and_reload"
	e18_filter_component2 "github.com/sunfmin/bran/x/examples/e18_filter_component"
	e19_stripeui_key_info2 "github.com/sunfmin/bran/x/examples/e19_stripeui_key_info"
	e20_vuetify_expansion_panels2 "github.com/sunfmin/bran/x/examples/e20_vuetify_expansion_panels"

	"github.com/gobuffalo/packr"
	"github.com/sunfmin/bran/web"
	"github.com/sunfmin/bran/x/codehighlight"
	m "github.com/sunfmin/bran/x/material"
	"github.com/sunfmin/bran/x/overlay"
	"github.com/sunfmin/bran/x/vuetify"
	"github.com/theplant/appkit/contexts"
	"github.com/theplant/appkit/server"
	. "github.com/theplant/htmlgo"
)

type pageItem struct {
	url        string
	renderFunc web.PageFunc
	vuetify    bool
}

func (p pageItem) Title() string {
	segs := strings.Split(p.url, "_")
	segs[1] = strings.Title(segs[1])
	return fmt.Sprintf("%s: %s", strings.ToUpper(segs[0]), strings.Join(segs[1:], " "))
}

func exampleLinks(prefix string, pages []pageItem) (comp HTMLComponent) {
	var links []HTMLComponent
	for _, p := range pages {
		links = append(links,
			Li(
				A().Href(fmt.Sprintf("%s/%s", prefix, p.url)).Text(p.Title()),
			),
		)
	}
	comp = Ul(links...)
	return
}

var exampleBox = packr.NewBox("../")

func layout(in web.PageFunc, pages []pageItem, prefix string, cp pageItem) (out web.PageFunc) {
	return func(ctx *web.EventContext) (pr web.PageResponse, err error) {

		tailScript := `<script src='/assets/main.js'></script>`
		if len(os.Getenv("DEV")) > 0 {
			fmt.Println("Using Dev environment, make sure you did: yarn start")
			tailScript = `
				<script src='http://localhost:3050/app.js'></script>
				<script src='http://localhost:3100/app.js'></script>
			`
		}

		ctx.Injector.Title(cp.Title())
		ctx.Injector.HeadHTML(`
			<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto+Mono" async>
			<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:300,400,500" async>
			<link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons" async>
			<link rel="stylesheet" href="/assets/main.css">
			<script src='/assets/vue.js'></script>
			<script src='/assets/codehighlight.js'></script>
		`)
		if cp.vuetify {
			ctx.Injector.HeadHTML(`
				<link rel="stylesheet" href="/assets/vuetify.css">
			`)

			vuetifyjs := `<script src='/assets/vuetify.js'></script>`
			if len(os.Getenv("DEV")) > 0 {
				vuetifyjs = `<script src='http://localhost:3080/app.js'></script>`
			}
			ctx.Injector.HeadHTML(vuetifyjs)
		}

		ctx.Injector.TailHTML(tailScript)

		var innerPr web.PageResponse
		innerPr, err = in(ctx)
		if err != nil {
			panic(err)
		}

		demo := innerPr.Schema

		var dacComps = []HTMLComponent{demo.(HTMLComponent)}

		var code string
		code, err = exampleBox.FindString(cp.url + "/page.go")
		if err != nil {
			return
		}
		if len(code) > 0 {
			dacComps = append(dacComps,
				Div(
					codehighlight.Code(code).Language("go"),
				).Class("exampleCode"),
			)
		}
		ctx.Injector.HeadHTML(`
		<style>
			body {
				margin: 0;
			}
			pre {
				padding: 0;
				margin: 0;
			}
			.exampleCode {
				margin-top: 20px;
			}
			[v-cloak] {
				display: none;
			}
		</style>
		`)

		pr.Schema = m.Grid(
			m.Cell(exampleLinks(prefix, pages)).Span(3, m.ScreenAll),
			m.Cell(dacComps...).Span(9, m.ScreenAll),
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
	ub := web.New()

	mux := http.NewServeMux()

	mux.Handle("/assets/main.js",
		ub.PacksHandler("text/javascript",
			overlay.JSComponentsPack(),
			web.JSComponentsPack(),
		),
	)

	mux.Handle("/assets/vue.js",
		ub.PacksHandler("text/javascript",
			web.JSVueComponentsPack(),
		),
	)

	mux.Handle("/assets/main.css",
		ub.PacksHandler("text/css",
			codehighlight.CSSComponentsPack(),
			overlay.CSSComponentsPack(),
			m.CSSComponentsPack(),
		),
	)

	mux.Handle("/assets/codehighlight.js",
		ub.PacksHandler("text/javascript",
			codehighlight.JSComponentsPack(),
		),
	)

	mux.Handle("/assets/vuetify.js",
		ub.PacksHandler("text/javascript",
			vuetify.JSComponentsPack(),
		),
	)

	mux.Handle("/favicon.ico", http.NotFoundHandler())

	mux.Handle("/assets/vuetify.css",
		ub.PacksHandler("text/css",
			vuetify.CSSComponentsPack(),
		),
	)

	var pages = []pageItem{
		{
			url:        "e01_hello_button",
			renderFunc: e01_hello_button2.HelloButton,
		},
		{
			url:        "e02_hello_material_button",
			renderFunc: e02_hello_material_button2.HelloButton,
		},
		//{
		//	url:        "e03_hello_card",
		//	renderFunc: e03_hello_card.HelloCard,
		//},
		{
			url:        "e04_hello_material_grid",
			renderFunc: e04_hello_material_grid2.HelloGrid,
		},
		//{
		//	url:        "e05_hello_customized_component",
		//	renderFunc: e05_hello_customized_component.HelloCustomziedComponent,
		//},
		{
			url:        "e06_hello_drawer",
			renderFunc: e06_hello_drawer2.HelloDrawer,
		},
		{
			url:        "e07_hello_lazy_portal_in_drawer",
			renderFunc: e07_hello_lazy_portal_in_drawer2.HelloLazyLoaderInDrawer,
		},
		{
			url:        "e08_hello_popover",
			renderFunc: e08_hello_popover2.HelloPopover,
		},
		{
			url:        "e09_hello_dialog",
			renderFunc: e09_hello_dialog2.HelloDialog,
		},
		{
			url:        "e10_hello_vuetify_autocomplete",
			renderFunc: e10_hello_vuetify_autocomplete2.HelloVuetifyAutocomplete,
			vuetify:    true,
		},
		{
			url:        "e11_hello_vuetify_text_field",
			renderFunc: e11_hello_vuetify_text_field2.HelloVuetifyTextField,
			vuetify:    true,
		},
		{
			url:        "e12_hello_vuetify_grid",
			renderFunc: e12_hello_vuetify_grid2.HelloVuetifyGrid,
			vuetify:    true,
		},
		{
			url:        "e13_hello_vuetify_list",
			renderFunc: e13_hello_vuetify_list2.HelloVuetifyList,
			vuetify:    true,
		},
		{
			url:        "e14_hello_vuetify_menu",
			renderFunc: e14_hello_vuetify_menu2.HelloVuetifyMenu,
			vuetify:    true,
		},
		{
			url:        "e15_hello_vuetify_navigation_drawer",
			renderFunc: e15_hello_vuetify_navigation_drawer2.HelloVuetifyNavigationDrawer,
			vuetify:    true,
		},
		{
			url:        "e16_hello_vuetify_simple_components",
			renderFunc: e16_hello_vuetify_simple_components2.HelloVuetifySimpleComponents,
			vuetify:    true,
		},
		{
			url:        "e17_hello_lazy_portals_and_reload",
			renderFunc: e17_hello_lazy_portals_and_reload2.HelloLazyPortalsAndReload,
			vuetify:    true,
		},
		{
			url:        "e18_filter_component",
			renderFunc: e18_filter_component2.FilterComponent,
			vuetify:    true,
		},
		{
			url:        "e19_stripeui_key_info",
			renderFunc: e19_stripeui_key_info2.KeyInfoDemo,
			vuetify:    true,
		},
		{
			url:        "e20_vuetify_expansion_panels",
			renderFunc: e20_vuetify_expansion_panels2.ExpansionPanelDemo,
			vuetify:    true,
		},
	}

	// l := log.Default()
	mw := server.Compose(
		// server.LogRequest,
		// log.WithLogger(l),
		contexts.WithHTTPStatus,
	)

	for _, p := range pages {
		mux.Handle(
			fmt.Sprintf("/%s", p.url),
			mw(ub.Page(layout(p.renderFunc, pages, prefix, p))),
		)
	}

	mux.Handle("/", home(prefix, pages))
	return mux
}
