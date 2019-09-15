package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VListItemIconBuilder struct {
	tag *h.HTMLTagBuilder
}

func VListItemIcon(children ...h.HTMLComponent) (r *VListItemIconBuilder) {
	r = &VListItemIconBuilder{
		tag: h.Tag("v-list-item-icon").Children(children...),
	}
	return
}

func (b *VListItemIconBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VListItemIconBuilder) Attr(vs ...interface{}) (r *VListItemIconBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VListItemIconBuilder) Children(children ...h.HTMLComponent) (r *VListItemIconBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VListItemIconBuilder) AppendChildren(children ...h.HTMLComponent) (r *VListItemIconBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VListItemIconBuilder) PrependChildren(children ...h.HTMLComponent) (r *VListItemIconBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VListItemIconBuilder) Class(names ...string) (r *VListItemIconBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VListItemIconBuilder) ClassIf(name string, add bool) (r *VListItemIconBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VListItemIconBuilder) On(name string, value string) (r *VListItemIconBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VListItemIconBuilder) Bind(name string, value string) (r *VListItemIconBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VListItemIconBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
