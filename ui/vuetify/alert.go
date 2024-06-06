package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VAlertBuilder struct {
	tag *h.HTMLTagBuilder
}

func VAlert(children ...h.HTMLComponent) (r *VAlertBuilder) {
	r = &VAlertBuilder{
		tag: h.Tag("v-alert").Children(children...),
	}
	return
}

func (b *VAlertBuilder) Title(v string) (r *VAlertBuilder) {
	b.tag.Attr("title", v)
	return b
}

func (b *VAlertBuilder) Text(v string) (r *VAlertBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VAlertBuilder) Border(v interface{}) (r *VAlertBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VAlertBuilder) BorderColor(v string) (r *VAlertBuilder) {
	b.tag.Attr("border-color", v)
	return b
}

func (b *VAlertBuilder) Closable(v bool) (r *VAlertBuilder) {
	b.tag.Attr(":closable", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) CloseIcon(v interface{}) (r *VAlertBuilder) {
	b.tag.Attr(":close-icon", h.JSONString(v))
	return b
}

func (b *VAlertBuilder) Type(v interface{}) (r *VAlertBuilder) {
	b.tag.Attr(":type", h.JSONString(v))
	return b
}

func (b *VAlertBuilder) CloseLabel(v string) (r *VAlertBuilder) {
	b.tag.Attr("close-label", v)
	return b
}

func (b *VAlertBuilder) Icon(v interface{}) (r *VAlertBuilder) {
	b.tag.Attr(":icon", h.JSONString(v))
	return b
}

func (b *VAlertBuilder) ModelValue(v bool) (r *VAlertBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) Prominent(v bool) (r *VAlertBuilder) {
	b.tag.Attr(":prominent", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) Density(v interface{}) (r *VAlertBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VAlertBuilder) Height(v interface{}) (r *VAlertBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VAlertBuilder) MaxHeight(v interface{}) (r *VAlertBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VAlertBuilder) MaxWidth(v interface{}) (r *VAlertBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VAlertBuilder) MinHeight(v interface{}) (r *VAlertBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VAlertBuilder) MinWidth(v interface{}) (r *VAlertBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VAlertBuilder) Width(v interface{}) (r *VAlertBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VAlertBuilder) Elevation(v interface{}) (r *VAlertBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VAlertBuilder) Location(v interface{}) (r *VAlertBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VAlertBuilder) Position(v interface{}) (r *VAlertBuilder) {
	b.tag.Attr(":position", h.JSONString(v))
	return b
}

func (b *VAlertBuilder) Rounded(v interface{}) (r *VAlertBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VAlertBuilder) Tile(v bool) (r *VAlertBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VAlertBuilder) Tag(v string) (r *VAlertBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VAlertBuilder) Theme(v string) (r *VAlertBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VAlertBuilder) Color(v string) (r *VAlertBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VAlertBuilder) Variant(v interface{}) (r *VAlertBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VAlertBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VAlertBuilder) Attr(vs ...interface{}) (r *VAlertBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VAlertBuilder) Children(children ...h.HTMLComponent) (r *VAlertBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VAlertBuilder) AppendChildren(children ...h.HTMLComponent) (r *VAlertBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VAlertBuilder) PrependChildren(children ...h.HTMLComponent) (r *VAlertBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VAlertBuilder) Class(names ...string) (r *VAlertBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VAlertBuilder) ClassIf(name string, add bool) (r *VAlertBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VAlertBuilder) On(name string, value string) (r *VAlertBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VAlertBuilder) Bind(name string, value string) (r *VAlertBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VAlertBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
