package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSheetBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSheet(children ...h.HTMLComponent) (r *VSheetBuilder) {
	r = &VSheetBuilder{
		tag: h.Tag("v-sheet").Children(children...),
	}
	return
}

func (b *VSheetBuilder) Color(v string) (r *VSheetBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VSheetBuilder) Dark(v bool) (r *VSheetBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VSheetBuilder) Elevation(v int) (r *VSheetBuilder) {
	b.tag.Attr(":elevation", fmt.Sprint(v))
	return b
}

func (b *VSheetBuilder) Height(v int) (r *VSheetBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VSheetBuilder) Light(v bool) (r *VSheetBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VSheetBuilder) MaxHeight(v int) (r *VSheetBuilder) {
	b.tag.Attr(":max-height", fmt.Sprint(v))
	return b
}

func (b *VSheetBuilder) MaxWidth(v int) (r *VSheetBuilder) {
	b.tag.Attr(":max-width", fmt.Sprint(v))
	return b
}

func (b *VSheetBuilder) MinHeight(v int) (r *VSheetBuilder) {
	b.tag.Attr(":min-height", fmt.Sprint(v))
	return b
}

func (b *VSheetBuilder) MinWidth(v int) (r *VSheetBuilder) {
	b.tag.Attr(":min-width", fmt.Sprint(v))
	return b
}

func (b *VSheetBuilder) Outlined(v bool) (r *VSheetBuilder) {
	b.tag.Attr(":outlined", fmt.Sprint(v))
	return b
}

func (b *VSheetBuilder) Rounded(v bool) (r *VSheetBuilder) {
	b.tag.Attr(":rounded", fmt.Sprint(v))
	return b
}

func (b *VSheetBuilder) Shaped(v bool) (r *VSheetBuilder) {
	b.tag.Attr(":shaped", fmt.Sprint(v))
	return b
}

func (b *VSheetBuilder) Tag(v string) (r *VSheetBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VSheetBuilder) Tile(v bool) (r *VSheetBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VSheetBuilder) Width(v int) (r *VSheetBuilder) {
	b.tag.Attr(":width", fmt.Sprint(v))
	return b
}

func (b *VSheetBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSheetBuilder) Attr(vs ...interface{}) (r *VSheetBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSheetBuilder) Children(children ...h.HTMLComponent) (r *VSheetBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSheetBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSheetBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSheetBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSheetBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSheetBuilder) Class(names ...string) (r *VSheetBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSheetBuilder) ClassIf(name string, add bool) (r *VSheetBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSheetBuilder) On(name string, value string) (r *VSheetBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSheetBuilder) Bind(name string, value string) (r *VSheetBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSheetBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
