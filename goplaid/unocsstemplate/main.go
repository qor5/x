package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/goplaid/web"
	. "github.com/theplant/htmlgo"
)

const doAction1 = "doAction1"

func Home(ctx *web.EventContext) (r web.PageResponse, err error) {
	r.Body = Div(
		Div(
			H1(time.Now().String()).Class("transition duration-500 transform hover:scale-105"),
			Button("Server side action").Attr("@click",
				web.Plaid().EventFunc(doAction1).Query("id", "1").Go(),
			).Class("px-5 py-3 rounded-xl text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-800 active:bg-grey-900 focus:outline-none border-4 border-white focus:border-purple-200 transition-all"),
			Div(
				Button("Browser side action").Attr("@click", "locals.count++").
					Class("px-5 py-3 rounded-xl text-sm font-medium text-indigo-600 bg-white outline-none focus:outline-none m-1 hover:m-0 focus:m-0 border border-indigo-600 hover:border-4 focus:border-4 hover:border-indigo-800 hover:text-indigo-800 focus:border-purple-200 active:border-grey-900 active:text-grey-900 transition-all"),
				Span("{{locals.count}}"),
			).Attr("v-init-context:locals", "{count: 0}"),
		).Class("p-5 w-80 bg-white rounded-lg shadow-lg hover:shadow-2xl cursor-pointer"),
	).Class("flex items-center justify-center")
	return
}

func DoAction1(ctx *web.EventContext) (r web.EventResponse, err error) {
	r.Reload = true
	return
}

func layout(in web.PageFunc) (out web.PageFunc) {
	return func(ctx *web.EventContext) (pr web.PageResponse, err error) {

		ctx.Injector.TailHTML(`
			<script src='/main.js'></script>
		`)

		ctx.Injector.HeadHTML(`
 		<link rel="stylesheet" href="/main.css" />
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

func main() {
	w := web.New()

	mux := http.NewServeMux()
	mux.Handle("/main.js",
		w.PacksHandler("text/javascript",
			web.JSVueComponentsPack(),
			web.JSComponentsPack()))
	mux.Handle("/main.css",
		w.PacksHandler("text/css",
			css()))

	mux.Handle("/", w.Page(layout(Home)).
		EventFuncs(doAction1, DoAction1))

	port := os.Getenv("PORT")
	if port == "" {
		port = "9010"
	}
	log.Printf("Listen on %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

//go:embed *.css
var front embed.FS

func css() web.ComponentsPack {
	v, err := front.ReadFile("uno.css")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}
