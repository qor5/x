package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VListItemSubtitleBuilder struct {
	tag *h.HTMLTagBuilder
}

func VListItemSubtitle(children ...h.HTMLComponent) (r *VListItemSubtitleBuilder) {
	r = &VListItemSubtitleBuilder{
		tag: h.Tag("v-list-item-subtitle").Children(children...),
	}
	return
}

func (b *VListItemSubtitleBuilder) Opacity(v interface{}) (r *VListItemSubtitleBuilder) {
	b.tag.Attr(":opacity", h.JSONString(v))
	return b
}

func (b *VListItemSubtitleBuilder) Tag(v string) (r *VListItemSubtitleBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VListItemSubtitleBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VListItemSubtitleBuilder) Attr(vs ...interface{}) (r *VListItemSubtitleBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VListItemSubtitleBuilder) Children(children ...h.HTMLComponent) (r *VListItemSubtitleBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VListItemSubtitleBuilder) AppendChildren(children ...h.HTMLComponent) (r *VListItemSubtitleBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VListItemSubtitleBuilder) PrependChildren(children ...h.HTMLComponent) (r *VListItemSubtitleBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VListItemSubtitleBuilder) Class(names ...string) (r *VListItemSubtitleBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VListItemSubtitleBuilder) ClassIf(name string, add bool) (r *VListItemSubtitleBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VListItemSubtitleBuilder) On(name string, value string) (r *VListItemSubtitleBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VListItemSubtitleBuilder) Bind(name string, value string) (r *VListItemSubtitleBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VListItemSubtitleBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
