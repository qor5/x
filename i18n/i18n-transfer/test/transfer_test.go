package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/qor5/x/v3/i18n/i18n-transfer/csv"

	"github.com/qor5/x/v3/i18n/i18n-transfer/parser"
	"github.com/stretchr/testify/assert"
)

func TestExport(t *testing.T) {
	projectPath, _ := os.Getwd()
	projectPath, _ = filepath.Abs(projectPath)
	filepath.Join(projectPath, "mock")
	translationsMap, err := parser.ExportToTranslationsMap(filepath.Join(projectPath, "mock"))
	assert.NoError(t, err)
	want := map[string]map[string]string{
		"Japanese": {
			"mock/messages/name":        "User JP",
			"mock/messages/Email":       "terry@theplant.jp",
			"mock/messages/PhoneNumber": "+100",
		},
		"Chinese": {
			"mock/messages/Email":       "terry@theplant.cn",
			"mock/messages/PhoneNumber": "+86",
			"mock/messages/name":        "User CN",
		},
	}
	assert.Equal(t, want, translationsMap)
}

func TestImport(t *testing.T) {
	projectPath, err := os.Getwd()
	projectPath = filepath.Join(projectPath, "mock")
	assert.NoError(t, err)

	// Store original data for restoration
	originalTranslations, err := parser.ExportToTranslationsMap(projectPath)
	assert.NoError(t, err)

	// Cleanup function to restore original data
	t.Cleanup(func() {
		err := parser.ImportFromTranslationsMap(projectPath, originalTranslations)
		assert.NoError(t, err)
	})

	translationMap, err := csv.CsvToTranslationsMap(projectPath + "/test_import.csv")
	assert.NoError(t, err)
	err = parser.ImportFromTranslationsMap(projectPath, translationMap)

	translationsMap, err := parser.ExportToTranslationsMap(projectPath)
	assert.NoError(t, err)
	want := map[string]map[string]string{
		"Japanese": {
			"mock/messages/name":        "New User JP",
			"mock/messages/Email":       "New JPEmail",
			"mock/messages/PhoneNumber": "+100000",
		},
		"Chinese": {
			"mock/messages/Email":       "New CNEmail",
			"mock/messages/PhoneNumber": "+8666",
			"mock/messages/name":        "New User CN",
		},
	}
	assert.Equal(t, want, translationsMap)
}
