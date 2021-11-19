package vuetify

import (
	"github.com/goplaid/web"
)

func (b *VAutocompleteBuilder) ErrorMessages(v ...string) (r *VAutocompleteBuilder) {
	SetErrorMessages(b.tag, v)
	return b
}

func (b *VAutocompleteBuilder) FieldName(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr(web.VFieldName(v)...)
	return b
}
