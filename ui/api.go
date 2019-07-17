package ui

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Component interface {
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
	Flash    interface{} // pass value from actions to index
}

type PageInjector interface {
	Title(title string)
	Meta(attrs ...string)
	MetaNameContent(name, content string)
	//PutScript(script string)
	//PutStyle(style string)
	PutHeadHTML(v string)
	PutTailHTML(v string)

	HeadString() string
}

const eventContextKey = iota

func WrapEventContext(parent context.Context, ctx *EventContext) (r context.Context) {
	r = context.WithValue(parent, eventContextKey, ctx)
	return
}

func MustGetEventContext(c context.Context) (r *EventContext) {
	r, _ = c.Value(eventContextKey).(*EventContext)
	if r == nil {
		panic("EventContext required")
	}
	return
}

func Injector(c context.Context) (r PageInjector) {
	ctx := MustGetEventContext(c)
	r = ctx.Injector
	return
}

type Styles struct {
	pairs [][]string
}

func (s *Styles) String() string {
	segs := []string{}
	for _, v := range s.pairs {
		segs = append(segs, fmt.Sprintf("%s:%s;", v[0], v[1]))
	}
	return strings.Join(segs, " ")
}

func (s *Styles) Put(name, value string) (r *Styles) {
	for _, el := range s.pairs {
		if el[0] == name {
			el[1] = value
			return s
		}
	}

	s.pairs = append(s.pairs, []string{name, value})
	return s
}
