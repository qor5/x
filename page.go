package pagui

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
	"github.com/sunfmin/pagui/templates"
)

type PageBuilder struct {
	b              *Builder
	eventFuncRefs  map[string]ui.EventFunc
	pageRenderFunc ui.PageRenderFunc
	h              http.Handler
	pageStateType  reflect.Type
	maxFormSize    int64
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

func (p *PageBuilder) RenderFunc(pslf ui.PageRenderFunc) (r *PageBuilder) {
	p.pageRenderFunc = pslf
	r = p
	return
}

func (p *PageBuilder) RefEventFunc(eventFuncId string, ef ui.EventFunc) (key string) {
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

type eventBody struct {
	EventFuncID EventFuncID `json:"eventFuncId,omitempty"`
	Event       Event       `json:"event,omitempty"`
	PageState   PageState   `json:"pageState,omitempty"`
}

type ServerSideData struct {
	Schema  Component  `json:"schema,omitempty"`
	States  url.Values `json:"states,omitempty"`
	Scripts string     `json:"scripts,omitempty"`
	Styles  string     `json:"styles,omitempty"`
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

func (p *PageBuilder) render(
	serverSideData *ServerSideData,
	w http.ResponseWriter,
	r *http.Request,
	ctx *EventContext,
	renderJSON bool,
) (pager *PageResponse, body *bytes.Buffer) {

	if p.pageRenderFunc == nil {
		return
	}

	ctx.Hub = p
	ctx.R = r
	ctx.W = w
	ctx.Head = &PageHeadBuilder{}
	pr, err := p.pageRenderFunc(ctx)
	if err != nil {
		panic(err)
	}
	pager = &pr

	if pager.Schema == nil {
		panic("page's RenderFunc returns nil schema, use pr.Schema = root to set it")
	}

	// fmt.Println("eventFuncRefs count: ", len(p.eventFuncRefs))
	if comp, ok := pr.Schema.(SchemaComponent); ok {
		serverSideData.Schema = WithContext(ctx, comp)
	}

	body = bytes.NewBuffer(nil)

	if comp, ok := pr.Schema.(HTMLComponent); ok {
		var schema []byte
		schema, err = comp.MarshalHTML(ctx)
		if err != nil {
			panic(err)
		}

		if pr.JSONOnly || renderJSON {
			serverSideData.Schema = string(schema)
			serverSideData.Scripts = ctx.Head.MainScripts(false)
			serverSideData.Styles = ctx.Head.MainStyles(false)
		} else {
			serverSideData.Schema = nil
			serverSideData.Scripts = ""
			serverSideData.Styles = ""
		}

		body.WriteString(ctx.Head.MainStyles(true))

		body.WriteString("<div id=\"app\">\n")
		body.Write(schema)
		body.WriteString("</div>\n")

		body.WriteString(ctx.Head.RealHTML())
		body.WriteString(ctx.Head.MainScripts(true))
	}

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
	ctx := new(EventContext)
	pr, body := p.render(serverSideData, w, r, ctx, false)

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
	resp, err = p.b.GetLayoutMiddleFunc()(nil, ctx.Head)(r, body.String())
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
	ctx.Head = &PageHeadBuilder{}

	if len(p.eventFuncRefs) == 0 {
		log.Println("Rerender because eventFuncs gone, might server restarted")
		ssd := &ServerSideData{}
		p.render(ssd, w, r, ctx, true)
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
		p.render(ssd, w, r, ctx, true)
		er.Schema = ssd.Schema
		er.Scripts = ssd.Scripts
		er.Styles = ssd.Styles
	}

	eventResponseWithContext(ctx, &er)

	err = json.NewEncoder(w).Encode(er)
	if err != nil {
		panic(err)
	}
	return
}

func eventResponseWithContext(ctx *EventContext, er *EventResponse) {
	if comp, ok := er.Alert.(SchemaComponent); ok {
		er.Alert = WithContext(ctx, comp)
	}
	if comp, ok := er.Confirm.(SchemaComponent); ok {
		er.Confirm = WithContext(ctx, comp)
	}
	if comp, ok := er.Schema.(SchemaComponent); ok {
		er.Schema = WithContext(ctx, comp)
	}
	if comp, ok := er.Dialog.(SchemaComponent); ok {
		er.Dialog = WithContext(ctx, comp)
	}
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
