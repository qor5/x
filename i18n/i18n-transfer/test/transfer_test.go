package test

import (
	"os"
	"path/filepath"
	"testing"

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
