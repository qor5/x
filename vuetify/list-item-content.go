package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VListItemContentBuilder struct {
	tag *h.HTMLTagBuilder
}

func VListItemContent(children ...h.HTMLComponent) (r *VListItemContentBuilder) {
	r = &VListItemContentBuilder{
		tag: h.Tag("v-list-item-content").Children(children...),
	}
	return
}

func (b *VListItemContentBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VListItemContentBuilder) Attr(vs ...interface{}) (r *VListItemContentBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VListItemContentBuilder) Children(children ...h.HTMLComponent) (r *VListItemContentBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VListItemContentBuilder) AppendChildren(children ...h.HTMLComponent) (r *VListItemContentBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VListItemContentBuilder) PrependChildren(children ...h.HTMLComponent) (r *VListItemContentBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VListItemContentBuilder) Class(names ...string) (r *VListItemContentBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VListItemContentBuilder) ClassIf(name string, add bool) (r *VListItemContentBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VListItemContentBuilder) On(name string, value string) (r *VListItemContentBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VListItemContentBuilder) Bind(name string, value string) (r *VListItemContentBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VListItemContentBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
