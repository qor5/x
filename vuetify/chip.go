package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VChipBuilder struct {
	tag *h.HTMLTagBuilder
}

func VChip(children ...h.HTMLComponent) (r *VChipBuilder) {
	r = &VChipBuilder{
		tag: h.Tag("v-chip").Children(children...),
	}
	return
}

func (b *VChipBuilder) Active(v bool) (r *VChipBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) ActiveClass(v string) (r *VChipBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VChipBuilder) Append(v bool) (r *VChipBuilder) {
	b.tag.Attr(":append", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Close(v bool) (r *VChipBuilder) {
	b.tag.Attr(":close", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) CloseIcon(v string) (r *VChipBuilder) {
	b.tag.Attr("close-icon", v)
	return b
}

func (b *VChipBuilder) CloseLabel(v string) (r *VChipBuilder) {
	b.tag.Attr("close-label", v)
	return b
}

func (b *VChipBuilder) Color(v string) (r *VChipBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VChipBuilder) Dark(v bool) (r *VChipBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Disabled(v bool) (r *VChipBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Draggable(v bool) (r *VChipBuilder) {
	b.tag.Attr(":draggable", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Exact(v bool) (r *VChipBuilder) {
	b.tag.Attr(":exact", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) ExactActiveClass(v string) (r *VChipBuilder) {
	b.tag.Attr("exact-active-class", v)
	return b
}

func (b *VChipBuilder) ExactPath(v bool) (r *VChipBuilder) {
	b.tag.Attr(":exact-path", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Filter(v bool) (r *VChipBuilder) {
	b.tag.Attr(":filter", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) FilterIcon(v string) (r *VChipBuilder) {
	b.tag.Attr("filter-icon", v)
	return b
}

func (b *VChipBuilder) Href(v interface{}) (r *VChipBuilder) {
	b.tag.Attr(":href", h.JSONString(v))
	return b
}

func (b *VChipBuilder) InputValue(v interface{}) (r *VChipBuilder) {
	b.tag.Attr(":input-value", h.JSONString(v))
	return b
}

func (b *VChipBuilder) Label(v bool) (r *VChipBuilder) {
	b.tag.Attr(":label", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Large(v bool) (r *VChipBuilder) {
	b.tag.Attr(":large", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Light(v bool) (r *VChipBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Link(v bool) (r *VChipBuilder) {
	b.tag.Attr(":link", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Nuxt(v bool) (r *VChipBuilder) {
	b.tag.Attr(":nuxt", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Outlined(v bool) (r *VChipBuilder) {
	b.tag.Attr(":outlined", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Pill(v bool) (r *VChipBuilder) {
	b.tag.Attr(":pill", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Replace(v bool) (r *VChipBuilder) {
	b.tag.Attr(":replace", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Ripple(v interface{}) (r *VChipBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VChipBuilder) Small(v bool) (r *VChipBuilder) {
	b.tag.Attr(":small", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) Tag(v string) (r *VChipBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VChipBuilder) Target(v string) (r *VChipBuilder) {
	b.tag.Attr("target", v)
	return b
}

func (b *VChipBuilder) TextColor(v string) (r *VChipBuilder) {
	b.tag.Attr("text-color", v)
	return b
}

func (b *VChipBuilder) To(v interface{}) (r *VChipBuilder) {
	b.tag.Attr(":to", h.JSONString(v))
	return b
}

func (b *VChipBuilder) Value(v interface{}) (r *VChipBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VChipBuilder) XLarge(v bool) (r *VChipBuilder) {
	b.tag.Attr(":x-large", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) XSmall(v bool) (r *VChipBuilder) {
	b.tag.Attr(":x-small", fmt.Sprint(v))
	return b
}

func (b *VChipBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VChipBuilder) Attr(vs ...interface{}) (r *VChipBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VChipBuilder) Children(children ...h.HTMLComponent) (r *VChipBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VChipBuilder) AppendChildren(children ...h.HTMLComponent) (r *VChipBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VChipBuilder) PrependChildren(children ...h.HTMLComponent) (r *VChipBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VChipBuilder) Class(names ...string) (r *VChipBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VChipBuilder) ClassIf(name string, add bool) (r *VChipBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VChipBuilder) On(name string, value string) (r *VChipBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VChipBuilder) Bind(name string, value string) (r *VChipBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VChipBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
