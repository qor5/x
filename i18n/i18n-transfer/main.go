package main

import (
	"fmt"
	"log"

	"github.com/goplaid/x/i18n/i18n-transfer/export_to_csv"
	"github.com/goplaid/x/i18n/i18n-transfer/parser"
)

func main() {
	translationsMap, err := parser.GetTranslationsMap("./")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", translationsMap)
	export_to_csv.ExportToCsv(translationsMap)
}
