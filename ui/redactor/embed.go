package redactor

import (
	"bytes"
	"embed"

	"github.com/qor5/web/v3"
)

//go:embed redactorjs
var box embed.FS

func JSComponentsPack() web.ComponentsPack {
	var js [][]byte
	j1, err := box.ReadFile("redactorjs/dist/redactorjs.umd.cjs")
	if err != nil {
		panic(err)
	}
	js = append(js, j1)
	return web.ComponentsPack(bytes.Join(js, []byte("\n\n")))
}

func CSSComponentsPack() web.ComponentsPack {
	c, err := box.ReadFile("redactorjs/dist/redactor.css")
	if err != nil {
		panic(err)
	}
	return web.ComponentsPack(c)
}
