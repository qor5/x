package main

import (
	"log"
	"net/http"

	"github.com/sunfmin/bran/presets/examples"
)

func main() {
	p := examples.Preset1()
	log.Println("serving on :7000")
	log.Fatal(http.ListenAndServe(":7000", p))
}
