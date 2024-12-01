package filesystem

import (
	"testing"

	"github.com/qor5/x/v3/oss/tests"
)

func TestAll(t *testing.T) {
	fileSystem := New("/tmp")
	tests.TestAll(fileSystem, t)
}
