package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VPaginationBuilder struct {
	tag *h.HTMLTagBuilder
}

func VPagination(children ...h.HTMLComponent) (r *VPaginationBuilder) {
	r = &VPaginationBuilder{
		tag: h.Tag("v-pagination").Children(children...),
	}
	return
}

func (b *VPaginationBuilder) Circle(v bool) (r *VPaginationBuilder) {
	b.tag.Attr(":circle", fmt.Sprint(v))
	return b
}

func (b *VPaginationBuilder) Color(v string) (r *VPaginationBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VPaginationBuilder) CurrentPageAriaLabel(v string) (r *VPaginationBuilder) {
	b.tag.Attr("current-page-aria-label", v)
	return b
}

func (b *VPaginationBuilder) Dark(v bool) (r *VPaginationBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VPaginationBuilder) Disabled(v bool) (r *VPaginationBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VPaginationBuilder) Length(v int) (r *VPaginationBuilder) {
	b.tag.Attr(":length", fmt.Sprint(v))
	return b
}

func (b *VPaginationBuilder) Light(v bool) (r *VPaginationBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VPaginationBuilder) NextAriaLabel(v string) (r *VPaginationBuilder) {
	b.tag.Attr("next-aria-label", v)
	return b
}

func (b *VPaginationBuilder) NextIcon(v string) (r *VPaginationBuilder) {
	b.tag.Attr("next-icon", v)
	return b
}

func (b *VPaginationBuilder) PageAriaLabel(v string) (r *VPaginationBuilder) {
	b.tag.Attr("page-aria-label", v)
	return b
}

func (b *VPaginationBuilder) PrevIcon(v string) (r *VPaginationBuilder) {
	b.tag.Attr("prev-icon", v)
	return b
}

func (b *VPaginationBuilder) PreviousAriaLabel(v string) (r *VPaginationBuilder) {
	b.tag.Attr("previous-aria-label", v)
	return b
}

func (b *VPaginationBuilder) TotalVisible(v int64) (r *VPaginationBuilder) {
	b.tag.Attr(":total-visible", fmt.Sprint(v))
	return b
}

func (b *VPaginationBuilder) Value(v int) (r *VPaginationBuilder) {
	b.tag.Attr(":value", fmt.Sprint(v))
	return b
}

func (b *VPaginationBuilder) WrapperAriaLabel(v string) (r *VPaginationBuilder) {
	b.tag.Attr("wrapper-aria-label", v)
	return b
}

func (b *VPaginationBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VPaginationBuilder) Attr(vs ...interface{}) (r *VPaginationBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VPaginationBuilder) Children(children ...h.HTMLComponent) (r *VPaginationBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VPaginationBuilder) AppendChildren(children ...h.HTMLComponent) (r *VPaginationBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VPaginationBuilder) PrependChildren(children ...h.HTMLComponent) (r *VPaginationBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VPaginationBuilder) Class(names ...string) (r *VPaginationBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VPaginationBuilder) ClassIf(name string, add bool) (r *VPaginationBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VPaginationBuilder) On(name string, value string) (r *VPaginationBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VPaginationBuilder) Bind(name string, value string) (r *VPaginationBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VPaginationBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
