package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/goplaid/x/docs"
)

func main() {
	mux := docs.Mux("/")
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9100"
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
