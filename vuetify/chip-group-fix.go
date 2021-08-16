package vuetify

import h "github.com/theplant/htmlgo"

func (b *VChipGroupBuilder) FieldName(v string) (r *VChipGroupBuilder) {
	b.tag.Attr("v-field-name", h.JSONString(v))
	return b
}
