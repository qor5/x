package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VFooterBuilder struct {
	tag *h.HTMLTagBuilder
}

func VFooter(children ...h.HTMLComponent) (r *VFooterBuilder) {
	r = &VFooterBuilder{
		tag: h.Tag("v-footer").Children(children...),
	}
	return
}

func (b *VFooterBuilder) App(v bool) (r *VFooterBuilder) {
	b.tag.Attr(":app", fmt.Sprint(v))
	return b
}

func (b *VFooterBuilder) Color(v string) (r *VFooterBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VFooterBuilder) Height(v interface{}) (r *VFooterBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VFooterBuilder) Border(v interface{}) (r *VFooterBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VFooterBuilder) Elevation(v interface{}) (r *VFooterBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VFooterBuilder) Name(v string) (r *VFooterBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VFooterBuilder) Order(v interface{}) (r *VFooterBuilder) {
	b.tag.Attr(":order", h.JSONString(v))
	return b
}

func (b *VFooterBuilder) Absolute(v bool) (r *VFooterBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VFooterBuilder) Rounded(v interface{}) (r *VFooterBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VFooterBuilder) Tile(v bool) (r *VFooterBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VFooterBuilder) Tag(v string) (r *VFooterBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VFooterBuilder) Theme(v string) (r *VFooterBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VFooterBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VFooterBuilder) Attr(vs ...interface{}) (r *VFooterBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VFooterBuilder) Children(children ...h.HTMLComponent) (r *VFooterBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VFooterBuilder) AppendChildren(children ...h.HTMLComponent) (r *VFooterBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VFooterBuilder) PrependChildren(children ...h.HTMLComponent) (r *VFooterBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VFooterBuilder) Class(names ...string) (r *VFooterBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VFooterBuilder) ClassIf(name string, add bool) (r *VFooterBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VFooterBuilder) On(name string, value string) (r *VFooterBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VFooterBuilder) Bind(name string, value string) (r *VFooterBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VFooterBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
