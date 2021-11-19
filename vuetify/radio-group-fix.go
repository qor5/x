package vuetify

import (
	"github.com/goplaid/web"
)

func (b *VRadioGroupBuilder) FieldName(v string) (r *VRadioGroupBuilder) {
	b.tag.Attr(web.VFieldName(v)...)
	return b
}

func (b *VRadioGroupBuilder) ErrorMessages(v ...string) (r *VRadioGroupBuilder) {
	SetErrorMessages(b.tag, v)
	return b
}
