package tiptap

// @snippet_begin(TipTapEditorHTMLComponent)
import (
	"context"

	"github.com/goplaid/web"
	h "github.com/theplant/htmlgo"
)

type TipTapEditorBuilder struct {
	tag *h.HTMLTagBuilder
}

func TipTapEditor() (r *TipTapEditorBuilder) {
	r = &TipTapEditorBuilder{
		tag: h.Tag("tiptap-editor"),
	}

	return
}

func (b *TipTapEditorBuilder) FieldName(v string) (r *TipTapEditorBuilder) {
	b.tag.Attr(web.VFieldName(v)...)
	return b
}

func (b *TipTapEditorBuilder) Value(v string) (r *TipTapEditorBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *TipTapEditorBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}

// @snippet_end
