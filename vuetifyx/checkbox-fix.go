package vuetifyx

import (
	"github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
)

func VXCheckbox(children ...h.HTMLComponent) (r *VXCheckboxBuilder) {
	r = &VXCheckboxBuilder{
		tag: h.Tag("vx-checkbox").Children(children...),
	}
	//r.FalseValue("false").TrueValue("true")
	return
}

func (b *VXCheckboxBuilder) FieldName(v string) (r *VXCheckboxBuilder) {
	b.tag.Attr("field-name", v)
	return b
}

func (b *VXCheckboxBuilder) LoadPageWithArrayOp(v bool) (r *VXCheckboxBuilder) {
	b.tag.Attr("load-page-with-array-op", v)
	return b
}

func (b *VXCheckboxBuilder) ErrorMessages(v ...string) (r *VXCheckboxBuilder) {
	vuetify.SetErrorMessages(b.tag, v)
	return b
}
