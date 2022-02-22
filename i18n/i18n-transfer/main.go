package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/goplaid/x/i18n/i18n-transfer/csv"
	"github.com/goplaid/x/i18n/i18n-transfer/parser"
)

func main() {
	importCmd := flag.NewFlagSet("import", flag.ExitOnError)
	importCsv := importCmd.String("csv", "", "input csv file path")

	if len(os.Args) < 2 {
		fmt.Println("expected 'import' or 'export' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "import":
		importCmd.Parse(os.Args[2:])
		if *importCsv == "" {
			fmt.Println(`flag needs an argument: -csv
Usage of import:
  -csv string
    	input csv file path`)
			os.Exit(1)
		}
		translationMap, err := csv.GetTranslationsMap(*importCsv)
		if err != nil {
			log.Fatalln(err)
		}
		err = parser.ImportFromTranslationsMap("./", translationMap)
		if err != nil {
			log.Fatalln(err)
		}
	case "export":
		translationsMap, err := parser.ExportToTranslationsMap("./")
		if err != nil {
			log.Fatalln(err)
		}
		for locale, translationMap := range translationsMap {
			fmt.Println(locale)
			for k, v := range translationMap {
				fmt.Printf("    %v: %v\n", k, v)
			}
		}
		err = csv.ExportToCsv(translationsMap)
		if err != nil {
			log.Fatalln(err)
		}
	default:
		fmt.Println("expected 'import' or 'export' subcommands")
		os.Exit(1)
	}
}
