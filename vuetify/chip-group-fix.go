package vuetify

import h "github.com/theplant/htmlgo"

func VChipGroup(children ...h.HTMLComponent) (r *VChipGroupBuilder) {
	r = &VChipGroupBuilder{
		tag: h.Tag("vw-chip-group").Children(children...),
	}
	return
}

func (b *VChipGroupBuilder) FieldName(v string) (r *VChipGroupBuilder) {
	b.tag.Attr("field-name", v)
	return b
}
