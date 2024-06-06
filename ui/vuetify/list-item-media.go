package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VListItemMediaBuilder struct {
	tag *h.HTMLTagBuilder
}

func VListItemMedia(children ...h.HTMLComponent) (r *VListItemMediaBuilder) {
	r = &VListItemMediaBuilder{
		tag: h.Tag("v-list-item-media").Children(children...),
	}
	return
}

func (b *VListItemMediaBuilder) Start(v bool) (r *VListItemMediaBuilder) {
	b.tag.Attr(":start", fmt.Sprint(v))
	return b
}

func (b *VListItemMediaBuilder) End(v bool) (r *VListItemMediaBuilder) {
	b.tag.Attr(":end", fmt.Sprint(v))
	return b
}

func (b *VListItemMediaBuilder) Tag(v string) (r *VListItemMediaBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VListItemMediaBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VListItemMediaBuilder) Attr(vs ...interface{}) (r *VListItemMediaBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VListItemMediaBuilder) Children(children ...h.HTMLComponent) (r *VListItemMediaBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VListItemMediaBuilder) AppendChildren(children ...h.HTMLComponent) (r *VListItemMediaBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VListItemMediaBuilder) PrependChildren(children ...h.HTMLComponent) (r *VListItemMediaBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VListItemMediaBuilder) Class(names ...string) (r *VListItemMediaBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VListItemMediaBuilder) ClassIf(name string, add bool) (r *VListItemMediaBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VListItemMediaBuilder) On(name string, value string) (r *VListItemMediaBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VListItemMediaBuilder) Bind(name string, value string) (r *VListItemMediaBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VListItemMediaBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
