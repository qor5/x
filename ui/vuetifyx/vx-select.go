package vuetifyx

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type VXSelectBuilder struct {
	tag           *h.HTMLTagBuilder
	items         interface{}
}

func VXSelect(children ...h.HTMLComponent) (r *VXSelectBuilder) {
	r = &VXSelectBuilder{
		tag: h.Tag("vx-select").Children(children...),
	}
	return
}

func (b *VXSelectBuilder) ErrorMessages(v ...string) (r *VXSelectBuilder) {
	b.tag.Attr("error-messages", v)
	return b
}

func (b *VXSelectBuilder) Items(v interface{}) (r *VXSelectBuilder) {
	b.tag.Attr(":items", b.items)
	return b
}

func (b *VXSelectBuilder) FieldName(v string) (r *VXSelectBuilder) {
	b.tag.Attr("field-name", v)
	return b
}

func (b *VXSelectBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
