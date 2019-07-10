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
}

func HelloVuetifyTextField(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	s := ctx.StateOrInit(&mystate{
		MyValue:       "This is my value",
		TextareaValue: "123",
		Gender:        "M",
		Agreed:        true,
	}).(*mystate)

	pr.Schema = vt.VApp(
		vt.VContent(
			vt.VContainer(
				vt.VTextField().Label("Default"),
				vt.VTextField().Solo(true).Label("Solo"),
				vt.VTextField().Box(true).Label("Mask").Mask("###-####").Hint("Input number like 123-1231"),
				vt.VTextField().Clearable(true).Label("Clearable"),
				vt.VTextField().Error(true).Label("Error"),
				vt.VTextField().Label("Default Value").Value("Hello"),
				vt.VTextField().Label("Form Value").FieldName("MyValue"),
				Pre(s.MyValue),
				vt.VTextarea().FieldName("TextareaValue").Solo(true),
				Pre(s.TextareaValue),
				vt.VRadioGroup(
					vt.VRadio().Value("F").Label("Female"),
					vt.VRadio().Value("M").Label("Male"),
				).FieldName("Gender"),
				Pre(s.Gender),
				vt.VCheckbox().FieldName("Agreed").Label("Agree"),
				Pre(fmt.Sprint(s.Agreed)),

				vt.VBtn("Update").OnClick(ctx.Hub, "update", update).Color("info").Round(true),
			),
		),
	).Id("mainapp")
	return
}

func update(ctx *ui.EventContext) (r ui.EventResponse, err error) {

	r.Reload = true

	return
}
