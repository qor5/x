package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/sunfmin/bran/examples/setup"
)

func main() {
	h := setup.Setup("")
	fmt.Println("Starting examples at :8080")
	http.Handle("/", h)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
