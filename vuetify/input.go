package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VInputBuilder struct {
	tag *h.HTMLTagBuilder
}

func VInput(children ...h.HTMLComponent) (r *VInputBuilder) {
	r = &VInputBuilder{
		tag: h.Tag("v-input").Children(children...),
	}
	return
}

func (b *VInputBuilder) AppendIcon(v string) (r *VInputBuilder) {
	b.tag.Attr("append-icon", v)
	return b
}

func (b *VInputBuilder) BackgroundColor(v string) (r *VInputBuilder) {
	b.tag.Attr("background-color", v)
	return b
}

func (b *VInputBuilder) Color(v string) (r *VInputBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VInputBuilder) Dark(v bool) (r *VInputBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VInputBuilder) Dense(v bool) (r *VInputBuilder) {
	b.tag.Attr(":dense", fmt.Sprint(v))
	return b
}

func (b *VInputBuilder) Disabled(v bool) (r *VInputBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VInputBuilder) Error(v bool) (r *VInputBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VInputBuilder) ErrorCount(v int) (r *VInputBuilder) {
	b.tag.Attr(":error-count", fmt.Sprint(v))
	return b
}

func (b *VInputBuilder) ErrorMessages(v string) (r *VInputBuilder) {
	b.tag.Attr("error-messages", v)
	return b
}

func (b *VInputBuilder) Height(v int) (r *VInputBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VInputBuilder) HideDetails(v bool) (r *VInputBuilder) {
	b.tag.Attr(":hide-details", fmt.Sprint(v))
	return b
}

func (b *VInputBuilder) Hint(v string) (r *VInputBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VInputBuilder) Id(v string) (r *VInputBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VInputBuilder) Label(v string) (r *VInputBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VInputBuilder) Light(v bool) (r *VInputBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VInputBuilder) Loading(v bool) (r *VInputBuilder) {
	b.tag.Attr(":loading", fmt.Sprint(v))
	return b
}

func (b *VInputBuilder) Messages(v string) (r *VInputBuilder) {
	b.tag.Attr("messages", v)
	return b
}

func (b *VInputBuilder) PersistentHint(v bool) (r *VInputBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VInputBuilder) PrependIcon(v string) (r *VInputBuilder) {
	b.tag.Attr("prepend-icon", v)
	return b
}

func (b *VInputBuilder) Readonly(v bool) (r *VInputBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VInputBuilder) Rules(v interface{}) (r *VInputBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VInputBuilder) Success(v bool) (r *VInputBuilder) {
	b.tag.Attr(":success", fmt.Sprint(v))
	return b
}

func (b *VInputBuilder) SuccessMessages(v string) (r *VInputBuilder) {
	b.tag.Attr("success-messages", v)
	return b
}

func (b *VInputBuilder) ValidateOnBlur(v bool) (r *VInputBuilder) {
	b.tag.Attr(":validate-on-blur", fmt.Sprint(v))
	return b
}

func (b *VInputBuilder) Value(v interface{}) (r *VInputBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VInputBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VInputBuilder) Attr(vs ...interface{}) (r *VInputBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VInputBuilder) Children(children ...h.HTMLComponent) (r *VInputBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VInputBuilder) AppendChildren(children ...h.HTMLComponent) (r *VInputBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VInputBuilder) PrependChildren(children ...h.HTMLComponent) (r *VInputBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VInputBuilder) Class(names ...string) (r *VInputBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VInputBuilder) ClassIf(name string, add bool) (r *VInputBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VInputBuilder) On(name string, value string) (r *VInputBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VInputBuilder) Bind(name string, value string) (r *VInputBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VInputBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
