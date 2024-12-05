package vuetifyx

import (
	"context"
	"fmt"

	"github.com/qor5/web/v3"
	h "github.com/theplant/htmlgo"
)

type VXChipBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXChip(children ...h.HTMLComponent) (r *VXChipBuilder) {
	r = &VXChipBuilder{
		tag: h.Tag("vx-btn").Children(children...),
	}
	return
}

func (b *VXChipBuilder) Presets(v string) (r *VXChipBuilder) {
	b.tag.Attr("presets", v)
	return b
}

func (b *VXChipBuilder) Round(v bool) (r *VXChipBuilder) {
	b.tag.Attr(":round", fmt.Sprint(v))
	return b
}

func (b *VXChipBuilder) Symbol(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":symbol", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Flat(v bool) (r *VXChipBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VXChipBuilder) Active(v bool) (r *VXChipBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VXChipBuilder) BaseColor(v string) (r *VXChipBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VXChipBuilder) PrependIcon(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) AppendIcon(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Block(v bool) (r *VXChipBuilder) {
	b.tag.Attr(":block", fmt.Sprint(v))
	return b
}

func (b *VXChipBuilder) Readonly(v bool) (r *VXChipBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VXChipBuilder) Slim(v bool) (r *VXChipBuilder) {
	b.tag.Attr(":slim", fmt.Sprint(v))
	return b
}

func (b *VXChipBuilder) Stacked(v bool) (r *VXChipBuilder) {
	b.tag.Attr(":stacked", fmt.Sprint(v))
	return b
}

func (b *VXChipBuilder) Ripple(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Value(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Text(v string) (r *VXChipBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VXChipBuilder) Border(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Density(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Height(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) MaxHeight(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) MaxWidth(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) MinHeight(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) MinWidth(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Width(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Elevation(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Disabled(v bool) (r *VXChipBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VXChipBuilder) SelectedClass(v string) (r *VXChipBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VXChipBuilder) Loading(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":loading", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Location(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Position(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":position", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Rounded(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Tile(v bool) (r *VXChipBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VXChipBuilder) Href(v string) (r *VXChipBuilder) {
	b.tag.Attr("href", v)
	return b
}

func (b *VXChipBuilder) Replace(v bool) (r *VXChipBuilder) {
	b.tag.Attr(":replace", fmt.Sprint(v))
	return b
}

func (b *VXChipBuilder) Exact(v bool) (r *VXChipBuilder) {
	b.tag.Attr(":exact", fmt.Sprint(v))
	return b
}

func (b *VXChipBuilder) To(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":to", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Size(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":size", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Tag(v string) (r *VXChipBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VXChipBuilder) Theme(v string) (r *VXChipBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VXChipBuilder) Color(v string) (r *VXChipBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VXChipBuilder) Variant(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) Icon(v interface{}) (r *VXChipBuilder) {
	b.tag.Attr(":icon", h.JSONString(v))
	return b
}

func (b *VXChipBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXChipBuilder) Attr(vs ...interface{}) (r *VXChipBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXChipBuilder) Children(children ...h.HTMLComponent) (r *VXChipBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXChipBuilder) AppendChildren(children ...h.HTMLComponent) (r *VXChipBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VXChipBuilder) PrependChildren(children ...h.HTMLComponent) (r *VXChipBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VXChipBuilder) Class(names ...string) (r *VXChipBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXChipBuilder) ClassIf(name string, add bool) (r *VXChipBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VXChipBuilder) On(name string, value string) (r *VXChipBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VXChipBuilder) Bind(name string, value string) (r *VXChipBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VXChipBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}

func (b *VXChipBuilder) OnClick(eventFuncId string) (r *VXChipBuilder) {
	b.tag.Attr("@click", web.POST().EventFunc(eventFuncId).Go())
	return b
}

func (b *VXChipBuilder) AttrIf(key, value interface{}, add bool) (r *VXChipBuilder) {
	b.tag.AttrIf(key, value, add)
	return b
}
