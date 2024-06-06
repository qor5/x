package tiptap

// @snippet_begin(TipTapEditorHTMLComponent)
import (
	"context"

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

func (b *TipTapEditorBuilder) Attr(vs ...interface{}) (r *TipTapEditorBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *TipTapEditorBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}

// @snippet_end
