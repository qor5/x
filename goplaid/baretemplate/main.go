package main

import (
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
		H1(time.Now().String()),
		Button("DoAction1").Attr("@click",
			web.Plaid().EventFunc(doAction1).Query("id", "1").Go(),
		))
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

	mux.Handle("/", w.Page(layout(Home)).
		EventFuncs(doAction1, DoAction1))

	port := os.Getenv("PORT")
	if port == "" {
		port = "9010"
	}
	log.Printf("Listen on %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
