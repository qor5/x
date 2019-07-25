package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VBtnToggleBuilder struct {
	tag *h.HTMLTagBuilder
}

func VBtnToggle(children ...h.HTMLComponent) (r *VBtnToggleBuilder) {
	r = &VBtnToggleBuilder{
		tag: h.Tag("v-btn-toggle").Children(children...),
	}
	return
}
func (b *VBtnToggleBuilder) ActiveClass(v string) (r *VBtnToggleBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VBtnToggleBuilder) Color(v string) (r *VBtnToggleBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VBtnToggleBuilder) Dark(v bool) (r *VBtnToggleBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VBtnToggleBuilder) Light(v bool) (r *VBtnToggleBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VBtnToggleBuilder) Mandatory(v bool) (r *VBtnToggleBuilder) {
	b.tag.Attr(":mandatory", fmt.Sprint(v))
	return b
}

func (b *VBtnToggleBuilder) Max(v int) (r *VBtnToggleBuilder) {
	b.tag.Attr(":max", fmt.Sprint(v))
	return b
}

func (b *VBtnToggleBuilder) Multiple(v bool) (r *VBtnToggleBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VBtnToggleBuilder) Rounded(v bool) (r *VBtnToggleBuilder) {
	b.tag.Attr(":rounded", fmt.Sprint(v))
	return b
}

func (b *VBtnToggleBuilder) Value(v interface{}) (r *VBtnToggleBuilder) {
	b.tag.Attr(":value", v)
	return b
}

func (b *VBtnToggleBuilder) Children(children ...h.HTMLComponent) (r *VBtnToggleBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VBtnToggleBuilder) AppendChildren(children ...h.HTMLComponent) (r *VBtnToggleBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VBtnToggleBuilder) PrependChildren(children ...h.HTMLComponent) (r *VBtnToggleBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VBtnToggleBuilder) Class(names ...string) (r *VBtnToggleBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VBtnToggleBuilder) ClassIf(name string, add bool) (r *VBtnToggleBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VBtnToggleBuilder) On(name string, value string) (r *VBtnToggleBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VBtnToggleBuilder) Bind(name string, value string) (r *VBtnToggleBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VBtnToggleBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
