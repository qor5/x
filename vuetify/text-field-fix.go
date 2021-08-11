package vuetify

import h "github.com/theplant/htmlgo"

func VTextField(children ...h.HTMLComponent) (r *VTextFieldBuilder) {
	r = &VTextFieldBuilder{
		tag: h.Tag("vw-text-field").Children(children...),
	}
	return
}

func (b *VTextFieldBuilder) FieldName(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("field-name", v)
	return b
}

func (b *VTextFieldBuilder) ErrorMessages(v ...string) (r *VTextFieldBuilder) {
	setErrorMessages(b.tag, v)
	return b
}
