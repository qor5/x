package vuetify

import h "github.com/theplant/htmlgo"

func (b *VTextareaBuilder) FieldName(v string) (r *VTextareaBuilder) {
	b.tag.Attr("v-field-name", h.JSONString(v))
	return b
}

func (b *VTextareaBuilder) ErrorMessages(v ...string) (r *VTextareaBuilder) {
	setErrorMessages(b.tag, v)
	return b
}
