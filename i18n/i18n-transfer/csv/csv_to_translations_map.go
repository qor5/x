package csv

import (
	"encoding/csv"
	"os"
)

func CsvToTranslationsMap(csvPath string) (translationsMap map[string]map[string]string, err error) {
	f, err := os.Open(csvPath)
	if err != nil {
		return
	}

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		return
	}

	keyValueMap := make(map[int]map[int]string)
	for i, record := range records {
		for j, value := range record {
			if keyValueMap[i] == nil {
				keyValueMap[i] = make(map[int]string)
			}
			keyValueMap[i][j] = value
		}
	}

	translationsMap = make(map[string]map[string]string)
	for i, record := range records {
		for j, value := range record {
			if translationsMap[records[0][j]] == nil {
				translationsMap[records[0][j]] = make(map[string]string)
			}
			translationsMap[records[0][j]][records[i][0]] = value
		}
	}
	return
}
