package e21_presents

import (
	"fmt"

	"github.com/goplaid/web"
	"github.com/goplaid/x/presets"
	"github.com/goplaid/x/presets/actions"
	. "github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
	"gorm.io/gorm"
)

// @snippet_begin(PresetsModelBuilderExtensionsSample)

func PresetsModelBuilderExtensions(b *presets.Builder) (
	mb *presets.ModelBuilder,
	db *gorm.DB,
) {
	mb, db = PresetsHelloWorld(b)
	b.URIPrefix(PresetsModelBuilderExtensionsPath)

	eb := mb.Editing("Actions", "Name").ActionsFunc(func(ctx *web.EventContext) h.HTMLComponent {
		return h.Components(
			VSpacer(),
			VBtn("Action 1"),
			VBtn("Action 2"),
			VBtn("Update").
				Color("primary").
				Attr("@click", web.Plaid().
					EventFunc(actions.Update, ctx.Event.Params...).
					URL(mb.Info().ListingHref()).
					Go()),
		)
	})

	eb.Field("Actions").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		cust := obj.(*Customer)
		return VBtn("Change Name").Attr("@click",
			web.Plaid().
				EventFunc("changeName", fmt.Sprint(cust.ID)).
				Go(),
		)
	})

	eb.ValidateFunc(func(obj interface{}, ctx *web.EventContext) (err web.ValidationErrors) {
		cust := obj.(*Customer)
		if len(cust.Name) < 5 {
			err.GlobalError("Name must be longer than 5")
		}
		return
	})

	mb.RegisterEventFunc("changeName", changeNameEventFunc(mb))

	return
}

func changeNameEventFunc(mb *presets.ModelBuilder) web.EventFunc {
	return func(ctx *web.EventContext) (r web.EventResponse, err error) {
		eb := mb.Editing()
		obj := mb.NewModel()
		id := ctx.Event.Params[0]
		obj, err = eb.Fetcher(obj, id, ctx)
		obj.(*Customer).Name = "Darwin"
		err = eb.Saver(obj, id, ctx)
		presets.ShowMessage(&r, "Nicely updated", "")
		eb.UpdateRightDrawerContent(ctx, &r, obj, "Good work", err)
		return
	}
}

const PresetsModelBuilderExtensionsPath = "/samples/presets-model-builder-extensions"

// @snippet_end
