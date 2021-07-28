package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VEditDialogBuilder struct {
	tag *h.HTMLTagBuilder
}

func VEditDialog(children ...h.HTMLComponent) (r *VEditDialogBuilder) {
	r = &VEditDialogBuilder{
		tag: h.Tag("v-edit-dialog").Children(children...),
	}
	return
}

func (b *VEditDialogBuilder) CancelText(v interface{}) (r *VEditDialogBuilder) {
	b.tag.Attr(":cancel-text", h.JSONString(v))
	return b
}

func (b *VEditDialogBuilder) Dark(v bool) (r *VEditDialogBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VEditDialogBuilder) Eager(v bool) (r *VEditDialogBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VEditDialogBuilder) Large(v bool) (r *VEditDialogBuilder) {
	b.tag.Attr(":large", fmt.Sprint(v))
	return b
}

func (b *VEditDialogBuilder) Light(v bool) (r *VEditDialogBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VEditDialogBuilder) Persistent(v bool) (r *VEditDialogBuilder) {
	b.tag.Attr(":persistent", fmt.Sprint(v))
	return b
}

func (b *VEditDialogBuilder) ReturnValue(v interface{}) (r *VEditDialogBuilder) {
	b.tag.Attr(":return-value", h.JSONString(v))
	return b
}

func (b *VEditDialogBuilder) SaveText(v interface{}) (r *VEditDialogBuilder) {
	b.tag.Attr(":save-text", h.JSONString(v))
	return b
}

func (b *VEditDialogBuilder) Transition(v string) (r *VEditDialogBuilder) {
	b.tag.Attr("transition", v)
	return b
}

func (b *VEditDialogBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VEditDialogBuilder) Attr(vs ...interface{}) (r *VEditDialogBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VEditDialogBuilder) Children(children ...h.HTMLComponent) (r *VEditDialogBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VEditDialogBuilder) AppendChildren(children ...h.HTMLComponent) (r *VEditDialogBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VEditDialogBuilder) PrependChildren(children ...h.HTMLComponent) (r *VEditDialogBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VEditDialogBuilder) Class(names ...string) (r *VEditDialogBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VEditDialogBuilder) ClassIf(name string, add bool) (r *VEditDialogBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VEditDialogBuilder) On(name string, value string) (r *VEditDialogBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VEditDialogBuilder) Bind(name string, value string) (r *VEditDialogBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VEditDialogBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
