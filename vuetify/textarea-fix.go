package vuetify

import (
	"github.com/goplaid/web"
)

func (b *VTextareaBuilder) FieldName(v string) (r *VTextareaBuilder) {
	b.tag.Attr(web.VFieldName(v)...)
	return b
}

func (b *VTextareaBuilder) ErrorMessages(v ...string) (r *VTextareaBuilder) {
	SetErrorMessages(b.tag, v)
	return b
}
