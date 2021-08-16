package e22_vuetify_variant_sub_form

// @snippet_begin(VuetifyVariantSubForm)

import (
	"github.com/goplaid/web"
	"github.com/goplaid/x/docs/utils"
	. "github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
)

type myFormValue struct {
	Type  string
	Form1 struct {
		Gender string
	}
	Form2 struct {
		Feature1 bool
		Slider1  int
	}
}

func VuetifyVariantSubForm(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("switchForm", switchForm)
	ctx.Hub.RegisterEventFunc("submit", submit)

	var fv myFormValue
	ctx.MustUnmarshalForm(&fv)
	if fv.Type == "" {
		fv.Type = "Type1"
	}
	var verr web.ValidationErrors

	pr.Body = VContainer(
		utils.PrettyFormAsJSON(ctx),

		web.Bind(
			VSelect().
				Items([]string{
					"Type1",
					"Type2",
				}).
				FieldName("Type").
				Value(fv.Type),
		).
			On("input").
			EventFunc("switchForm"),
		web.Portal(
			h.If(fv.Type == "Type1",
				form1(ctx, &fv),
			).Else(
				form2(ctx, &fv, &verr),
			),
		).Name("subform"),

		VBtn("Submit").OnClick("submit"),
	)
	return
}

func form1(ctx *web.EventContext, fv *myFormValue) h.HTMLComponent {

	return VContainer(
		h.H1("Form1"),
		VRadioGroup(
			VRadio().Value("F").Label("Female"),
			VRadio().Value("M").Label("Male"),
		).FieldName("Form1.Gender").
			Value(fv.Form1.Gender).
			Label("Gender"),
	)
}
func form2(ctx *web.EventContext, fv *myFormValue, verr *web.ValidationErrors) h.HTMLComponent {

	return VContainer(
		h.H1("Form2"),

		VSwitch().
			FieldName("Form2.Feature1").
			InputValue(fv.Form2.Feature1).
			Label("Feature1"),

		VSlider().FieldName("Form2.Slider1").
			ErrorMessages(verr.GetFieldErrors("Slider1")...).
			Value(fv.Form2.Slider1).
			Label("Slider1"),
	)
}

func submit(ctx *web.EventContext) (r web.EventResponse, err error) {
	r.Reload = true
	return
}

func switchForm(ctx *web.EventContext) (r web.EventResponse, err error) {
	var verr web.ValidationErrors

	var fv myFormValue
	ctx.MustUnmarshalForm(&fv)
	form := form1(ctx, &fv)
	if fv.Type == "Type2" {
		form = form2(ctx, &fv, &verr)
	}

	r.UpdatePortals = append(r.UpdatePortals, &web.PortalUpdate{
		Name: "subform",
		Body: form,
	})

	return
}

// @snippet_end

const VuetifyVariantSubFormPath = "/samples/vuetify-variant-sub-form"
