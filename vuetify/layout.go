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

func (b *VLayoutBuilder) AlignBaseline(v bool) (r *VLayoutBuilder) {
	b.tag.Attr(":align-baseline", fmt.Sprint(v))
	return b
}

func (b *VLayoutBuilder) AlignCenter(v bool) (r *VLayoutBuilder) {
	b.tag.Attr(":align-center", fmt.Sprint(v))
	return b
}

func (b *VLayoutBuilder) AlignContentCenter(v bool) (r *VLayoutBuilder) {
	b.tag.Attr(":align-content-center", fmt.Sprint(v))
	return b
}

func (b *VLayoutBuilder) AlignContentEnd(v bool) (r *VLayoutBuilder) {
	b.tag.Attr(":align-content-end", fmt.Sprint(v))
	return b
}

func (b *VLayoutBuilder) AlignContentSpaceAround(v bool) (r *VLayoutBuilder) {
	b.tag.Attr(":align-content-space-around", fmt.Sprint(v))
	return b
}

func (b *VLayoutBuilder) AlignContentSpaceBetween(v bool) (r *VLayoutBuilder) {
	b.tag.Attr(":align-content-space-between", fmt.Sprint(v))
	return b
}

func (b *VLayoutBuilder) AlignContentStart(v bool) (r *VLayoutBuilder) {
	b.tag.Attr(":align-content-start", fmt.Sprint(v))
	return b
}

func (b *VLayoutBuilder) AlignEnd(v bool) (r *VLayoutBuilder) {
	b.tag.Attr(":align-end", fmt.Sprint(v))
	return b
}

func (b *VLayoutBuilder) AlignStart(v bool) (r *VLayoutBuilder) {
	b.tag.Attr(":align-start", fmt.Sprint(v))
	return b
}

func (b *VLayoutBuilder) Column(v bool) (r *VLayoutBuilder) {
	b.tag.Attr(":column", fmt.Sprint(v))
	return b
}

func (b *VLayoutBuilder) Dtype(v bool) (r *VLayoutBuilder) {
	b.tag.Attr(":d-{type}", fmt.Sprint(v))
	return b
}

func (b *VLayoutBuilder) FillHeight(v bool) (r *VLayoutBuilder) {
	b.tag.Attr(":fill-height", fmt.Sprint(v))
	return b
}

func (b *VLayoutBuilder) Id(v string) (r *VLayoutBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VLayoutBuilder) JustifyCenter(v bool) (r *VLayoutBuilder) {
	b.tag.Attr(":justify-center", fmt.Sprint(v))
	return b
}

func (b *VLayoutBuilder) JustifyEnd(v bool) (r *VLayoutBuilder) {
	b.tag.Attr(":justify-end", fmt.Sprint(v))
	return b
}

func (b *VLayoutBuilder) JustifySpaceAround(v bool) (r *VLayoutBuilder) {
	b.tag.Attr(":justify-space-around", fmt.Sprint(v))
	return b
}

func (b *VLayoutBuilder) JustifySpaceBetween(v bool) (r *VLayoutBuilder) {
	b.tag.Attr(":justify-space-between", fmt.Sprint(v))
	return b
}

func (b *VLayoutBuilder) JustifyStart(v bool) (r *VLayoutBuilder) {
	b.tag.Attr(":justify-start", fmt.Sprint(v))
	return b
}

func (b *VLayoutBuilder) Reverse(v bool) (r *VLayoutBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VLayoutBuilder) Row(v bool) (r *VLayoutBuilder) {
	b.tag.Attr(":row", fmt.Sprint(v))
	return b
}

func (b *VLayoutBuilder) Tag(v string) (r *VLayoutBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VLayoutBuilder) Wrap(v bool) (r *VLayoutBuilder) {
	b.tag.Attr(":wrap", fmt.Sprint(v))
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
