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

func (b *VToolbarBuilder) Absolute(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Bottom(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":bottom", fmt.Sprint(v))
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

func (b *VToolbarBuilder) Dark(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Dense(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Elevation(v int) (r *VToolbarBuilder) {
	b.tag.Attr(":elevation", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Extended(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":extended", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) ExtensionHeight(v int) (r *VToolbarBuilder) {
	b.tag.Attr(":extension-height", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Flat(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Floating(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":floating", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Height(v int) (r *VToolbarBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Light(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) MaxHeight(v int) (r *VToolbarBuilder) {
	b.tag.Attr(":max-height", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) MaxWidth(v int) (r *VToolbarBuilder) {
	b.tag.Attr(":max-width", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) MinHeight(v int) (r *VToolbarBuilder) {
	b.tag.Attr(":min-height", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) MinWidth(v int) (r *VToolbarBuilder) {
	b.tag.Attr(":min-width", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Outlined(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":outlined", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Prominent(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":prominent", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Rounded(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":rounded", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Shaped(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":shaped", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Short(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":short", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Src(v interface{}) (r *VToolbarBuilder) {
	b.tag.Attr(":src", h.JSONString(v))
	return b
}

func (b *VToolbarBuilder) Tag(v string) (r *VToolbarBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VToolbarBuilder) Tile(v bool) (r *VToolbarBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VToolbarBuilder) Width(v int) (r *VToolbarBuilder) {
	b.tag.Attr(":width", fmt.Sprint(v))
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
