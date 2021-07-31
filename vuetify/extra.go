package vuetify

import (
	"fmt"
	"strings"

	"github.com/goplaid/web"
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

func (b *VCheckboxBuilder) LoadPageWithArrayOp(v bool) (r *VCheckboxBuilder) {
	b.tag.Attr("load-page-with-array-op", v)
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

func (b *VChipGroupBuilder) FieldName(v string) (r *VChipGroupBuilder) {
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

func (b *VAutocompleteBuilder) ItemsEventFunc(eventFuncId string, params ...string) (r *VAutocompleteBuilder) {

	b.tag.Attr(":items-event-func-id", &web.EventFuncID{
		ID:     eventFuncId,
		Params: params,
	})
	return b
}

func (b *VBtnBuilder) OnClick(eventFuncId string, params ...string) (r *VBtnBuilder) {
	web.Bind(b.tag).OnClick(eventFuncId, params...).Update()
	return b
}

func (b *VListItemBuilder) Slot(v string) (r *VListItemBuilder) {
	b.tag.Attr("slot", v)
	return b
}

func stringsTrim(vs ...string) (r []string) {
	for _, v := range vs {
		if cv := strings.TrimSpace(v); len(cv) > 0 {
			r = append(r, cv)
		}
	}
	return
}

func setErrorMessages(t h.MutableAttrHTMLComponent, vs []string) {
	cvs := stringsTrim(vs...)
	if len(cvs) == 0 {
		return
	}
	t.SetAttr(":error-messages", h.JSONString(cvs))
}
