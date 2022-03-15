package vuetify

import (
	"github.com/goplaid/web"

	h "github.com/theplant/htmlgo"
)

func VSelect(children ...h.HTMLComponent) (r *VSelectBuilder) {
	r = &VSelectBuilder{
		tag: h.Tag("v-select").Children(children...),
	}
	r.Attach("")
	return
}

func (b *VSelectBuilder) ErrorMessages(v ...string) (r *VSelectBuilder) {
	SetErrorMessages(b.tag, v)
	return b
}

func (b *VSelectBuilder) FieldName(v string) (r *VSelectBuilder) {
	b.tag.Attr(web.VFieldName(v)...)
	return b
}
