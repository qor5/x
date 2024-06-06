package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VProgressLinearBuilder struct {
	tag *h.HTMLTagBuilder
}

func VProgressLinear(children ...h.HTMLComponent) (r *VProgressLinearBuilder) {
	r = &VProgressLinearBuilder{
		tag: h.Tag("v-progress-linear").Children(children...),
	}
	return
}

func (b *VProgressLinearBuilder) Absolute(v bool) (r *VProgressLinearBuilder) {
	b.tag.Attr(":absolute", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) Active(v bool) (r *VProgressLinearBuilder) {
	b.tag.Attr(":active", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) BgColor(v string) (r *VProgressLinearBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VProgressLinearBuilder) BgOpacity(v interface{}) (r *VProgressLinearBuilder) {
	b.tag.Attr(":bg-opacity", h.JSONString(v))
	return b
}

func (b *VProgressLinearBuilder) BufferValue(v interface{}) (r *VProgressLinearBuilder) {
	b.tag.Attr(":buffer-value", h.JSONString(v))
	return b
}

func (b *VProgressLinearBuilder) BufferColor(v string) (r *VProgressLinearBuilder) {
	b.tag.Attr("buffer-color", v)
	return b
}

func (b *VProgressLinearBuilder) BufferOpacity(v interface{}) (r *VProgressLinearBuilder) {
	b.tag.Attr(":buffer-opacity", h.JSONString(v))
	return b
}

func (b *VProgressLinearBuilder) Clickable(v bool) (r *VProgressLinearBuilder) {
	b.tag.Attr(":clickable", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) Color(v string) (r *VProgressLinearBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VProgressLinearBuilder) Height(v interface{}) (r *VProgressLinearBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VProgressLinearBuilder) Indeterminate(v bool) (r *VProgressLinearBuilder) {
	b.tag.Attr(":indeterminate", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) Max(v interface{}) (r *VProgressLinearBuilder) {
	b.tag.Attr(":max", h.JSONString(v))
	return b
}

func (b *VProgressLinearBuilder) ModelValue(v interface{}) (r *VProgressLinearBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VProgressLinearBuilder) Opacity(v interface{}) (r *VProgressLinearBuilder) {
	b.tag.Attr(":opacity", h.JSONString(v))
	return b
}

func (b *VProgressLinearBuilder) Reverse(v bool) (r *VProgressLinearBuilder) {
	b.tag.Attr(":reverse", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) Stream(v bool) (r *VProgressLinearBuilder) {
	b.tag.Attr(":stream", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) Striped(v bool) (r *VProgressLinearBuilder) {
	b.tag.Attr(":striped", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) RoundedBar(v bool) (r *VProgressLinearBuilder) {
	b.tag.Attr(":rounded-bar", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) Location(v interface{}) (r *VProgressLinearBuilder) {
	b.tag.Attr(":location", h.JSONString(v))
	return b
}

func (b *VProgressLinearBuilder) Rounded(v interface{}) (r *VProgressLinearBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VProgressLinearBuilder) Tile(v bool) (r *VProgressLinearBuilder) {
	b.tag.Attr(":tile", fmt.Sprint(v))
	return b
}

func (b *VProgressLinearBuilder) Tag(v string) (r *VProgressLinearBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VProgressLinearBuilder) Theme(v string) (r *VProgressLinearBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VProgressLinearBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VProgressLinearBuilder) Attr(vs ...interface{}) (r *VProgressLinearBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VProgressLinearBuilder) Children(children ...h.HTMLComponent) (r *VProgressLinearBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VProgressLinearBuilder) AppendChildren(children ...h.HTMLComponent) (r *VProgressLinearBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VProgressLinearBuilder) PrependChildren(children ...h.HTMLComponent) (r *VProgressLinearBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VProgressLinearBuilder) Class(names ...string) (r *VProgressLinearBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VProgressLinearBuilder) ClassIf(name string, add bool) (r *VProgressLinearBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VProgressLinearBuilder) On(name string, value string) (r *VProgressLinearBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VProgressLinearBuilder) Bind(name string, value string) (r *VProgressLinearBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VProgressLinearBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
