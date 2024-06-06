package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTimePickerBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTimePicker(children ...h.HTMLComponent) (r *VTimePickerBuilder) {
	r = &VTimePickerBuilder{
		tag: h.Tag("v-time-picker").Children(children...),
	}
	return
}

func (b *VTimePickerBuilder) Title(v string) (r *VTimePickerBuilder) {
	b.tag.Attr("title", v)
	return b
}

func (b *VTimePickerBuilder) AmpmInTitle(v bool) (r *VTimePickerBuilder) {
	b.tag.Attr(":ampm-in-title", fmt.Sprint(v))
	return b
}

func (b *VTimePickerBuilder) Disabled(v bool) (r *VTimePickerBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VTimePickerBuilder) Format(v interface{}) (r *VTimePickerBuilder) {
	b.tag.Attr(":format", h.JSONString(v))
	return b
}

func (b *VTimePickerBuilder) Max(v string) (r *VTimePickerBuilder) {
	b.tag.Attr("max", v)
	return b
}

func (b *VTimePickerBuilder) Min(v string) (r *VTimePickerBuilder) {
	b.tag.Attr("min", v)
	return b
}

func (b *VTimePickerBuilder) Readonly(v bool) (r *VTimePickerBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VTimePickerBuilder) Scrollable(v bool) (r *VTimePickerBuilder) {
	b.tag.Attr(":scrollable", fmt.Sprint(v))
	return b
}

func (b *VTimePickerBuilder) UseSeconds(v bool) (r *VTimePickerBuilder) {
	b.tag.Attr(":use-seconds", fmt.Sprint(v))
	return b
}

func (b *VTimePickerBuilder) BgColor(v string) (r *VTimePickerBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VTimePickerBuilder) HideHeader(v bool) (r *VTimePickerBuilder) {
	b.tag.Attr(":hide-header", fmt.Sprint(v))
	return b
}

func (b *VTimePickerBuilder) Color(v string) (r *VTimePickerBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VTimePickerBuilder) Border(v interface{}) (r *VTimePickerBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VTimePickerBuilder) Height(v interface{}) (r *VTimePickerBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VTimePickerBuilder) MaxHeight(v interface{}) (r *VTimePickerBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VTimePickerBuilder) MaxWidth(v interface{}) (r *VTimePickerBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VTimePickerBuilder) MinHeight(v interface{}) (r *VTimePickerBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VTimePickerBuilder) MinWidth(v interface{}) (r *VTimePickerBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VTimePickerBuilder) Width(v interface{}) (r *VTimePickerBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VTimePickerBuilder) Elevation(v interface{}) (r *VTimePickerBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VTimePickerBuilder) Location(v interface{}) (r *VTimePickerBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VTimePickerBuilder) Position(v interface{}) (r *VTimePickerBuilder) {
	b.tag.Attr(":position", h.JSONString(v))
	return b
}

func (b *VTimePickerBuilder) Rounded(v interface{}) (r *VTimePickerBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VTimePickerBuilder) Tile(v bool) (r *VTimePickerBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VTimePickerBuilder) Tag(v string) (r *VTimePickerBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VTimePickerBuilder) Theme(v string) (r *VTimePickerBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VTimePickerBuilder) AllowedHours(v interface{}) (r *VTimePickerBuilder) {
	b.tag.Attr(":allowed-hours", h.JSONString(v))
	return b
}

func (b *VTimePickerBuilder) AllowedMinutes(v interface{}) (r *VTimePickerBuilder) {
	b.tag.Attr(":allowed-minutes", h.JSONString(v))
	return b
}

func (b *VTimePickerBuilder) AllowedSeconds(v interface{}) (r *VTimePickerBuilder) {
	b.tag.Attr(":allowed-seconds", h.JSONString(v))
	return b
}

func (b *VTimePickerBuilder) ModelValue(v interface{}) (r *VTimePickerBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VTimePickerBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VTimePickerBuilder) Attr(vs ...interface{}) (r *VTimePickerBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VTimePickerBuilder) Children(children ...h.HTMLComponent) (r *VTimePickerBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VTimePickerBuilder) AppendChildren(children ...h.HTMLComponent) (r *VTimePickerBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VTimePickerBuilder) PrependChildren(children ...h.HTMLComponent) (r *VTimePickerBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VTimePickerBuilder) Class(names ...string) (r *VTimePickerBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTimePickerBuilder) ClassIf(name string, add bool) (r *VTimePickerBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTimePickerBuilder) On(name string, value string) (r *VTimePickerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTimePickerBuilder) Bind(name string, value string) (r *VTimePickerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTimePickerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
