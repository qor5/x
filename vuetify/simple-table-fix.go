package vuetify

import h "github.com/theplant/htmlgo"

func VSimpleTable(children ...h.HTMLComponent) (r *VSimpleTableBuilder) {
	r = &VSimpleTableBuilder{
		tag: h.Tag("table").Children(children...).Attr("is", "v-simple-table"),
	}
	return
}
