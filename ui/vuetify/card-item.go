package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCardItemBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCardItem(children ...h.HTMLComponent) (r *VCardItemBuilder) {
	r = &VCardItemBuilder{
		tag: h.Tag("v-card-item").Children(children...),
	}
	return
}

func (b *VCardItemBuilder) Title(v interface{}) (r *VCardItemBuilder) {
	b.tag.Attr(":title", h.JSONString(v))
	return b
}

func (b *VCardItemBuilder) Subtitle(v interface{}) (r *VCardItemBuilder) {
	b.tag.Attr(":subtitle", h.JSONString(v))
	return b
}

func (b *VCardItemBuilder) AppendAvatar(v string) (r *VCardItemBuilder) {
	b.tag.Attr("append-avatar", v)
	return b
}

func (b *VCardItemBuilder) AppendIcon(v interface{}) (r *VCardItemBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VCardItemBuilder) PrependAvatar(v string) (r *VCardItemBuilder) {
	b.tag.Attr("prepend-avatar", v)
	return b
}

func (b *VCardItemBuilder) PrependIcon(v interface{}) (r *VCardItemBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VCardItemBuilder) Density(v interface{}) (r *VCardItemBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VCardItemBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VCardItemBuilder) Attr(vs ...interface{}) (r *VCardItemBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VCardItemBuilder) Children(children ...h.HTMLComponent) (r *VCardItemBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VCardItemBuilder) AppendChildren(children ...h.HTMLComponent) (r *VCardItemBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VCardItemBuilder) PrependChildren(children ...h.HTMLComponent) (r *VCardItemBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VCardItemBuilder) Class(names ...string) (r *VCardItemBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCardItemBuilder) ClassIf(name string, add bool) (r *VCardItemBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCardItemBuilder) On(name string, value string) (r *VCardItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCardItemBuilder) Bind(name string, value string) (r *VCardItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCardItemBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
