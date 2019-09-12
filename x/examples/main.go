package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	setup2 "github.com/sunfmin/bran/x/examples/setup"
)

func main() {
	h := setup2.Setup("")
	fmt.Println("Starting examples at :8080")
	http.Handle("/", h)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
