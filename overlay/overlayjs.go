package branoverlay

import (
	"github.com/sunfmin/bran"
	"github.com/gobuffalo/packr"
)

var overlayjsassetsbox = packr.NewBox("./overlayjs/dist/")

func JSComponentsPack() bran.ComponentsPack {
	v, err := overlayjsassetsbox.FindString("branoverlay.umd.min.js")
	if err != nil {
		panic(err)
	}

	return bran.ComponentsPack(v)
}

func CSSComponentsPack() bran.ComponentsPack {
	v, err := overlayjsassetsbox.FindString("branoverlay.css")
	if err != nil {
		panic(err)
	}

	return bran.ComponentsPack(v)
}
