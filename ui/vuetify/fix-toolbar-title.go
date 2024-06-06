package vuetify

import h "github.com/theplant/htmlgo"

func VToolbarTitle(text string) (r *VToolbarTitleBuilder) {
	r = &VToolbarTitleBuilder{
		tag: h.Tag("v-toolbar-title").Text(text),
	}
	return
}
