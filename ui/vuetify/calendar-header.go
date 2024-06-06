package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCalendarHeaderBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCalendarHeader(children ...h.HTMLComponent) (r *VCalendarHeaderBuilder) {
	r = &VCalendarHeaderBuilder{
		tag: h.Tag("v-calendar-header").Children(children...),
	}
	return
}

func (b *VCalendarHeaderBuilder) NextIcon(v string) (r *VCalendarHeaderBuilder) {
	b.tag.Attr("next-icon", v)
	return b
}

func (b *VCalendarHeaderBuilder) PrevIcon(v string) (r *VCalendarHeaderBuilder) {
	b.tag.Attr("prev-icon", v)
	return b
}

func (b *VCalendarHeaderBuilder) Title(v string) (r *VCalendarHeaderBuilder) {
	b.tag.Attr("title", v)
	return b
}

func (b *VCalendarHeaderBuilder) Text(v string) (r *VCalendarHeaderBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VCalendarHeaderBuilder) ViewMode(v interface{}) (r *VCalendarHeaderBuilder) {
	b.tag.Attr(":view-mode", h.JSONString(v))
	return b
}

func (b *VCalendarHeaderBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VCalendarHeaderBuilder) Attr(vs ...interface{}) (r *VCalendarHeaderBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VCalendarHeaderBuilder) Children(children ...h.HTMLComponent) (r *VCalendarHeaderBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VCalendarHeaderBuilder) AppendChildren(children ...h.HTMLComponent) (r *VCalendarHeaderBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VCalendarHeaderBuilder) PrependChildren(children ...h.HTMLComponent) (r *VCalendarHeaderBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VCalendarHeaderBuilder) Class(names ...string) (r *VCalendarHeaderBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCalendarHeaderBuilder) ClassIf(name string, add bool) (r *VCalendarHeaderBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCalendarHeaderBuilder) On(name string, value string) (r *VCalendarHeaderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCalendarHeaderBuilder) Bind(name string, value string) (r *VCalendarHeaderBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCalendarHeaderBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
