package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VBottomSheetTransitionBuilder struct {
	tag *h.HTMLTagBuilder
}

func VBottomSheetTransition() (r *VBottomSheetTransitionBuilder) {
	r = &VBottomSheetTransitionBuilder{
		tag: h.Tag("v-bottom-sheet-transition"),
	}
	return
}

func (b *VBottomSheetTransitionBuilder) Group(v bool) (r *VBottomSheetTransitionBuilder) {
	b.tag.Attr(":group", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetTransitionBuilder) HideOnLeave(v bool) (r *VBottomSheetTransitionBuilder) {
	b.tag.Attr(":hide-on-leave", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetTransitionBuilder) LeaveAbsolute(v bool) (r *VBottomSheetTransitionBuilder) {
	b.tag.Attr(":leave-absolute", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetTransitionBuilder) Mode(v string) (r *VBottomSheetTransitionBuilder) {
	b.tag.Attr("mode", v)
	return b
}

func (b *VBottomSheetTransitionBuilder) Origin(v string) (r *VBottomSheetTransitionBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VBottomSheetTransitionBuilder) Class(names ...string) (r *VBottomSheetTransitionBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VBottomSheetTransitionBuilder) ClassIf(name string, add bool) (r *VBottomSheetTransitionBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VBottomSheetTransitionBuilder) On(name string, value string) (r *VBottomSheetTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VBottomSheetTransitionBuilder) Bind(name string, value string) (r *VBottomSheetTransitionBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VBottomSheetTransitionBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
