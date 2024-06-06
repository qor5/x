package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VAvatarBuilder struct {
	tag *h.HTMLTagBuilder
}

func VAvatar(children ...h.HTMLComponent) (r *VAvatarBuilder) {
	r = &VAvatarBuilder{
		tag: h.Tag("v-avatar").Children(children...),
	}
	return
}

func (b *VAvatarBuilder) Start(v bool) (r *VAvatarBuilder) {
	b.tag.Attr(":start", fmt.Sprint(v))
	return b
}

func (b *VAvatarBuilder) End(v bool) (r *VAvatarBuilder) {
	b.tag.Attr(":end", fmt.Sprint(v))
	return b
}

func (b *VAvatarBuilder) Icon(v interface{}) (r *VAvatarBuilder) {
	b.tag.Attr(":icon", h.JSONString(v))
	return b
}

func (b *VAvatarBuilder) Image(v string) (r *VAvatarBuilder) {
	b.tag.Attr("image", v)
	return b
}

func (b *VAvatarBuilder) Text(v string) (r *VAvatarBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VAvatarBuilder) Density(v interface{}) (r *VAvatarBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VAvatarBuilder) Rounded(v interface{}) (r *VAvatarBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VAvatarBuilder) Tile(v bool) (r *VAvatarBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VAvatarBuilder) Size(v interface{}) (r *VAvatarBuilder) {
	b.tag.Attr(":size", h.JSONString(v))
	return b
}

func (b *VAvatarBuilder) Tag(v string) (r *VAvatarBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VAvatarBuilder) Theme(v string) (r *VAvatarBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VAvatarBuilder) Color(v string) (r *VAvatarBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VAvatarBuilder) Variant(v interface{}) (r *VAvatarBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VAvatarBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VAvatarBuilder) Attr(vs ...interface{}) (r *VAvatarBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VAvatarBuilder) Children(children ...h.HTMLComponent) (r *VAvatarBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VAvatarBuilder) AppendChildren(children ...h.HTMLComponent) (r *VAvatarBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VAvatarBuilder) PrependChildren(children ...h.HTMLComponent) (r *VAvatarBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VAvatarBuilder) Class(names ...string) (r *VAvatarBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VAvatarBuilder) ClassIf(name string, add bool) (r *VAvatarBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VAvatarBuilder) On(name string, value string) (r *VAvatarBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VAvatarBuilder) Bind(name string, value string) (r *VAvatarBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VAvatarBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
