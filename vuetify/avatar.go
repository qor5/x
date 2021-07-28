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

func (b *VAvatarBuilder) Color(v string) (r *VAvatarBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VAvatarBuilder) Height(v int) (r *VAvatarBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VAvatarBuilder) Left(v bool) (r *VAvatarBuilder) {
	b.tag.Attr(":left", fmt.Sprint(v))
	return b
}

func (b *VAvatarBuilder) MaxHeight(v int) (r *VAvatarBuilder) {
	b.tag.Attr(":max-height", fmt.Sprint(v))
	return b
}

func (b *VAvatarBuilder) MaxWidth(v int) (r *VAvatarBuilder) {
	b.tag.Attr(":max-width", fmt.Sprint(v))
	return b
}

func (b *VAvatarBuilder) MinHeight(v int) (r *VAvatarBuilder) {
	b.tag.Attr(":min-height", fmt.Sprint(v))
	return b
}

func (b *VAvatarBuilder) MinWidth(v int) (r *VAvatarBuilder) {
	b.tag.Attr(":min-width", fmt.Sprint(v))
	return b
}

func (b *VAvatarBuilder) Right(v bool) (r *VAvatarBuilder) {
	b.tag.Attr(":right", fmt.Sprint(v))
	return b
}

func (b *VAvatarBuilder) Rounded(v bool) (r *VAvatarBuilder) {
	b.tag.Attr(":rounded", fmt.Sprint(v))
	return b
}

func (b *VAvatarBuilder) Size(v int) (r *VAvatarBuilder) {
	b.tag.Attr(":size", fmt.Sprint(v))
	return b
}

func (b *VAvatarBuilder) Tile(v bool) (r *VAvatarBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VAvatarBuilder) Width(v int) (r *VAvatarBuilder) {
	b.tag.Attr(":width", fmt.Sprint(v))
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
