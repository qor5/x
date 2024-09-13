package vuetifyx

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

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

func (b *VXTiptapEditorBuilder) Attr(vs ...interface{}) (r *VXTiptapEditorBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXTiptapEditorBuilder) Disabled(v bool) (r *VXTiptapEditorBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VXTiptapEditorBuilder) Readonly(v bool) (r *VXTiptapEditorBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VXTiptapEditorBuilder) Value(v string) (r *VXTiptapEditorBuilder) {
	b.Attr(":model-value", h.JSONString(v))
	return b
}

func (b *VXTiptapEditorBuilder) MarkdownTheme(v string) (r *VXTiptapEditorBuilder) {
	b.Attr(":markdown-theme", h.JSONString(v))
	return b
}

func (b *VXTiptapEditorBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
