package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VListItemTitleBuilder struct {
	tag *h.HTMLTagBuilder
}

func VListItemTitle(children ...h.HTMLComponent) (r *VListItemTitleBuilder) {
	r = &VListItemTitleBuilder{
		tag: h.Tag("v-list-item-title").Children(children...),
	}
	return
}

func (b *VListItemTitleBuilder) Tag(v string) (r *VListItemTitleBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VListItemTitleBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VListItemTitleBuilder) Attr(vs ...interface{}) (r *VListItemTitleBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VListItemTitleBuilder) Children(children ...h.HTMLComponent) (r *VListItemTitleBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VListItemTitleBuilder) AppendChildren(children ...h.HTMLComponent) (r *VListItemTitleBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VListItemTitleBuilder) PrependChildren(children ...h.HTMLComponent) (r *VListItemTitleBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VListItemTitleBuilder) Class(names ...string) (r *VListItemTitleBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VListItemTitleBuilder) ClassIf(name string, add bool) (r *VListItemTitleBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VListItemTitleBuilder) On(name string, value string) (r *VListItemTitleBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VListItemTitleBuilder) Bind(name string, value string) (r *VListItemTitleBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VListItemTitleBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
