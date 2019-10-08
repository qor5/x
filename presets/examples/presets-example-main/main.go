package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/jinzhu/gorm"
	examples2 "github.com/goplaid/x/presets/examples"
)

func main() {
	db, err := gorm.Open("postgres", os.Getenv("DBString"))
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	p := examples2.Preset1(db)

	log.Println("serving on :7000")
	log.Fatal(http.ListenAndServe(":7000", middleware.Logger(
		middleware.RequestID(p))))
}
