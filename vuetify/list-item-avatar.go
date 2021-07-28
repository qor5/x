package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VListItemAvatarBuilder struct {
	tag *h.HTMLTagBuilder
}

func VListItemAvatar(children ...h.HTMLComponent) (r *VListItemAvatarBuilder) {
	r = &VListItemAvatarBuilder{
		tag: h.Tag("v-list-item-avatar").Children(children...),
	}
	return
}

func (b *VListItemAvatarBuilder) Color(v string) (r *VListItemAvatarBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VListItemAvatarBuilder) Height(v int) (r *VListItemAvatarBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VListItemAvatarBuilder) Horizontal(v bool) (r *VListItemAvatarBuilder) {
	b.tag.Attr(":horizontal", fmt.Sprint(v))
	return b
}

func (b *VListItemAvatarBuilder) Left(v bool) (r *VListItemAvatarBuilder) {
	b.tag.Attr(":left", fmt.Sprint(v))
	return b
}

func (b *VListItemAvatarBuilder) MaxHeight(v int) (r *VListItemAvatarBuilder) {
	b.tag.Attr(":max-height", fmt.Sprint(v))
	return b
}

func (b *VListItemAvatarBuilder) MaxWidth(v int) (r *VListItemAvatarBuilder) {
	b.tag.Attr(":max-width", fmt.Sprint(v))
	return b
}

func (b *VListItemAvatarBuilder) MinHeight(v int) (r *VListItemAvatarBuilder) {
	b.tag.Attr(":min-height", fmt.Sprint(v))
	return b
}

func (b *VListItemAvatarBuilder) MinWidth(v int) (r *VListItemAvatarBuilder) {
	b.tag.Attr(":min-width", fmt.Sprint(v))
	return b
}

func (b *VListItemAvatarBuilder) Right(v bool) (r *VListItemAvatarBuilder) {
	b.tag.Attr(":right", fmt.Sprint(v))
	return b
}

func (b *VListItemAvatarBuilder) Rounded(v bool) (r *VListItemAvatarBuilder) {
	b.tag.Attr(":rounded", fmt.Sprint(v))
	return b
}

func (b *VListItemAvatarBuilder) Size(v int) (r *VListItemAvatarBuilder) {
	b.tag.Attr(":size", fmt.Sprint(v))
	return b
}

func (b *VListItemAvatarBuilder) Tile(v bool) (r *VListItemAvatarBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VListItemAvatarBuilder) Width(v int) (r *VListItemAvatarBuilder) {
	b.tag.Attr(":width", fmt.Sprint(v))
	return b
}

func (b *VListItemAvatarBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VListItemAvatarBuilder) Attr(vs ...interface{}) (r *VListItemAvatarBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VListItemAvatarBuilder) Children(children ...h.HTMLComponent) (r *VListItemAvatarBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VListItemAvatarBuilder) AppendChildren(children ...h.HTMLComponent) (r *VListItemAvatarBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VListItemAvatarBuilder) PrependChildren(children ...h.HTMLComponent) (r *VListItemAvatarBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VListItemAvatarBuilder) Class(names ...string) (r *VListItemAvatarBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VListItemAvatarBuilder) ClassIf(name string, add bool) (r *VListItemAvatarBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VListItemAvatarBuilder) On(name string, value string) (r *VListItemAvatarBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VListItemAvatarBuilder) Bind(name string, value string) (r *VListItemAvatarBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VListItemAvatarBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
