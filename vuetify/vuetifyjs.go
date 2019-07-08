package vuetify

import (
	"github.com/gobuffalo/packr"
	"github.com/sunfmin/bran"
)

var assetsbox = packr.NewBox("./vuetifyjs/dist/")

func JSComponentsPack() bran.ComponentsPack {
	v, err := assetsbox.FindString("vuetifyjs.umd.min.js")
	if err != nil {
		panic(err)
	}

	return bran.ComponentsPack(v)
}

func CSSComponentsPack() bran.ComponentsPack {
	v, err := assetsbox.FindString("vuetifyjs.css")
	if err != nil {
		panic(err)
	}

	return bran.ComponentsPack(v)
}
