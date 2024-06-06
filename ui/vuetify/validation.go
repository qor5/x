package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VValidationBuilder struct {
	tag *h.HTMLTagBuilder
}

func VValidation(children ...h.HTMLComponent) (r *VValidationBuilder) {
	r = &VValidationBuilder{
		tag: h.Tag("v-validation").Children(children...),
	}
	return
}

func (b *VValidationBuilder) Disabled(v bool) (r *VValidationBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VValidationBuilder) Error(v bool) (r *VValidationBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VValidationBuilder) ErrorMessages(v interface{}) (r *VValidationBuilder) {
	b.tag.Attr(":error-messages", h.JSONString(v))
	return b
}

func (b *VValidationBuilder) MaxErrors(v interface{}) (r *VValidationBuilder) {
	b.tag.Attr(":max-errors", h.JSONString(v))
	return b
}

func (b *VValidationBuilder) Name(v string) (r *VValidationBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VValidationBuilder) Label(v string) (r *VValidationBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VValidationBuilder) Readonly(v bool) (r *VValidationBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VValidationBuilder) Rules(v interface{}) (r *VValidationBuilder) {
	b.tag.Attr(":rules", h.JSONString(v))
	return b
}

func (b *VValidationBuilder) ModelValue(v interface{}) (r *VValidationBuilder) {
	b.tag.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VValidationBuilder) ValidateOn(v interface{}) (r *VValidationBuilder) {
	b.tag.Attr(":validate-on", h.JSONString(v))
	return b
}

func (b *VValidationBuilder) ValidationValue(v interface{}) (r *VValidationBuilder) {
	b.tag.Attr(":validation-value", h.JSONString(v))
	return b
}

func (b *VValidationBuilder) Focused(v bool) (r *VValidationBuilder) {
	b.tag.Attr(":focused", fmt.Sprint(v))
	return b
}

func (b *VValidationBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VValidationBuilder) Attr(vs ...interface{}) (r *VValidationBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VValidationBuilder) Children(children ...h.HTMLComponent) (r *VValidationBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VValidationBuilder) AppendChildren(children ...h.HTMLComponent) (r *VValidationBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VValidationBuilder) PrependChildren(children ...h.HTMLComponent) (r *VValidationBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VValidationBuilder) Class(names ...string) (r *VValidationBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VValidationBuilder) ClassIf(name string, add bool) (r *VValidationBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VValidationBuilder) On(name string, value string) (r *VValidationBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VValidationBuilder) Bind(name string, value string) (r *VValidationBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VValidationBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
