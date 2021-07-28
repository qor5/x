package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VListItemGroupBuilder struct {
	tag *h.HTMLTagBuilder
}

func VListItemGroup(children ...h.HTMLComponent) (r *VListItemGroupBuilder) {
	r = &VListItemGroupBuilder{
		tag: h.Tag("v-list-item-group").Children(children...),
	}
	return
}

func (b *VListItemGroupBuilder) ActiveClass(v string) (r *VListItemGroupBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VListItemGroupBuilder) Color(v string) (r *VListItemGroupBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VListItemGroupBuilder) Dark(v bool) (r *VListItemGroupBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VListItemGroupBuilder) Light(v bool) (r *VListItemGroupBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VListItemGroupBuilder) Mandatory(v bool) (r *VListItemGroupBuilder) {
	b.tag.Attr(":mandatory", fmt.Sprint(v))
	return b
}

func (b *VListItemGroupBuilder) Max(v int) (r *VListItemGroupBuilder) {
	b.tag.Attr(":max", fmt.Sprint(v))
	return b
}

func (b *VListItemGroupBuilder) Multiple(v bool) (r *VListItemGroupBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VListItemGroupBuilder) Tag(v string) (r *VListItemGroupBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VListItemGroupBuilder) Value(v interface{}) (r *VListItemGroupBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VListItemGroupBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VListItemGroupBuilder) Attr(vs ...interface{}) (r *VListItemGroupBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VListItemGroupBuilder) Children(children ...h.HTMLComponent) (r *VListItemGroupBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VListItemGroupBuilder) AppendChildren(children ...h.HTMLComponent) (r *VListItemGroupBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VListItemGroupBuilder) PrependChildren(children ...h.HTMLComponent) (r *VListItemGroupBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VListItemGroupBuilder) Class(names ...string) (r *VListItemGroupBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VListItemGroupBuilder) ClassIf(name string, add bool) (r *VListItemGroupBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VListItemGroupBuilder) On(name string, value string) (r *VListItemGroupBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VListItemGroupBuilder) Bind(name string, value string) (r *VListItemGroupBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VListItemGroupBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
