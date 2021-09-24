package vuetify

import h "github.com/theplant/htmlgo"

func (b *VSliderBuilder) ErrorMessages(v ...string) (r *VSliderBuilder) {
	SetErrorMessages(b.tag, v)
	return b
}

func (b *VSliderBuilder) FieldName(v string) (r *VSliderBuilder) {
	b.tag.Attr("v-field-name", h.JSONString(v))
	return b
}
