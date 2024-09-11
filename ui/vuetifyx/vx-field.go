package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXFieldBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXField(children ...h.HTMLComponent) (r *VXFieldBuilder) {
	r = &VXFieldBuilder{
		tag: h.Tag("vx-field").Children(children...),
	}
	return
}

func (b *VXFieldBuilder) Label(v string) (r *VXFieldBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VXFieldBuilder) Id(v string) (r *VXFieldBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VXFieldBuilder) Focused(v bool) (r *VXFieldBuilder) {
	b.tag.Attr(":focused", fmt.Sprint(v))
	return b
}

func (b *VXFieldBuilder) Reverse(v bool) (r *VXFieldBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VXFieldBuilder) Flat(v bool) (r *VXFieldBuilder) {
	b.tag.Attr(":flat", fmt.Sprint(v))
	return b
}

func (b *VXFieldBuilder) AppendInnerIcon(v interface{}) (r *VXFieldBuilder) {
	b.tag.Attr(":append-inner-icon", h.JSONString(v))
	return b
}

func (b *VXFieldBuilder) BgColor(v string) (r *VXFieldBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VXFieldBuilder) Clearable(v bool) (r *VXFieldBuilder) {
	b.tag.Attr(":clearable", fmt.Sprint(v))
	return b
}

func (b *VXFieldBuilder) ClearIcon(v interface{}) (r *VXFieldBuilder) {
	b.tag.Attr(":clear-icon", h.JSONString(v))
	return b
}

func (b *VXFieldBuilder) Active(v bool) (r *VXFieldBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VXFieldBuilder) CenterAffix(v bool) (r *VXFieldBuilder) {
	b.tag.Attr(":center-affix", fmt.Sprint(v))
	return b
}

func (b *VXFieldBuilder) Color(v string) (r *VXFieldBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VXFieldBuilder) BaseColor(v string) (r *VXFieldBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VXFieldBuilder) Dirty(v bool) (r *VXFieldBuilder) {
	b.tag.Attr(":dirty", fmt.Sprint(v))
	return b
}

func (b *VXFieldBuilder) Disabled(v bool) (r *VXFieldBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VXFieldBuilder) Error(v bool) (r *VXFieldBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VXFieldBuilder) PersistentClear(v bool) (r *VXFieldBuilder) {
	b.tag.Attr(":persistent-clear", fmt.Sprint(v))
	return b
}

func (b *VXFieldBuilder) PrependInnerIcon(v interface{}) (r *VXFieldBuilder) {
	b.tag.Attr(":prepend-inner-icon", h.JSONString(v))
	return b
}

func (b *VXFieldBuilder) SingleLine(v bool) (r *VXFieldBuilder) {
	b.tag.Attr(":single-line", fmt.Sprint(v))
	return b
}

func (b *VXFieldBuilder) Variant(v interface{}) (r *VXFieldBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VXFieldBuilder) Loading(v interface{}) (r *VXFieldBuilder) {
	b.tag.Attr(":loading", h.JSONString(v))
	return b
}

func (b *VXFieldBuilder) Rounded(v interface{}) (r *VXFieldBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VXFieldBuilder) Tile(v bool) (r *VXFieldBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VXFieldBuilder) Theme(v string) (r *VXFieldBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VXFieldBuilder) ModelValue(v interface{}) (r *VXFieldBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VXFieldBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXFieldBuilder) Attr(vs ...interface{}) (r *VXFieldBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXFieldBuilder) Children(children ...h.HTMLComponent) (r *VXFieldBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXFieldBuilder) AppendChildren(children ...h.HTMLComponent) (r *VXFieldBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VXFieldBuilder) PrependChildren(children ...h.HTMLComponent) (r *VXFieldBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VXFieldBuilder) Class(names ...string) (r *VXFieldBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXFieldBuilder) ClassIf(name string, add bool) (r *VXFieldBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VXFieldBuilder) On(name string, value string) (r *VXFieldBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VXFieldBuilder) Bind(name string, value string) (r *VXFieldBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VXFieldBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
