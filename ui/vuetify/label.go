package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VLabelBuilder struct {
	tag *h.HTMLTagBuilder
}

func VLabel(children ...h.HTMLComponent) (r *VLabelBuilder) {
	r = &VLabelBuilder{
		tag: h.Tag("v-label").Children(children...),
	}
	return
}

func (b *VLabelBuilder) Text(v string) (r *VLabelBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VLabelBuilder) Theme(v string) (r *VLabelBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VLabelBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VLabelBuilder) Attr(vs ...interface{}) (r *VLabelBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VLabelBuilder) Children(children ...h.HTMLComponent) (r *VLabelBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VLabelBuilder) AppendChildren(children ...h.HTMLComponent) (r *VLabelBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VLabelBuilder) PrependChildren(children ...h.HTMLComponent) (r *VLabelBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VLabelBuilder) Class(names ...string) (r *VLabelBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VLabelBuilder) ClassIf(name string, add bool) (r *VLabelBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VLabelBuilder) On(name string, value string) (r *VLabelBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VLabelBuilder) Bind(name string, value string) (r *VLabelBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VLabelBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
