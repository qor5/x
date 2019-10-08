package overlay

import (
	"github.com/gobuffalo/packr"
	"github.com/goplaid/web"
)

var overlayjsassetsbox = packr.NewBox("./overlayjs/dist/")

func JSComponentsPack() web.ComponentsPack {
	v, err := overlayjsassetsbox.FindString("branoverlay.umd.min.js")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}

func CSSComponentsPack() web.ComponentsPack {
	v, err := overlayjsassetsbox.FindString("branoverlay.css")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}
