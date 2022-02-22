package import_from_csv

import (
	"encoding/csv"
	"os"
)

func GetTranslationsMap(csvPath string) (translationMap map[string]map[string]string, err error) {
	f, err := os.Open(csvPath)
	if err != nil {
		return
	}

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		return
	}

	var keyValueMap = make(map[int]map[int]string)
	for i, record := range records {
		for j, value := range record {
			if keyValueMap[i] == nil {
				keyValueMap[i] = make(map[int]string)
			}
			keyValueMap[i][j] = value
		}
	}

	translationMap = make(map[string]map[string]string)
	for i, record := range records {
		for j, value := range record {
			if translationMap[records[0][j]] == nil {
				translationMap[records[0][j]] = make(map[string]string)
			}
			translationMap[records[0][j]][records[i][0]] = value
		}
	}
	return
}
