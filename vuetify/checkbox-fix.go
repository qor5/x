package vuetify

import h "github.com/theplant/htmlgo"

func VCheckbox(children ...h.HTMLComponent) (r *VCheckboxBuilder) {
	r = &VCheckboxBuilder{
		tag: h.Tag("vw-checkbox").Children(children...),
	}
	//r.FalseValue("false").TrueValue("true")
	return
}

func (b *VCheckboxBuilder) FieldName(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("field-name", v)
	return b
}

func (b *VCheckboxBuilder) LoadPageWithArrayOp(v bool) (r *VCheckboxBuilder) {
	b.tag.Attr("load-page-with-array-op", v)
	return b
}

func (b *VCheckboxBuilder) ErrorMessages(v ...string) (r *VCheckboxBuilder) {
	setErrorMessages(b.tag, v)
	return b
}
