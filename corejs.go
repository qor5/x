package bran

import (
	"github.com/gobuffalo/packr"
	"github.com/sunfmin/pagui"
)

var corejsassetsbox = packr.NewBox("./corejs/dist/")

func ComponentsPacks() []pagui.ComponentsPack {
	v, err := corejsassetsbox.FindString("js/chunk-vendors.js")
	if err != nil {
		panic(err)
	}

	r, err := corejsassetsbox.FindString("js/app.js")
	if err != nil {
		panic(err)
	}

	return []pagui.ComponentsPack{pagui.ComponentsPack(v), pagui.ComponentsPack(r)}
}
