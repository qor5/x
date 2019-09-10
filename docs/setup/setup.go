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
				Nav(
					Div(
						A().Href("/").Text("Learn").Class("nav-item"),
						A().Href("/").Text("Github").Class("nav-item"),
					).Class("nav-links"),
				).Class("g-layout"),
			).Class("global-nav"),
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
		}

		ctx.Injector.Title(cp.title)
		ctx.Injector.HeadHTML(`
			<link rel="stylesheet" href="/assets/main.css">
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
			title: "Overview",
			slug:  "overview",
			items: []*pageItem{
				{
					title:      "The Go HTML builder",
					slug:       "go-builder.html",
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
