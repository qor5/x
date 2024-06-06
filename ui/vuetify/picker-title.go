package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VPickerTitleBuilder struct {
	tag *h.HTMLTagBuilder
}

func VPickerTitle(children ...h.HTMLComponent) (r *VPickerTitleBuilder) {
	r = &VPickerTitleBuilder{
		tag: h.Tag("v-picker-title").Children(children...),
	}
	return
}

func (b *VPickerTitleBuilder) Tag(v string) (r *VPickerTitleBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VPickerTitleBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VPickerTitleBuilder) Attr(vs ...interface{}) (r *VPickerTitleBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VPickerTitleBuilder) Children(children ...h.HTMLComponent) (r *VPickerTitleBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VPickerTitleBuilder) AppendChildren(children ...h.HTMLComponent) (r *VPickerTitleBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VPickerTitleBuilder) PrependChildren(children ...h.HTMLComponent) (r *VPickerTitleBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VPickerTitleBuilder) Class(names ...string) (r *VPickerTitleBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VPickerTitleBuilder) ClassIf(name string, add bool) (r *VPickerTitleBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VPickerTitleBuilder) On(name string, value string) (r *VPickerTitleBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VPickerTitleBuilder) Bind(name string, value string) (r *VPickerTitleBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VPickerTitleBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
