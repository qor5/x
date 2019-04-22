package page

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/sunfmin/reflectutils"

	"github.com/go-playground/form"
	"github.com/sunfmin/page/templates"
)

type PageBuilder struct {
	b			  *Builder
	eventFuncRefs  map[string]EventFunc
	pageRenderFunc PageRenderFunc
	h			  http.Handler
	pageStateType  reflect.Type
	maxFormSize	int64
}

func (b *Builder) NewPage() (pb *PageBuilder) {
	pb = &PageBuilder{}
	pb.eventFuncRefs = make(map[string]EventFunc)
	pb.b = b
	return
}

func (p *PageBuilder) MaxFormSize(v int64) (r *PageBuilder) {
	p.maxFormSize = v
	r = p
	return
}

func (p *PageBuilder) RenderFunc(pslf PageRenderFunc) (r *PageBuilder) {
	p.pageRenderFunc = pslf
	r = p
	return
}

func (p *PageBuilder) RefEventFunc(eventFuncId string, ef EventFunc) (key string) {
	key = eventFuncId
	if f, ok := p.eventFuncRefs[eventFuncId]; ok {
		funcAddress := fmt.Sprint(ef)
		if fmt.Sprint(f) == funcAddress {
			return
		}

		panic(fmt.Sprintf("%s already registered in this page", eventFuncId))
	}
	p.eventFuncRefs[eventFuncId] = ef
	return
}

type Component interface {
}

type ComponentJSON interface {
	MarshalJSON() ([]byte, error)
}

type ComponentHTML interface {
	MarshalHTML(phb *PageHeadBuilder) ([]byte, error)
}

type StringComponentHTML string

func (s StringComponentHTML) MarshalHTML(phb *PageHeadBuilder) (r []byte, err error) {
	r = []byte(s)
	return
}

type StringComponentJSON string

func (s StringComponentJSON) MarshalJSON() (r []byte, err error) {
	r = []byte(s)
	return
}

type EventContext struct {
	R	 *http.Request
	W	 http.ResponseWriter
	Hub   EventFuncHub
	State PageState
	Event *Event
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

type PageState interface{}

type PageResponse struct {
	Schema   Component
	State	PageState
	Head	 PageHeadBuilder
	JSONOnly bool
}

type PageRenderFunc func(ctx *EventContext) (r PageResponse, err error)

type EventResponse struct {
	Alert	   Component   `json:"alert,omitempty"`
	Confirm	 Component   `json:"confirm,omitempty"`
	Dialog	  Component   `json:"dialog,omitempty"`
	CloseDialog bool		`json:"closeDialog,omitempty"`
	Schema	  Component   `json:"schema,omitempty"`
	State	   PageState   `json:"states,omitempty"`
	Reload	  bool		`json:"reload,omitempty"`
	RedirectURL string	  `json:"redirectURL,omitempty"`
	Data		interface{} `json:"data,omitempty"` // used for return collection data like TagsInput data source
}

type EventFunc func(ctx *EventContext) (r EventResponse, err error)

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
	ID		string	 `json:"id,omitempty"`
	Params	[]string   `json:"params,omitempty"`
	PushState url.Values `json:"pushState"` // This we don't omitempty, So that {} can be keeped when use url.Values{}
}

/*
	Event is for an individual component like checkbox, input, data picker etc's onChange callback
	will pass the Event to server side. use ctx.Event.Checked etc to get the value.
*/
type Event struct {
	Checked bool	 `json:"checked,omitempty"` // For Checkbox
	From	string   `json:"from,omitempty"`	// For DatePicker
	To	  string   `json:"to,omitempty"`	  // For DatePicker
	Value   string   `json:"value,omitempty"`   // For Input, DatePicker
	Params  []string `json:"-"`
}

type eventBody struct {
	EventFuncID EventFuncID `json:"eventFuncId,omitempty"`
	Event	   Event	   `json:"event,omitempty"`
	PageState   PageState   `json:"pageState,omitempty"`
}

type ServerSideData struct {
	Schema Component  `json:"schema,omitempty"`
	States url.Values `json:"states,omitempty"`
}

func (p *PageBuilder) render(serverSideData *ServerSideData, w http.ResponseWriter, r *http.Request, ctx *EventContext) (pager *PageResponse) {

	if p.pageRenderFunc == nil {
		return
	}
	if ctx == nil {
		ctx = new(EventContext)
		ctx.Hub = p
	}
	ctx.R = r
	ctx.W = w
	pr, err := p.pageRenderFunc(ctx)
	if err != nil {
		panic(err)
	}
	pager = &pr

	// fmt.Println("eventFuncRefs count: ", len(p.eventFuncRefs))
	serverSideData.Schema = pr.Schema

	// default page response state to ctx state if not set
	if ctx.State != nil && pr.State == nil {
		pr.State = ctx.State
	}

	if pr.State == nil {
		return
	}
	p.pageStateType, serverSideData.States = encodePageState(pr.State)
	return
}

