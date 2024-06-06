package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VOtpInputBuilder struct {
	tag *h.HTMLTagBuilder
}

func VOtpInput(children ...h.HTMLComponent) (r *VOtpInputBuilder) {
	r = &VOtpInputBuilder{
		tag: h.Tag("v-otp-input").Children(children...),
	}
	return
}

func (b *VOtpInputBuilder) Length(v interface{}) (r *VOtpInputBuilder) {
	b.tag.Attr(":length", h.JSONString(v))
	return b
}

func (b *VOtpInputBuilder) Autofocus(v bool) (r *VOtpInputBuilder) {
	b.tag.Attr(":autofocus", fmt.Sprint(v))
	return b
}

func (b *VOtpInputBuilder) Divider(v string) (r *VOtpInputBuilder) {
	b.tag.Attr("divider", v)
	return b
}

func (b *VOtpInputBuilder) FocusAll(v bool) (r *VOtpInputBuilder) {
	b.tag.Attr(":focus-all", fmt.Sprint(v))
	return b
}

func (b *VOtpInputBuilder) Label(v string) (r *VOtpInputBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VOtpInputBuilder) Type(v interface{}) (r *VOtpInputBuilder) {
	b.tag.Attr(":type", h.JSONString(v))
	return b
}

func (b *VOtpInputBuilder) ModelValue(v interface{}) (r *VOtpInputBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VOtpInputBuilder) Placeholder(v string) (r *VOtpInputBuilder) {
	b.tag.Attr("placeholder", v)
	return b
}

func (b *VOtpInputBuilder) Height(v interface{}) (r *VOtpInputBuilder) {
	b.tag.Attr(":height", h.JSONString(v))
	return b
}

func (b *VOtpInputBuilder) MaxHeight(v interface{}) (r *VOtpInputBuilder) {
	b.tag.Attr(":max-height", h.JSONString(v))
	return b
}

func (b *VOtpInputBuilder) MaxWidth(v interface{}) (r *VOtpInputBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VOtpInputBuilder) MinHeight(v interface{}) (r *VOtpInputBuilder) {
	b.tag.Attr(":min-height", h.JSONString(v))
	return b
}

func (b *VOtpInputBuilder) MinWidth(v interface{}) (r *VOtpInputBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VOtpInputBuilder) Width(v interface{}) (r *VOtpInputBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VOtpInputBuilder) Focused(v bool) (r *VOtpInputBuilder) {
	b.tag.Attr(":focused", fmt.Sprint(v))
	return b
}

func (b *VOtpInputBuilder) BgColor(v string) (r *VOtpInputBuilder) {
	b.tag.Attr("bg-color", v)
	return b
}

func (b *VOtpInputBuilder) Color(v string) (r *VOtpInputBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VOtpInputBuilder) BaseColor(v string) (r *VOtpInputBuilder) {
	b.tag.Attr("base-color", v)
	return b
}

func (b *VOtpInputBuilder) Disabled(v bool) (r *VOtpInputBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VOtpInputBuilder) Error(v bool) (r *VOtpInputBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VOtpInputBuilder) Variant(v interface{}) (r *VOtpInputBuilder) {
	b.tag.Attr(":variant", h.JSONString(v))
	return b
}

func (b *VOtpInputBuilder) Loading(v interface{}) (r *VOtpInputBuilder) {
	b.tag.Attr(":loading", h.JSONString(v))
	return b
}

func (b *VOtpInputBuilder) Rounded(v interface{}) (r *VOtpInputBuilder) {
	b.tag.Attr(":rounded", h.JSONString(v))
	return b
}

func (b *VOtpInputBuilder) Theme(v string) (r *VOtpInputBuilder) {
	b.tag.Attr("theme", v)
	return b
}

func (b *VOtpInputBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VOtpInputBuilder) Attr(vs ...interface{}) (r *VOtpInputBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VOtpInputBuilder) Children(children ...h.HTMLComponent) (r *VOtpInputBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VOtpInputBuilder) AppendChildren(children ...h.HTMLComponent) (r *VOtpInputBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VOtpInputBuilder) PrependChildren(children ...h.HTMLComponent) (r *VOtpInputBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VOtpInputBuilder) Class(names ...string) (r *VOtpInputBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VOtpInputBuilder) ClassIf(name string, add bool) (r *VOtpInputBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VOtpInputBuilder) On(name string, value string) (r *VOtpInputBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VOtpInputBuilder) Bind(name string, value string) (r *VOtpInputBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VOtpInputBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
