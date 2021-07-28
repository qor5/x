package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTabsItemsBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTabsItems(children ...h.HTMLComponent) (r *VTabsItemsBuilder) {
	r = &VTabsItemsBuilder{
		tag: h.Tag("v-tabs-items").Children(children...),
	}
	return
}

func (b *VTabsItemsBuilder) ActiveClass(v string) (r *VTabsItemsBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VTabsItemsBuilder) Continuous(v bool) (r *VTabsItemsBuilder) {
	b.tag.Attr(":continuous", fmt.Sprint(v))
	return b
}

func (b *VTabsItemsBuilder) Dark(v bool) (r *VTabsItemsBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VTabsItemsBuilder) Light(v bool) (r *VTabsItemsBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VTabsItemsBuilder) Mandatory(v bool) (r *VTabsItemsBuilder) {
	b.tag.Attr(":mandatory", fmt.Sprint(v))
	return b
}

func (b *VTabsItemsBuilder) Max(v int) (r *VTabsItemsBuilder) {
	b.tag.Attr(":max", fmt.Sprint(v))
	return b
}

func (b *VTabsItemsBuilder) Multiple(v bool) (r *VTabsItemsBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VTabsItemsBuilder) NextIcon(v bool) (r *VTabsItemsBuilder) {
	b.tag.Attr(":next-icon", fmt.Sprint(v))
	return b
}

func (b *VTabsItemsBuilder) PrevIcon(v bool) (r *VTabsItemsBuilder) {
	b.tag.Attr(":prev-icon", fmt.Sprint(v))
	return b
}

func (b *VTabsItemsBuilder) Reverse(v bool) (r *VTabsItemsBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VTabsItemsBuilder) ShowArrows(v bool) (r *VTabsItemsBuilder) {
	b.tag.Attr(":show-arrows", fmt.Sprint(v))
	return b
}

func (b *VTabsItemsBuilder) ShowArrowsOnHover(v bool) (r *VTabsItemsBuilder) {
	b.tag.Attr(":show-arrows-on-hover", fmt.Sprint(v))
	return b
}

func (b *VTabsItemsBuilder) Tag(v string) (r *VTabsItemsBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VTabsItemsBuilder) Touch(v interface{}) (r *VTabsItemsBuilder) {
	b.tag.Attr(":touch", h.JSONString(v))
	return b
}

func (b *VTabsItemsBuilder) Touchless(v bool) (r *VTabsItemsBuilder) {
	b.tag.Attr(":touchless", fmt.Sprint(v))
	return b
}

func (b *VTabsItemsBuilder) Value(v interface{}) (r *VTabsItemsBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VTabsItemsBuilder) Vertical(v bool) (r *VTabsItemsBuilder) {
	b.tag.Attr(":vertical", fmt.Sprint(v))
	return b
}

func (b *VTabsItemsBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VTabsItemsBuilder) Attr(vs ...interface{}) (r *VTabsItemsBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VTabsItemsBuilder) Children(children ...h.HTMLComponent) (r *VTabsItemsBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VTabsItemsBuilder) AppendChildren(children ...h.HTMLComponent) (r *VTabsItemsBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VTabsItemsBuilder) PrependChildren(children ...h.HTMLComponent) (r *VTabsItemsBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VTabsItemsBuilder) Class(names ...string) (r *VTabsItemsBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTabsItemsBuilder) ClassIf(name string, add bool) (r *VTabsItemsBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTabsItemsBuilder) On(name string, value string) (r *VTabsItemsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTabsItemsBuilder) Bind(name string, value string) (r *VTabsItemsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTabsItemsBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
