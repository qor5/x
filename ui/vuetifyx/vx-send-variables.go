package vuetifyx

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type VXSendVariablesBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXSendVariables(children ...h.HTMLComponent) (r *VXSendVariablesBuilder) {
	r = &VXSendVariablesBuilder{
		tag: h.Tag("vx-send-variables").Children(children...),
	}
	return
}

func (b *VXSendVariablesBuilder) Value(v string) (r *VXSendVariablesBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VXSendVariablesBuilder) Placeholder(v string) (r *VXSendVariablesBuilder) {
	b.tag.Attr(":placeholder", h.JSONString(v))
	return b
}

func (b *VXSendVariablesBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXSendVariablesBuilder) Attr(vs ...interface{}) (r *VXSendVariablesBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXSendVariablesBuilder) Children(children ...h.HTMLComponent) (r *VXSendVariablesBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXSendVariablesBuilder) AppendChildren(children ...h.HTMLComponent) (r *VXSendVariablesBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VXSendVariablesBuilder) PrependChildren(children ...h.HTMLComponent) (r *VXSendVariablesBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VXSendVariablesBuilder) Class(names ...string) (r *VXSendVariablesBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXSendVariablesBuilder) ClassIf(name string, add bool) (r *VXSendVariablesBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VXSendVariablesBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
