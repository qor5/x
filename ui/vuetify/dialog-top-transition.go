package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDialogTopTransitionBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDialogTopTransition(children ...h.HTMLComponent) (r *VDialogTopTransitionBuilder) {
	r = &VDialogTopTransitionBuilder{
		tag: h.Tag("v-dialog-top-transition").Children(children...),
	}
	return
}

func (b *VDialogTopTransitionBuilder) Disabled(v bool) (r *VDialogTopTransitionBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VDialogTopTransitionBuilder) Group(v bool) (r *VDialogTopTransitionBuilder) {
	b.tag.Attr(":group", fmt.Sprint(v))
	return b
}

func (b *VDialogTopTransitionBuilder) HideOnLeave(v bool) (r *VDialogTopTransitionBuilder) {
	b.tag.Attr(":hide-on-leave", fmt.Sprint(v))
	return b
}

func (b *VDialogTopTransitionBuilder) LeaveAbsolute(v bool) (r *VDialogTopTransitionBuilder) {
	b.tag.Attr(":leave-absolute", fmt.Sprint(v))
	return b
}

func (b *VDialogTopTransitionBuilder) Mode(v string) (r *VDialogTopTransitionBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VDialogTopTransitionBuilder) Origin(v string) (r *VDialogTopTransitionBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VDialogTopTransitionBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDialogTopTransitionBuilder) Attr(vs ...interface{}) (r *VDialogTopTransitionBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDialogTopTransitionBuilder) Children(children ...h.HTMLComponent) (r *VDialogTopTransitionBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDialogTopTransitionBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDialogTopTransitionBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDialogTopTransitionBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDialogTopTransitionBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDialogTopTransitionBuilder) Class(names ...string) (r *VDialogTopTransitionBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDialogTopTransitionBuilder) ClassIf(name string, add bool) (r *VDialogTopTransitionBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDialogTopTransitionBuilder) On(name string, value string) (r *VDialogTopTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDialogTopTransitionBuilder) Bind(name string, value string) (r *VDialogTopTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDialogTopTransitionBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
