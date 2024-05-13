package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

func TranslationsMapToCsv(translationsMap map[string]map[string]string) (err error) {
	var (
		locales             []string
		translationKeys     []string
		translationsKeysMap = map[string]bool{}
		now                 = time.Now().Format("20060102150405")
		filename            = fmt.Sprintf("translations.%v.csv", now)
	)

	// Sort locales
	for locale := range translationsMap {
		locales = append(locales, locale)
	}
	sort.Strings(locales)

	csvFile, err := os.Create(filename)
	defer csvFile.Close()
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	writer := csv.NewWriter(csvFile)

	// Append Headers
	if err = writer.Write(append([]string{"Translation Keys"}, locales...)); err != nil {
		log.Fatalln(err)
	}

	// Sort translation keys
	for _, locale := range locales {
		for key := range translationsMap[locale] {
			translationsKeysMap[key] = true
		}
	}

	for key := range translationsKeysMap {
		translationKeys = append(translationKeys, key)
	}
	sort.Strings(translationKeys)

	// Write CSV file

	index := 0

	for _, translationKey := range translationKeys {
		// Filter out translation by scope
		index++
		translations := []string{translationKey}
		for _, locale := range locales {
			translations = append(translations, translationsMap[locale][translationKey])
		}
		writer.Write(translations)
	}
	writer.Flush()

	fmt.Printf(`
----------------------------------------------
generate translation csv: 
	%s
`, filename)

	return
}
