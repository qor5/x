package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VListItemIconBuilder struct {
	tag *h.HTMLTagBuilder
}

func VListItemIcon() (r *VListItemIconBuilder) {
	r = &VListItemIconBuilder{
		tag: h.Tag("v-list-item-icon"),
	}
	return
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
