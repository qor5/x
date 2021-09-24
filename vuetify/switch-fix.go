package vuetify

import h "github.com/theplant/htmlgo"

func (b *VSwitchBuilder) ErrorMessages(v ...string) (r *VSwitchBuilder) {
	SetErrorMessages(b.tag, v)
	return b
}

func (b *VSwitchBuilder) FieldName(v string) (r *VSwitchBuilder) {
	b.tag.Attr("v-field-name", h.JSONString(v))
	return b
}
