package docs

import (
	"embed"
	"github.com/goplaid/web"
)

//go:embed docsjs/dist
var box embed.FS

func JSComponentsPack() web.ComponentsPack {
	v, err := box.ReadFile("docsjs/dist/docs.umd.min.js")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}

func CSSComponentsPack() web.ComponentsPack {
	v, err := box.ReadFile("docsjs/dist/docs.css")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}
