package vuetify

import h "github.com/theplant/htmlgo"

func VItemGroup(children ...h.HTMLComponent) (r *VItemGroupBuilder) {
	r = &VItemGroupBuilder{
		tag: h.Tag("vw-item-group").Children(children...),
	}
	return
}

func (b *VItemGroupBuilder) FieldName(v string) (r *VItemGroupBuilder) {
	b.tag.Attr("field-name", v)
	return b
}
