package vuetify

import (
	"github.com/goplaid/web"

	h "github.com/theplant/htmlgo"
)

func VAutocomplete(children ...h.HTMLComponent) (r *VAutocompleteBuilder) {
	r = &VAutocompleteBuilder{
		tag: h.Tag("v-autocomplete").Children(children...),
	}
	r.Attach("")
	return
}

func (b *VAutocompleteBuilder) ErrorMessages(v ...string) (r *VAutocompleteBuilder) {
	SetErrorMessages(b.tag, v)
	return b
}

func (b *VAutocompleteBuilder) FieldName(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr(web.VFieldName(v)...)
	return b
}
