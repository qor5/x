package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTabBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTab(children ...h.HTMLComponent) (r *VTabBuilder) {
	r = &VTabBuilder{
		tag: h.Tag("v-tab").Children(children...),
	}
	return
}

func (b *VTabBuilder) Fixed(v bool) (r *VTabBuilder) {
	b.tag.Attr(":fixed", fmt.Sprint(v))
	return b
}

func (b *VTabBuilder) SliderColor(v string) (r *VTabBuilder) {
	b.tag.Attr("slider-color", v)
	return b
}

func (b *VTabBuilder) HideSlider(v bool) (r *VTabBuilder) {
	b.tag.Attr(":hide-slider", fmt.Sprint(v))
	return b
}

func (b *VTabBuilder) Direction(v interface{}) (r *VTabBuilder) {
	b.tag.Attr(":direction", h.JSONString(v))
	return b
}

func (b *VTabBuilder) BaseColor(v string) (r *VTabBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VTabBuilder) PrependIcon(v interface{}) (r *VTabBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VTabBuilder) AppendIcon(v interface{}) (r *VTabBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VTabBuilder) Readonly(v bool) (r *VTabBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VTabBuilder) Slim(v bool) (r *VTabBuilder) {
	b.tag.Attr(":slim", fmt.Sprint(v))
	return b
}

func (b *VTabBuilder) Stacked(v bool) (r *VTabBuilder) {
	b.tag.Attr(":stacked", fmt.Sprint(v))
	return b
}

func (b *VTabBuilder) Ripple(v interface{}) (r *VTabBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VTabBuilder) Value(v interface{}) (r *VTabBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VTabBuilder) Text(v string) (r *VTabBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VTabBuilder) Border(v interface{}) (r *VTabBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VTabBuilder) Density(v interface{}) (r *VTabBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VTabBuilder) Height(v interface{}) (r *VTabBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VTabBuilder) MaxHeight(v interface{}) (r *VTabBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VTabBuilder) MaxWidth(v interface{}) (r *VTabBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VTabBuilder) MinHeight(v interface{}) (r *VTabBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VTabBuilder) MinWidth(v interface{}) (r *VTabBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VTabBuilder) Width(v interface{}) (r *VTabBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VTabBuilder) Elevation(v interface{}) (r *VTabBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VTabBuilder) Disabled(v bool) (r *VTabBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VTabBuilder) SelectedClass(v string) (r *VTabBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VTabBuilder) Loading(v interface{}) (r *VTabBuilder) {
	b.tag.Attr(":loading", h.JSONString(v))
	return b
}

func (b *VTabBuilder) Rounded(v interface{}) (r *VTabBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VTabBuilder) Tile(v bool) (r *VTabBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VTabBuilder) Href(v string) (r *VTabBuilder) {
	b.tag.Attr("href", v)
	return b
}

func (b *VTabBuilder) Replace(v bool) (r *VTabBuilder) {
	b.tag.Attr(":replace", fmt.Sprint(v))
	return b
}

func (b *VTabBuilder) Exact(v bool) (r *VTabBuilder) {
	b.tag.Attr(":exact", fmt.Sprint(v))
	return b
}

func (b *VTabBuilder) To(v interface{}) (r *VTabBuilder) {
	b.tag.Attr(":to", h.JSONString(v))
	return b
}

func (b *VTabBuilder) Size(v interface{}) (r *VTabBuilder) {
	b.tag.Attr(":size", h.JSONString(v))
	return b
}

func (b *VTabBuilder) Tag(v string) (r *VTabBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VTabBuilder) Theme(v string) (r *VTabBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VTabBuilder) Color(v string) (r *VTabBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VTabBuilder) Variant(v interface{}) (r *VTabBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VTabBuilder) Icon(v interface{}) (r *VTabBuilder) {
	b.tag.Attr(":icon", h.JSONString(v))
	return b
}

func (b *VTabBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VTabBuilder) Attr(vs ...interface{}) (r *VTabBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VTabBuilder) Children(children ...h.HTMLComponent) (r *VTabBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VTabBuilder) AppendChildren(children ...h.HTMLComponent) (r *VTabBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VTabBuilder) PrependChildren(children ...h.HTMLComponent) (r *VTabBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VTabBuilder) Class(names ...string) (r *VTabBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTabBuilder) ClassIf(name string, add bool) (r *VTabBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTabBuilder) On(name string, value string) (r *VTabBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTabBuilder) Bind(name string, value string) (r *VTabBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTabBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
