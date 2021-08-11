package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSimpleTableBuilder struct {
	tag *h.HTMLTagBuilder
}

func (b *VSimpleTableBuilder) Dark(v bool) (r *VSimpleTableBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VSimpleTableBuilder) Dense(v bool) (r *VSimpleTableBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
	return b
}

func (b *VSimpleTableBuilder) FixedHeader(v bool) (r *VSimpleTableBuilder) {
	b.tag.Attr(":fixed-header", fmt.Sprint(v))
	return b
}

func (b *VSimpleTableBuilder) Height(v int) (r *VSimpleTableBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VSimpleTableBuilder) Light(v bool) (r *VSimpleTableBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VSimpleTableBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSimpleTableBuilder) Attr(vs ...interface{}) (r *VSimpleTableBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSimpleTableBuilder) Children(children ...h.HTMLComponent) (r *VSimpleTableBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSimpleTableBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSimpleTableBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSimpleTableBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSimpleTableBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSimpleTableBuilder) Class(names ...string) (r *VSimpleTableBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSimpleTableBuilder) ClassIf(name string, add bool) (r *VSimpleTableBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSimpleTableBuilder) On(name string, value string) (r *VSimpleTableBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSimpleTableBuilder) Bind(name string, value string) (r *VSimpleTableBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSimpleTableBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
