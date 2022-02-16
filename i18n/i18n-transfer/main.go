package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/goplaid/x/i18n/i18n-transfer/export_to_csv"
	"github.com/goplaid/x/i18n/i18n-transfer/import_from_csv"
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
		translationMap, err := import_from_csv.GetTranslationsMap(*importCsv)
		if err != nil {
			log.Fatalln(err)
		}
		err = import_from_csv.ImportFromCsv("./", translationMap)
		if err != nil {
			log.Fatalln(err)
		}
	case "export":
		translationsMap, err := export_to_csv.GetTranslationsMap("./")
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%+v\n", translationsMap)
		err = export_to_csv.ExportToCsv(translationsMap)
		if err != nil {
			log.Fatalln(err)
		}
	default:
		fmt.Println("expected 'import' or 'export' subcommands")
		os.Exit(1)
	}
}
