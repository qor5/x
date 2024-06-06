package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VToolbarTitleBuilder struct {
	tag *h.HTMLTagBuilder
}

func (b *VToolbarTitleBuilder) Text(v string) (r *VToolbarTitleBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VToolbarTitleBuilder) Tag(v string) (r *VToolbarTitleBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VToolbarTitleBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VToolbarTitleBuilder) Attr(vs ...interface{}) (r *VToolbarTitleBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VToolbarTitleBuilder) Children(children ...h.HTMLComponent) (r *VToolbarTitleBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VToolbarTitleBuilder) AppendChildren(children ...h.HTMLComponent) (r *VToolbarTitleBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VToolbarTitleBuilder) PrependChildren(children ...h.HTMLComponent) (r *VToolbarTitleBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VToolbarTitleBuilder) Class(names ...string) (r *VToolbarTitleBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VToolbarTitleBuilder) ClassIf(name string, add bool) (r *VToolbarTitleBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VToolbarTitleBuilder) On(name string, value string) (r *VToolbarTitleBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VToolbarTitleBuilder) Bind(name string, value string) (r *VToolbarTitleBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VToolbarTitleBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
