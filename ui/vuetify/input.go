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

func (b *VInputBuilder) Id(v string) (r *VInputBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VInputBuilder) AppendIcon(v interface{}) (r *VInputBuilder) {
	b.tag.Attr(":append-icon", h.JSONString(v))
	return b
}

func (b *VInputBuilder) CenterAffix(v bool) (r *VInputBuilder) {
	b.tag.Attr(":center-affix", fmt.Sprint(v))
	return b
}

func (b *VInputBuilder) PrependIcon(v interface{}) (r *VInputBuilder) {
	b.tag.Attr(":prepend-icon", h.JSONString(v))
	return b
}

func (b *VInputBuilder) HideSpinButtons(v bool) (r *VInputBuilder) {
	b.tag.Attr(":hide-spin-buttons", fmt.Sprint(v))
	return b
}

func (b *VInputBuilder) Hint(v string) (r *VInputBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VInputBuilder) PersistentHint(v bool) (r *VInputBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VInputBuilder) Messages(v interface{}) (r *VInputBuilder) {
	b.tag.Attr(":messages", h.JSONString(v))
	return b
}

func (b *VInputBuilder) Direction(v interface{}) (r *VInputBuilder) {
	b.tag.Attr(":direction", h.JSONString(v))
	return b
}

func (b *VInputBuilder) Density(v interface{}) (r *VInputBuilder) {
	b.tag.Attr(":density", h.JSONString(v))
	return b
}

func (b *VInputBuilder) MaxWidth(v interface{}) (r *VInputBuilder) {
	b.tag.Attr(":max-width", h.JSONString(v))
	return b
}

func (b *VInputBuilder) MinWidth(v interface{}) (r *VInputBuilder) {
	b.tag.Attr(":min-width", h.JSONString(v))
	return b
}

func (b *VInputBuilder) Width(v interface{}) (r *VInputBuilder) {
	b.tag.Attr(":width", h.JSONString(v))
	return b
}

func (b *VInputBuilder) Theme(v string) (r *VInputBuilder) {
	b.tag.Attr("theme", v)
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

func (b *VInputBuilder) ErrorMessages(v interface{}) (r *VInputBuilder) {
	b.tag.Attr(":error-messages", h.JSONString(v))
	return b
}

func (b *VInputBuilder) MaxErrors(v interface{}) (r *VInputBuilder) {
	b.tag.Attr(":max-errors", h.JSONString(v))
	return b
}

func (b *VInputBuilder) Name(v string) (r *VInputBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VInputBuilder) Label(v string) (r *VInputBuilder) {
	b.tag.Attr("label", v)
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

func (b *VInputBuilder) ModelValue(v interface{}) (r *VInputBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VInputBuilder) ValidateOn(v interface{}) (r *VInputBuilder) {
	b.tag.Attr(":validate-on", h.JSONString(v))
	return b
}

func (b *VInputBuilder) ValidationValue(v interface{}) (r *VInputBuilder) {
	b.tag.Attr(":validation-value", h.JSONString(v))
	return b
}

func (b *VInputBuilder) Focused(v bool) (r *VInputBuilder) {
	b.tag.Attr(":focused", fmt.Sprint(v))
	return b
}

func (b *VInputBuilder) HideDetails(v interface{}) (r *VInputBuilder) {
	b.tag.Attr(":hide-details", h.JSONString(v))
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
