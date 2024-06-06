package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCarouselBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCarousel(children ...h.HTMLComponent) (r *VCarouselBuilder) {
	r = &VCarouselBuilder{
		tag: h.Tag("v-carousel").Children(children...),
	}
	return
}

func (b *VCarouselBuilder) Color(v string) (r *VCarouselBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VCarouselBuilder) Cycle(v bool) (r *VCarouselBuilder) {
	b.tag.Attr(":cycle", fmt.Sprint(v))
	return b
}

func (b *VCarouselBuilder) DelimiterIcon(v interface{}) (r *VCarouselBuilder) {
	b.tag.Attr(":delimiter-icon", h.JSONString(v))
	return b
}

func (b *VCarouselBuilder) Height(v interface{}) (r *VCarouselBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VCarouselBuilder) HideDelimiters(v bool) (r *VCarouselBuilder) {
	b.tag.Attr(":hide-delimiters", fmt.Sprint(v))
	return b
}

func (b *VCarouselBuilder) HideDelimiterBackground(v bool) (r *VCarouselBuilder) {
	b.tag.Attr(":hide-delimiter-background", fmt.Sprint(v))
	return b
}

func (b *VCarouselBuilder) Interval(v interface{}) (r *VCarouselBuilder) {
	b.tag.Attr(":interval", h.JSONString(v))
	return b
}

func (b *VCarouselBuilder) Progress(v interface{}) (r *VCarouselBuilder) {
	b.tag.Attr(":progress", h.JSONString(v))
	return b
}

func (b *VCarouselBuilder) Continuous(v bool) (r *VCarouselBuilder) {
	b.tag.Attr(":continuous", fmt.Sprint(v))
	return b
}

func (b *VCarouselBuilder) NextIcon(v interface{}) (r *VCarouselBuilder) {
	b.tag.Attr(":next-icon", h.JSONString(v))
	return b
}

func (b *VCarouselBuilder) PrevIcon(v interface{}) (r *VCarouselBuilder) {
	b.tag.Attr(":prev-icon", h.JSONString(v))
	return b
}

func (b *VCarouselBuilder) Reverse(v bool) (r *VCarouselBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VCarouselBuilder) ShowArrows(v interface{}) (r *VCarouselBuilder) {
	b.tag.Attr(":show-arrows", h.JSONString(v))
	return b
}

func (b *VCarouselBuilder) Touch(v interface{}) (r *VCarouselBuilder) {
	b.tag.Attr(":touch", h.JSONString(v))
	return b
}

func (b *VCarouselBuilder) Direction(v interface{}) (r *VCarouselBuilder) {
	b.tag.Attr(":direction", h.JSONString(v))
	return b
}

func (b *VCarouselBuilder) ModelValue(v interface{}) (r *VCarouselBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VCarouselBuilder) Disabled(v bool) (r *VCarouselBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VCarouselBuilder) SelectedClass(v string) (r *VCarouselBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VCarouselBuilder) Mandatory(v interface{}) (r *VCarouselBuilder) {
	b.tag.Attr(":mandatory", h.JSONString(v))
	return b
}

func (b *VCarouselBuilder) Tag(v string) (r *VCarouselBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VCarouselBuilder) Theme(v string) (r *VCarouselBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VCarouselBuilder) VerticalDelimiters(v interface{}) (r *VCarouselBuilder) {
	b.tag.Attr(":vertical-delimiters", h.JSONString(v))
	return b
}

func (b *VCarouselBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VCarouselBuilder) Attr(vs ...interface{}) (r *VCarouselBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VCarouselBuilder) Children(children ...h.HTMLComponent) (r *VCarouselBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VCarouselBuilder) AppendChildren(children ...h.HTMLComponent) (r *VCarouselBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VCarouselBuilder) PrependChildren(children ...h.HTMLComponent) (r *VCarouselBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VCarouselBuilder) Class(names ...string) (r *VCarouselBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCarouselBuilder) ClassIf(name string, add bool) (r *VCarouselBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCarouselBuilder) On(name string, value string) (r *VCarouselBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCarouselBuilder) Bind(name string, value string) (r *VCarouselBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCarouselBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
