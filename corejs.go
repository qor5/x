package bran

import (
	"github.com/gobuffalo/packr"
)

var corejsassetsbox = packr.NewBox("./corejs/dist/")

func JSComponentsPack() ComponentsPack {
	v, err := corejsassetsbox.FindString("bran.umd.min.js")
	if err != nil {
		panic(err)
	}

	return ComponentsPack(v)
}

func JSVueComponentsPack() ComponentsPack {
	v, err := corejsassetsbox.FindString("vue.min.js")
	if err != nil {
		panic(err)
	}

	return ComponentsPack(v)
}

