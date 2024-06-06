package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VToolbarBuilder struct {
	tag *h.HTMLTagBuilder
}

func VToolbar(children ...h.HTMLComponent) (r *VToolbarBuilder) {
	r = &VToolbarBuilder{
		tag: h.Tag("v-toolbar").Children(children...),
	}
	return
}

func (b *VToolbarBuilder) Image(v string) (r *VToolbarBuilder) {
	b.tag.Attr("image", v)
	return b
}

func (b *VToolbarBuilder) Title(v string) (r *VToolbarBuilder) {
	b.tag.Attr("title", v)
	return b
}

func (b *VToolbarBuilder) Flat(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Absolute(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Collapse(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":collapse", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Color(v string) (r *VToolbarBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VToolbarBuilder) Density(v interface{}) (r *VToolbarBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VToolbarBuilder) Extended(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":extended", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) ExtensionHeight(v interface{}) (r *VToolbarBuilder) {
	b.tag.Attr(":extension-height", h.JSONString(v))
	return b
}

func (b *VToolbarBuilder) Floating(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":floating", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Height(v interface{}) (r *VToolbarBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VToolbarBuilder) Border(v interface{}) (r *VToolbarBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VToolbarBuilder) Elevation(v interface{}) (r *VToolbarBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VToolbarBuilder) Rounded(v interface{}) (r *VToolbarBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VToolbarBuilder) Tile(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Tag(v string) (r *VToolbarBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VToolbarBuilder) Theme(v string) (r *VToolbarBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VToolbarBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VToolbarBuilder) Attr(vs ...interface{}) (r *VToolbarBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VToolbarBuilder) Children(children ...h.HTMLComponent) (r *VToolbarBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VToolbarBuilder) AppendChildren(children ...h.HTMLComponent) (r *VToolbarBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VToolbarBuilder) PrependChildren(children ...h.HTMLComponent) (r *VToolbarBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VToolbarBuilder) Class(names ...string) (r *VToolbarBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VToolbarBuilder) ClassIf(name string, add bool) (r *VToolbarBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VToolbarBuilder) On(name string, value string) (r *VToolbarBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VToolbarBuilder) Bind(name string, value string) (r *VToolbarBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VToolbarBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
