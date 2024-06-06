package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VBtnBuilder struct {
	tag *h.HTMLTagBuilder
}

func (b *VBtnBuilder) Symbol(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":symbol", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) Flat(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Active(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) BaseColor(v string) (r *VBtnBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VBtnBuilder) PrependIcon(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) AppendIcon(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) Block(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":block", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Readonly(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Slim(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":slim", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Stacked(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":stacked", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Ripple(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) Value(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) Text(v string) (r *VBtnBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VBtnBuilder) Border(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) Density(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) Height(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) MaxHeight(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) MaxWidth(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) MinHeight(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) MinWidth(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) Width(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) Elevation(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) Disabled(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) SelectedClass(v string) (r *VBtnBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VBtnBuilder) Loading(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":loading", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) Location(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) Position(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":position", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) Rounded(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) Tile(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Href(v string) (r *VBtnBuilder) {
	b.tag.Attr("href", v)
	return b
}

func (b *VBtnBuilder) Replace(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":replace", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) Exact(v bool) (r *VBtnBuilder) {
	b.tag.Attr(":exact", fmt.Sprint(v))
	return b
}

func (b *VBtnBuilder) To(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":to", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) Size(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":size", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) Tag(v string) (r *VBtnBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VBtnBuilder) Theme(v string) (r *VBtnBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VBtnBuilder) Color(v string) (r *VBtnBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VBtnBuilder) Variant(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) Icon(v interface{}) (r *VBtnBuilder) {
	b.tag.Attr(":icon", h.JSONString(v))
	return b
}

func (b *VBtnBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VBtnBuilder) Attr(vs ...interface{}) (r *VBtnBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VBtnBuilder) Children(children ...h.HTMLComponent) (r *VBtnBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VBtnBuilder) AppendChildren(children ...h.HTMLComponent) (r *VBtnBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VBtnBuilder) PrependChildren(children ...h.HTMLComponent) (r *VBtnBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VBtnBuilder) Class(names ...string) (r *VBtnBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VBtnBuilder) ClassIf(name string, add bool) (r *VBtnBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VBtnBuilder) On(name string, value string) (r *VBtnBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VBtnBuilder) Bind(name string, value string) (r *VBtnBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VBtnBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
