package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTabsSliderBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTabsSlider(children ...h.HTMLComponent) (r *VTabsSliderBuilder) {
	r = &VTabsSliderBuilder{
		tag: h.Tag("v-tabs-slider").Children(children...),
	}
	return
}

func (b *VTabsSliderBuilder) Color(v string) (r *VTabsSliderBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VTabsSliderBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VTabsSliderBuilder) Attr(vs ...interface{}) (r *VTabsSliderBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VTabsSliderBuilder) Children(children ...h.HTMLComponent) (r *VTabsSliderBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VTabsSliderBuilder) AppendChildren(children ...h.HTMLComponent) (r *VTabsSliderBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VTabsSliderBuilder) PrependChildren(children ...h.HTMLComponent) (r *VTabsSliderBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VTabsSliderBuilder) Class(names ...string) (r *VTabsSliderBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTabsSliderBuilder) ClassIf(name string, add bool) (r *VTabsSliderBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTabsSliderBuilder) On(name string, value string) (r *VTabsSliderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTabsSliderBuilder) Bind(name string, value string) (r *VTabsSliderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTabsSliderBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
