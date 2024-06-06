package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VColBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCol(children ...h.HTMLComponent) (r *VColBuilder) {
	r = &VColBuilder{
		tag: h.Tag("v-col").Children(children...),
	}
	return
}

func (b *VColBuilder) Cols(v interface{}) (r *VColBuilder) {
	b.tag.Attr(":cols", h.JSONString(v))
	return b
}

func (b *VColBuilder) Sm(v interface{}) (r *VColBuilder) {
	b.tag.Attr(":sm", h.JSONString(v))
	return b
}

func (b *VColBuilder) Md(v interface{}) (r *VColBuilder) {
	b.tag.Attr(":md", h.JSONString(v))
	return b
}

func (b *VColBuilder) Lg(v interface{}) (r *VColBuilder) {
	b.tag.Attr(":lg", h.JSONString(v))
	return b
}

func (b *VColBuilder) Xl(v interface{}) (r *VColBuilder) {
	b.tag.Attr(":xl", h.JSONString(v))
	return b
}

func (b *VColBuilder) Xxl(v interface{}) (r *VColBuilder) {
	b.tag.Attr(":xxl", h.JSONString(v))
	return b
}

func (b *VColBuilder) Offset(v interface{}) (r *VColBuilder) {
	b.tag.Attr(":offset", h.JSONString(v))
	return b
}

func (b *VColBuilder) OffsetSm(v interface{}) (r *VColBuilder) {
	b.tag.Attr(":offset-sm", h.JSONString(v))
	return b
}

func (b *VColBuilder) OffsetMd(v interface{}) (r *VColBuilder) {
	b.tag.Attr(":offset-md", h.JSONString(v))
	return b
}

func (b *VColBuilder) OffsetLg(v interface{}) (r *VColBuilder) {
	b.tag.Attr(":offset-lg", h.JSONString(v))
	return b
}

func (b *VColBuilder) OffsetXl(v interface{}) (r *VColBuilder) {
	b.tag.Attr(":offset-xl", h.JSONString(v))
	return b
}

func (b *VColBuilder) OffsetXxl(v interface{}) (r *VColBuilder) {
	b.tag.Attr(":offset-xxl", h.JSONString(v))
	return b
}

func (b *VColBuilder) Order(v interface{}) (r *VColBuilder) {
	b.tag.Attr(":order", h.JSONString(v))
	return b
}

func (b *VColBuilder) OrderSm(v interface{}) (r *VColBuilder) {
	b.tag.Attr(":order-sm", h.JSONString(v))
	return b
}

func (b *VColBuilder) OrderMd(v interface{}) (r *VColBuilder) {
	b.tag.Attr(":order-md", h.JSONString(v))
	return b
}

func (b *VColBuilder) OrderLg(v interface{}) (r *VColBuilder) {
	b.tag.Attr(":order-lg", h.JSONString(v))
	return b
}

func (b *VColBuilder) OrderXl(v interface{}) (r *VColBuilder) {
	b.tag.Attr(":order-xl", h.JSONString(v))
	return b
}

func (b *VColBuilder) OrderXxl(v interface{}) (r *VColBuilder) {
	b.tag.Attr(":order-xxl", h.JSONString(v))
	return b
}

func (b *VColBuilder) AlignSelf(v interface{}) (r *VColBuilder) {
	b.tag.Attr(":align-self", h.JSONString(v))
	return b
}

func (b *VColBuilder) Tag(v string) (r *VColBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VColBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VColBuilder) Attr(vs ...interface{}) (r *VColBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VColBuilder) Children(children ...h.HTMLComponent) (r *VColBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VColBuilder) AppendChildren(children ...h.HTMLComponent) (r *VColBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VColBuilder) PrependChildren(children ...h.HTMLComponent) (r *VColBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VColBuilder) Class(names ...string) (r *VColBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VColBuilder) ClassIf(name string, add bool) (r *VColBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VColBuilder) On(name string, value string) (r *VColBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VColBuilder) Bind(name string, value string) (r *VColBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VColBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
