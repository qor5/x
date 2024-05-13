package parser

import (
	go_path "path"
	"strings"
)

func getTranslationMapKey(path, keyName, projectPath string) string {
	temp := strings.Split(projectPath, "/")
	return strings.TrimPrefix(go_path.Join(path, keyName), strings.TrimSuffix(strings.Join(temp[:len(temp)-1], "/"), "/")+"/")
}
