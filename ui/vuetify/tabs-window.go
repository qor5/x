package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTabsWindowBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTabsWindow(children ...h.HTMLComponent) (r *VTabsWindowBuilder) {
	r = &VTabsWindowBuilder{
		tag: h.Tag("v-tabs-window").Children(children...),
	}
	return
}

func (b *VTabsWindowBuilder) Reverse(v bool) (r *VTabsWindowBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VTabsWindowBuilder) Direction(v interface{}) (r *VTabsWindowBuilder) {
	b.tag.Attr(":direction", h.JSONString(v))
	return b
}

func (b *VTabsWindowBuilder) ModelValue(v interface{}) (r *VTabsWindowBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VTabsWindowBuilder) Disabled(v bool) (r *VTabsWindowBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VTabsWindowBuilder) SelectedClass(v string) (r *VTabsWindowBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VTabsWindowBuilder) Tag(v string) (r *VTabsWindowBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VTabsWindowBuilder) Theme(v string) (r *VTabsWindowBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VTabsWindowBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VTabsWindowBuilder) Attr(vs ...interface{}) (r *VTabsWindowBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VTabsWindowBuilder) Children(children ...h.HTMLComponent) (r *VTabsWindowBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VTabsWindowBuilder) AppendChildren(children ...h.HTMLComponent) (r *VTabsWindowBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VTabsWindowBuilder) PrependChildren(children ...h.HTMLComponent) (r *VTabsWindowBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VTabsWindowBuilder) Class(names ...string) (r *VTabsWindowBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTabsWindowBuilder) ClassIf(name string, add bool) (r *VTabsWindowBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTabsWindowBuilder) On(name string, value string) (r *VTabsWindowBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTabsWindowBuilder) Bind(name string, value string) (r *VTabsWindowBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTabsWindowBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
