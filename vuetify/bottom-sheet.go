package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VBottomSheetBuilder struct {
	tag *h.HTMLTagBuilder
}

func VBottomSheet(children ...h.HTMLComponent) (r *VBottomSheetBuilder) {
	r = &VBottomSheetBuilder{
		tag: h.Tag("v-bottom-sheet").Children(children...),
	}
	return
}

func (b *VBottomSheetBuilder) Disabled(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) FullWidth(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":full-width", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) HideOverlay(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":hide-overlay", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) Inset(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":inset", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) Lazy(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":lazy", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) MaxWidth(v string) (r *VBottomSheetBuilder) {
	b.tag.Attr("max-width", v)
	return b
}

func (b *VBottomSheetBuilder) Persistent(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":persistent", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) Value(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":value", v)
	return b
}

func (b *VBottomSheetBuilder) Class(names ...string) (r *VBottomSheetBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VBottomSheetBuilder) ClassIf(name string, add bool) (r *VBottomSheetBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VBottomSheetBuilder) On(name string, value string) (r *VBottomSheetBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VBottomSheetBuilder) Bind(name string, value string) (r *VBottomSheetBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VBottomSheetBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
