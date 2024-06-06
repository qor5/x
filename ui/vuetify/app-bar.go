package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VAppBarBuilder struct {
	tag *h.HTMLTagBuilder
}

func VAppBar(children ...h.HTMLComponent) (r *VAppBarBuilder) {
	r = &VAppBarBuilder{
		tag: h.Tag("v-app-bar").Children(children...),
	}
	return
}

func (b *VAppBarBuilder) Image(v string) (r *VAppBarBuilder) {
	b.tag.Attr("image", v)
	return b
}

func (b *VAppBarBuilder) Title(v string) (r *VAppBarBuilder) {
	b.tag.Attr("title", v)
	return b
}

func (b *VAppBarBuilder) Flat(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) Collapse(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":collapse", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) ModelValue(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) Location(v interface{}) (r *VAppBarBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VAppBarBuilder) Absolute(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) Color(v string) (r *VAppBarBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VAppBarBuilder) Density(v interface{}) (r *VAppBarBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VAppBarBuilder) Extended(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":extended", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) ExtensionHeight(v interface{}) (r *VAppBarBuilder) {
	b.tag.Attr(":extension-height", h.JSONString(v))
	return b
}

func (b *VAppBarBuilder) Floating(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":floating", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) Height(v interface{}) (r *VAppBarBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VAppBarBuilder) Border(v interface{}) (r *VAppBarBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VAppBarBuilder) Elevation(v interface{}) (r *VAppBarBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VAppBarBuilder) Rounded(v interface{}) (r *VAppBarBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VAppBarBuilder) Tile(v bool) (r *VAppBarBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VAppBarBuilder) Tag(v string) (r *VAppBarBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VAppBarBuilder) Theme(v string) (r *VAppBarBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VAppBarBuilder) Name(v string) (r *VAppBarBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VAppBarBuilder) Order(v interface{}) (r *VAppBarBuilder) {
	b.tag.Attr(":order", h.JSONString(v))
	return b
}

func (b *VAppBarBuilder) ScrollTarget(v string) (r *VAppBarBuilder) {
	b.tag.Attr("scroll-target", v)
	return b
}

func (b *VAppBarBuilder) ScrollThreshold(v interface{}) (r *VAppBarBuilder) {
	b.tag.Attr(":scroll-threshold", h.JSONString(v))
	return b
}

func (b *VAppBarBuilder) ScrollBehavior(v interface{}) (r *VAppBarBuilder) {
	b.tag.Attr(":scroll-behavior", h.JSONString(v))
	return b
}

func (b *VAppBarBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VAppBarBuilder) Attr(vs ...interface{}) (r *VAppBarBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VAppBarBuilder) Children(children ...h.HTMLComponent) (r *VAppBarBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VAppBarBuilder) AppendChildren(children ...h.HTMLComponent) (r *VAppBarBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VAppBarBuilder) PrependChildren(children ...h.HTMLComponent) (r *VAppBarBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VAppBarBuilder) Class(names ...string) (r *VAppBarBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VAppBarBuilder) ClassIf(name string, add bool) (r *VAppBarBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VAppBarBuilder) On(name string, value string) (r *VAppBarBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VAppBarBuilder) Bind(name string, value string) (r *VAppBarBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VAppBarBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
