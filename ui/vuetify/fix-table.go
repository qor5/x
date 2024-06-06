package vuetify

import h "github.com/theplant/htmlgo"

func VTable(children ...h.HTMLComponent) (r *VTableBuilder) {
	r = &VTableBuilder{
		tag: h.Tag("v-table").Children(
			h.Template(
				children...,
			).Attr("#default", true),
		),
	}
	return
}
