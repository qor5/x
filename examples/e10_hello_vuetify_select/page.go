package e10_hello_vuetify_select

import (
	"github.com/sunfmin/bran/ui"
	vt "github.com/sunfmin/bran/vuetify"
	. "github.com/theplant/htmlgo"
)

type mystate struct {
	MyValues []string
}

func HelloVuetifySelect(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	s := ctx.StateOrInit(&mystate{
		MyValues: []string{
			"react",
			"jquery",
		},
	}).(*mystate)

	result := Ul()
	for _, v := range s.MyValues {
		result.AppendChildren(Li().Text(v))
	}

	pr.Schema = vt.VApp(
		vt.VContent(
			vt.VContainer(
				vt.VAutoComplete([]*vt.TextValue{
					{"Vue", "vue"},
					{"React", "react"},
					{"Angular", "angular"},
					{"Amber", "amber"},
					{"jQuery", "jquery"},
				}).FieldName("MyValues"),
				result,
				vt.VBtn("Update").
					Color("success").
					OnClick(ctx.Hub, "update", update),
			),
		),
	)
	return
}

func update(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	r.Reload = true

	return
}
