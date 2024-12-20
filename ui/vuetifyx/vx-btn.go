package vuetifyx

import (
	"context"
	"fmt"

	"github.com/qor5/web/v3"
	h "github.com/theplant/htmlgo"
)

type VXBtnBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXBtn(name string) (r *VXBtnBuilder) {
	r = &VXBtnBuilder{
		tag: h.Tag("vx-btn").Text(name),
	}

	return
}

func (b *VXBtnBuilder) Presets(v string) (r *VXBtnBuilder) {
	b.tag.Attr("presets", v)
	return b
}

func (b *VXBtnBuilder) Symbol(v interface{}) (r *VXBtnBuilder) {
	b.tag.Attr(":symbol", h.JSONString(v))
	return b
}

func (b *VXBtnBuilder) Flat(v bool) (r *VXBtnBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VXBtnBuilder) Active(v bool) (r *VXBtnBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VXBtnBuilder) BaseColor(v string) (r *VXBtnBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VXBtnBuilder) PrependIcon(v interface{}) (r *VXBtnBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VXBtnBuilder) AppendIcon(v interface{}) (r *VXBtnBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VXBtnBuilder) Block(v bool) (r *VXBtnBuilder) {
	b.tag.Attr(":block", fmt.Sprint(v))
	return b
}

func (b *VXBtnBuilder) Readonly(v bool) (r *VXBtnBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VXBtnBuilder) Slim(v bool) (r *VXBtnBuilder) {
	b.tag.Attr(":slim", fmt.Sprint(v))
	return b
}

func (b *VXBtnBuilder) Stacked(v bool) (r *VXBtnBuilder) {
	b.tag.Attr(":stacked", fmt.Sprint(v))
	return b
}

func (b *VXBtnBuilder) Ripple(v interface{}) (r *VXBtnBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VXBtnBuilder) Value(v interface{}) (r *VXBtnBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VXBtnBuilder) Text(v string) (r *VXBtnBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VXBtnBuilder) Border(v interface{}) (r *VXBtnBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VXBtnBuilder) Density(v interface{}) (r *VXBtnBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VXBtnBuilder) Height(v interface{}) (r *VXBtnBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VXBtnBuilder) MaxHeight(v interface{}) (r *VXBtnBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VXBtnBuilder) MaxWidth(v interface{}) (r *VXBtnBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VXBtnBuilder) MinHeight(v interface{}) (r *VXBtnBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VXBtnBuilder) MinWidth(v interface{}) (r *VXBtnBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VXBtnBuilder) Width(v interface{}) (r *VXBtnBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VXBtnBuilder) Elevation(v interface{}) (r *VXBtnBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VXBtnBuilder) Disabled(v bool) (r *VXBtnBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VXBtnBuilder) SelectedClass(v string) (r *VXBtnBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VXBtnBuilder) Loading(v interface{}) (r *VXBtnBuilder) {
	b.tag.Attr(":loading", h.JSONString(v))
	return b
}

func (b *VXBtnBuilder) Location(v interface{}) (r *VXBtnBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VXBtnBuilder) Position(v interface{}) (r *VXBtnBuilder) {
	b.tag.Attr(":position", h.JSONString(v))
	return b
}

func (b *VXBtnBuilder) Rounded(v interface{}) (r *VXBtnBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VXBtnBuilder) Tile(v bool) (r *VXBtnBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VXBtnBuilder) Href(v string) (r *VXBtnBuilder) {
	b.tag.Attr("href", v)
	return b
}

func (b *VXBtnBuilder) Replace(v bool) (r *VXBtnBuilder) {
	b.tag.Attr(":replace", fmt.Sprint(v))
	return b
}

func (b *VXBtnBuilder) Exact(v bool) (r *VXBtnBuilder) {
	b.tag.Attr(":exact", fmt.Sprint(v))
	return b
}

func (b *VXBtnBuilder) To(v interface{}) (r *VXBtnBuilder) {
	b.tag.Attr(":to", h.JSONString(v))
	return b
}

func (b *VXBtnBuilder) Size(v interface{}) (r *VXBtnBuilder) {
	b.tag.Attr(":size", h.JSONString(v))
	return b
}

func (b *VXBtnBuilder) Tag(v string) (r *VXBtnBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VXBtnBuilder) Theme(v string) (r *VXBtnBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VXBtnBuilder) Color(v string) (r *VXBtnBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VXBtnBuilder) Variant(v interface{}) (r *VXBtnBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VXBtnBuilder) Icon(v interface{}) (r *VXBtnBuilder) {
	b.tag.Attr(":icon", h.JSONString(v))
	return b
}

func (b *VXBtnBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXBtnBuilder) Attr(vs ...interface{}) (r *VXBtnBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXBtnBuilder) Children(children ...h.HTMLComponent) (r *VXBtnBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXBtnBuilder) AppendChildren(children ...h.HTMLComponent) (r *VXBtnBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VXBtnBuilder) PrependChildren(children ...h.HTMLComponent) (r *VXBtnBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VXBtnBuilder) Class(names ...string) (r *VXBtnBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXBtnBuilder) ClassIf(name string, add bool) (r *VXBtnBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VXBtnBuilder) On(name string, value string) (r *VXBtnBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VXBtnBuilder) Bind(name string, value string) (r *VXBtnBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VXBtnBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}

func (b *VXBtnBuilder) OnClick(eventFuncId string) (r *VXBtnBuilder) {
	b.tag.Attr("@click", web.POST().EventFunc(eventFuncId).Go())
	return b
}

func (b *VXBtnBuilder) AttrIf(key, value interface{}, add bool) (r *VXBtnBuilder) {
	b.tag.AttrIf(key, value, add)
	return b
}
