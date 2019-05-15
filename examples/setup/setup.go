package setup

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"

	"github.com/sunfmin/branoverlay"

	"github.com/gobuffalo/packr"
	"github.com/sunfmin/bran"
	"github.com/sunfmin/bran/examples/e01_hello_button"
	"github.com/sunfmin/bran/examples/e02_hello_material_button"
	"github.com/sunfmin/bran/examples/e03_hello_card"
	"github.com/sunfmin/bran/examples/e04_hello_material_grid"
	"github.com/sunfmin/bran/examples/e05_hello_customized_component"
	"github.com/sunfmin/bran/examples/e06_hello_drawer"
	"github.com/sunfmin/bran/ui"
	m "github.com/sunfmin/material"
	"github.com/theplant/appkit/contexts"
	"github.com/theplant/appkit/server"
	. "github.com/theplant/htmlgo"
)

type pageItem struct {
	url         string
	renderFunc  ui.PageFunc
	mui         bool
	withoutCard bool
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
				A().Href(fmt.Sprintf("%s/%s/", prefix, p.url)).Text(p.Title()),
			),
		)
	}
	comp = Ul(links...)
	return
}

var exampleBox = packr.NewBox("../")

func layout(in ui.PageFunc, pages []pageItem, prefix string, cp pageItem) (out ui.PageFunc) {
	return func(ctx *ui.EventContext) (pr ui.PageResponse, err error) {

		tailScript := `<script src='/assets/main.js'></script>`
		if len(os.Getenv("DEV")) > 0 {
			fmt.Println("Using Dev environment, make sure you did: yarn start")
			tailScript = `<script src='http://localhost:3100/app.js'></script>`
		}

		ctx.Injector.Title(cp.Title())
		ctx.Injector.PutHeadHTML(`
			<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto+Mono">
			<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:300,400,500">
			<link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">
			<script src="https://unpkg.com/vue"></script>
			<link rel="stylesheet" href="/assets/overlay.css">
		`)
		ctx.Injector.PutTailHTML(tailScript)

		var innerPr ui.PageResponse
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
				Pre(code).Style("font-family: monospace"),
			)
		}
		ctx.Injector.PutStyle(`
			pre {
				padding: 24px;
				background-color: #eee;
			}
		`)

		pr.Schema = m.Grid(
			m.Cell(exampleLinks(prefix, pages)).Span(3, m.ScreenAll),
			m.Cell(dacComps...).Span(9, m.ScreenAll),
		)

		pr.State = innerPr.State

		return
	}
}

func home(prefix string, pages []pageItem) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "e01_hello_button/", 302)
		return
	}
}

func Setup(prefix string) http.Handler {
	ub := bran.New()

	mux := http.NewServeMux()

	mux.Handle("/assets/main.js",
		ub.PacksHandler("text/javascript",
			branoverlay.JSComponentsPack(),
			bran.JSComponentsPack(),
		),
	)

	mux.Handle("/assets/overlay.css",
		ub.PacksHandler("text/css", branoverlay.CSSComponentsPack()))

	var pages = []pageItem{
		{
			url:        "e01_hello_button",
			renderFunc: e01_hello_button.HelloButton,
		},
		{
			url:        "e02_hello_material_button",
			renderFunc: e02_hello_material_button.HelloButton,
		},
		{
			url:        "e03_hello_card",
			renderFunc: e03_hello_card.HelloCard,
		},
		{
			url:        "e04_hello_material_grid",
			renderFunc: e04_hello_material_grid.HelloGrid,
		},
		{
			url:        "e05_hello_customized_component",
			renderFunc: e05_hello_customized_component.HelloCustomziedComponent,
		},
		{
			url:        "e06_hello_drawer",
			renderFunc: e06_hello_drawer.HelloDrawer,
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
			fmt.Sprintf("/%s/", p.url),
			bran.StripPrefix(
				fmt.Sprintf("/%s", p.url), mw(ub.Page(layout(p.renderFunc, pages, prefix, p))),
			),
			// mw(ub.NewPage().RenderFunc(layout(p.renderFunc, pages, prefix, p)).Handler()),
		)
	}

	mux.Handle("/", home(prefix, pages))
	return mux
}
