package vuetify

import (
	"github.com/gobuffalo/packr"
	"github.com/sunfmin/bran/web"
)

var assetsbox = packr.NewBox("./vuetifyjs/dist/")

func JSComponentsPack() web.ComponentsPack {
	v, err := assetsbox.FindString("vuetifyjs.umd.min.js")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}

func CSSComponentsPack() web.ComponentsPack {
	v, err := assetsbox.FindString("vuetifyjs.css")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}
