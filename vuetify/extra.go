package vuetify

import (
	"github.com/sunfmin/bran/ui"
)

func (b *VTextFieldBuilder) FieldName(v string) (r *VTextFieldBuilder) {
	b.tag.Attr("field-name", v)
	return b
}

func (b *VTextareaBuilder) FieldName(v string) (r *VTextareaBuilder) {
	b.tag.Attr("field-name", v)
	return b
}

func (b *VCheckboxBuilder) FieldName(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("field-name", v)
	return b
}

func (b *VRadioGroupBuilder) FieldName(v string) (r *VRadioGroupBuilder) {
	b.tag.Attr("field-name", v)
	return b
}

func (b *VSwitchBuilder) FieldName(v string) (r *VSwitchBuilder) {
	b.tag.Attr("field-name", v)
	return b
}

func (b *VSliderBuilder) FieldName(v string) (r *VSliderBuilder) {
	b.tag.Attr("field-name", v)
	return b
}

func (b *VSelectBuilder) FieldName(v string) (r *VSelectBuilder) {
	b.tag.Attr("field-name", v)
	return b
}

func (b *VSelectBuilder) SelectedItems(v interface{}) (r *VSelectBuilder) {
	b.selectedItems = v
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

func (b *VAutocompleteBuilder) ItemsEventFunc(hub ui.EventFuncHub, eventFuncId string, ef ui.EventFunc, params ...string) (r *VAutocompleteBuilder) {

	if ef == nil {
		return b
	}

	b.tag.Attr(":items-event-func-id", &ui.EventFuncID{
		ID:     hub.RefEventFunc(eventFuncId, ef),
		Params: params,
	})
	return b
}
