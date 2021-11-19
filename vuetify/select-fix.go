package vuetify

import (
	"github.com/goplaid/web"
)

func (b *VSelectBuilder) ErrorMessages(v ...string) (r *VSelectBuilder) {
	SetErrorMessages(b.tag, v)
	return b
}

func (b *VSelectBuilder) FieldName(v string) (r *VSelectBuilder) {
	b.tag.Attr(web.VFieldName(v)...)
	return b
}
