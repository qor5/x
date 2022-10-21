package admin

import (
	"github.com/goplaid/web"
	"github.com/goplaid/x/goplaid/admintemplate/models"
	"github.com/goplaid/x/presets"
	"github.com/goplaid/x/presets/gorm2op"
	"github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
	"net/http"
)

func Initialize() *http.ServeMux {
	b := initializeProject()
	mux := SetupRouter(b)

	return mux
}

func initializeProject() (b *presets.Builder) {
	db := ConnectDB()

	// Initialize the builder of GoPlaid
	b = presets.New()

	// Setup the project name, ORM and Homepage
	b.URIPrefix("/admin").
		BrandTitle("GoplaidPackageName").
		DataOperator(gorm2op.DataOperator(db)).
		HomePageFunc(func(ctx *web.EventContext) (r web.PageResponse, err error) {
			r.Body = vuetify.VContainer(
				h.H1("Home"),
				h.P().Text("Change your home page here"))
			return
		})

	// Register Post into the builder
	// Use m to customize the model, Or config more models here.
	m := b.Model(&models.Post{})
	_ = m

	return
}
