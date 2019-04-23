package setup

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/sunfmin/bran"
	"github.com/sunfmin/bran/examples/e01_hello_button"
	"github.com/sunfmin/bran/examples/e02_hello_card"
	. "github.com/sunfmin/bran/html"
	"github.com/sunfmin/bran/material"
	ui "github.com/sunfmin/page"
	"github.com/theplant/appkit/contexts"
	"github.com/theplant/appkit/server"
)

type pageItem struct {
	url         string
	renderFunc  ui.PageRenderFunc
	mui         bool
	withoutCard bool
}

func (p pageItem) Title() string {
	segs := strings.Split(p.url, "_")
	segs[1] = strings.Title(segs[1])
	return fmt.Sprintf("%s: %s", strings.ToUpper(segs[0]), strings.Join(segs[1:], " "))
}

var exampleBox = packr.NewBox("../")

func layout(in ui.PageRenderFunc, pages []pageItem, prefix string, cp pageItem) (out ui.PageRenderFunc) {
	return func(ctx *ui.EventContext) (pr ui.PageResponse, err error) {

		ctx.Head.Title(cp.Title())
		ctx.Head.HTML(`<link rel="stylesheet" href="/assets/material.css">`)
		ctx.Head.HTML(`<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:300,400,500">`)

		var innerPr ui.PageResponse
		innerPr, err = in(ctx)
		if err != nil {
			panic(err)
		}

		root := innerPr.Schema

		pr.Schema = root

		pr.State = innerPr.State

		return
	}
}

func home(prefix string, pages []pageItem) ui.PageRenderFunc {
	return func(ctx *ui.EventContext) (pr ui.PageResponse, err error) {

		pr.Schema = Button("Home")

		return
	}
}

func Setup(prefix string) http.Handler {
	ub := ui.New().Prefix(
		prefix + "/assets",
	)

	if len(os.Getenv("DEV")) > 0 {
		fmt.Println("Using Dev environment, make sure you did: yarn start")
		ub.FrontDev(true)
	}

	mux := http.NewServeMux()

	mux.Handle("/assets/main.js",
		ub.PacksHandler("text/javascript", bran.ComponentsPacks()...))
	mux.Handle("/assets/material.css",
		ub.PacksHandler("text/css", material.ComponentsPacks()...))

	var pages = []pageItem{
		{
			url:        "e01_hello_button",
			renderFunc: e01_hello_button.HelloButton,
		},
		{
			url:        "e02_hello_card",
			renderFunc: e02_hello_card.HelloCard,
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
			ui.StripPrefix(
				fmt.Sprintf("/%s", p.url), mw(ub.NewPage().RenderFunc(layout(p.renderFunc, pages, prefix, p)).Handler()),
			),
			// mw(ub.NewPage().RenderFunc(layout(p.renderFunc, pages, prefix, p)).Handler()),
		)
	}

	mux.Handle("/", ub.NewPage().RenderFunc(home(prefix, pages)).Handler())
	return mux
}
