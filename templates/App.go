package templates

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
)

func App(dev bool, prefix string, heads string, body string) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n\n<!DOCTYPE html>\n<html>\n\t<head>\n\t")
	_buffer.WriteString((heads))
	_buffer.WriteString("\n\t</head>\n\t<body class=\"front\">\n\t")
	_buffer.WriteString((body))
	_buffer.WriteString("\n\t")
	if dev {

		_buffer.WriteString("<script type=\"text/javascript\" src=\"http://localhost:3001/main.js\"></script>")

		_buffer.WriteString("<script type=\"text/javascript\" src=\"http://localhost:3002/main.js\"></script>")

		_buffer.WriteString("<script type=\"text/javascript\" src=\"http://localhost:3000/main.js\"></script>")

	} else {

		_buffer.WriteString("<script type=\"text/javascript\" src=\"")
		_buffer.WriteString(gorazor.HTMLEscape(prefix))
		_buffer.WriteString("/main.js\"></script>")

	}
	_buffer.WriteString("\n\t</body>\n</html>")

	return _buffer.String()
}
