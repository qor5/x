package ui

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/sunfmin/reflectutils"
)

type Component interface {
}

type SchemaComponent interface {
	MarshalSchema(ctx *EventContext) ([]byte, error)
}

type HTMLComponent interface {
	MarshalHTML(ctx *EventContext) ([]byte, error)
}

type RawHTML string

func (s RawHTML) MarshalHTML(ctx *EventContext) (r []byte, err error) {
	r = []byte(s)
	return
}

type RawSchema string

func (s RawSchema) MarshalSchema(ctx *EventContext) (r []byte, err error) {
	r = []byte(s)
	return
}

type ComponentFunc func(ctx *EventContext) (r []byte, err error)

func (f ComponentFunc) MarshalHTML(ctx *EventContext) (r []byte, err error) {
	return f(ctx)
}

func (f ComponentFunc) MarshalSchema(ctx *EventContext) (r []byte, err error) {
	return f(ctx)
}

type HTMLComponents []HTMLComponent

func (hcs HTMLComponents) MarshalHTML(ctx *EventContext) (r []byte, err error) {
	buf := bytes.NewBuffer(nil)
	for _, h := range hcs {
		var b []byte
		b, err = h.MarshalHTML(ctx)
		if err != nil {
			return
		}
		buf.Write(b)
	}
	r = buf.Bytes()
	return
}

type PageState interface{}

type PageResponse struct {
	Schema   Component
	State    PageState
	JSONOnly bool
}

type EventResponse struct {
	Alert       Component   `json:"alert,omitempty"`
	Confirm     Component   `json:"confirm,omitempty"`
	Dialog      Component   `json:"dialog,omitempty"`
	CloseDialog bool        `json:"closeDialog,omitempty"`
	Schema      Component   `json:"schema,omitempty"`
	State       PageState   `json:"states,omitempty"`
	Reload      bool        `json:"reload,omitempty"`
	RedirectURL string      `json:"redirectURL,omitempty"`
	Data        interface{} `json:"data,omitempty"` // used for return collection data like TagsInput data source
	Scripts     string      `json:"scripts,omitempty"`
	Styles      string      `json:"styles,omitempty"`
}

type PageFunc func(ctx *EventContext) (r PageResponse, err error)

type EventFunc func(ctx *EventContext) (r EventResponse, err error)

type LayoutFunc func(r *http.Request, body string) (output string, err error)

type LayoutMiddleFunc func(in LayoutFunc, injector PageInjector) (out LayoutFunc)

type EventFuncHub interface {
	RefEventFunc(eventFuncId string, ef EventFunc) (key string)
}

/*
	PushState: Whatever put into this, will do window.history.pushState to the current page url with
	it as query string, for example: /my-page-url/?key=name&value=felix. and It also pass the query string along
	to the /my-page-url/__execute_event__/?key=name&value=felix, Mostly used for setting EventResponse: `er.Reload = true` case.
	So that you can refresh the page with different query string in pushState manner, without doing a Browser redirect or refresh.
	It is used in Pager (Pagination) component.
*/
type EventFuncID struct {
	ID        string     `json:"id,omitempty"`
	Params    []string   `json:"params,omitempty"`
	PushState url.Values `json:"pushState"` // This we don't omitempty, So that {} can be keeped when use url.Values{}
}

/*
	Event is for an individual component like checkbox, input, data picker etc's onChange callback
	will pass the Event to server side. use ctx.Event.Checked etc to get the value.
*/
type Event struct {
	Checked bool     `json:"checked,omitempty"` // For Checkbox
	From    string   `json:"from,omitempty"`    // For DatePicker
	To      string   `json:"to,omitempty"`      // For DatePicker
	Value   string   `json:"value,omitempty"`   // For Input, DatePicker
	Params  []string `json:"-"`
}

type EventContext struct {
	R        *http.Request
	W        http.ResponseWriter
	Hub      EventFuncHub
	Injector PageInjector
	State    PageState
	Event    *Event
}

type PageInjector interface {
	Title(title string)
	Meta(attrs ...string)
	MetaNameContent(name, content string)
	PutScript(script string)
	PutStyle(style string)
	PutHeadHTML(v string)
	PutTailHTML(v string)

	HeadString() string
}

func (ctx *EventContext) StateOrInit(v PageState) (r PageState) {
	if ctx.State == nil {
		ctx.State = v
	}
	r = ctx.State
	return
}

func (ctx *EventContext) SubStateOrInit(reflectPath string, v interface{}) (r interface{}) {

	r = reflectutils.MustGet(ctx.State, reflectPath)
	if r == nil {
		err := reflectutils.Set(ctx.State, reflectPath, v)
		if err != nil {
			panic(err)
		}
		r = reflectutils.MustGet(ctx.State, reflectPath)
	}

	return
}

func WithContext(ctx *EventContext, comp SchemaComponent) json.Marshaler {
	return &withCtx{ctx, comp}
}

type withCtx struct {
	ctx  *EventContext
	body SchemaComponent
}

func (wc *withCtx) MarshalJSON() ([]byte, error) {
	return wc.body.MarshalSchema(wc.ctx)
}
