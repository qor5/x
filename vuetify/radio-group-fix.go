package vuetify

import h "github.com/theplant/htmlgo"

func (b *VRadioGroupBuilder) FieldName(v string) (r *VRadioGroupBuilder) {
	b.tag.Attr("v-field-name", h.JSONString(v))
	return b
}

func (b *VRadioGroupBuilder) ErrorMessages(v ...string) (r *VRadioGroupBuilder) {
	setErrorMessages(b.tag, v)
	return b
}
