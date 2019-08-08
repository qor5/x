package e11_hello_vuetify_text_field

import (
	"fmt"

	"github.com/sunfmin/bran/ui"
	vt "github.com/sunfmin/bran/vuetify"
	. "github.com/theplant/htmlgo"
)

type mystate struct {
	MyValue       string
	TextareaValue string
	Gender        string
	Agreed        bool
	Feature1      bool
	Slider1       int
}

var s = &mystate{
	MyValue:       "This is my value",
	TextareaValue: "123",
	Gender:        "M",
	Agreed:        true,
	Feature1:      true,
	Slider1:       10,
}

func HelloVuetifyTextField(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("update", update)

	pr.Schema = vt.VApp(
		vt.VContent(
			vt.VContainer(
				vt.VTextField().Label("Default"),
				vt.VTextField().Solo(true).Label("Solo"),
				vt.VTextField().Box(true).Label("Mask").Mask("###-####").Hint("Input number like 123-1231"),
				vt.VTextField().Clearable(true).Label("Clearable"),
				vt.VTextField().Error(true).Label("Error"),
				vt.VTextField().Label("Default ValueIs").Value("Hello"),
				vt.VTextField().Label("Form ValueIs").FieldName("MyValue").Value(s.MyValue),
				Pre(s.MyValue),
				vt.VTextarea().FieldName("TextareaValue").Solo(true).Value(s.TextareaValue),
				Pre(s.TextareaValue),
				vt.VRadioGroup(
					vt.VRadio().Value("F").Label("Female"),
					vt.VRadio().Value("M").Label("Male"),
				).FieldName("Gender").Value(s.Gender),
				Pre(s.Gender),
				vt.VCheckbox().FieldName("Agreed").Label("Agree").InputValue(s.Agreed),
				Pre(fmt.Sprint(s.Agreed)),
				vt.VSwitch().FieldName("Feature1").InputValue(s.Feature1),
				Pre(fmt.Sprint(s.Feature1)),

				vt.VSlider().FieldName("Slider1").Value(s.Slider1),
				Pre(fmt.Sprint(s.Slider1)),

				vt.VSlider().Step(10).Ticks(true).ThumbLabel("always").FieldName("Slider1").Value(s.Slider1),

				vt.VBtn("Update").OnClick("update"),
			),
		),
	).Id("mainapp")
	return
}

func update(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	s = &mystate{}
	ctx.UnmarshalForm(s)
	r.Reload = true

	return
}
