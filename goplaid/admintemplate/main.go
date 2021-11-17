package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/goplaid/x/goplaid/admintemplate/admin"
)

func main() {
	mux := admin.Router()
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	fmt.Println("Served at http://localhost:" + port + "/admin")
	http.Handle("/", mux)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
