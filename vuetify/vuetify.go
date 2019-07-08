package vuetify

import (
	h "github.com/theplant/htmlgo"
)

func VApp(children ...h.HTMLComponent) (r h.HTMLComponent) {
	r = h.Tag("v-app").Children(children...)
	return
}

func VContent(children ...h.HTMLComponent) (r h.HTMLComponent) {
	r = h.Tag("v-content").Children(children...)
	return
}
