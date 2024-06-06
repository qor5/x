package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDialogTransitionBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDialogTransition(children ...h.HTMLComponent) (r *VDialogTransitionBuilder) {
	r = &VDialogTransitionBuilder{
		tag: h.Tag("v-dialog-transition").Children(children...),
	}
	return
}

func (b *VDialogTransitionBuilder) Target(v interface{}) (r *VDialogTransitionBuilder) {
	b.tag.Attr(":target", h.JSONString(v))
	return b
}

func (b *VDialogTransitionBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDialogTransitionBuilder) Attr(vs ...interface{}) (r *VDialogTransitionBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDialogTransitionBuilder) Children(children ...h.HTMLComponent) (r *VDialogTransitionBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDialogTransitionBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDialogTransitionBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDialogTransitionBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDialogTransitionBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDialogTransitionBuilder) Class(names ...string) (r *VDialogTransitionBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDialogTransitionBuilder) ClassIf(name string, add bool) (r *VDialogTransitionBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDialogTransitionBuilder) On(name string, value string) (r *VDialogTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDialogTransitionBuilder) Bind(name string, value string) (r *VDialogTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDialogTransitionBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
