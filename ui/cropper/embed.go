package cropper

import (
	"embed"

	"github.com/qor5/web/v3"
)

//go:embed cropperjs/dist
var box embed.FS

func JSComponentsPack() web.ComponentsPack {
	v, err := box.ReadFile("cropperjs/dist/cropperjs.umd.cjs")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}

func CSSComponentsPack() web.ComponentsPack {
	v, err := box.ReadFile("cropperjs/dist/cropperjs.css")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}
