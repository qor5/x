package vuetify

import (
	"context"

	"github.com/sunfmin/bran/ui"

	h "github.com/theplant/htmlgo"
)

type VBtnBuilder struct {
	tag *h.HTMLTagBuilder
}

func VBtn(text string) (r *VBtnBuilder) {
	r = &VBtnBuilder{
		tag: h.Tag("v-btn").Text(text),
	}
	return
}

func (b *VBtnBuilder) OnClick(hub ui.EventFuncHub, eventFuncId string, ef ui.EventFunc, params ...string) (r *VBtnBuilder) {
	ui.Bind(b.tag).OnClick(hub, eventFuncId, ef, params...)
	return b
}

func (b *VBtnBuilder) Color(v string) (r *VBtnBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VBtnBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
