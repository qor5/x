package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"

	"github.com/sunfmin/bran/presets/examples"
)

func main() {
	db, err := gorm.Open("postgres", os.Getenv("TEST_DB"))
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	p := examples.Preset1(db)

	log.Println("serving on :7000")
	log.Fatal(http.ListenAndServe(":7000", p))
}
