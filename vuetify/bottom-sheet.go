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

func (b *VBottomSheetBuilder) Activator(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":activator", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) Attach(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":attach", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) CloseDelay(v int) (r *VBottomSheetBuilder) {
	b.tag.Attr(":close-delay", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) ContentClass(v string) (r *VBottomSheetBuilder) {
	b.tag.Attr("content-class", v)
	return b
}

func (b *VBottomSheetBuilder) Dark(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) Disabled(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) Eager(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) Fullscreen(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":fullscreen", fmt.Sprint(v))
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

func (b *VBottomSheetBuilder) InternalActivator(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":internal-activator", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) Light(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) MaxWidth(v string) (r *VBottomSheetBuilder) {
	b.tag.Attr("max-width", v)
	return b
}

func (b *VBottomSheetBuilder) NoClickAnimation(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":no-click-animation", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) OpenDelay(v int) (r *VBottomSheetBuilder) {
	b.tag.Attr(":open-delay", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) OpenOnFocus(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":open-on-focus", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) OpenOnHover(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":open-on-hover", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) Origin(v string) (r *VBottomSheetBuilder) {
	b.tag.Attr("origin", v)
	return b
}

func (b *VBottomSheetBuilder) OverlayColor(v string) (r *VBottomSheetBuilder) {
	b.tag.Attr("overlay-color", v)
	return b
}

func (b *VBottomSheetBuilder) OverlayOpacity(v int) (r *VBottomSheetBuilder) {
	b.tag.Attr(":overlay-opacity", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) Persistent(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":persistent", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) RetainFocus(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":retain-focus", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) ReturnValue(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":return-value", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) Scrollable(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":scrollable", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) Transition(v string) (r *VBottomSheetBuilder) {
	b.tag.Attr("transition", v)
	return b
}

func (b *VBottomSheetBuilder) Value(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) Width(v string) (r *VBottomSheetBuilder) {
	b.tag.Attr("width", v)
	return b
}

func (b *VBottomSheetBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VBottomSheetBuilder) Attr(vs ...interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VBottomSheetBuilder) Children(children ...h.HTMLComponent) (r *VBottomSheetBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VBottomSheetBuilder) AppendChildren(children ...h.HTMLComponent) (r *VBottomSheetBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VBottomSheetBuilder) PrependChildren(children ...h.HTMLComponent) (r *VBottomSheetBuilder) {
	b.tag.PrependChildren(children...)
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
