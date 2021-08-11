package vuetify

import (
	"github.com/goplaid/web"
	h "github.com/theplant/htmlgo"
)

func VBtn(text string) (r *VBtnBuilder) {
	r = &VBtnBuilder{
		tag: h.Tag("v-btn").Text(text),
	}
	return
}

func (b *VBtnBuilder) OnClick(eventFuncId string, params ...string) (r *VBtnBuilder) {
	web.Bind(b.tag).OnClick(eventFuncId, params...).Update()
	return b
}
