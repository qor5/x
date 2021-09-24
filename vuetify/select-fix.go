package vuetify

import (
	h "github.com/theplant/htmlgo"
)

func (b *VSelectBuilder) ErrorMessages(v ...string) (r *VSelectBuilder) {
	SetErrorMessages(b.tag, v)
	return b
}

func (b *VSelectBuilder) FieldName(v string) (r *VSelectBuilder) {
	b.tag.Attr("v-field-name", h.JSONString(v))
	return b
}
