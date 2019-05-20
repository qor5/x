package codehighlight

import (
	"github.com/gobuffalo/packr"
	"github.com/sunfmin/bran"
)

var assetsbox = packr.NewBox("./codehighlightjs/dist/")

func JSComponentsPack() bran.ComponentsPack {
	v, err := assetsbox.FindString("codehighlight.umd.min.js")
	if err != nil {
		panic(err)
	}

	return bran.ComponentsPack(v)
}

func CSSComponentsPack() bran.ComponentsPack {
	v, err := assetsbox.FindString("codehighlight.css")
	if err != nil {
		panic(err)
	}

	return bran.ComponentsPack(v)
}
