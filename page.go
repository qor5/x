package vuibuilder

import (
	"github.com/gobuffalo/packr"
	ui "github.com/sunfmin/page"
)

var corejsassetsbox = packr.NewBox("./corejs/dist/")

func ComponentsPack() ui.ComponentsPack {
	r, err := corejsassetsbox.FindString("js/app.js")
	if err != nil {
		panic(err)
	}
	return ui.ComponentsPack(r)
}
