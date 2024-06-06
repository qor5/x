package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VHoverBuilder struct {
	tag *h.HTMLTagBuilder
}

func VHover(children ...h.HTMLComponent) (r *VHoverBuilder) {
	r = &VHoverBuilder{
		tag: h.Tag("v-hover").Children(children...),
	}
	return
}

func (b *VHoverBuilder) Disabled(v bool) (r *VHoverBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VHoverBuilder) ModelValue(v bool) (r *VHoverBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VHoverBuilder) CloseDelay(v interface{}) (r *VHoverBuilder) {
	b.tag.Attr(":close-delay", h.JSONString(v))
	return b
}

func (b *VHoverBuilder) OpenDelay(v interface{}) (r *VHoverBuilder) {
	b.tag.Attr(":open-delay", h.JSONString(v))
	return b
}

func (b *VHoverBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VHoverBuilder) Attr(vs ...interface{}) (r *VHoverBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VHoverBuilder) Children(children ...h.HTMLComponent) (r *VHoverBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VHoverBuilder) AppendChildren(children ...h.HTMLComponent) (r *VHoverBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VHoverBuilder) PrependChildren(children ...h.HTMLComponent) (r *VHoverBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VHoverBuilder) Class(names ...string) (r *VHoverBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VHoverBuilder) ClassIf(name string, add bool) (r *VHoverBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VHoverBuilder) On(name string, value string) (r *VHoverBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VHoverBuilder) Bind(name string, value string) (r *VHoverBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VHoverBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
