package vuetifyx

import (
	"embed"

	"github.com/qor5/web/v3"
)

//go:embed vuetifyxjs/dist
var assetsbox embed.FS

func JSComponentsPack() web.ComponentsPack {
	v, err := assetsbox.ReadFile("vuetifyxjs/dist/vuetifyxjs.umd.js")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}
