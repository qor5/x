package slogx

import "io"

// SetDefaultWriter overrides the destination used by subsequent
// SetupDefaultLogger calls. Test-only — visible only at `go test`
// build time via the export_test.go convention.
func SetDefaultWriter(w io.Writer) (previous io.Writer) {
	previous = defaultWriter
	defaultWriter = w
	return previous
}
