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
	b.tag.Attr("@click", web.Plaid().EventFunc(eventFuncId, params...).Go())
	return b
}
