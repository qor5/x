package vuetify

import h "github.com/theplant/htmlgo"

func VRadioGroup(children ...h.HTMLComponent) (r *VRadioGroupBuilder) {
	r = &VRadioGroupBuilder{
		tag: h.Tag("vw-radio-group").Children(children...),
	}
	return
}

func (b *VRadioGroupBuilder) FieldName(v string) (r *VRadioGroupBuilder) {
	b.tag.Attr("field-name", v)
	return b
}

func (b *VRadioGroupBuilder) ErrorMessages(v ...string) (r *VRadioGroupBuilder) {
	setErrorMessages(b.tag, v)
	return b
}
