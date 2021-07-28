package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VWindowBuilder struct {
	tag *h.HTMLTagBuilder
}

func VWindow(children ...h.HTMLComponent) (r *VWindowBuilder) {
	r = &VWindowBuilder{
		tag: h.Tag("v-window").Children(children...),
	}
	return
}

func (b *VWindowBuilder) ActiveClass(v string) (r *VWindowBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VWindowBuilder) Continuous(v bool) (r *VWindowBuilder) {
	b.tag.Attr(":continuous", fmt.Sprint(v))
	return b
}

func (b *VWindowBuilder) Dark(v bool) (r *VWindowBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VWindowBuilder) Light(v bool) (r *VWindowBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VWindowBuilder) NextIcon(v bool) (r *VWindowBuilder) {
	b.tag.Attr(":next-icon", fmt.Sprint(v))
	return b
}

func (b *VWindowBuilder) PrevIcon(v bool) (r *VWindowBuilder) {
	b.tag.Attr(":prev-icon", fmt.Sprint(v))
	return b
}

func (b *VWindowBuilder) Reverse(v bool) (r *VWindowBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VWindowBuilder) ShowArrows(v bool) (r *VWindowBuilder) {
	b.tag.Attr(":show-arrows", fmt.Sprint(v))
	return b
}

func (b *VWindowBuilder) ShowArrowsOnHover(v bool) (r *VWindowBuilder) {
	b.tag.Attr(":show-arrows-on-hover", fmt.Sprint(v))
	return b
}

func (b *VWindowBuilder) Tag(v string) (r *VWindowBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VWindowBuilder) Touch(v interface{}) (r *VWindowBuilder) {
	b.tag.Attr(":touch", h.JSONString(v))
	return b
}

func (b *VWindowBuilder) Touchless(v bool) (r *VWindowBuilder) {
	b.tag.Attr(":touchless", fmt.Sprint(v))
	return b
}

func (b *VWindowBuilder) Value(v interface{}) (r *VWindowBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VWindowBuilder) Vertical(v bool) (r *VWindowBuilder) {
	b.tag.Attr(":vertical", fmt.Sprint(v))
	return b
}

func (b *VWindowBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VWindowBuilder) Attr(vs ...interface{}) (r *VWindowBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VWindowBuilder) Children(children ...h.HTMLComponent) (r *VWindowBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VWindowBuilder) AppendChildren(children ...h.HTMLComponent) (r *VWindowBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VWindowBuilder) PrependChildren(children ...h.HTMLComponent) (r *VWindowBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VWindowBuilder) Class(names ...string) (r *VWindowBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VWindowBuilder) ClassIf(name string, add bool) (r *VWindowBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VWindowBuilder) On(name string, value string) (r *VWindowBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VWindowBuilder) Bind(name string, value string) (r *VWindowBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VWindowBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
