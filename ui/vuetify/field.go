package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VFieldBuilder struct {
	tag *h.HTMLTagBuilder
}

func VField(children ...h.HTMLComponent) (r *VFieldBuilder) {
	r = &VFieldBuilder{
		tag: h.Tag("v-field").Children(children...),
	}
	return
}

func (b *VFieldBuilder) Label(v string) (r *VFieldBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VFieldBuilder) Id(v string) (r *VFieldBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VFieldBuilder) Focused(v bool) (r *VFieldBuilder) {
	b.tag.Attr(":focused", fmt.Sprint(v))
	return b
}

func (b *VFieldBuilder) Reverse(v bool) (r *VFieldBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VFieldBuilder) Flat(v bool) (r *VFieldBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VFieldBuilder) AppendInnerIcon(v interface{}) (r *VFieldBuilder) {
	b.tag.Attr(":append-inner-icon", h.JSONString(v))
	return b
}

func (b *VFieldBuilder) BgColor(v string) (r *VFieldBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VFieldBuilder) Clearable(v bool) (r *VFieldBuilder) {
	b.tag.Attr(":clearable", fmt.Sprint(v))
	return b
}

func (b *VFieldBuilder) ClearIcon(v interface{}) (r *VFieldBuilder) {
	b.tag.Attr(":clear-icon", h.JSONString(v))
	return b
}

func (b *VFieldBuilder) Active(v bool) (r *VFieldBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VFieldBuilder) CenterAffix(v bool) (r *VFieldBuilder) {
	b.tag.Attr(":center-affix", fmt.Sprint(v))
	return b
}

func (b *VFieldBuilder) Color(v string) (r *VFieldBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VFieldBuilder) BaseColor(v string) (r *VFieldBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VFieldBuilder) Dirty(v bool) (r *VFieldBuilder) {
	b.tag.Attr(":dirty", fmt.Sprint(v))
	return b
}

func (b *VFieldBuilder) Disabled(v bool) (r *VFieldBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VFieldBuilder) Error(v bool) (r *VFieldBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VFieldBuilder) PersistentClear(v bool) (r *VFieldBuilder) {
	b.tag.Attr(":persistent-clear", fmt.Sprint(v))
	return b
}

func (b *VFieldBuilder) PrependInnerIcon(v interface{}) (r *VFieldBuilder) {
	b.tag.Attr(":prepend-inner-icon", h.JSONString(v))
	return b
}

func (b *VFieldBuilder) SingleLine(v bool) (r *VFieldBuilder) {
	b.tag.Attr(":single-line", fmt.Sprint(v))
	return b
}

func (b *VFieldBuilder) Variant(v interface{}) (r *VFieldBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VFieldBuilder) Loading(v interface{}) (r *VFieldBuilder) {
	b.tag.Attr(":loading", h.JSONString(v))
	return b
}

func (b *VFieldBuilder) Rounded(v interface{}) (r *VFieldBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VFieldBuilder) Tile(v bool) (r *VFieldBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VFieldBuilder) Theme(v string) (r *VFieldBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VFieldBuilder) ModelValue(v interface{}) (r *VFieldBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VFieldBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VFieldBuilder) Attr(vs ...interface{}) (r *VFieldBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VFieldBuilder) Children(children ...h.HTMLComponent) (r *VFieldBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VFieldBuilder) AppendChildren(children ...h.HTMLComponent) (r *VFieldBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VFieldBuilder) PrependChildren(children ...h.HTMLComponent) (r *VFieldBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VFieldBuilder) Class(names ...string) (r *VFieldBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VFieldBuilder) ClassIf(name string, add bool) (r *VFieldBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VFieldBuilder) On(name string, value string) (r *VFieldBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VFieldBuilder) Bind(name string, value string) (r *VFieldBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VFieldBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
