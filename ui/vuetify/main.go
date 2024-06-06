package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VMainBuilder struct {
	tag *h.HTMLTagBuilder
}

func VMain(children ...h.HTMLComponent) (r *VMainBuilder) {
	r = &VMainBuilder{
		tag: h.Tag("v-main").Children(children...),
	}
	return
}

func (b *VMainBuilder) Scrollable(v bool) (r *VMainBuilder) {
	b.tag.Attr(":scrollable", fmt.Sprint(v))
	return b
}

func (b *VMainBuilder) Height(v interface{}) (r *VMainBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VMainBuilder) MaxHeight(v interface{}) (r *VMainBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VMainBuilder) MaxWidth(v interface{}) (r *VMainBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VMainBuilder) MinHeight(v interface{}) (r *VMainBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VMainBuilder) MinWidth(v interface{}) (r *VMainBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VMainBuilder) Width(v interface{}) (r *VMainBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VMainBuilder) Tag(v string) (r *VMainBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VMainBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VMainBuilder) Attr(vs ...interface{}) (r *VMainBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VMainBuilder) Children(children ...h.HTMLComponent) (r *VMainBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VMainBuilder) AppendChildren(children ...h.HTMLComponent) (r *VMainBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VMainBuilder) PrependChildren(children ...h.HTMLComponent) (r *VMainBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VMainBuilder) Class(names ...string) (r *VMainBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VMainBuilder) ClassIf(name string, add bool) (r *VMainBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VMainBuilder) On(name string, value string) (r *VMainBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VMainBuilder) Bind(name string, value string) (r *VMainBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VMainBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
