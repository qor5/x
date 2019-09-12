package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	setup2 "github.com/goplaid/x/docs/setup"
)

func main() {
	h := setup2.Setup("")
	fmt.Println("Starting docs at :9000")
	http.Handle("/", h)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}
}
