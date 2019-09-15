package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/goplaid/x/docs/setup"
)

func main() {
	mux := setup.Setup("")
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9000"
	}
	// @snippet_begin(HelloWorldMainSample)
	fmt.Println("Starting docs at :" + port)
	http.Handle("/", mux)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
	// @snippet_end
}
