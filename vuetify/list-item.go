package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VListItemBuilder struct {
	tag *h.HTMLTagBuilder
}

func VListItem(children ...h.HTMLComponent) (r *VListItemBuilder) {
	r = &VListItemBuilder{
		tag: h.Tag("v-list-item").Children(children...),
	}
	return
}

func (b *VListItemBuilder) ActiveClass(v string) (r *VListItemBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VListItemBuilder) Append(v bool) (r *VListItemBuilder) {
	b.tag.Attr(":append", fmt.Sprint(v))
	return b
}

func (b *VListItemBuilder) Color(v string) (r *VListItemBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VListItemBuilder) Dark(v bool) (r *VListItemBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VListItemBuilder) Dense(v bool) (r *VListItemBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
	return b
}

func (b *VListItemBuilder) Disabled(v bool) (r *VListItemBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VListItemBuilder) Exact(v bool) (r *VListItemBuilder) {
	b.tag.Attr(":exact", fmt.Sprint(v))
	return b
}

func (b *VListItemBuilder) ExactActiveClass(v string) (r *VListItemBuilder) {
	b.tag.Attr("exact-active-class", v)
	return b
}

func (b *VListItemBuilder) ExactPath(v bool) (r *VListItemBuilder) {
	b.tag.Attr(":exact-path", fmt.Sprint(v))
	return b
}

func (b *VListItemBuilder) Href(v interface{}) (r *VListItemBuilder) {
	b.tag.Attr(":href", h.JSONString(v))
	return b
}

func (b *VListItemBuilder) Inactive(v bool) (r *VListItemBuilder) {
	b.tag.Attr(":inactive", fmt.Sprint(v))
	return b
}

func (b *VListItemBuilder) InputValue(v interface{}) (r *VListItemBuilder) {
	b.tag.Attr(":input-value", h.JSONString(v))
	return b
}

func (b *VListItemBuilder) Light(v bool) (r *VListItemBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VListItemBuilder) Link(v bool) (r *VListItemBuilder) {
	b.tag.Attr(":link", fmt.Sprint(v))
	return b
}

func (b *VListItemBuilder) Nuxt(v bool) (r *VListItemBuilder) {
	b.tag.Attr(":nuxt", fmt.Sprint(v))
	return b
}

func (b *VListItemBuilder) Replace(v bool) (r *VListItemBuilder) {
	b.tag.Attr(":replace", fmt.Sprint(v))
	return b
}

func (b *VListItemBuilder) Ripple(v interface{}) (r *VListItemBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VListItemBuilder) Selectable(v bool) (r *VListItemBuilder) {
	b.tag.Attr(":selectable", fmt.Sprint(v))
	return b
}

func (b *VListItemBuilder) Tag(v string) (r *VListItemBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VListItemBuilder) Target(v string) (r *VListItemBuilder) {
	b.tag.Attr("target", v)
	return b
}

func (b *VListItemBuilder) ThreeLine(v bool) (r *VListItemBuilder) {
	b.tag.Attr(":three-line", fmt.Sprint(v))
	return b
}

func (b *VListItemBuilder) To(v interface{}) (r *VListItemBuilder) {
	b.tag.Attr(":to", h.JSONString(v))
	return b
}

func (b *VListItemBuilder) TwoLine(v bool) (r *VListItemBuilder) {
	b.tag.Attr(":two-line", fmt.Sprint(v))
	return b
}

func (b *VListItemBuilder) Value(v interface{}) (r *VListItemBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VListItemBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VListItemBuilder) Attr(vs ...interface{}) (r *VListItemBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VListItemBuilder) Children(children ...h.HTMLComponent) (r *VListItemBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VListItemBuilder) AppendChildren(children ...h.HTMLComponent) (r *VListItemBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VListItemBuilder) PrependChildren(children ...h.HTMLComponent) (r *VListItemBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VListItemBuilder) Class(names ...string) (r *VListItemBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VListItemBuilder) ClassIf(name string, add bool) (r *VListItemBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VListItemBuilder) On(name string, value string) (r *VListItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VListItemBuilder) Bind(name string, value string) (r *VListItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VListItemBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
