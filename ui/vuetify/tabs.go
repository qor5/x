package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTabsBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTabs(children ...h.HTMLComponent) (r *VTabsBuilder) {
	r = &VTabsBuilder{
		tag: h.Tag("v-tabs").Children(children...),
	}
	return
}

func (b *VTabsBuilder) Symbol(v interface{}) (r *VTabsBuilder) {
	b.tag.Attr(":symbol", h.JSONString(v))
	return b
}

func (b *VTabsBuilder) AlignTabs(v interface{}) (r *VTabsBuilder) {
	b.tag.Attr(":align-tabs", h.JSONString(v))
	return b
}

func (b *VTabsBuilder) Color(v string) (r *VTabsBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VTabsBuilder) FixedTabs(v bool) (r *VTabsBuilder) {
	b.tag.Attr(":fixed-tabs", fmt.Sprint(v))
	return b
}

func (b *VTabsBuilder) Items(v interface{}) (r *VTabsBuilder) {
	b.tag.Attr(":items", h.JSONString(v))
	return b
}

func (b *VTabsBuilder) Stacked(v bool) (r *VTabsBuilder) {
	b.tag.Attr(":stacked", fmt.Sprint(v))
	return b
}

func (b *VTabsBuilder) BgColor(v string) (r *VTabsBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VTabsBuilder) Grow(v bool) (r *VTabsBuilder) {
	b.tag.Attr(":grow", fmt.Sprint(v))
	return b
}

func (b *VTabsBuilder) Height(v interface{}) (r *VTabsBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VTabsBuilder) HideSlider(v bool) (r *VTabsBuilder) {
	b.tag.Attr(":hide-slider", fmt.Sprint(v))
	return b
}

func (b *VTabsBuilder) SliderColor(v string) (r *VTabsBuilder) {
	b.tag.Attr("slider-color", v)
	return b
}

func (b *VTabsBuilder) CenterActive(v bool) (r *VTabsBuilder) {
	b.tag.Attr(":center-active", fmt.Sprint(v))
	return b
}

func (b *VTabsBuilder) Direction(v interface{}) (r *VTabsBuilder) {
	b.tag.Attr(":direction", h.JSONString(v))
	return b
}

func (b *VTabsBuilder) NextIcon(v interface{}) (r *VTabsBuilder) {
	b.tag.Attr(":next-icon", h.JSONString(v))
	return b
}

func (b *VTabsBuilder) PrevIcon(v interface{}) (r *VTabsBuilder) {
	b.tag.Attr(":prev-icon", h.JSONString(v))
	return b
}

func (b *VTabsBuilder) ShowArrows(v interface{}) (r *VTabsBuilder) {
	b.tag.Attr(":show-arrows", h.JSONString(v))
	return b
}

func (b *VTabsBuilder) Mobile(v bool) (r *VTabsBuilder) {
	b.tag.Attr(":mobile", fmt.Sprint(v))
	return b
}

func (b *VTabsBuilder) MobileBreakpoint(v interface{}) (r *VTabsBuilder) {
	b.tag.Attr(":mobile-breakpoint", h.JSONString(v))
	return b
}

func (b *VTabsBuilder) Tag(v string) (r *VTabsBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VTabsBuilder) ModelValue(v interface{}) (r *VTabsBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VTabsBuilder) Multiple(v bool) (r *VTabsBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VTabsBuilder) Max(v int) (r *VTabsBuilder) {
	b.tag.Attr(":max", fmt.Sprint(v))
	return b
}

func (b *VTabsBuilder) SelectedClass(v string) (r *VTabsBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VTabsBuilder) Disabled(v bool) (r *VTabsBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VTabsBuilder) Mandatory(v interface{}) (r *VTabsBuilder) {
	b.tag.Attr(":mandatory", h.JSONString(v))
	return b
}

func (b *VTabsBuilder) Density(v interface{}) (r *VTabsBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VTabsBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VTabsBuilder) Attr(vs ...interface{}) (r *VTabsBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VTabsBuilder) Children(children ...h.HTMLComponent) (r *VTabsBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VTabsBuilder) AppendChildren(children ...h.HTMLComponent) (r *VTabsBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VTabsBuilder) PrependChildren(children ...h.HTMLComponent) (r *VTabsBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VTabsBuilder) Class(names ...string) (r *VTabsBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTabsBuilder) ClassIf(name string, add bool) (r *VTabsBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTabsBuilder) On(name string, value string) (r *VTabsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTabsBuilder) Bind(name string, value string) (r *VTabsBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTabsBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
