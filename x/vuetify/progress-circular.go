package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VProgressCircularBuilder struct {
	tag *h.HTMLTagBuilder
}

func VProgressCircular(children ...h.HTMLComponent) (r *VProgressCircularBuilder) {
	r = &VProgressCircularBuilder{
		tag: h.Tag("v-progress-circular").Children(children...),
	}
	return
}

func (b *VProgressCircularBuilder) Button(v bool) (r *VProgressCircularBuilder) {
	b.tag.Attr(":button", fmt.Sprint(v))
	return b
}

func (b *VProgressCircularBuilder) Color(v string) (r *VProgressCircularBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VProgressCircularBuilder) Indeterminate(v bool) (r *VProgressCircularBuilder) {
	b.tag.Attr(":indeterminate", fmt.Sprint(v))
	return b
}

func (b *VProgressCircularBuilder) Rotate(v int) (r *VProgressCircularBuilder) {
	b.tag.Attr(":rotate", fmt.Sprint(v))
	return b
}

func (b *VProgressCircularBuilder) Size(v int) (r *VProgressCircularBuilder) {
	b.tag.Attr(":size", fmt.Sprint(v))
	return b
}

func (b *VProgressCircularBuilder) Value(v int) (r *VProgressCircularBuilder) {
	b.tag.Attr(":value", fmt.Sprint(v))
	return b
}

func (b *VProgressCircularBuilder) Width(v int) (r *VProgressCircularBuilder) {
	b.tag.Attr(":width", fmt.Sprint(v))
	return b
}

func (b *VProgressCircularBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VProgressCircularBuilder) Attr(vs ...interface{}) (r *VProgressCircularBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VProgressCircularBuilder) Children(children ...h.HTMLComponent) (r *VProgressCircularBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VProgressCircularBuilder) AppendChildren(children ...h.HTMLComponent) (r *VProgressCircularBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VProgressCircularBuilder) PrependChildren(children ...h.HTMLComponent) (r *VProgressCircularBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VProgressCircularBuilder) Class(names ...string) (r *VProgressCircularBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VProgressCircularBuilder) ClassIf(name string, add bool) (r *VProgressCircularBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VProgressCircularBuilder) On(name string, value string) (r *VProgressCircularBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VProgressCircularBuilder) Bind(name string, value string) (r *VProgressCircularBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VProgressCircularBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
