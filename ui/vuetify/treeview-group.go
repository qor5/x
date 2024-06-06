package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTreeviewGroupBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTreeviewGroup(children ...h.HTMLComponent) (r *VTreeviewGroupBuilder) {
	r = &VTreeviewGroupBuilder{
		tag: h.Tag("v-treeview-group").Children(children...),
	}
	return
}

func (b *VTreeviewGroupBuilder) ActiveColor(v string) (r *VTreeviewGroupBuilder) {
	b.tag.Attr("active-color", v)
	return b
}

func (b *VTreeviewGroupBuilder) BaseColor(v string) (r *VTreeviewGroupBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VTreeviewGroupBuilder) Color(v string) (r *VTreeviewGroupBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VTreeviewGroupBuilder) CollapseIcon(v interface{}) (r *VTreeviewGroupBuilder) {
	b.tag.Attr(":collapse-icon", h.JSONString(v))
	return b
}

func (b *VTreeviewGroupBuilder) ExpandIcon(v interface{}) (r *VTreeviewGroupBuilder) {
	b.tag.Attr(":expand-icon", h.JSONString(v))
	return b
}

func (b *VTreeviewGroupBuilder) PrependIcon(v interface{}) (r *VTreeviewGroupBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VTreeviewGroupBuilder) AppendIcon(v interface{}) (r *VTreeviewGroupBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VTreeviewGroupBuilder) Fluid(v bool) (r *VTreeviewGroupBuilder) {
	b.tag.Attr(":fluid", fmt.Sprint(v))
	return b
}

func (b *VTreeviewGroupBuilder) Title(v string) (r *VTreeviewGroupBuilder) {
	b.tag.Attr("title", v)
	return b
}

func (b *VTreeviewGroupBuilder) Value(v interface{}) (r *VTreeviewGroupBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VTreeviewGroupBuilder) Tag(v string) (r *VTreeviewGroupBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VTreeviewGroupBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VTreeviewGroupBuilder) Attr(vs ...interface{}) (r *VTreeviewGroupBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VTreeviewGroupBuilder) Children(children ...h.HTMLComponent) (r *VTreeviewGroupBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VTreeviewGroupBuilder) AppendChildren(children ...h.HTMLComponent) (r *VTreeviewGroupBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VTreeviewGroupBuilder) PrependChildren(children ...h.HTMLComponent) (r *VTreeviewGroupBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VTreeviewGroupBuilder) Class(names ...string) (r *VTreeviewGroupBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTreeviewGroupBuilder) ClassIf(name string, add bool) (r *VTreeviewGroupBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTreeviewGroupBuilder) On(name string, value string) (r *VTreeviewGroupBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTreeviewGroupBuilder) Bind(name string, value string) (r *VTreeviewGroupBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTreeviewGroupBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
