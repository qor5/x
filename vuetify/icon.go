package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VIconBuilder struct {
	tag *h.HTMLTagBuilder
}

func (b *VIconBuilder) Color(v string) (r *VIconBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VIconBuilder) Dark(v bool) (r *VIconBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VIconBuilder) Dense(v bool) (r *VIconBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
	return b
}

func (b *VIconBuilder) Disabled(v bool) (r *VIconBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VIconBuilder) Large(v bool) (r *VIconBuilder) {
	b.tag.Attr(":large", fmt.Sprint(v))
	return b
}

func (b *VIconBuilder) Left(v bool) (r *VIconBuilder) {
	b.tag.Attr(":left", fmt.Sprint(v))
	return b
}

func (b *VIconBuilder) Light(v bool) (r *VIconBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VIconBuilder) Right(v bool) (r *VIconBuilder) {
	b.tag.Attr(":right", fmt.Sprint(v))
	return b
}

func (b *VIconBuilder) Size(v int) (r *VIconBuilder) {
	b.tag.Attr(":size", fmt.Sprint(v))
	return b
}

func (b *VIconBuilder) Small(v bool) (r *VIconBuilder) {
	b.tag.Attr(":small", fmt.Sprint(v))
	return b
}

func (b *VIconBuilder) Tag(v string) (r *VIconBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VIconBuilder) XLarge(v bool) (r *VIconBuilder) {
	b.tag.Attr(":x-large", fmt.Sprint(v))
	return b
}

func (b *VIconBuilder) XSmall(v bool) (r *VIconBuilder) {
	b.tag.Attr(":x-small", fmt.Sprint(v))
	return b
}

func (b *VIconBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VIconBuilder) Attr(vs ...interface{}) (r *VIconBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VIconBuilder) Children(children ...h.HTMLComponent) (r *VIconBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VIconBuilder) AppendChildren(children ...h.HTMLComponent) (r *VIconBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VIconBuilder) PrependChildren(children ...h.HTMLComponent) (r *VIconBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VIconBuilder) Class(names ...string) (r *VIconBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VIconBuilder) ClassIf(name string, add bool) (r *VIconBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VIconBuilder) On(name string, value string) (r *VIconBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VIconBuilder) Bind(name string, value string) (r *VIconBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VIconBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
