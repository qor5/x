package vuetify

import h "github.com/theplant/htmlgo"

func VSwitch(children ...h.HTMLComponent) (r *VSwitchBuilder) {
	r = &VSwitchBuilder{
		tag: h.Tag("vw-switch").Children(children...),
	}
	return
}

func (b *VSwitchBuilder) ErrorMessages(v ...string) (r *VSwitchBuilder) {
	setErrorMessages(b.tag, v)
	return b
}

func (b *VSwitchBuilder) FieldName(v string) (r *VSwitchBuilder) {
	b.tag.Attr("field-name", v)
	return b
}
