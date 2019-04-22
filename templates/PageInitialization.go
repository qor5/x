package templates

import (
	"bytes"
)

func PageInitialization(json string) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n\n<script type=\"text/javascript\">\n\nwindow.__serverSideData__=")
	_buffer.WriteString((json))
	_buffer.WriteString("\n\n</script>")

	return _buffer.String()
}
