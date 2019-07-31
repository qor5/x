package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VVirtualTableBuilder struct {
	tag *h.HTMLTagBuilder
}

func VVirtualTable() (r *VVirtualTableBuilder) {
	r = &VVirtualTableBuilder{
		tag: h.Tag("v-virtual-table"),
	}
	return
}

func (b *VVirtualTableBuilder) ChunkSize(v int) (r *VVirtualTableBuilder) {
	b.tag.Attr(":chunk-size", fmt.Sprint(v))
	return b
}

func (b *VVirtualTableBuilder) Dark(v bool) (r *VVirtualTableBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VVirtualTableBuilder) Dense(v bool) (r *VVirtualTableBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
	return b
}

func (b *VVirtualTableBuilder) FixedHeader(v bool) (r *VVirtualTableBuilder) {
	b.tag.Attr(":fixed-header", fmt.Sprint(v))
	return b
}

func (b *VVirtualTableBuilder) HeaderHeight(v int) (r *VVirtualTableBuilder) {
	b.tag.Attr(":header-height", fmt.Sprint(v))
	return b
}

func (b *VVirtualTableBuilder) Height(v int) (r *VVirtualTableBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VVirtualTableBuilder) Items(v []string) (r *VVirtualTableBuilder) {
	b.tag.Attr(":items", v)
	return b
}

func (b *VVirtualTableBuilder) Light(v bool) (r *VVirtualTableBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VVirtualTableBuilder) RowHeight(v int) (r *VVirtualTableBuilder) {
	b.tag.Attr(":row-height", fmt.Sprint(v))
	return b
}

func (b *VVirtualTableBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VVirtualTableBuilder) Attr(vs ...interface{}) (r *VVirtualTableBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VVirtualTableBuilder) Children(children ...h.HTMLComponent) (r *VVirtualTableBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VVirtualTableBuilder) AppendChildren(children ...h.HTMLComponent) (r *VVirtualTableBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VVirtualTableBuilder) PrependChildren(children ...h.HTMLComponent) (r *VVirtualTableBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VVirtualTableBuilder) Class(names ...string) (r *VVirtualTableBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VVirtualTableBuilder) ClassIf(name string, add bool) (r *VVirtualTableBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VVirtualTableBuilder) On(name string, value string) (r *VVirtualTableBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VVirtualTableBuilder) Bind(name string, value string) (r *VVirtualTableBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VVirtualTableBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
