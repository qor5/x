package codehighlight

import (
	"github.com/gobuffalo/packr"
	"github.com/sunfmin/bran/web"
)

var assetsbox = packr.NewBox("./codehighlightjs/dist/")

func JSComponentsPack() web.ComponentsPack {
	v, err := assetsbox.FindString("codehighlight.umd.min.js")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}

func CSSComponentsPack() web.ComponentsPack {
	v, err := assetsbox.FindString("codehighlight.css")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}
