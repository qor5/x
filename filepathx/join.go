package filepathx

import (
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

func Join(base string, path string) (string, error) {
	fullpath := filepath.Join(base, path)
	relPath, err := filepath.Rel(base, fullpath)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get relative path")
	}
	if filepath.IsAbs(relPath) || strings.HasPrefix(relPath, "..") {
		return "", errors.Errorf("illegal file path: %s", path)
	}
	return fullpath, nil
}
