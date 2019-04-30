package bran

import (
	"github.com/gobuffalo/packr"
)

var corejsassetsbox = packr.NewBox("./corejs/dist/")

func ComponentsPacks() []ComponentsPack {
	v, err := corejsassetsbox.FindString("js/chunk-vendors.js")
	if err != nil {
		panic(err)
	}

	r, err := corejsassetsbox.FindString("js/app.js")
	if err != nil {
		panic(err)
	}

	return []ComponentsPack{ComponentsPack(v), ComponentsPack(r)}
}
