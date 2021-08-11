package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VRadioBuilder struct {
	tag *h.HTMLTagBuilder
}

func VRadio(children ...h.HTMLComponent) (r *VRadioBuilder) {
	r = &VRadioBuilder{
		tag: h.Tag("v-radio").Children(children...),
	}
	return
}

func (b *VRadioBuilder) ActiveClass(v string) (r *VRadioBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VRadioBuilder) Color(v string) (r *VRadioBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VRadioBuilder) Dark(v bool) (r *VRadioBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VRadioBuilder) Disabled(v bool) (r *VRadioBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VRadioBuilder) Id(v string) (r *VRadioBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VRadioBuilder) Label(v string) (r *VRadioBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VRadioBuilder) Light(v bool) (r *VRadioBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VRadioBuilder) Name(v string) (r *VRadioBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VRadioBuilder) OffIcon(v string) (r *VRadioBuilder) {
	b.tag.Attr("off-icon", v)
	return b
}

func (b *VRadioBuilder) OnIcon(v string) (r *VRadioBuilder) {
	b.tag.Attr("on-icon", v)
	return b
}

func (b *VRadioBuilder) Readonly(v bool) (r *VRadioBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VRadioBuilder) Ripple(v interface{}) (r *VRadioBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VRadioBuilder) Value(v interface{}) (r *VRadioBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VRadioBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VRadioBuilder) Attr(vs ...interface{}) (r *VRadioBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VRadioBuilder) Children(children ...h.HTMLComponent) (r *VRadioBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VRadioBuilder) AppendChildren(children ...h.HTMLComponent) (r *VRadioBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VRadioBuilder) PrependChildren(children ...h.HTMLComponent) (r *VRadioBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VRadioBuilder) Class(names ...string) (r *VRadioBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VRadioBuilder) ClassIf(name string, add bool) (r *VRadioBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VRadioBuilder) On(name string, value string) (r *VRadioBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VRadioBuilder) Bind(name string, value string) (r *VRadioBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VRadioBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
