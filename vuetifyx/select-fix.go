package vuetifyx

import (
	"context"

	"github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
)

type VXSelectBuilder struct {
	tag           *h.HTMLTagBuilder
	selectedItems interface{}
	items         interface{}
}

func VXSelect(children ...h.HTMLComponent) (r *VXSelectBuilder) {
	r = &VXSelectBuilder{
		tag: h.Tag("vx-select").Children(children...),
	}
	return
}

func (b *VXSelectBuilder) ErrorMessages(v ...string) (r *VXSelectBuilder) {
	vuetify.SetErrorMessages(b.tag, v)
	return b
}

func (b *VXSelectBuilder) Items(v interface{}) (r *VXSelectBuilder) {
	b.items = v
	return b
}

func (b *VXSelectBuilder) SelectedItems(v interface{}) (r *VXSelectBuilder) {
	b.selectedItems = v
	return b
}

func (b *VXSelectBuilder) FieldName(v string) (r *VXSelectBuilder) {
	b.tag.Attr("field-name", v)
	return b
}

func (b *VXSelectBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	if b.items == nil {
		b.items = b.selectedItems
	}
	b.tag.Attr(":items", b.items)
	b.tag.Attr(":selected-items", b.selectedItems)

	return b.tag.MarshalHTML(ctx)
}
