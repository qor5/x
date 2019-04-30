package bran

import (
	"github.com/gobuffalo/packr"
	ui "github.com/sunfmin/pagui"
)

var corejsassetsbox = packr.NewBox("./corejs/dist/")

func ComponentsPacks() []ui.ComponentsPack {
	v, err := corejsassetsbox.FindString("js/chunk-vendors.js")
	if err != nil {
		panic(err)
	}

	r, err := corejsassetsbox.FindString("js/app.js")
	if err != nil {
		panic(err)
	}

	return []ui.ComponentsPack{ui.ComponentsPack(v), ui.ComponentsPack(r)}
}
