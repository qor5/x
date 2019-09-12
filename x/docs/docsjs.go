package docs

import (
	"github.com/gobuffalo/packr"
	"github.com/goplaid/web"
)

var box = packr.NewBox("./docsjs/dist/")

func JSComponentsPack() web.ComponentsPack {
	v, err := box.FindString("docs.umd.min.js")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}

func CSSComponentsPack() web.ComponentsPack {
	v, err := box.FindString("docs.css")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}

func JSVueComponentsPack() web.ComponentsPack {
	v, err := box.FindString("vue.min.js")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}
