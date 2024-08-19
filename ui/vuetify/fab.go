package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VFabBuilder struct {
	tag *h.HTMLTagBuilder
}

func VFab(children ...h.HTMLComponent) (r *VFabBuilder) {
	r = &VFabBuilder{
		tag: h.Tag("v-fab").Children(children...),
	}
	return
}

func (b *VFabBuilder) Symbol(v interface{}) (r *VFabBuilder) {
	b.tag.Attr(":symbol", h.JSONString(v))
	return b
}

func (b *VFabBuilder) Flat(v bool) (r *VFabBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VFabBuilder) App(v bool) (r *VFabBuilder) {
	b.tag.Attr(":app", fmt.Sprint(v))
	return b
}

func (b *VFabBuilder) Appear(v bool) (r *VFabBuilder) {
	b.tag.Attr(":appear", fmt.Sprint(v))
	return b
}

func (b *VFabBuilder) Extended(v bool) (r *VFabBuilder) {
	b.tag.Attr(":extended", fmt.Sprint(v))
	return b
}

func (b *VFabBuilder) Layout(v bool) (r *VFabBuilder) {
	b.tag.Attr(":layout", fmt.Sprint(v))
	return b
}

func (b *VFabBuilder) Offset(v bool) (r *VFabBuilder) {
	b.tag.Attr(":offset", fmt.Sprint(v))
	return b
}

func (b *VFabBuilder) ModelValue(v bool) (r *VFabBuilder) {
	b.tag.Attr(":model-value", fmt.Sprint(v))
	return b
}

func (b *VFabBuilder) Active(v bool) (r *VFabBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VFabBuilder) BaseColor(v string) (r *VFabBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VFabBuilder) PrependIcon(v interface{}) (r *VFabBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VFabBuilder) AppendIcon(v interface{}) (r *VFabBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VFabBuilder) Block(v bool) (r *VFabBuilder) {
	b.tag.Attr(":block", fmt.Sprint(v))
	return b
}

func (b *VFabBuilder) Readonly(v bool) (r *VFabBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VFabBuilder) Slim(v bool) (r *VFabBuilder) {
	b.tag.Attr(":slim", fmt.Sprint(v))
	return b
}

func (b *VFabBuilder) Stacked(v bool) (r *VFabBuilder) {
	b.tag.Attr(":stacked", fmt.Sprint(v))
	return b
}

func (b *VFabBuilder) Ripple(v interface{}) (r *VFabBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VFabBuilder) Value(v interface{}) (r *VFabBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VFabBuilder) Text(v string) (r *VFabBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VFabBuilder) Border(v interface{}) (r *VFabBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VFabBuilder) Density(v interface{}) (r *VFabBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VFabBuilder) Height(v interface{}) (r *VFabBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VFabBuilder) MaxHeight(v interface{}) (r *VFabBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VFabBuilder) MaxWidth(v interface{}) (r *VFabBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VFabBuilder) MinHeight(v interface{}) (r *VFabBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VFabBuilder) MinWidth(v interface{}) (r *VFabBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VFabBuilder) Width(v interface{}) (r *VFabBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VFabBuilder) Elevation(v interface{}) (r *VFabBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VFabBuilder) Disabled(v bool) (r *VFabBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VFabBuilder) SelectedClass(v string) (r *VFabBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VFabBuilder) Loading(v interface{}) (r *VFabBuilder) {
	b.tag.Attr(":loading", h.JSONString(v))
	return b
}

func (b *VFabBuilder) Location(v interface{}) (r *VFabBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VFabBuilder) Position(v interface{}) (r *VFabBuilder) {
	b.tag.Attr(":position", h.JSONString(v))
	return b
}

func (b *VFabBuilder) Absolute(v bool) (r *VFabBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VFabBuilder) Rounded(v interface{}) (r *VFabBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VFabBuilder) Tile(v bool) (r *VFabBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VFabBuilder) Href(v string) (r *VFabBuilder) {
	b.tag.Attr("href", v)
	return b
}

func (b *VFabBuilder) Replace(v bool) (r *VFabBuilder) {
	b.tag.Attr(":replace", fmt.Sprint(v))
	return b
}

func (b *VFabBuilder) Exact(v bool) (r *VFabBuilder) {
	b.tag.Attr(":exact", fmt.Sprint(v))
	return b
}

func (b *VFabBuilder) To(v interface{}) (r *VFabBuilder) {
	b.tag.Attr(":to", h.JSONString(v))
	return b
}

func (b *VFabBuilder) Size(v interface{}) (r *VFabBuilder) {
	b.tag.Attr(":size", h.JSONString(v))
	return b
}

func (b *VFabBuilder) Tag(v string) (r *VFabBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VFabBuilder) Theme(v string) (r *VFabBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VFabBuilder) Color(v string) (r *VFabBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VFabBuilder) Variant(v interface{}) (r *VFabBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VFabBuilder) Icon(v interface{}) (r *VFabBuilder) {
	b.tag.Attr(":icon", h.JSONString(v))
	return b
}

func (b *VFabBuilder) Name(v string) (r *VFabBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VFabBuilder) Order(v interface{}) (r *VFabBuilder) {
	b.tag.Attr(":order", h.JSONString(v))
	return b
}

func (b *VFabBuilder) Transition(v interface{}) (r *VFabBuilder) {
	b.tag.Attr(":transition", h.JSONString(v))
	return b
}

func (b *VFabBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VFabBuilder) Attr(vs ...interface{}) (r *VFabBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VFabBuilder) Children(children ...h.HTMLComponent) (r *VFabBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VFabBuilder) AppendChildren(children ...h.HTMLComponent) (r *VFabBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VFabBuilder) PrependChildren(children ...h.HTMLComponent) (r *VFabBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VFabBuilder) Class(names ...string) (r *VFabBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VFabBuilder) ClassIf(name string, add bool) (r *VFabBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VFabBuilder) On(name string, value string) (r *VFabBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VFabBuilder) Bind(name string, value string) (r *VFabBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VFabBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
