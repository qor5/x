package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/goplaid/x/i18n/i18n-transfer/csv"
	"github.com/goplaid/x/i18n/i18n-transfer/parser"
	"github.com/manifoldco/promptui"
)

func main() {
	prompt := promptui.Select{
		Label: "Import Or Export",
		Items: []string{"Import", "Export"},
	}

	_, result, err := prompt.Run()

	if err != nil || (result != "Import" && result != "Export") {
		fmt.Printf("Please select \"Import\" or \"Export\"\n")
		return
	}

	if result == "Import" {
		validate := func(input string) error {
			f, err := os.Open(input)
			if err != nil {
				return errors.New("Please input correct csv file path")
			}
			f.Close()
			return nil
		}

		prompt := promptui.Prompt{
			Label:    "Csv File Path",
			Validate: validate,
		}

		result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Please input correct csv file path\n")
			return
		}

		translationMap, err := csv.GetTranslationsMap(result)
		if err != nil {
			log.Fatalln(err)
		}
		err = parser.ImportFromTranslationsMap("./", translationMap)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if result == "Export" {
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
	}

}
