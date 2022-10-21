package admin

import (
	"github.com/goplaid/x/presets"
	"net/http"
)

func SetupRouter(b *presets.Builder) (mux *http.ServeMux) {
	mux = http.NewServeMux()
	mux.Handle("/", b)
	return
}
