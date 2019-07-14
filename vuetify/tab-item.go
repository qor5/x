package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTabItemBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTabItem(children ...h.HTMLComponent) (r *VTabItemBuilder) {
	r = &VTabItemBuilder{
		tag: h.Tag("v-tab-item").Children(children...),
	}
	return
}

func (b *VTabItemBuilder) ActiveClass(v string) (r *VTabItemBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VTabItemBuilder) Disabled(v bool) (r *VTabItemBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VTabItemBuilder) Id(v string) (r *VTabItemBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VTabItemBuilder) Lazy(v bool) (r *VTabItemBuilder) {
	b.tag.Attr(":lazy", fmt.Sprint(v))
	return b
}

func (b *VTabItemBuilder) ReverseTransition(v bool) (r *VTabItemBuilder) {
	b.tag.Attr(":reverse-transition", fmt.Sprint(v))
	return b
}

func (b *VTabItemBuilder) Transition(v bool) (r *VTabItemBuilder) {
	b.tag.Attr(":transition", fmt.Sprint(v))
	return b
}

func (b *VTabItemBuilder) Value(v interface{}) (r *VTabItemBuilder) {
	b.tag.Attr(":value", v)
	return b
}

func (b *VTabItemBuilder) Class(names ...string) (r *VTabItemBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTabItemBuilder) ClassIf(name string, add bool) (r *VTabItemBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTabItemBuilder) On(name string, value string) (r *VTabItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTabItemBuilder) Bind(name string, value string) (r *VTabItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTabItemBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
