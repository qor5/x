package vuetify

import (
	"fmt"

	"github.com/sunfmin/bran/ui"
	h "github.com/theplant/htmlgo"
)

type DType string

const (
	DTypeFlex       DType = "flex"
	DTypeInlineFlex DType = "inline-flex"
	DTypeBlock      DType = "block"
)

type SizeType string

const (
	Xs SizeType = "xs"
	Sm SizeType = "sm"
	Md SizeType = "md"
	Lg SizeType = "lg"
	Xl SizeType = "xl"
)

type AlignType string

const (
	Left    AlignType = "left"
	Center  AlignType = "center"
	Right   AlignType = "right"
	Justify AlignType = "justify"
)

func (b *VContainerBuilder) DType(v DType) (r *VContainerBuilder) {
	b.tag.Attr(fmt.Sprintf(":d-%s", v), fmt.Sprint(true))
	return b
}

func (b *VContainerBuilder) TextAlign(s SizeType, a AlignType) (r *VContainerBuilder) {
	b.tag.Attr(fmt.Sprintf(":text-%s-%s", s, a), fmt.Sprint(true))
	return b
}

func (b *VLayoutBuilder) DType(v DType) (r *VLayoutBuilder) {
	b.tag.Attr(fmt.Sprintf(":d-%s", v), fmt.Sprint(true))
	return b
}

func (b *VContainerBuilder) GridList(s SizeType) (r *VContainerBuilder) {
	b.tag.Attr(fmt.Sprintf(":grid-list-%s", s), fmt.Sprint(true))
	return b
}

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
		ID:     hub.RegisterEventFunc(eventFuncId, ef),
		Params: params,
	})
	return b
}

func (b *VBtnBuilder) On(v string) (r *VBtnBuilder) {
	b.tag.Attr("v-on", v)
	return b
}

func (b *VBtnBuilder) Children(comps ...h.HTMLComponent) (r *VBtnBuilder) {
	b.tag.Children(comps...)
	return b
}
