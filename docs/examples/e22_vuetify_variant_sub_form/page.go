package e22_vuetify_variant_sub_form

// @snippet_begin(VuetifyVariantSubForm)

import (
	"github.com/goplaid/web"
	. "github.com/goplaid/x/vuetify"
)

type myFormValue struct {
	Type  string
}


func VuetifyVariantSubForm(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("update", update)

	var fv myFormValue
	ctx.MustUnmarshalForm(&fv)

	pr.Body = VContainer(

		web.Bind(

		VSelect().
			Items([]string{
					"Type1",
					"Type2",
				}).
			FieldName("Type").
			Value(fv.Type),
		).
			OnInput("update"),

	)
	return
}

func update(ctx *web.EventContext) (r web.EventResponse, err error) {
	r.Reload = true

	return
}

// @snippet_end

const VuetifyVariantSubFormPath = "/samples/vuetify-variant-sub-form"
