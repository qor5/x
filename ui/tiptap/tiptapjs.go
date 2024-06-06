package tiptap

// @snippet_begin(TipTapPackrSample)
import (
	"embed"

	"github.com/qor5/web/v3"
)

//go:embed tiptapjs/dist
var box embed.FS

func JSComponentsPack() web.ComponentsPack {
	v, err := box.ReadFile("tiptapjs/dist/tiptapjs.umd.cjs")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}

func CSSComponentsPack() web.ComponentsPack {
	v, err := box.ReadFile("tiptapjs/dist/tiptap.css")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}

// @snippet_end