func encodePageState(pageState PageState) (pageStateType reflect.Type, values url.Values) {
	var cantEncode bool
	pageStateType, cantEncode = getPageStateType(pageState)
	if cantEncode {
		return
	}

	fe := form.NewEncoder()
	var err error
	values, err = fe.Encode(pageState)
	if err != nil {
		panic(err)
	}
	return
}

func (p *PageBuilder) index(w http.ResponseWriter, r *http.Request) {
	var err error

	var serverSideData = &ServerSideData{}
	pr := p.render(serverSideData, w, r, nil)
	if serverSideData.Schema == nil {
		panic("page's RenderFunc returns nil schema, use pr.Schema = root to set it")
	}
	head := pr.Head

	var body = bytes.NewBuffer(nil)

	if hcomp, ok := serverSideData.Schema.(ComponentHTML); ok {
		var schema []byte
		schema, err = hcomp.MarshalHTML(&head)
		if err != nil {
			panic(err)
		}

		if pr.JSONOnly {
			serverSideData.Schema = string(schema)
		} else {
			serverSideData.Schema = nil
		}

		body.Write(schema)
		body.WriteString("\n")
	}

	var serverSideDataJSON []byte
	serverSideDataJSON, err = json.MarshalIndent(serverSideData, "", "\t")
	if err != nil {
		panic(err)
	}

	if pr.JSONOnly {
		fmt.Fprintln(w, string(serverSideDataJSON))
		return
	}

	body.WriteString(templates.PageInitialization(string(serverSideDataJSON)))

	var resp string
	resp, err = p.b.GetLayoutMiddleFunc()(nil, &head)(r, body.String())
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, resp)
	return
}

func getPageStateType(pageState PageState) (t reflect.Type, cantEncode bool) {
	v := reflect.ValueOf(pageState)

	if v.Kind() != reflect.Ptr {
		t = v.Type()
		return
	}

	if v.IsNil() {
		cantEncode = true
		t = v.Type()
		for t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		return
	}

	for v.Elem().Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t = v.Elem().Type()
	return
}

func (p *PageBuilder) eventBodyFromRequest(r *http.Request) *eventBody {
	maxSize := p.maxFormSize
	if maxSize == 0 {
		maxSize = 128 << 20 // 128MB
	}

	err := r.ParseMultipartForm(maxSize)
	if err != nil {
		panic(err)
	}

	multif := r.MultipartForm

	var eb eventBody
	err = json.NewDecoder(strings.NewReader(multif.Value["__event_data__"][0])).Decode(&eb)
	if err != nil {
		panic(err)
	}

	if p.pageStateType != nil {
		pageState := p.NewPageState()
		dec := form.NewDecoder()

		err = dec.Decode(pageState, multif.Value)
		if err != nil {
			panic(err)
		}

		if len(multif.File) > 0 {
			for k, vs := range multif.File {
				err = reflectutils.Set(pageState, k, vs)
				if err != nil {
					panic(err)
				}
			}
		}

		eb.PageState = pageState
	}
	return &eb
}

func (p *PageBuilder) executeEvent(w http.ResponseWriter, r *http.Request) {
	// for server side restart and lost all the eventFuncs, but user keep clicking page without refresh page to call p.render to fill up eventFuncs

	ctx := new(EventContext)
	ctx.R = r
	ctx.W = w
	ctx.Hub = p

	if len(p.eventFuncRefs) == 0 {
		log.Println("Rerender because eventFuncs gone, might server restarted")
		ssd := &ServerSideData{}
		p.render(ssd, w, r, ctx)
		json.Marshal(ssd) // to fill in event funcs that setup inside a component
	}

	eb := p.eventBodyFromRequest(r)
	ctx.Event = &eb.Event
	ctx.State = eb.PageState

	ef, ok := p.eventFuncRefs[eb.EventFuncID.ID]
	if !ok {
		panic(fmt.Errorf("event %s not found", eb.EventFuncID.ID))
	}

	eb.Event.Params = eb.EventFuncID.Params

	er, err := ef(ctx)
	if err != nil {
		panic(err)
	}

	// for TagsInput like needs datasource to return data, don't need to return State
	if er.Data != nil {
		ctx.State = nil
	}

	// default event response state to ctx state if not set
	if er.State == nil && ctx.State != nil {
		er.State = ctx.State
	}

	if er.State != nil {
		_, er.State = encodePageState(er.State)
	}

	if er.Reload {
		ssd := &ServerSideData{}
		p.render(ssd, w, r, ctx)
		er.Schema = ssd.Schema
	}

	err = json.NewEncoder(w).Encode(er)
	if err != nil {
		panic(err)
	}
	return
}

func (p *PageBuilder) NewPageState() interface{} {
	if p.pageStateType == nil {
		return nil
	}

	return reflect.New(p.pageStateType).Interface()
}

func (p *PageBuilder) Handler() http.Handler {
	if p.h == nil {
		m := http.NewServeMux()
		m.HandleFunc("/__execute_event__/", p.executeEvent)
		m.HandleFunc("/", p.index)
		p.h = m
	}
	return p.h
}
