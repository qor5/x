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

func (b *VBottomSheetBuilder) Inset(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":inset", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) Fullscreen(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":fullscreen", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) RetainFocus(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":retain-focus", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) Scrollable(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":scrollable", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) Absolute(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) CloseOnBack(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":close-on-back", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) Contained(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":contained", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) ContentClass(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":content-class", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) ContentProps(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":content-props", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) Disabled(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) Opacity(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":opacity", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) NoClickAnimation(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":no-click-animation", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) ModelValue(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) Persistent(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":persistent", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) Scrim(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":scrim", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) ZIndex(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":z-index", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) Target(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":target", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) ActivatorProps(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":activator-props", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) OpenOnClick(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":open-on-click", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) OpenOnHover(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":open-on-hover", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) OpenOnFocus(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":open-on-focus", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) CloseOnContentClick(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":close-on-content-click", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) CloseDelay(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":close-delay", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) OpenDelay(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":open-delay", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) Height(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) MaxHeight(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) MaxWidth(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) MinHeight(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) MinWidth(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) Width(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) Eager(v bool) (r *VBottomSheetBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VBottomSheetBuilder) LocationStrategy(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":location-strategy", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) Location(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) Origin(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":origin", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) Offset(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":offset", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) ScrollStrategy(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":scroll-strategy", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) Theme(v string) (r *VBottomSheetBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VBottomSheetBuilder) Transition(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":transition", h.JSONString(v))
	return b
}

func (b *VBottomSheetBuilder) Attach(v interface{}) (r *VBottomSheetBuilder) {
	b.tag.Attr(":attach", h.JSONString(v))
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
