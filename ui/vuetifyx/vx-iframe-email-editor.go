package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VXIframeEmailEditorBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXIframeEmailEditor(children ...h.HTMLComponent) (r *VXIframeEmailEditorBuilder) {
	r = &VXIframeEmailEditorBuilder{
		tag: h.Tag("vx-iframe-email-editor").Children(children...),
	}
	return
}

func (b *VXIframeEmailEditorBuilder) Ref(v interface{}) (r *VXIframeEmailEditorBuilder) {
	b.tag.Attr(":ref", h.JSONString(v))
	return b
}

func (b *VXIframeEmailEditorBuilder) Src(v interface{}) (r *VXIframeEmailEditorBuilder) {
	b.tag.Attr(":src", h.JSONString(v))
	return b
}

func (b *VXIframeEmailEditorBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXIframeEmailEditorBuilder) Attr(vs ...interface{}) (r *VXIframeEmailEditorBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXIframeEmailEditorBuilder) Children(children ...h.HTMLComponent) (r *VXIframeEmailEditorBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXIframeEmailEditorBuilder) AppendChildren(children ...h.HTMLComponent) (r *VXIframeEmailEditorBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VXIframeEmailEditorBuilder) PrependChildren(children ...h.HTMLComponent) (r *VXIframeEmailEditorBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VXIframeEmailEditorBuilder) Class(names ...string) (r *VXIframeEmailEditorBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXIframeEmailEditorBuilder) ClassIf(name string, add bool) (r *VXIframeEmailEditorBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VXIframeEmailEditorBuilder) On(name string, value string) (r *VXIframeEmailEditorBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VXIframeEmailEditorBuilder) Bind(name string, value string) (r *VXIframeEmailEditorBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VXIframeEmailEditorBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
