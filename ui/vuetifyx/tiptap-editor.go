package vuetifyx

import (
	"context"

	h "github.com/theplant/htmlgo"
)

// TODO: not ready
type VXTiptapEditorBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXTiptapEditor() (r *VXTiptapEditorBuilder) {
	r = &VXTiptapEditorBuilder{
		tag: h.Tag("vx-tiptap-editor"),
	}
	return
}

type VXTiptapEditorExtension struct {
	Name    string      `json:"name"`
	Options interface{} `json:"options"`
}

func (b *VXTiptapEditorBuilder) Extensions(v []VXTiptapEditorExtension) (r *VXTiptapEditorBuilder) {
	b.tag.Attr(":extensions", h.JSONString(v))
	return b
}

func (b *VXTiptapEditorBuilder) Label(v string) (r *VXTiptapEditorBuilder) {
	b.tag.Attr(":label", h.JSONString(v))
	return b
}

func (b *VXTiptapEditorBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXTiptapEditorBuilder) Attr(vs ...interface{}) (r *VXTiptapEditorBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXTiptapEditorBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
