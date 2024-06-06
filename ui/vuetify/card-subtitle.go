package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCardSubtitleBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCardSubtitle(children ...h.HTMLComponent) (r *VCardSubtitleBuilder) {
	r = &VCardSubtitleBuilder{
		tag: h.Tag("v-card-subtitle").Children(children...),
	}
	return
}

func (b *VCardSubtitleBuilder) Opacity(v interface{}) (r *VCardSubtitleBuilder) {
	b.tag.Attr(":opacity", h.JSONString(v))
	return b
}

func (b *VCardSubtitleBuilder) Tag(v string) (r *VCardSubtitleBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VCardSubtitleBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VCardSubtitleBuilder) Attr(vs ...interface{}) (r *VCardSubtitleBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VCardSubtitleBuilder) Children(children ...h.HTMLComponent) (r *VCardSubtitleBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VCardSubtitleBuilder) AppendChildren(children ...h.HTMLComponent) (r *VCardSubtitleBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VCardSubtitleBuilder) PrependChildren(children ...h.HTMLComponent) (r *VCardSubtitleBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VCardSubtitleBuilder) Class(names ...string) (r *VCardSubtitleBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCardSubtitleBuilder) ClassIf(name string, add bool) (r *VCardSubtitleBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCardSubtitleBuilder) On(name string, value string) (r *VCardSubtitleBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCardSubtitleBuilder) Bind(name string, value string) (r *VCardSubtitleBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCardSubtitleBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
