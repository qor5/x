package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXDraggableBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXDraggable(children ...h.HTMLComponent) (r *VXDraggableBuilder) {
	r = &VXDraggableBuilder{
		tag: h.Tag("vx-draggable").Children(children...),
	}
	r.Handle(".handle")
	return
}

func (b *VXDraggableBuilder) ItemKey(v string) (r *VXDraggableBuilder) {
	b.tag.SetAttr(":item-key", h.JSONString(v))
	return b
}

func (b *VXDraggableBuilder) Handle(v string) (r *VXDraggableBuilder) {
	b.tag.SetAttr(":handle", h.JSONString(v))
	return b
}

func (b *VXDraggableBuilder) Animation(v int) (r *VXDraggableBuilder) {
	b.tag.SetAttr(":animation", h.JSONString(v))
	return b
}

func (b *VXDraggableBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXDraggableBuilder) Attr(vs ...interface{}) (r *VXDraggableBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXDraggableBuilder) Children(children ...h.HTMLComponent) (r *VXDraggableBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXDraggableBuilder) AppendChildren(children ...h.HTMLComponent) (r *VXDraggableBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VXDraggableBuilder) PrependChildren(children ...h.HTMLComponent) (r *VXDraggableBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VXDraggableBuilder) Class(names ...string) (r *VXDraggableBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXDraggableBuilder) ClassIf(name string, add bool) (r *VXDraggableBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VXDraggableBuilder) On(name string, value string) (r *VXDraggableBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VXDraggableBuilder) Bind(name string, value string) (r *VXDraggableBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VXDraggableBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
