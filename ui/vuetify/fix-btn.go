package vuetify

import (
	"github.com/qor5/web/v3"
	h "github.com/theplant/htmlgo"
)

func VBtn(text string) (r *VBtnBuilder) {
	r = &VBtnBuilder{
		tag: h.Tag("v-btn").Text(text),
	}
	return
}

func (b *VBtnBuilder) OnClick(eventFuncId string) (r *VBtnBuilder) {
	b.tag.Attr("@click", web.POST().EventFunc(eventFuncId).Go())
	return b
}

func (b *VBtnBuilder) AttrIf(key, value interface{}, add bool) (r *VBtnBuilder) {
	b.tag.AttrIf(key, value, add)
	return b
}
