package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSimpleCheckboxBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSimpleCheckbox(children ...h.HTMLComponent) (r *VSimpleCheckboxBuilder) {
	r = &VSimpleCheckboxBuilder{
		tag: h.Tag("v-simple-checkbox").Children(children...),
	}
	return
}

func (b *VSimpleCheckboxBuilder) Color(v string) (r *VSimpleCheckboxBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VSimpleCheckboxBuilder) Dark(v bool) (r *VSimpleCheckboxBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VSimpleCheckboxBuilder) Disabled(v bool) (r *VSimpleCheckboxBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VSimpleCheckboxBuilder) Indeterminate(v bool) (r *VSimpleCheckboxBuilder) {
	b.tag.Attr(":indeterminate", fmt.Sprint(v))
	return b
}

func (b *VSimpleCheckboxBuilder) IndeterminateIcon(v string) (r *VSimpleCheckboxBuilder) {
	b.tag.Attr("indeterminate-icon", v)
	return b
}

func (b *VSimpleCheckboxBuilder) Light(v bool) (r *VSimpleCheckboxBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VSimpleCheckboxBuilder) OffIcon(v string) (r *VSimpleCheckboxBuilder) {
	b.tag.Attr("off-icon", v)
	return b
}

func (b *VSimpleCheckboxBuilder) OnIcon(v string) (r *VSimpleCheckboxBuilder) {
	b.tag.Attr("on-icon", v)
	return b
}

func (b *VSimpleCheckboxBuilder) Ripple(v bool) (r *VSimpleCheckboxBuilder) {
	b.tag.Attr(":ripple", fmt.Sprint(v))
	return b
}

func (b *VSimpleCheckboxBuilder) Value(v bool) (r *VSimpleCheckboxBuilder) {
	b.tag.Attr(":value", fmt.Sprint(v))
	return b
}

func (b *VSimpleCheckboxBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSimpleCheckboxBuilder) Attr(vs ...interface{}) (r *VSimpleCheckboxBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSimpleCheckboxBuilder) Children(children ...h.HTMLComponent) (r *VSimpleCheckboxBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSimpleCheckboxBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSimpleCheckboxBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSimpleCheckboxBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSimpleCheckboxBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSimpleCheckboxBuilder) Class(names ...string) (r *VSimpleCheckboxBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSimpleCheckboxBuilder) ClassIf(name string, add bool) (r *VSimpleCheckboxBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSimpleCheckboxBuilder) On(name string, value string) (r *VSimpleCheckboxBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSimpleCheckboxBuilder) Bind(name string, value string) (r *VSimpleCheckboxBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSimpleCheckboxBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
