package vuetify

import h "github.com/theplant/htmlgo"

func (b *VComboboxBuilder) FieldName(v string) (r *VComboboxBuilder) {
	b.tag.Attr("v-field-name", h.JSONString(v))
	return b
}
