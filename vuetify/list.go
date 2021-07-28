package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VListBuilder struct {
	tag *h.HTMLTagBuilder
}

func VList(children ...h.HTMLComponent) (r *VListBuilder) {
	r = &VListBuilder{
		tag: h.Tag("v-list").Children(children...),
	}
	return
}

func (b *VListBuilder) Color(v string) (r *VListBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VListBuilder) Dark(v bool) (r *VListBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) Dense(v bool) (r *VListBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) Disabled(v bool) (r *VListBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) Elevation(v int) (r *VListBuilder) {
	b.tag.Attr(":elevation", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) Expand(v bool) (r *VListBuilder) {
	b.tag.Attr(":expand", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) Flat(v bool) (r *VListBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) Height(v int) (r *VListBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) Light(v bool) (r *VListBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) MaxHeight(v int) (r *VListBuilder) {
	b.tag.Attr(":max-height", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) MaxWidth(v int) (r *VListBuilder) {
	b.tag.Attr(":max-width", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) MinHeight(v int) (r *VListBuilder) {
	b.tag.Attr(":min-height", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) MinWidth(v int) (r *VListBuilder) {
	b.tag.Attr(":min-width", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) Nav(v bool) (r *VListBuilder) {
	b.tag.Attr(":nav", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) Outlined(v bool) (r *VListBuilder) {
	b.tag.Attr(":outlined", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) Rounded(v bool) (r *VListBuilder) {
	b.tag.Attr(":rounded", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) Shaped(v bool) (r *VListBuilder) {
	b.tag.Attr(":shaped", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) Subheader(v bool) (r *VListBuilder) {
	b.tag.Attr(":subheader", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) Tag(v string) (r *VListBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VListBuilder) ThreeLine(v bool) (r *VListBuilder) {
	b.tag.Attr(":three-line", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) Tile(v bool) (r *VListBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) TwoLine(v bool) (r *VListBuilder) {
	b.tag.Attr(":two-line", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) Width(v int) (r *VListBuilder) {
	b.tag.Attr(":width", fmt.Sprint(v))
	return b
}

func (b *VListBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VListBuilder) Attr(vs ...interface{}) (r *VListBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VListBuilder) Children(children ...h.HTMLComponent) (r *VListBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VListBuilder) AppendChildren(children ...h.HTMLComponent) (r *VListBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VListBuilder) PrependChildren(children ...h.HTMLComponent) (r *VListBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VListBuilder) Class(names ...string) (r *VListBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VListBuilder) ClassIf(name string, add bool) (r *VListBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VListBuilder) On(name string, value string) (r *VListBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VListBuilder) Bind(name string, value string) (r *VListBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VListBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
