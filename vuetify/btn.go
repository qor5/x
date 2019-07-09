package vuetify

import (
	"context"
	"fmt"

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

func (b *VBtnBuilder) Absolute(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) ActiveClass(v string) (r *VBtnBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VBtnBuilder) Append(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":append", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Block(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":block", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Bottom(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":bottom", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Color(v string) (r *VBtnBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VBtnBuilder) Dark(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Depressed(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":depressed", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Disabled(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Exact(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":exact", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) ExactActiveClass(v string) (r *VBtnBuilder) {
	b.tag.Attr("exact-active-class", v)
	return b
}

func (b *VBtnBuilder) Fab(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":fab", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Fixed(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":fixed", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Flat(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Href(v string) (r *VBtnBuilder) {
	b.tag.Attr("href", v)
	return b
}

func (b *VBtnBuilder) Icon(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":icon", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Large(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":large", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Left(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":left", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Light(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Loading(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":loading", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Outline(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":outline", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Replace(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":replace", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Right(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":right", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Ripple(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":ripple", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Round(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":round", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Small(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":small", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Tag(v string) (r *VBtnBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VBtnBuilder) Target(v string) (r *VBtnBuilder) {
	b.tag.Attr("target", v)
	return b
}

func (b *VBtnBuilder) To(v string) (r *VBtnBuilder) {
	b.tag.Attr("to", v)
	return b
}

func (b *VBtnBuilder) Top(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":top", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Type(v string) (r *VBtnBuilder) {
	b.tag.Attr("type", v)
	return b
}

func (b *VBtnBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}

func (b *VBtnBuilder) OnClick(hub ui.EventFuncHub, eventFuncId string, ef ui.EventFunc, params ...string) (r *VBtnBuilder) {
	ui.Bind(b.tag).OnClick(hub, eventFuncId, ef, params...)
	return b
}
