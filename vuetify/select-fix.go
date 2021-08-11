package vuetify

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type VSelectBuilder struct {
	tag           *h.HTMLTagBuilder
	selectedItems interface{}
	items         interface{}
}

func VSelect(children ...h.HTMLComponent) (r *VSelectBuilder) {
	r = &VSelectBuilder{
		tag: h.Tag("vw-select").Children(children...),
	}
	return
}

func (b *VSelectBuilder) ErrorMessages(v ...string) (r *VSelectBuilder) {
	setErrorMessages(b.tag, v)
	return b
}

func (b *VSelectBuilder) Items(v interface{}) (r *VSelectBuilder) {
	b.items = v
	return b
}

func (b *VSelectBuilder) SelectedItems(v interface{}) (r *VSelectBuilder) {
	b.selectedItems = v
	return b
}

func (b *VSelectBuilder) FieldName(v string) (r *VSelectBuilder) {
	b.tag.Attr("field-name", v)
	return b
}

func (b *VSelectBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	if b.items == nil {
		b.items = b.selectedItems
	}
	b.tag.Attr(":items", b.items)
	b.tag.Attr(":selected-items", b.selectedItems)

	return b.tag.MarshalHTML(ctx)
}
