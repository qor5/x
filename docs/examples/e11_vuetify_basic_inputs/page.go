package e11_vuetify_basic_inputs

// @snippet_begin(VuetifyBasicInputsSample)
import (
	"mime/multipart"

	"github.com/goplaid/web"
	"github.com/goplaid/x/docs/utils"
	. "github.com/goplaid/x/vuetify"
)

type myFormValue struct {
	MyValue       string
	TextareaValue string
	Gender        string
	Agreed        bool
	Feature1      bool
	Slider1       int
	Files1        []*multipart.FileHeader
}

var s = &myFormValue{
	MyValue:       "123",
	TextareaValue: "This is textarea value",
	Gender:        "M",
	Agreed:        false,
	Feature1:      true,
	Slider1:       60,
}

func VuetifyBasicInputs(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("update", update)

	var verr web.ValidationErrors
	if ve, ok := ctx.Flash.(web.ValidationErrors); ok {
		verr = ve
	}

	pr.Body = VContainer(
		utils.PrettyFormAsJSON(ctx),
		VTextField().
			Label("Form ValueIs").
			Solo(true).
			Clearable(true).
			FieldName("MyValue").
			ErrorMessages(verr.GetFieldErrors("MyValue")...).
			Value(s.MyValue),
		VTextarea().FieldName("TextareaValue").
			ErrorMessages(verr.GetFieldErrors("TextareaValue")...).
			Solo(true).Value(s.TextareaValue),
		VRadioGroup(
			VRadio().Value("F").Label("Female"),
			VRadio().Value("M").Label("Male"),
		).FieldName("Gender").Value(s.Gender),
		VCheckbox().FieldName("Agreed").
			ErrorMessages(verr.GetFieldErrors("Agreed")...).
			Label("Agree").InputValue(s.Agreed),
		VSwitch().FieldName("Feature1").InputValue(s.Feature1),

		VSlider().FieldName("Slider1").
			ErrorMessages(verr.GetFieldErrors("Slider1")...).
			Value(s.Slider1),

		VFileInput().FieldName("Files1"),

		VBtn("Update").OnClick("update"),
	)

	return
}

func update(ctx *web.EventContext) (r web.EventResponse, err error) {
	s = &myFormValue{}
	ctx.MustUnmarshalForm(s)
	verr := web.ValidationErrors{}
	if len(s.MyValue) < 10 {
		verr.FieldError("MyValue", "my value is too small")
	}

	if len(s.TextareaValue) > 5 {
		verr.FieldError("TextareaValue", "textarea value is too large")
	}

	if !s.Agreed {
		verr.FieldError("Agreed", "You must agree the terms")
	}

	if s.Slider1 > 50 {
		verr.FieldError("Slider1", "You slide too much")
	}

	ctx.Flash = verr
	r.Reload = true

	return
}

// @snippet_end

const VuetifyBasicInputsPath = "/samples/vuetify-basic-inputs"
