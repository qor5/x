package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VMenuBuilder struct {
	tag *h.HTMLTagBuilder
}

func VMenu(children ...h.HTMLComponent) (r *VMenuBuilder) {
	r = &VMenuBuilder{
		tag: h.Tag("v-menu").Children(children...),
	}
	return
}

func (b *VMenuBuilder) Activator(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":activator", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) Id(v string) (r *VMenuBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VMenuBuilder) CloseOnBack(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":close-on-back", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) Contained(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":contained", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) ContentClass(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":content-class", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) ContentProps(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":content-props", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) Disabled(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) Opacity(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":opacity", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) NoClickAnimation(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":no-click-animation", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) ModelValue(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) Persistent(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":persistent", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) Scrim(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":scrim", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) ZIndex(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":z-index", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) Target(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":target", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) ActivatorProps(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":activator-props", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) OpenOnClick(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":open-on-click", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) OpenOnHover(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":open-on-hover", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) OpenOnFocus(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":open-on-focus", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) CloseOnContentClick(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":close-on-content-click", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) CloseDelay(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":close-delay", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) OpenDelay(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":open-delay", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) Height(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) MaxHeight(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) MaxWidth(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) MinHeight(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) MinWidth(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) Width(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) Eager(v bool) (r *VMenuBuilder) {
	b.tag.Attr(":eager", fmt.Sprint(v))
	return b
}

func (b *VMenuBuilder) LocationStrategy(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":location-strategy", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) Location(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) Origin(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":origin", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) Offset(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":offset", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) ScrollStrategy(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":scroll-strategy", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) Theme(v string) (r *VMenuBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VMenuBuilder) Transition(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":transition", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) Attach(v interface{}) (r *VMenuBuilder) {
	b.tag.Attr(":attach", h.JSONString(v))
	return b
}

func (b *VMenuBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VMenuBuilder) Attr(vs ...interface{}) (r *VMenuBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VMenuBuilder) Children(children ...h.HTMLComponent) (r *VMenuBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VMenuBuilder) AppendChildren(children ...h.HTMLComponent) (r *VMenuBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VMenuBuilder) PrependChildren(children ...h.HTMLComponent) (r *VMenuBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VMenuBuilder) Class(names ...string) (r *VMenuBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VMenuBuilder) ClassIf(name string, add bool) (r *VMenuBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VMenuBuilder) On(name string, value string) (r *VMenuBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VMenuBuilder) Bind(name string, value string) (r *VMenuBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VMenuBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
