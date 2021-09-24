package vuetify

import h "github.com/theplant/htmlgo"

func (b *VAutocompleteBuilder) ErrorMessages(v ...string) (r *VAutocompleteBuilder) {
	SetErrorMessages(b.tag, v)
	return b
}

func (b *VAutocompleteBuilder) FieldName(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("v-field-name", h.JSONString(v))
	return b
}
