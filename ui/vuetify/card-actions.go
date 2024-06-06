package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCardActionsBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCardActions(children ...h.HTMLComponent) (r *VCardActionsBuilder) {
	r = &VCardActionsBuilder{
		tag: h.Tag("v-card-actions").Children(children...),
	}
	return
}

func (b *VCardActionsBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VCardActionsBuilder) Attr(vs ...interface{}) (r *VCardActionsBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VCardActionsBuilder) Children(children ...h.HTMLComponent) (r *VCardActionsBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VCardActionsBuilder) AppendChildren(children ...h.HTMLComponent) (r *VCardActionsBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VCardActionsBuilder) PrependChildren(children ...h.HTMLComponent) (r *VCardActionsBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VCardActionsBuilder) Class(names ...string) (r *VCardActionsBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCardActionsBuilder) ClassIf(name string, add bool) (r *VCardActionsBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCardActionsBuilder) On(name string, value string) (r *VCardActionsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCardActionsBuilder) Bind(name string, value string) (r *VCardActionsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCardActionsBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
