package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VToolbarItemsBuilder struct {
	tag *h.HTMLTagBuilder
}

func VToolbarItems(children ...h.HTMLComponent) (r *VToolbarItemsBuilder) {
	r = &VToolbarItemsBuilder{
		tag: h.Tag("v-toolbar-items").Children(children...),
	}
	return
}

func (b *VToolbarItemsBuilder) Color(v string) (r *VToolbarItemsBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VToolbarItemsBuilder) Variant(v interface{}) (r *VToolbarItemsBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VToolbarItemsBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VToolbarItemsBuilder) Attr(vs ...interface{}) (r *VToolbarItemsBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VToolbarItemsBuilder) Children(children ...h.HTMLComponent) (r *VToolbarItemsBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VToolbarItemsBuilder) AppendChildren(children ...h.HTMLComponent) (r *VToolbarItemsBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VToolbarItemsBuilder) PrependChildren(children ...h.HTMLComponent) (r *VToolbarItemsBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VToolbarItemsBuilder) Class(names ...string) (r *VToolbarItemsBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VToolbarItemsBuilder) ClassIf(name string, add bool) (r *VToolbarItemsBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VToolbarItemsBuilder) On(name string, value string) (r *VToolbarItemsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VToolbarItemsBuilder) Bind(name string, value string) (r *VToolbarItemsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VToolbarItemsBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
