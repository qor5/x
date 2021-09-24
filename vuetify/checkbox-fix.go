package vuetify

import h "github.com/theplant/htmlgo"

func (b *VCheckboxBuilder) FieldName(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("v-field-name", h.JSONString(v))
	return b
}

func (b *VCheckboxBuilder) ErrorMessages(v ...string) (r *VCheckboxBuilder) {
	SetErrorMessages(b.tag, v)
	return b
}
