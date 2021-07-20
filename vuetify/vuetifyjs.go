package vuetify

import (
	"embed"
	"github.com/goplaid/web"
)

//go:embed vuetifyjs/dist
var assetsbox embed.FS

func JSComponentsPack() web.ComponentsPack {
	v, err := assetsbox.ReadFile("vuetifyjs/dist/vuetifyjs.umd.min.js")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}

func CSSComponentsPack() web.ComponentsPack {
	v, err := assetsbox.ReadFile("vuetifyjs/dist/vuetifyjs.css")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}
