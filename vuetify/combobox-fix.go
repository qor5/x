package vuetify

import (
	"github.com/goplaid/web"
)

func (b *VComboboxBuilder) FieldName(v string) (r *VComboboxBuilder) {
	b.tag.Attr(web.VFieldName(v)...)
	return b
}
