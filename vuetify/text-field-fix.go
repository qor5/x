package vuetify

import h "github.com/theplant/htmlgo"

func (b *VTextFieldBuilder) FieldName(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("v-field-name", h.JSONString(v))
	return b
}

func (b *VTextFieldBuilder) ErrorMessages(v ...string) (r *VTextFieldBuilder) {
	setErrorMessages(b.tag, v)
	return b
}
