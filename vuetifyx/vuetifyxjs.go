package vuetifyx

import (
	"embed"

	"github.com/goplaid/web"
)

//go:embed vuetifyxjs/dist
var assetsbox embed.FS

func JSComponentsPack() web.ComponentsPack {
	v, err := assetsbox.ReadFile("vuetifyxjs/dist/vuetifyxjs.umd.min.js")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}

//func CSSComponentsPack() web.ComponentsPack {
//	v, err := assetsbox.ReadFile("vuetifyxjs/dist/vuetifyxjs.css")
//	if err != nil {
//		panic(err)
//	}
//
//	return web.ComponentsPack(v)
//}
