package vuetify

import h "github.com/theplant/htmlgo"

func VIcon(name string) (r *VIconBuilder) {
	r = &VIconBuilder{
		tag: h.Tag("v-icon").Text(name),
	}
	return
}
