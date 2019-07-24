package overlay

import (
	"github.com/gobuffalo/packr"
	"github.com/sunfmin/bran"
)

var overlayjsassetsbox = packr.NewBox("./overlayjs/dist/")

func JSComponentsPack() bran.ComponentsPack {
	v, err := overlayjsassetsbox.FindString("overlay.umd.min.js")
	if err != nil {
		panic(err)
	}

	return bran.ComponentsPack(v)
}

func CSSComponentsPack() bran.ComponentsPack {
	v, err := overlayjsassetsbox.FindString("overlay.css")
	if err != nil {
		panic(err)
	}

	return bran.ComponentsPack(v)
}
