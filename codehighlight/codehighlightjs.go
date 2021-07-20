package codehighlight

import (
	"embed"
	"github.com/goplaid/web"
)

//go:embed codehighlightjs/dist
var assetsbox embed.FS

func JSComponentsPack() web.ComponentsPack {
	v, err := assetsbox.ReadFile("codehighlightjs/dist/codehighlight.umd.min.js")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}

func CSSComponentsPack() web.ComponentsPack {
	v, err := assetsbox.ReadFile("codehighlightjs/dist/codehighlight.css")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}
