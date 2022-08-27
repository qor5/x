package parser

import (
	go_path "path"
	"strings"
)

func getTranslationMapKey(path, keyName, projectDir string) string {
	return strings.TrimPrefix(go_path.Join(path, keyName), strings.TrimSuffix(projectDir, "/")+"/")
}
