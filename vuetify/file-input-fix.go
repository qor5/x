package vuetify

import h "github.com/theplant/htmlgo"

func VFileInput(children ...h.HTMLComponent) (r *VFileInputBuilder) {
	r = &VFileInputBuilder{
		tag: h.Tag("vw-file-input").Children(children...),
	}
	return
}

func (b *VFileInputBuilder) FieldName(v string) (r *VFileInputBuilder) {
	b.tag.Attr("field-name", v)
	return b
}
