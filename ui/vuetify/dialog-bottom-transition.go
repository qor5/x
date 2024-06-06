package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VDialogBottomTransitionBuilder struct {
	tag *h.HTMLTagBuilder
}

func VDialogBottomTransition(children ...h.HTMLComponent) (r *VDialogBottomTransitionBuilder) {
	r = &VDialogBottomTransitionBuilder{
		tag: h.Tag("v-dialog-bottom-transition").Children(children...),
	}
	return
}

func (b *VDialogBottomTransitionBuilder) Disabled(v bool) (r *VDialogBottomTransitionBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VDialogBottomTransitionBuilder) Group(v bool) (r *VDialogBottomTransitionBuilder) {
	b.tag.Attr(":group", fmt.Sprint(v))
	return b
}

func (b *VDialogBottomTransitionBuilder) HideOnLeave(v bool) (r *VDialogBottomTransitionBuilder) {
	b.tag.Attr(":hide-on-leave", fmt.Sprint(v))
	return b
}

func (b *VDialogBottomTransitionBuilder) LeaveAbsolute(v bool) (r *VDialogBottomTransitionBuilder) {
	b.tag.Attr(":leave-absolute", fmt.Sprint(v))
	return b
}

func (b *VDialogBottomTransitionBuilder) Mode(v string) (r *VDialogBottomTransitionBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VDialogBottomTransitionBuilder) Origin(v string) (r *VDialogBottomTransitionBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VDialogBottomTransitionBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VDialogBottomTransitionBuilder) Attr(vs ...interface{}) (r *VDialogBottomTransitionBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VDialogBottomTransitionBuilder) Children(children ...h.HTMLComponent) (r *VDialogBottomTransitionBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VDialogBottomTransitionBuilder) AppendChildren(children ...h.HTMLComponent) (r *VDialogBottomTransitionBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VDialogBottomTransitionBuilder) PrependChildren(children ...h.HTMLComponent) (r *VDialogBottomTransitionBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VDialogBottomTransitionBuilder) Class(names ...string) (r *VDialogBottomTransitionBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VDialogBottomTransitionBuilder) ClassIf(name string, add bool) (r *VDialogBottomTransitionBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VDialogBottomTransitionBuilder) On(name string, value string) (r *VDialogBottomTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VDialogBottomTransitionBuilder) Bind(name string, value string) (r *VDialogBottomTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VDialogBottomTransitionBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
