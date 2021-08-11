package vuetify

import h "github.com/theplant/htmlgo"

func VTextarea(children ...h.HTMLComponent) (r *VTextareaBuilder) {
	r = &VTextareaBuilder{
		tag: h.Tag("vw-textarea").Children(children...),
	}
	return
}

func (b *VTextareaBuilder) FieldName(v string) (r *VTextareaBuilder) {
	b.tag.Attr("field-name", v)
	return b
}

func (b *VTextareaBuilder) ErrorMessages(v ...string) (r *VTextareaBuilder) {
	setErrorMessages(b.tag, v)
	return b
}
