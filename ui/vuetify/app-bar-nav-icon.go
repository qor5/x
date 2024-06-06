package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VAppBarNavIconBuilder struct {
	tag *h.HTMLTagBuilder
}

func VAppBarNavIcon(children ...h.HTMLComponent) (r *VAppBarNavIconBuilder) {
	r = &VAppBarNavIconBuilder{
		tag: h.Tag("v-app-bar-nav-icon").Children(children...),
	}
	return
}

func (b *VAppBarNavIconBuilder) Symbol(v interface{}) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":symbol", h.JSONString(v))
	return b
}

func (b *VAppBarNavIconBuilder) Flat(v bool) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VAppBarNavIconBuilder) Active(v bool) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VAppBarNavIconBuilder) BaseColor(v string) (r *VAppBarNavIconBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VAppBarNavIconBuilder) PrependIcon(v interface{}) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VAppBarNavIconBuilder) AppendIcon(v interface{}) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VAppBarNavIconBuilder) Block(v bool) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":block", fmt.Sprint(v))
	return b
}

func (b *VAppBarNavIconBuilder) Readonly(v bool) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VAppBarNavIconBuilder) Slim(v bool) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":slim", fmt.Sprint(v))
	return b
}

func (b *VAppBarNavIconBuilder) Stacked(v bool) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":stacked", fmt.Sprint(v))
	return b
}

func (b *VAppBarNavIconBuilder) Ripple(v interface{}) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":ripple", h.JSONString(v))
	return b
}

func (b *VAppBarNavIconBuilder) Value(v interface{}) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VAppBarNavIconBuilder) Text(v string) (r *VAppBarNavIconBuilder) {
	b.tag.Attr("text", v)
	return b
}

func (b *VAppBarNavIconBuilder) Border(v interface{}) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":border", h.JSONString(v))
	return b
}

func (b *VAppBarNavIconBuilder) Density(v interface{}) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VAppBarNavIconBuilder) Height(v interface{}) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VAppBarNavIconBuilder) MaxHeight(v interface{}) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VAppBarNavIconBuilder) MaxWidth(v interface{}) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VAppBarNavIconBuilder) MinHeight(v interface{}) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VAppBarNavIconBuilder) MinWidth(v interface{}) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VAppBarNavIconBuilder) Width(v interface{}) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VAppBarNavIconBuilder) Elevation(v interface{}) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":elevation", h.JSONString(v))
	return b
}

func (b *VAppBarNavIconBuilder) Disabled(v bool) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VAppBarNavIconBuilder) SelectedClass(v string) (r *VAppBarNavIconBuilder) {
	b.tag.Attr("selected-class", v)
	return b
}

func (b *VAppBarNavIconBuilder) Loading(v interface{}) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":loading", h.JSONString(v))
	return b
}

func (b *VAppBarNavIconBuilder) Location(v interface{}) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VAppBarNavIconBuilder) Position(v interface{}) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":position", h.JSONString(v))
	return b
}

func (b *VAppBarNavIconBuilder) Rounded(v interface{}) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VAppBarNavIconBuilder) Tile(v bool) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VAppBarNavIconBuilder) Href(v string) (r *VAppBarNavIconBuilder) {
	b.tag.Attr("href", v)
	return b
}

func (b *VAppBarNavIconBuilder) Replace(v bool) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":replace", fmt.Sprint(v))
	return b
}

func (b *VAppBarNavIconBuilder) Exact(v bool) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":exact", fmt.Sprint(v))
	return b
}

func (b *VAppBarNavIconBuilder) To(v interface{}) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":to", h.JSONString(v))
	return b
}

func (b *VAppBarNavIconBuilder) Size(v interface{}) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":size", h.JSONString(v))
	return b
}

func (b *VAppBarNavIconBuilder) Tag(v string) (r *VAppBarNavIconBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VAppBarNavIconBuilder) Theme(v string) (r *VAppBarNavIconBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VAppBarNavIconBuilder) Color(v string) (r *VAppBarNavIconBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VAppBarNavIconBuilder) Variant(v interface{}) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VAppBarNavIconBuilder) Icon(v interface{}) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(":icon", h.JSONString(v))
	return b
}

func (b *VAppBarNavIconBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VAppBarNavIconBuilder) Attr(vs ...interface{}) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VAppBarNavIconBuilder) Children(children ...h.HTMLComponent) (r *VAppBarNavIconBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VAppBarNavIconBuilder) AppendChildren(children ...h.HTMLComponent) (r *VAppBarNavIconBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VAppBarNavIconBuilder) PrependChildren(children ...h.HTMLComponent) (r *VAppBarNavIconBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VAppBarNavIconBuilder) Class(names ...string) (r *VAppBarNavIconBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VAppBarNavIconBuilder) ClassIf(name string, add bool) (r *VAppBarNavIconBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VAppBarNavIconBuilder) On(name string, value string) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VAppBarNavIconBuilder) Bind(name string, value string) (r *VAppBarNavIconBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VAppBarNavIconBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
