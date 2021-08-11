package vuetify

import (
	"context"

	"github.com/goplaid/web"
	h "github.com/theplant/htmlgo"
)

type VAutocompleteBuilder struct {
	tag           *h.HTMLTagBuilder
	selectedItems interface{}
	items         interface{}
}

func VAutocomplete(children ...h.HTMLComponent) (r *VAutocompleteBuilder) {
	r = &VAutocompleteBuilder{
		tag: h.Tag("vw-autocomplete").Children(children...),
	}
	r.Multiple(true)

	return
}

func (b *VAutocompleteBuilder) ErrorMessages(v ...string) (r *VAutocompleteBuilder) {
	setErrorMessages(b.tag, v)
	return b
}

func (b *VAutocompleteBuilder) Items(v interface{}) (r *VAutocompleteBuilder) {
	b.items = v
	return b
}

func (b *VAutocompleteBuilder) FieldName(v string) (r *VAutocompleteBuilder) {
	b.tag.Attr("field-name", v)
	return b
}

func (b *VAutocompleteBuilder) SelectedItems(v interface{}) (r *VAutocompleteBuilder) {
	b.selectedItems = v
	return b
}

func (b *VAutocompleteBuilder) ItemsEventFunc(eventFuncId string, params ...string) (r *VAutocompleteBuilder) {

	b.tag.Attr(":items-event-func-id", &web.EventFuncID{
		ID:     eventFuncId,
		Params: params,
	})
	return b
}

func (b *VAutocompleteBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	if b.items == nil {
		b.items = b.selectedItems
	}
	b.tag.Attr(":items", b.items)
	b.tag.Attr(":selected-items", b.selectedItems)
	return b.tag.MarshalHTML(ctx)
}
