package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VListGroupBuilder struct {
	tag *h.HTMLTagBuilder
}

func VListGroup(children ...h.HTMLComponent) (r *VListGroupBuilder) {
	r = &VListGroupBuilder{
		tag: h.Tag("v-list-group").Children(children...),
	}
	return
}

func (b *VListGroupBuilder) ActiveColor(v string) (r *VListGroupBuilder) {
	b.tag.Attr("active-color", v)
	return b
}

func (b *VListGroupBuilder) BaseColor(v string) (r *VListGroupBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VListGroupBuilder) Color(v string) (r *VListGroupBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VListGroupBuilder) CollapseIcon(v interface{}) (r *VListGroupBuilder) {
	b.tag.Attr(":collapse-icon", h.JSONString(v))
	return b
}

func (b *VListGroupBuilder) ExpandIcon(v interface{}) (r *VListGroupBuilder) {
	b.tag.Attr(":expand-icon", h.JSONString(v))
	return b
}

func (b *VListGroupBuilder) PrependIcon(v interface{}) (r *VListGroupBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VListGroupBuilder) AppendIcon(v interface{}) (r *VListGroupBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VListGroupBuilder) Fluid(v bool) (r *VListGroupBuilder) {
	b.tag.Attr(":fluid", fmt.Sprint(v))
	return b
}

func (b *VListGroupBuilder) Subgroup(v bool) (r *VListGroupBuilder) {
	b.tag.Attr(":subgroup", fmt.Sprint(v))
	return b
}

func (b *VListGroupBuilder) Title(v string) (r *VListGroupBuilder) {
	b.tag.Attr("title", v)
	return b
}

func (b *VListGroupBuilder) Value(v interface{}) (r *VListGroupBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VListGroupBuilder) Tag(v string) (r *VListGroupBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VListGroupBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VListGroupBuilder) Attr(vs ...interface{}) (r *VListGroupBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VListGroupBuilder) Children(children ...h.HTMLComponent) (r *VListGroupBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VListGroupBuilder) AppendChildren(children ...h.HTMLComponent) (r *VListGroupBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VListGroupBuilder) PrependChildren(children ...h.HTMLComponent) (r *VListGroupBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VListGroupBuilder) Class(names ...string) (r *VListGroupBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VListGroupBuilder) ClassIf(name string, add bool) (r *VListGroupBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VListGroupBuilder) On(name string, value string) (r *VListGroupBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VListGroupBuilder) Bind(name string, value string) (r *VListGroupBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VListGroupBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
