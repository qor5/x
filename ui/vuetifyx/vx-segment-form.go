package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXSegmentFormBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXSegmentForm(text string) (r *VXSegmentFormBuilder) {
	r = &VXSegmentFormBuilder{
		tag: h.Tag("vx-segment-form").Children(h.Text(text)),
	}
	return
}

func (b *VXSegmentFormBuilder) Options(v interface{}) (r *VXSegmentFormBuilder) {
	b.tag.Attr(":options", h.JSONString(v))
	return b
}

func (b *VXSegmentFormBuilder) Readonly(v bool) (r *VXSegmentFormBuilder) {
	b.tag.Attr(":readonly", fmt.Sprintf("%v", v))
	return b
}

func (b *VXSegmentFormBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXSegmentFormBuilder) Attr(vs ...interface{}) (r *VXSegmentFormBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXSegmentFormBuilder) Children(children ...h.HTMLComponent) (r *VXSegmentFormBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXSegmentFormBuilder) AppendChildren(children ...h.HTMLComponent) (r *VXSegmentFormBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VXSegmentFormBuilder) PrependChildren(children ...h.HTMLComponent) (r *VXSegmentFormBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VXSegmentFormBuilder) Class(names ...string) (r *VXSegmentFormBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXSegmentFormBuilder) ClassIf(name string, add bool) (r *VXSegmentFormBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VXSegmentFormBuilder) On(name string, value string) (r *VXSegmentFormBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VXSegmentFormBuilder) Bind(name string, value string) (r *VXSegmentFormBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VXSegmentFormBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
