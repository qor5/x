package core

import (
	"github.com/gobuffalo/packr"
	"github.com/sunfmin/bran"
)

var box = packr.NewBox("./corejs/dist/")

func JSComponentsPack() bran.ComponentsPack {
	v, err := box.FindString("bran.umd.min.js")
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
