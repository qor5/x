package vuetify

import (
	"github.com/goplaid/web"
)

func (b *VTextFieldBuilder) FieldName(v string) (r *VTextFieldBuilder) {
	b.tag.Attr(web.VFieldName(v)...)
	return b
}

func (b *VTextFieldBuilder) ErrorMessages(v ...string) (r *VTextFieldBuilder) {
	SetErrorMessages(b.tag, v)
	return b
}
