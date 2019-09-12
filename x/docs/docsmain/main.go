package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/goplaid/x/docs/setup"
)

func main() {
	mux := setup.Setup("")
	// @snippet_begin(HelloWorldMainSample)
	fmt.Println("Starting docs at :9000")
	http.Handle("/", mux)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}
	// @snippet_end
}
