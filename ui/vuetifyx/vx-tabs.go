package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXTabsBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXTabs(children ...h.HTMLComponent) (r *VXTabsBuilder) {
	r = &VXTabsBuilder{
		tag: h.Tag("vx-tabs").Children(children...),
	}
	return
}

func (b *VXTabsBuilder) Symbol(v interface{}) (r *VXTabsBuilder) {
	b.tag.Attr(":symbol", h.JSONString(v))
	return b
}

func (b *VXTabsBuilder) AlignTabs(v interface{}) (r *VXTabsBuilder) {
	b.tag.Attr(":align-tabs", h.JSONString(v))
	return b
}

func (b *VXTabsBuilder) UnderlineBorder(v string) (r *VXTabsBuilder) {
	b.tag.Attr("underline-border", v)
	return b
}

func (b *VXTabsBuilder) Color(v string) (r *VXTabsBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VXTabsBuilder) FixedTabs(v bool) (r *VXTabsBuilder) {
	b.tag.Attr(":fixed-tabs", fmt.Sprint(v))
	return b
}

func (b *VXTabsBuilder) Pill(v bool) (r *VXTabsBuilder) {
	b.tag.Attr(":pill", fmt.Sprint(v))
	return b
}

func (b *VXTabsBuilder) Items(v interface{}) (r *VXTabsBuilder) {
	b.tag.Attr(":items", h.JSONString(v))
	return b
}

func (b *VXTabsBuilder) Stacked(v bool) (r *VXTabsBuilder) {
	b.tag.Attr(":stacked", fmt.Sprint(v))
	return b
}

func (b *VXTabsBuilder) BgColor(v string) (r *VXTabsBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VXTabsBuilder) Grow(v bool) (r *VXTabsBuilder) {
	b.tag.Attr(":grow", fmt.Sprint(v))
	return b
}

func (b *VXTabsBuilder) Height(v interface{}) (r *VXTabsBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VXTabsBuilder) HideSlider(v bool) (r *VXTabsBuilder) {
	b.tag.Attr(":hide-slider", fmt.Sprint(v))
	return b
}

func (b *VXTabsBuilder) SliderColor(v string) (r *VXTabsBuilder) {
	b.tag.Attr("slider-color", v)
	return b
}

func (b *VXTabsBuilder) CenterActive(v bool) (r *VXTabsBuilder) {
	b.tag.Attr(":center-active", fmt.Sprint(v))
	return b
}

func (b *VXTabsBuilder) Direction(v interface{}) (r *VXTabsBuilder) {
	b.tag.Attr(":direction", h.JSONString(v))
	return b
}

func (b *VXTabsBuilder) NextIcon(v interface{}) (r *VXTabsBuilder) {
	b.tag.Attr(":next-icon", h.JSONString(v))
	return b
}

func (b *VXTabsBuilder) PrevIcon(v interface{}) (r *VXTabsBuilder) {
	b.tag.Attr(":prev-icon", h.JSONString(v))
	return b
}

func (b *VXTabsBuilder) ShowArrows(v interface{}) (r *VXTabsBuilder) {
	b.tag.Attr(":show-arrows", h.JSONString(v))
	return b
}

func (b *VXTabsBuilder) Mobile(v bool) (r *VXTabsBuilder) {
	b.tag.Attr(":mobile", fmt.Sprint(v))
	return b
}

func (b *VXTabsBuilder) MobileBreakpoint(v interface{}) (r *VXTabsBuilder) {
	b.tag.Attr(":mobile-breakpoint", h.JSONString(v))
	return b
}

func (b *VXTabsBuilder) Tag(v string) (r *VXTabsBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VXTabsBuilder) ModelValue(v interface{}) (r *VXTabsBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VXTabsBuilder) Multiple(v bool) (r *VXTabsBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VXTabsBuilder) Max(v int) (r *VXTabsBuilder) {
	b.tag.Attr(":max", fmt.Sprint(v))
	return b
}

func (b *VXTabsBuilder) SelectedClass(v string) (r *VXTabsBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VXTabsBuilder) Disabled(v bool) (r *VXTabsBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VXTabsBuilder) Mandatory(v interface{}) (r *VXTabsBuilder) {
	b.tag.Attr(":mandatory", h.JSONString(v))
	return b
}

func (b *VXTabsBuilder) Density(v interface{}) (r *VXTabsBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VXTabsBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXTabsBuilder) Attr(vs ...interface{}) (r *VXTabsBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXTabsBuilder) Children(children ...h.HTMLComponent) (r *VXTabsBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXTabsBuilder) AppendChildren(children ...h.HTMLComponent) (r *VXTabsBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VXTabsBuilder) PrependChildren(children ...h.HTMLComponent) (r *VXTabsBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VXTabsBuilder) Class(names ...string) (r *VXTabsBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXTabsBuilder) ClassIf(name string, add bool) (r *VXTabsBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VXTabsBuilder) On(name string, value string) (r *VXTabsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VXTabsBuilder) Bind(name string, value string) (r *VXTabsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VXTabsBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
