package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VOverlayBuilder struct {
	tag *h.HTMLTagBuilder
}

func VOverlay(children ...h.HTMLComponent) (r *VOverlayBuilder) {
	r = &VOverlayBuilder{
		tag: h.Tag("v-overlay").Children(children...),
	}
	return
}

func (b *VOverlayBuilder) Activator(v interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(":activator", h.JSONString(v))
	return b
}

func (b *VOverlayBuilder) Absolute(v bool) (r *VOverlayBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VOverlayBuilder) CloseOnBack(v bool) (r *VOverlayBuilder) {
	b.tag.Attr(":close-on-back", fmt.Sprint(v))
	return b
}

func (b *VOverlayBuilder) Contained(v bool) (r *VOverlayBuilder) {
	b.tag.Attr(":contained", fmt.Sprint(v))
	return b
}

func (b *VOverlayBuilder) ContentClass(v interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(":content-class", h.JSONString(v))
	return b
}

func (b *VOverlayBuilder) ContentProps(v interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(":content-props", h.JSONString(v))
	return b
}

func (b *VOverlayBuilder) Disabled(v bool) (r *VOverlayBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VOverlayBuilder) Opacity(v interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(":opacity", h.JSONString(v))
	return b
}

func (b *VOverlayBuilder) NoClickAnimation(v bool) (r *VOverlayBuilder) {
	b.tag.Attr(":no-click-animation", fmt.Sprint(v))
	return b
}

func (b *VOverlayBuilder) ModelValue(v bool) (r *VOverlayBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VOverlayBuilder) Persistent(v bool) (r *VOverlayBuilder) {
	b.tag.Attr(":persistent", fmt.Sprint(v))
	return b
}

func (b *VOverlayBuilder) Scrim(v interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(":scrim", h.JSONString(v))
	return b
}

func (b *VOverlayBuilder) ZIndex(v interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(":z-index", h.JSONString(v))
	return b
}

func (b *VOverlayBuilder) Target(v interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(":target", h.JSONString(v))
	return b
}

func (b *VOverlayBuilder) ActivatorProps(v interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(":activator-props", h.JSONString(v))
	return b
}

func (b *VOverlayBuilder) OpenOnClick(v bool) (r *VOverlayBuilder) {
	b.tag.Attr(":open-on-click", fmt.Sprint(v))
	return b
}

func (b *VOverlayBuilder) OpenOnHover(v bool) (r *VOverlayBuilder) {
	b.tag.Attr(":open-on-hover", fmt.Sprint(v))
	return b
}

func (b *VOverlayBuilder) OpenOnFocus(v bool) (r *VOverlayBuilder) {
	b.tag.Attr(":open-on-focus", fmt.Sprint(v))
	return b
}

func (b *VOverlayBuilder) CloseOnContentClick(v bool) (r *VOverlayBuilder) {
	b.tag.Attr(":close-on-content-click", fmt.Sprint(v))
	return b
}

func (b *VOverlayBuilder) CloseDelay(v interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(":close-delay", h.JSONString(v))
	return b
}

func (b *VOverlayBuilder) OpenDelay(v interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(":open-delay", h.JSONString(v))
	return b
}

func (b *VOverlayBuilder) Height(v interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VOverlayBuilder) MaxHeight(v interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VOverlayBuilder) MaxWidth(v interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VOverlayBuilder) MinHeight(v interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VOverlayBuilder) MinWidth(v interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VOverlayBuilder) Width(v interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VOverlayBuilder) Eager(v bool) (r *VOverlayBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VOverlayBuilder) LocationStrategy(v interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(":location-strategy", h.JSONString(v))
	return b
}

func (b *VOverlayBuilder) Location(v interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VOverlayBuilder) Origin(v interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(":origin", h.JSONString(v))
	return b
}

func (b *VOverlayBuilder) Offset(v interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(":offset", h.JSONString(v))
	return b
}

func (b *VOverlayBuilder) ScrollStrategy(v interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(":scroll-strategy", h.JSONString(v))
	return b
}

func (b *VOverlayBuilder) Theme(v string) (r *VOverlayBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VOverlayBuilder) Transition(v interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(":transition", h.JSONString(v))
	return b
}

func (b *VOverlayBuilder) Attach(v interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(":attach", h.JSONString(v))
	return b
}

func (b *VOverlayBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VOverlayBuilder) Attr(vs ...interface{}) (r *VOverlayBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VOverlayBuilder) Children(children ...h.HTMLComponent) (r *VOverlayBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VOverlayBuilder) AppendChildren(children ...h.HTMLComponent) (r *VOverlayBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VOverlayBuilder) PrependChildren(children ...h.HTMLComponent) (r *VOverlayBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VOverlayBuilder) Class(names ...string) (r *VOverlayBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VOverlayBuilder) ClassIf(name string, add bool) (r *VOverlayBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VOverlayBuilder) On(name string, value string) (r *VOverlayBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VOverlayBuilder) Bind(name string, value string) (r *VOverlayBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VOverlayBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
