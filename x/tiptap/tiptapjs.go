package tiptap

// @snippet_begin(TipTapPackrSample)
import (
	"github.com/gobuffalo/packr"
	"github.com/goplaid/web"
)

var box = packr.NewBox("./tiptapjs/dist/")

func JSComponentsPack() web.ComponentsPack {
	v, err := box.FindString("tiptap.umd.min.js")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}

func CSSComponentsPack() web.ComponentsPack {
	v, err := box.FindString("tiptap.css")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}

// @snippet_end
