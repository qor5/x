package vuetify

import (
	"context"
	"fmt"

	"github.com/sunfmin/bran/ui"
	h "github.com/theplant/htmlgo"
)

type VAutoCompleteBuilder struct {
	tag           *h.HTMLTagBuilder
	selectedItems interface{}
	items         interface{}
}

func VAutoComplete() (r *VAutoCompleteBuilder) {
	r = &VAutoCompleteBuilder{
		tag: h.Tag("vw-autocomplete"),
	}
	r.Multiple(true)
	return
}

func (b *VAutoCompleteBuilder) Items(v interface{}) (r *VAutoCompleteBuilder) {
	b.items = v
	return b
}

func (b *VAutoCompleteBuilder) SelectedItems(v interface{}) (r *VAutoCompleteBuilder) {
	b.selectedItems = v
	return b
}

func (b *VAutoCompleteBuilder) ItemsEventFunc(hub ui.EventFuncHub, eventFuncId string, ef ui.EventFunc, params ...string) (r *VAutoCompleteBuilder) {

	if ef == nil {
		return b
	}

	b.tag.Attr(":items-event-func-id", &ui.EventFuncID{
		ID:     hub.RefEventFunc(eventFuncId, ef),
		Params: params,
	})
	return b
}

func (b *VAutoCompleteBuilder) Label(v string) (r *VAutoCompleteBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VAutoCompleteBuilder) FieldName(v string) (r *VAutoCompleteBuilder) {
	b.tag.Attr("field-name", v)
	return b
}

func (b *VAutoCompleteBuilder) ItemText(v string) (r *VAutoCompleteBuilder) {
	b.tag.Attr("item-text", v)
	return b
}

func (b *VAutoCompleteBuilder) ItemValue(v string) (r *VAutoCompleteBuilder) {
	b.tag.Attr("item-value", v)
	return b
}

func (b *VAutoCompleteBuilder) Multiple(v bool) (r *VAutoCompleteBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VAutoCompleteBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	if b.items == nil {
		b.items = b.selectedItems
	}
	b.tag.Attr(":items", b.items)
	b.tag.Attr(":selected-items", b.selectedItems)

	return b.tag.MarshalHTML(ctx)
}
