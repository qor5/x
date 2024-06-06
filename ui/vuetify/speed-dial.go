package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VSpeedDialBuilder struct {
	tag *h.HTMLTagBuilder
}

func VSpeedDial(children ...h.HTMLComponent) (r *VSpeedDialBuilder) {
	r = &VSpeedDialBuilder{
		tag: h.Tag("v-speed-dial").Children(children...),
	}
	return
}

func (b *VSpeedDialBuilder) Activator(v interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(":activator", h.JSONString(v))
	return b
}

func (b *VSpeedDialBuilder) Id(v string) (r *VSpeedDialBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VSpeedDialBuilder) CloseOnBack(v bool) (r *VSpeedDialBuilder) {
	b.tag.Attr(":close-on-back", fmt.Sprint(v))
	return b
}

func (b *VSpeedDialBuilder) Contained(v bool) (r *VSpeedDialBuilder) {
	b.tag.Attr(":contained", fmt.Sprint(v))
	return b
}

func (b *VSpeedDialBuilder) ContentClass(v interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(":content-class", h.JSONString(v))
	return b
}

func (b *VSpeedDialBuilder) ContentProps(v interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(":content-props", h.JSONString(v))
	return b
}

func (b *VSpeedDialBuilder) Disabled(v bool) (r *VSpeedDialBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VSpeedDialBuilder) Opacity(v interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(":opacity", h.JSONString(v))
	return b
}

func (b *VSpeedDialBuilder) NoClickAnimation(v bool) (r *VSpeedDialBuilder) {
	b.tag.Attr(":no-click-animation", fmt.Sprint(v))
	return b
}

func (b *VSpeedDialBuilder) ModelValue(v bool) (r *VSpeedDialBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VSpeedDialBuilder) Persistent(v bool) (r *VSpeedDialBuilder) {
	b.tag.Attr(":persistent", fmt.Sprint(v))
	return b
}

func (b *VSpeedDialBuilder) Scrim(v interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(":scrim", h.JSONString(v))
	return b
}

func (b *VSpeedDialBuilder) ZIndex(v interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(":z-index", h.JSONString(v))
	return b
}

func (b *VSpeedDialBuilder) Target(v interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(":target", h.JSONString(v))
	return b
}

func (b *VSpeedDialBuilder) ActivatorProps(v interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(":activator-props", h.JSONString(v))
	return b
}

func (b *VSpeedDialBuilder) OpenOnClick(v bool) (r *VSpeedDialBuilder) {
	b.tag.Attr(":open-on-click", fmt.Sprint(v))
	return b
}

func (b *VSpeedDialBuilder) OpenOnHover(v bool) (r *VSpeedDialBuilder) {
	b.tag.Attr(":open-on-hover", fmt.Sprint(v))
	return b
}

func (b *VSpeedDialBuilder) OpenOnFocus(v bool) (r *VSpeedDialBuilder) {
	b.tag.Attr(":open-on-focus", fmt.Sprint(v))
	return b
}

func (b *VSpeedDialBuilder) CloseOnContentClick(v bool) (r *VSpeedDialBuilder) {
	b.tag.Attr(":close-on-content-click", fmt.Sprint(v))
	return b
}

func (b *VSpeedDialBuilder) CloseDelay(v interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(":close-delay", h.JSONString(v))
	return b
}

func (b *VSpeedDialBuilder) OpenDelay(v interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(":open-delay", h.JSONString(v))
	return b
}

func (b *VSpeedDialBuilder) Height(v interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VSpeedDialBuilder) MaxHeight(v interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VSpeedDialBuilder) MaxWidth(v interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VSpeedDialBuilder) MinHeight(v interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VSpeedDialBuilder) MinWidth(v interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VSpeedDialBuilder) Width(v interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VSpeedDialBuilder) Eager(v bool) (r *VSpeedDialBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VSpeedDialBuilder) LocationStrategy(v interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(":location-strategy", h.JSONString(v))
	return b
}

func (b *VSpeedDialBuilder) Location(v interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VSpeedDialBuilder) Origin(v interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(":origin", h.JSONString(v))
	return b
}

func (b *VSpeedDialBuilder) Offset(v interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(":offset", h.JSONString(v))
	return b
}

func (b *VSpeedDialBuilder) ScrollStrategy(v interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(":scroll-strategy", h.JSONString(v))
	return b
}

func (b *VSpeedDialBuilder) Theme(v string) (r *VSpeedDialBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VSpeedDialBuilder) Transition(v interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(":transition", h.JSONString(v))
	return b
}

func (b *VSpeedDialBuilder) Attach(v interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(":attach", h.JSONString(v))
	return b
}

func (b *VSpeedDialBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VSpeedDialBuilder) Attr(vs ...interface{}) (r *VSpeedDialBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VSpeedDialBuilder) Children(children ...h.HTMLComponent) (r *VSpeedDialBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VSpeedDialBuilder) AppendChildren(children ...h.HTMLComponent) (r *VSpeedDialBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VSpeedDialBuilder) PrependChildren(children ...h.HTMLComponent) (r *VSpeedDialBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VSpeedDialBuilder) Class(names ...string) (r *VSpeedDialBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VSpeedDialBuilder) ClassIf(name string, add bool) (r *VSpeedDialBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VSpeedDialBuilder) On(name string, value string) (r *VSpeedDialBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VSpeedDialBuilder) Bind(name string, value string) (r *VSpeedDialBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VSpeedDialBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
