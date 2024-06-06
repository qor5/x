package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTabsWindowItemBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTabsWindowItem(children ...h.HTMLComponent) (r *VTabsWindowItemBuilder) {
	r = &VTabsWindowItemBuilder{
		tag: h.Tag("v-tabs-window-item").Children(children...),
	}
	return
}

func (b *VTabsWindowItemBuilder) ReverseTransition(v interface{}) (r *VTabsWindowItemBuilder) {
	b.tag.Attr(":reverse-transition", h.JSONString(v))
	return b
}

func (b *VTabsWindowItemBuilder) Transition(v interface{}) (r *VTabsWindowItemBuilder) {
	b.tag.Attr(":transition", h.JSONString(v))
	return b
}

func (b *VTabsWindowItemBuilder) Value(v interface{}) (r *VTabsWindowItemBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VTabsWindowItemBuilder) Disabled(v bool) (r *VTabsWindowItemBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VTabsWindowItemBuilder) SelectedClass(v string) (r *VTabsWindowItemBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VTabsWindowItemBuilder) Eager(v bool) (r *VTabsWindowItemBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VTabsWindowItemBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VTabsWindowItemBuilder) Attr(vs ...interface{}) (r *VTabsWindowItemBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VTabsWindowItemBuilder) Children(children ...h.HTMLComponent) (r *VTabsWindowItemBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VTabsWindowItemBuilder) AppendChildren(children ...h.HTMLComponent) (r *VTabsWindowItemBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VTabsWindowItemBuilder) PrependChildren(children ...h.HTMLComponent) (r *VTabsWindowItemBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VTabsWindowItemBuilder) Class(names ...string) (r *VTabsWindowItemBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTabsWindowItemBuilder) ClassIf(name string, add bool) (r *VTabsWindowItemBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTabsWindowItemBuilder) On(name string, value string) (r *VTabsWindowItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTabsWindowItemBuilder) Bind(name string, value string) (r *VTabsWindowItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTabsWindowItemBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
