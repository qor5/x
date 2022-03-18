package vuetify

import (
	"github.com/goplaid/web"
)

func (b *VFileInputBuilder) FieldName(v string) (r *VFileInputBuilder) {
	b.tag.Attr(web.VFieldName(v)...)
	return b
}

func (b *VFileInputBuilder) ErrorMessages(v ...string) (r *VFileInputBuilder) {
	SetErrorMessages(b.tag, v)
	return b
}
