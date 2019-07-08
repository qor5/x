package material

import (
	"github.com/gobuffalo/packr"
	"github.com/sunfmin/bran"
)

var box = packr.NewBox("./dist/")

func CSSComponentsPack() bran.ComponentsPack {
	v, err := box.FindString("material-components-web.min.css")
	if err != nil {
		panic(err)
	}

	return bran.ComponentsPack(v)
}

