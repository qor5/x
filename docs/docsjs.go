package docs

import (
	"github.com/gobuffalo/packr"
	"github.com/sunfmin/bran"
)

var box = packr.NewBox("./docsjs/dist/")

func JSComponentsPack() bran.ComponentsPack {
	v, err := box.FindString("docs.umd.min.js")
	if err != nil {
		panic(err)
	}

	return bran.ComponentsPack(v)
}

func CSSComponentsPack() bran.ComponentsPack {
	v, err := box.FindString("docs.css")
	if err != nil {
		panic(err)
	}

	return bran.ComponentsPack(v)
}

func JSVueComponentsPack() bran.ComponentsPack {
	v, err := box.FindString("vue.min.js")
	if err != nil {
		panic(err)
	}

	return bran.ComponentsPack(v)
}
