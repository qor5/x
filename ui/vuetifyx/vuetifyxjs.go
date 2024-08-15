package vuetifyx

import (
	"embed"
	"net/http"
	"strings"

	"github.com/qor5/web/v3"
	"github.com/theplant/osenv"
)

//go:embed vuetifyxjs/dist
var vuetifyx embed.FS

//go:embed buildinAsserts
var assetsbox embed.FS

func JSComponentsPack() web.ComponentsPack {
	v, err := vuetifyx.ReadFile("vuetifyxjs/dist/vuetifyx.min.js")
	if err != nil {
		panic(err)
	}

	return web.ComponentsPack(v)
}

var customizeVuetifyCSS = osenv.GetBool("CUSTOMIZE_VUETIFY_CSS", "Use customized styles for vuetify", true)

func cssComponentsPack() web.ComponentsPack {
	var v []byte
	var err error
	if customizeVuetifyCSS {
		v, err = vuetifyx.ReadFile("vuetifyxjs/dist/assets/vuetifyx.min.css")
	} else {
		v, err = assetsbox.ReadFile("buildinAsserts/vuetify.min.3.6.14.css")
	}
	if err != nil {
		panic(err)
	}
	return web.ComponentsPack(v)
}

func fontEot() web.ComponentsPack {
	v, err := vuetifyx.ReadFile("vuetifyxjs/dist/assets/materialdesignicons-webfont.eot")
	if err != nil {
		panic(err)
	}
	return web.ComponentsPack(v)
}

func fontTtf() web.ComponentsPack {
	v, err := vuetifyx.ReadFile("vuetifyxjs/dist/assets/materialdesignicons-webfont.ttf")
	if err != nil {
		panic(err)
	}
	return web.ComponentsPack(v)
}

func fontWoff() web.ComponentsPack {
	v, err := vuetifyx.ReadFile("vuetifyxjs/dist/assets/materialdesignicons-webfont.woff")
	if err != nil {
		panic(err)
	}
	return web.ComponentsPack(v)
}

func fontWoff2() web.ComponentsPack {
	v, err := vuetifyx.ReadFile("vuetifyxjs/dist/assets/materialdesignicons-webfont.woff2")
	if err != nil {
		panic(err)
	}
	return web.ComponentsPack(v)
}

type muxer interface {
	Handle(pattern string, handler http.Handler)
}

func HandleMaterialDesignIcons(prefix string, mux muxer) {
	mux.Handle(prefix+"/vuetify/assets/index.css", web.PacksHandler(
		"text/css",
		web.ComponentsPack(
			strings.ReplaceAll(string(cssComponentsPack()), "/assets/materialdesignicons", prefix+"/assets/materialdesignicons")),
	))
	mux.Handle(prefix+"/assets/materialdesignicons-webfont.eot", web.PacksHandler("application/vnd.ms-fontobject",
		fontEot()))
	mux.Handle(prefix+"/assets/materialdesignicons-webfont.ttf", web.PacksHandler("font/ttf", fontTtf()))
	mux.Handle(prefix+"/assets/materialdesignicons-webfont.woff", web.PacksHandler("font/woff", fontWoff()))
	mux.Handle(prefix+"/assets/materialdesignicons-webfont.woff2", web.PacksHandler("font/woff2", fontWoff2()))
}
