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

func (b *VTabsBuilder) ActiveClass(v string) (r *VTabsBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VTabsBuilder) AlignWithTitle(v bool) (r *VTabsBuilder) {
	b.tag.Attr(":align-with-title", fmt.Sprint(v))
	return b
}

func (b *VTabsBuilder) BackgroundColor(v string) (r *VTabsBuilder) {
	b.tag.Attr("background-color", v)
	return b
}

func (b *VTabsBuilder) CenterActive(v bool) (r *VTabsBuilder) {
	b.tag.Attr(":center-active", fmt.Sprint(v))
	return b
}

func (b *VTabsBuilder) Centered(v bool) (r *VTabsBuilder) {
	b.tag.Attr(":centered", fmt.Sprint(v))
	return b
}

func (b *VTabsBuilder) Color(v string) (r *VTabsBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VTabsBuilder) Dark(v bool) (r *VTabsBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VTabsBuilder) FixedTabs(v bool) (r *VTabsBuilder) {
	b.tag.Attr(":fixed-tabs", fmt.Sprint(v))
	return b
}

func (b *VTabsBuilder) Grow(v bool) (r *VTabsBuilder) {
	b.tag.Attr(":grow", fmt.Sprint(v))
	return b
}

func (b *VTabsBuilder) Height(v int) (r *VTabsBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VTabsBuilder) HideSlider(v bool) (r *VTabsBuilder) {
	b.tag.Attr(":hide-slider", fmt.Sprint(v))
	return b
}

func (b *VTabsBuilder) IconsAndText(v bool) (r *VTabsBuilder) {
	b.tag.Attr(":icons-and-text", fmt.Sprint(v))
	return b
}

func (b *VTabsBuilder) Light(v bool) (r *VTabsBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VTabsBuilder) MobileBreakpoint(v string) (r *VTabsBuilder) {
	b.tag.Attr("mobile-breakpoint", v)
	return b
}

func (b *VTabsBuilder) NextIcon(v string) (r *VTabsBuilder) {
	b.tag.Attr("next-icon", v)
	return b
}

func (b *VTabsBuilder) Optional(v bool) (r *VTabsBuilder) {
	b.tag.Attr(":optional", fmt.Sprint(v))
	return b
}

func (b *VTabsBuilder) PrevIcon(v string) (r *VTabsBuilder) {
	b.tag.Attr("prev-icon", v)
	return b
}

func (b *VTabsBuilder) Right(v bool) (r *VTabsBuilder) {
	b.tag.Attr(":right", fmt.Sprint(v))
	return b
}

func (b *VTabsBuilder) ShowArrows(v bool) (r *VTabsBuilder) {
	b.tag.Attr(":show-arrows", fmt.Sprint(v))
	return b
}

func (b *VTabsBuilder) SliderColor(v string) (r *VTabsBuilder) {
	b.tag.Attr("slider-color", v)
	return b
}

func (b *VTabsBuilder) SliderSize(v int) (r *VTabsBuilder) {
	b.tag.Attr(":slider-size", fmt.Sprint(v))
	return b
}

func (b *VTabsBuilder) Value(v interface{}) (r *VTabsBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VTabsBuilder) Vertical(v bool) (r *VTabsBuilder) {
	b.tag.Attr(":vertical", fmt.Sprint(v))
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
