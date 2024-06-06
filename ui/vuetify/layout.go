package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VLayoutBuilder struct {
	tag *h.HTMLTagBuilder
}

func VLayout(children ...h.HTMLComponent) (r *VLayoutBuilder) {
	r = &VLayoutBuilder{
		tag: h.Tag("v-layout").Children(children...),
	}
	return
}

func (b *VLayoutBuilder) Height(v interface{}) (r *VLayoutBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VLayoutBuilder) MaxHeight(v interface{}) (r *VLayoutBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VLayoutBuilder) MaxWidth(v interface{}) (r *VLayoutBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VLayoutBuilder) MinHeight(v interface{}) (r *VLayoutBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VLayoutBuilder) MinWidth(v interface{}) (r *VLayoutBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VLayoutBuilder) Width(v interface{}) (r *VLayoutBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VLayoutBuilder) FullHeight(v bool) (r *VLayoutBuilder) {
	b.tag.Attr(":full-height", fmt.Sprint(v))
	return b
}

func (b *VLayoutBuilder) Overlaps(v interface{}) (r *VLayoutBuilder) {
	b.tag.Attr(":overlaps", h.JSONString(v))
	return b
}

func (b *VLayoutBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VLayoutBuilder) Attr(vs ...interface{}) (r *VLayoutBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VLayoutBuilder) Children(children ...h.HTMLComponent) (r *VLayoutBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VLayoutBuilder) AppendChildren(children ...h.HTMLComponent) (r *VLayoutBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VLayoutBuilder) PrependChildren(children ...h.HTMLComponent) (r *VLayoutBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VLayoutBuilder) Class(names ...string) (r *VLayoutBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VLayoutBuilder) ClassIf(name string, add bool) (r *VLayoutBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VLayoutBuilder) On(name string, value string) (r *VLayoutBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VLayoutBuilder) Bind(name string, value string) (r *VLayoutBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VLayoutBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
