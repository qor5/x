package e11_hello_vuetify_text_field

import (
	"github.com/sunfmin/bran/ui"
	vt "github.com/sunfmin/bran/vuetify"
	. "github.com/theplant/htmlgo"
)

type mystate struct {
	MyValue string
}

func HelloVuetifyTextField(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	s := ctx.StateOrInit(&mystate{
		MyValue: "This is my value",
	}).(*mystate)

	pr.Schema = vt.VApp(
		vt.VContent(
			vt.VContainer(
				vt.VTextField().Label("Default"),
				vt.VTextField().Solo(true).Label("Solo"),
				vt.VTextField().Box(true).Label("Mask").Mask("###-####").Hint("Input number like 123-1231"),
				vt.VTextField().Clearable(true).Label("Clearable"),
				vt.VTextField().Error(true).Label("Error"),
				vt.VTextField().Label("Default Value").Value("Hello").Autofocus(true),
				vt.VTextField().Label("Form Value").Autofocus(true).FieldName("MyValue"),
				Pre(s.MyValue),
				vt.VBtn("Update").OnClick(ctx.Hub, "update", update).Color("info"),
			),
		),
	)
	return
}

func update(ctx *ui.EventContext) (r ui.EventResponse, err error) {

	r.Reload = true

	return
}
