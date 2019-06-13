package codehighlight

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type CodeBuilder struct {
	tag  *h.HTMLTagBuilder
	code string
}

func Code(code string) (r *CodeBuilder) {
	r = &CodeBuilder{
		tag: h.Tag("bran-code"),
	}
	r.Code(code)
	return
}

func (b *CodeBuilder) Code(v string) (r *CodeBuilder) {
	b.code = v
	return b
}

func (b *CodeBuilder) Language(v string) (r *CodeBuilder) {
	b.tag.Attr("language", v)
	return b
}

func (b *CodeBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	b.tag.Children(
		h.Template().Text(b.code).Attr("slot", "sourcecode"),
	)
	return b.tag.MarshalHTML(ctx)
}
