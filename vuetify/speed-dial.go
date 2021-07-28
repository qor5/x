package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSpeedDialBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSpeedDial(children ...h.HTMLComponent) (r *VSpeedDialBuilder) {
	r = &VSpeedDialBuilder{
		tag: h.Tag("v-speed-dial").Children(children...),
	}
	return
}

func (b *VSpeedDialBuilder) Absolute(v bool) (r *VSpeedDialBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VSpeedDialBuilder) Bottom(v bool) (r *VSpeedDialBuilder) {
	b.tag.Attr(":bottom", fmt.Sprint(v))
	return b
}

func (b *VSpeedDialBuilder) Direction(v string) (r *VSpeedDialBuilder) {
	b.tag.Attr("direction", v)
	return b
}

func (b *VSpeedDialBuilder) Fixed(v bool) (r *VSpeedDialBuilder) {
	b.tag.Attr(":fixed", fmt.Sprint(v))
	return b
}

func (b *VSpeedDialBuilder) Left(v bool) (r *VSpeedDialBuilder) {
	b.tag.Attr(":left", fmt.Sprint(v))
	return b
}

func (b *VSpeedDialBuilder) Mode(v string) (r *VSpeedDialBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VSpeedDialBuilder) OpenOnHover(v bool) (r *VSpeedDialBuilder) {
	b.tag.Attr(":open-on-hover", fmt.Sprint(v))
	return b
}

func (b *VSpeedDialBuilder) Origin(v string) (r *VSpeedDialBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VSpeedDialBuilder) Right(v bool) (r *VSpeedDialBuilder) {
	b.tag.Attr(":right", fmt.Sprint(v))
	return b
}

func (b *VSpeedDialBuilder) Top(v bool) (r *VSpeedDialBuilder) {
	b.tag.Attr(":top", fmt.Sprint(v))
	return b
}

func (b *VSpeedDialBuilder) Transition(v string) (r *VSpeedDialBuilder) {
	b.tag.Attr("transition", v)
	return b
}

func (b *VSpeedDialBuilder) Value(v interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VSpeedDialBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSpeedDialBuilder) Attr(vs ...interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSpeedDialBuilder) Children(children ...h.HTMLComponent) (r *VSpeedDialBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSpeedDialBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSpeedDialBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSpeedDialBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSpeedDialBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSpeedDialBuilder) Class(names ...string) (r *VSpeedDialBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSpeedDialBuilder) ClassIf(name string, add bool) (r *VSpeedDialBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSpeedDialBuilder) On(name string, value string) (r *VSpeedDialBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSpeedDialBuilder) Bind(name string, value string) (r *VSpeedDialBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSpeedDialBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
