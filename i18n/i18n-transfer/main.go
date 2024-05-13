package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/qor5/x/v3/i18n/i18n-transfer/csv"
	"github.com/qor5/x/v3/i18n/i18n-transfer/parser"
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
			s, err := os.Stat(input)
			if err != nil || s.IsDir() {
				return errors.New("Please input correct csv file path")
			}
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

		translationMap, err := csv.CsvToTranslationsMap(result)
		if err != nil {
			log.Fatalln(err)
		}

		projectPath, err := os.Getwd()
		if err != nil {
			fmt.Printf("Please run i18n-transfer in a correct project path\n")
			return
		}

		err = parser.ImportFromTranslationsMap(projectPath, translationMap)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if result == "Export" {
		projectPath, _ := os.Getwd()
		projectPath, _ = filepath.Abs(projectPath)

		translationsMap, err := parser.ExportToTranslationsMap(projectPath)
		if err != nil {
			log.Fatalln(err)
		}
		for locale, translationMap := range translationsMap {
			fmt.Println(locale)
			for k, v := range translationMap {
				fmt.Printf("    %v: %v\n", k, v)
			}
		}
		err = csv.TranslationsMapToCsv(translationsMap)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
