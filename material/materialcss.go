package material

import (
	"github.com/gobuffalo/packr"
	"github.com/goplaid/web"
)

var box = packr.NewBox("./dist/")

func CSSComponentsPack() web.ComponentsPack {
	v, err := box.FindString("material-components-web.min.css")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}
