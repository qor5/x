package vuetify

import h "github.com/theplant/htmlgo"

func (b *VFileInputBuilder) FieldName(v string) (r *VFileInputBuilder) {
	b.tag.Attr("v-field-name", h.JSONString(v))
	return b
}
