package vuetify

import h "github.com/theplant/htmlgo"

func (b *VItemGroupBuilder) FieldName(v string) (r *VItemGroupBuilder) {
	b.tag.Attr("v-field-name", h.JSONString(v))
	return b
}
