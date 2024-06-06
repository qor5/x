package vuetify

import (
	h "github.com/theplant/htmlgo"
)

func VAutocomplete(children ...h.HTMLComponent) (r *VAutocompleteBuilder) {
	r = &VAutocompleteBuilder{
		tag: h.Tag("v-autocomplete").Children(children...),
	}
	return
}

func (b *VAutocompleteBuilder) ErrorMessages(v ...string) (r *VAutocompleteBuilder) {
	SetErrorMessages(b.tag, v)
	return b
}
