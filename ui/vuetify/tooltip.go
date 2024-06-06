package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTooltipBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTooltip(children ...h.HTMLComponent) (r *VTooltipBuilder) {
	r = &VTooltipBuilder{
		tag: h.Tag("v-tooltip").Children(children...),
	}
	return
}

func (b *VTooltipBuilder) Activator(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":activator", h.JSONString(v))
	return b
}

func (b *VTooltipBuilder) Id(v string) (r *VTooltipBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VTooltipBuilder) Text(v string) (r *VTooltipBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VTooltipBuilder) CloseOnBack(v bool) (r *VTooltipBuilder) {
	b.tag.Attr(":close-on-back", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) Contained(v bool) (r *VTooltipBuilder) {
	b.tag.Attr(":contained", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) ContentClass(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":content-class", h.JSONString(v))
	return b
}

func (b *VTooltipBuilder) ContentProps(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":content-props", h.JSONString(v))
	return b
}

func (b *VTooltipBuilder) Disabled(v bool) (r *VTooltipBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) Opacity(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":opacity", h.JSONString(v))
	return b
}

func (b *VTooltipBuilder) NoClickAnimation(v bool) (r *VTooltipBuilder) {
	b.tag.Attr(":no-click-animation", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) ModelValue(v bool) (r *VTooltipBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) Scrim(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":scrim", h.JSONString(v))
	return b
}

func (b *VTooltipBuilder) ZIndex(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":z-index", h.JSONString(v))
	return b
}

func (b *VTooltipBuilder) Target(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":target", h.JSONString(v))
	return b
}

func (b *VTooltipBuilder) ActivatorProps(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":activator-props", h.JSONString(v))
	return b
}

func (b *VTooltipBuilder) OpenOnClick(v bool) (r *VTooltipBuilder) {
	b.tag.Attr(":open-on-click", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) OpenOnHover(v bool) (r *VTooltipBuilder) {
	b.tag.Attr(":open-on-hover", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) OpenOnFocus(v bool) (r *VTooltipBuilder) {
	b.tag.Attr(":open-on-focus", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) CloseOnContentClick(v bool) (r *VTooltipBuilder) {
	b.tag.Attr(":close-on-content-click", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) CloseDelay(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":close-delay", h.JSONString(v))
	return b
}

func (b *VTooltipBuilder) OpenDelay(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":open-delay", h.JSONString(v))
	return b
}

func (b *VTooltipBuilder) Height(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VTooltipBuilder) MaxHeight(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VTooltipBuilder) MaxWidth(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VTooltipBuilder) MinHeight(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VTooltipBuilder) MinWidth(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VTooltipBuilder) Width(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VTooltipBuilder) Eager(v bool) (r *VTooltipBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VTooltipBuilder) LocationStrategy(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":location-strategy", h.JSONString(v))
	return b
}

func (b *VTooltipBuilder) Location(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VTooltipBuilder) Origin(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":origin", h.JSONString(v))
	return b
}

func (b *VTooltipBuilder) Offset(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":offset", h.JSONString(v))
	return b
}

func (b *VTooltipBuilder) ScrollStrategy(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":scroll-strategy", h.JSONString(v))
	return b
}

func (b *VTooltipBuilder) Theme(v string) (r *VTooltipBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VTooltipBuilder) Transition(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":transition", h.JSONString(v))
	return b
}

func (b *VTooltipBuilder) Attach(v interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(":attach", h.JSONString(v))
	return b
}

func (b *VTooltipBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VTooltipBuilder) Attr(vs ...interface{}) (r *VTooltipBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VTooltipBuilder) Children(children ...h.HTMLComponent) (r *VTooltipBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VTooltipBuilder) AppendChildren(children ...h.HTMLComponent) (r *VTooltipBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VTooltipBuilder) PrependChildren(children ...h.HTMLComponent) (r *VTooltipBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VTooltipBuilder) Class(names ...string) (r *VTooltipBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTooltipBuilder) ClassIf(name string, add bool) (r *VTooltipBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTooltipBuilder) On(name string, value string) (r *VTooltipBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTooltipBuilder) Bind(name string, value string) (r *VTooltipBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTooltipBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
