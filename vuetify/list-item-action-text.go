package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VListItemActionTextBuilder struct {
	tag *h.HTMLTagBuilder
}

func VListItemActionText(children ...h.HTMLComponent) (r *VListItemActionTextBuilder) {
	r = &VListItemActionTextBuilder{
		tag: h.Tag("v-list-item-action-text").Children(children...),
	}
	return
}

func (b *VListItemActionTextBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VListItemActionTextBuilder) Attr(vs ...interface{}) (r *VListItemActionTextBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VListItemActionTextBuilder) Children(children ...h.HTMLComponent) (r *VListItemActionTextBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VListItemActionTextBuilder) AppendChildren(children ...h.HTMLComponent) (r *VListItemActionTextBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VListItemActionTextBuilder) PrependChildren(children ...h.HTMLComponent) (r *VListItemActionTextBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VListItemActionTextBuilder) Class(names ...string) (r *VListItemActionTextBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VListItemActionTextBuilder) ClassIf(name string, add bool) (r *VListItemActionTextBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VListItemActionTextBuilder) On(name string, value string) (r *VListItemActionTextBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VListItemActionTextBuilder) Bind(name string, value string) (r *VListItemActionTextBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VListItemActionTextBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
