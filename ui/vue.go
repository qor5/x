package ui

import (
	"context"
	"fmt"
	"net/url"

	h "github.com/theplant/htmlgo"
)

type VueEventTagBuilder struct {
	tag           h.MutableAttrHTMLComponent
	fieldName     *string
	onInputFuncID *EventFuncID
}

func Bind(b h.MutableAttrHTMLComponent) (r *VueEventTagBuilder) {
	r = &VueEventTagBuilder{}
	r.tag = b
	return
}

func (b *VueEventTagBuilder) OnInput(eventFuncId string, params ...string) (r *VueEventTagBuilder) {

	b.onInputFuncID = &EventFuncID{
		ID:     eventFuncId,
		Params: params,
	}

	return b
}

func (b *VueEventTagBuilder) PushStateLink(pageURL string) (r *VueEventTagBuilder) {
	b.tag.SetAttr("v-on:click", fmt.Sprintf("topage(%s, %s)", h.JSONString(url.Values{}), h.JSONString(pageURL)))
	return b
}

func (b *VueEventTagBuilder) OnClick(eventFuncId string, params ...string) (r *VueEventTagBuilder) {

	fid := &EventFuncID{
		ID:     eventFuncId,
		Params: params,
	}

	b.tag.SetAttr("v-on:click", fmt.Sprintf("onclick(%s, $event)", h.JSONString(fid)))
	return b
}

func (b *VueEventTagBuilder) FieldName(v string) (r *VueEventTagBuilder) {
	b.fieldName = &v
	return b
}

func (b *VueEventTagBuilder) setupChange() {
	if b.fieldName == nil && b.onInputFuncID == nil {
		return
	}

	b.tag.SetAttr("v-on:input",
		fmt.Sprintf(`oninput(%s, %s, $event)`, h.JSONString(b.onInputFuncID), h.JSONString(b.fieldName)))
}

func (b *VueEventTagBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	b.setupChange()
	return b.tag.MarshalHTML(ctx)
}
