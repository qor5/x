package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VRowBuilder struct {
	tag *h.HTMLTagBuilder
}

func VRow(children ...h.HTMLComponent) (r *VRowBuilder) {
	r = &VRowBuilder{
		tag: h.Tag("v-row").Children(children...),
	}
	return
}

func (b *VRowBuilder) Align(v string) (r *VRowBuilder) {
	b.tag.Attr("align", v)
	return b
}

func (b *VRowBuilder) AlignContent(v string) (r *VRowBuilder) {
	b.tag.Attr("align-content", v)
	return b
}

func (b *VRowBuilder) AlignContentLg(v string) (r *VRowBuilder) {
	b.tag.Attr("align-content-lg", v)
	return b
}

func (b *VRowBuilder) AlignContentMd(v string) (r *VRowBuilder) {
	b.tag.Attr("align-content-md", v)
	return b
}

func (b *VRowBuilder) AlignContentSm(v string) (r *VRowBuilder) {
	b.tag.Attr("align-content-sm", v)
	return b
}

func (b *VRowBuilder) AlignContentXl(v string) (r *VRowBuilder) {
	b.tag.Attr("align-content-xl", v)
	return b
}

func (b *VRowBuilder) AlignLg(v string) (r *VRowBuilder) {
	b.tag.Attr("align-lg", v)
	return b
}

func (b *VRowBuilder) AlignMd(v string) (r *VRowBuilder) {
	b.tag.Attr("align-md", v)
	return b
}

func (b *VRowBuilder) AlignSm(v string) (r *VRowBuilder) {
	b.tag.Attr("align-sm", v)
	return b
}

func (b *VRowBuilder) AlignXl(v string) (r *VRowBuilder) {
	b.tag.Attr("align-xl", v)
	return b
}

func (b *VRowBuilder) Dense(v bool) (r *VRowBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
	return b
}

func (b *VRowBuilder) Justify(v string) (r *VRowBuilder) {
	b.tag.Attr("justify", v)
	return b
}

func (b *VRowBuilder) JustifyLg(v string) (r *VRowBuilder) {
	b.tag.Attr("justify-lg", v)
	return b
}

func (b *VRowBuilder) JustifyMd(v string) (r *VRowBuilder) {
	b.tag.Attr("justify-md", v)
	return b
}

func (b *VRowBuilder) JustifySm(v string) (r *VRowBuilder) {
	b.tag.Attr("justify-sm", v)
	return b
}

func (b *VRowBuilder) JustifyXl(v string) (r *VRowBuilder) {
	b.tag.Attr("justify-xl", v)
	return b
}

func (b *VRowBuilder) NoGutters(v bool) (r *VRowBuilder) {
	b.tag.Attr(":no-gutters", fmt.Sprint(v))
	return b
}

func (b *VRowBuilder) Tag(v string) (r *VRowBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VRowBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VRowBuilder) Attr(vs ...interface{}) (r *VRowBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VRowBuilder) Children(children ...h.HTMLComponent) (r *VRowBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VRowBuilder) AppendChildren(children ...h.HTMLComponent) (r *VRowBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VRowBuilder) PrependChildren(children ...h.HTMLComponent) (r *VRowBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VRowBuilder) Class(names ...string) (r *VRowBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VRowBuilder) ClassIf(name string, add bool) (r *VRowBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VRowBuilder) On(name string, value string) (r *VRowBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VRowBuilder) Bind(name string, value string) (r *VRowBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VRowBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
