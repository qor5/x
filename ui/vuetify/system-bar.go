package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSystemBarBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSystemBar(children ...h.HTMLComponent) (r *VSystemBarBuilder) {
	r = &VSystemBarBuilder{
		tag: h.Tag("v-system-bar").Children(children...),
	}
	return
}

func (b *VSystemBarBuilder) Color(v string) (r *VSystemBarBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VSystemBarBuilder) Height(v interface{}) (r *VSystemBarBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VSystemBarBuilder) Window(v bool) (r *VSystemBarBuilder) {
	b.tag.Attr(":window", fmt.Sprint(v))
	return b
}

func (b *VSystemBarBuilder) Elevation(v interface{}) (r *VSystemBarBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VSystemBarBuilder) Name(v string) (r *VSystemBarBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VSystemBarBuilder) Order(v interface{}) (r *VSystemBarBuilder) {
	b.tag.Attr(":order", h.JSONString(v))
	return b
}

func (b *VSystemBarBuilder) Absolute(v bool) (r *VSystemBarBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VSystemBarBuilder) Rounded(v interface{}) (r *VSystemBarBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VSystemBarBuilder) Tile(v bool) (r *VSystemBarBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VSystemBarBuilder) Tag(v string) (r *VSystemBarBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VSystemBarBuilder) Theme(v string) (r *VSystemBarBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VSystemBarBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSystemBarBuilder) Attr(vs ...interface{}) (r *VSystemBarBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSystemBarBuilder) Children(children ...h.HTMLComponent) (r *VSystemBarBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSystemBarBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSystemBarBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSystemBarBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSystemBarBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSystemBarBuilder) Class(names ...string) (r *VSystemBarBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSystemBarBuilder) ClassIf(name string, add bool) (r *VSystemBarBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSystemBarBuilder) On(name string, value string) (r *VSystemBarBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSystemBarBuilder) Bind(name string, value string) (r *VSystemBarBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSystemBarBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
