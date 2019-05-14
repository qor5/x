package bran

import (
	"github.com/gobuffalo/packr"
)

var corejsassetsbox = packr.NewBox("./corejs/dist/")

func ComponentsPacks() []ComponentsPack {
	v, err := corejsassetsbox.FindString("brancore.umd.min.js")
	if err != nil {
		panic(err)
	}

	return []ComponentsPack{ComponentsPack(v)}
}
