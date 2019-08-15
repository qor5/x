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
	eventType     string
	eventFunc     *EventFuncID
	pageURL       string
	toPage        bool
}

func Bind(b h.MutableAttrHTMLComponent) (r *VueEventTagBuilder) {
	r = &VueEventTagBuilder{
		eventType: "click",
		eventFunc: &EventFuncID{},
	}
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
	b.pageURL = pageURL
	b.toPage = true
	return b
}

func (b *VueEventTagBuilder) OnClick(eventFuncId string, params ...string) (r *VueEventTagBuilder) {
	return b.EventFunc(eventFuncId, params...)
}

func (b *VueEventTagBuilder) On(eventType string) (r *VueEventTagBuilder) {
	b.eventType = eventType
	return b
}

func (b *VueEventTagBuilder) EventFunc(eventFuncId string, params ...string) (r *VueEventTagBuilder) {
	b.eventFunc.ID = eventFuncId
	b.eventFunc.Params = params
	return b
}

func (b *VueEventTagBuilder) PageURL(pageURL string) (r *VueEventTagBuilder) {
	b.pageURL = pageURL
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

func (b *VueEventTagBuilder) Update() {
	b.setupChange()

	callFunc := ""

	if len(b.eventFunc.ID) > 0 {
		if len(b.pageURL) > 0 {
			callFunc = fmt.Sprintf("triggerEventFunc(%s, $event, %s)",
				h.JSONString(b.eventFunc),
				h.JSONString(b.pageURL),
			)
		} else {
			callFunc = fmt.Sprintf("triggerEventFunc(%s, $event)",
				h.JSONString(b.eventFunc),
			)
		}
	}

	if b.toPage {
		callFunc = fmt.Sprintf("topage(%s, %s)", h.JSONString(url.Values{}), h.JSONString(b.pageURL))
	}

	if len(callFunc) > 0 {
		b.tag.SetAttr(fmt.Sprintf("v-on:%s", b.eventType), callFunc)
	}
}

func (b *VueEventTagBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	b.Update()
	return b.tag.MarshalHTML(ctx)
}
