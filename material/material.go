package material

import (
	"github.com/gobuffalo/packr"
	ui "github.com/sunfmin/page"
)

var assetsbox = packr.NewBox("./css/")

func ComponentsPacks() []ui.ComponentsPack {
	v, err := assetsbox.FindString("material.css")
	if err != nil {
		panic(err)
	}

	return []ui.ComponentsPack{ui.ComponentsPack(v)}
}
