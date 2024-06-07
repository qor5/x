package vuetify

import (
	"embed"
	"fmt"
	"net/http"
	"strings"

	"github.com/qor5/web/v3"
	"github.com/theplant/osenv"
)

//go:embed dist
var assetsbox embed.FS

//go:embed vuetifyjs/dist
var vuetifyjs embed.FS

//go:embed vuetifyjs/src/plugins/theme.js
var themeFs embed.FS

var customizeVuetifyCSS = osenv.GetBool("CUSTOMIZE_VUETIFY_CSS", "Use customized styles for vuetify", true)

func JSComponentsPack() web.ComponentsPack {
	v, err := assetsbox.ReadFile("dist/vuetify.min.js")
	if err != nil {
		panic(err)
	}
	return web.ComponentsPack(v)
}

func theme() string {
	v, err := themeFs.ReadFile("vuetifyjs/src/plugins/theme.js")
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf(`{ theme: %s }`, strings.Replace(string(v), "export default", "", 1))
}

func cssComponentsPack() web.ComponentsPack {
	var v []byte
	var err error
	if customizeVuetifyCSS {
		v, err = vuetifyjs.ReadFile("vuetifyjs/dist/vuetify/assets/index.css")
	} else {
		v, err = assetsbox.ReadFile("dist/vuetify.min.css")
	}
	if err != nil {
		panic(err)
	}
	return web.ComponentsPack(v)
}

func fontEot() web.ComponentsPack {
	v, err := vuetifyjs.ReadFile("vuetifyjs/dist/vuetify/assets/materialdesignicons-webfont.eot")
	if err != nil {
		panic(err)
	}
	return web.ComponentsPack(v)
}

func fontTtf() web.ComponentsPack {
	v, err := vuetifyjs.ReadFile("vuetifyjs/dist/vuetify/assets/materialdesignicons-webfont.ttf")
	if err != nil {
		panic(err)
	}
	return web.ComponentsPack(v)
}

func fontWoff() web.ComponentsPack {
	v, err := vuetifyjs.ReadFile("vuetifyjs/dist/vuetify/assets/materialdesignicons-webfont.woff")
	if err != nil {
		panic(err)
	}
	return web.ComponentsPack(v)
}

func fontWoff2() web.ComponentsPack {
	v, err := vuetifyjs.ReadFile("vuetifyjs/dist/vuetify/assets/materialdesignicons-webfont.woff2")
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
			strings.ReplaceAll(string(cssComponentsPack()), "/vuetify/assets/materialdesignicons", prefix+"/vuetify/assets/materialdesignicons")),
	))
	mux.Handle(prefix+"/vuetify/assets/materialdesignicons-webfont.eot", web.PacksHandler("application/vnd.ms-fontobject",
		fontEot()))
	mux.Handle(prefix+"/vuetify/assets/materialdesignicons-webfont.ttf", web.PacksHandler("font/ttf", fontTtf()))
	mux.Handle(prefix+"/vuetify/assets/materialdesignicons-webfont.woff", web.PacksHandler("font/woff", fontWoff()))
	mux.Handle(prefix+"/vuetify/assets/materialdesignicons-webfont.woff2", web.PacksHandler("font/woff2", fontWoff()))
}

const initVuetify = `
window.__goplaidVueComponentRegisters = window.__goplaidVueComponentRegisters || [];
window.__goplaidVueComponentRegisters.push(function(app, vueOptions) {
		app.use(Vuetify.createVuetify({{vuetifyOpts}}));
	});
`

func Vuetify() web.ComponentsPack {
	return web.ComponentsPack(
		strings.NewReplacer("{{vuetifyOpts}}", theme()).
			Replace(initVuetify),
	)
}
