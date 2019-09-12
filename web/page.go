package web

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	h "github.com/theplant/htmlgo"
)

type PageBuilder struct {
	b              *Builder
	eventFuncs     []*idEventFunc
	pageRenderFunc PageFunc
	maxFormSize    int64
}

func (b *Builder) Page(pf PageFunc) (p *PageBuilder) {
	p = &PageBuilder{}
	p.b = b
	p.pageRenderFunc = pf
	p.RegisterEventFunc("__reload__", reload)
	return
}

func (p *PageBuilder) MaxFormSize(v int64) (r *PageBuilder) {
	p.maxFormSize = v
	r = p
	return
}

func (p *PageBuilder) RegisterEventFunc(eventFuncId string, ef EventFunc) (key string) {
	key = eventFuncId
	if p.eventFuncById(eventFuncId) != nil {
		return
	}

	p.eventFuncs = append(p.eventFuncs, &idEventFunc{eventFuncId, ef})
	return
}

func (p *PageBuilder) eventFuncById(id string) (r EventFunc) {
	for _, ne := range p.eventFuncs {
		if ne.id == id {
			r = ne.ef
			return
		}
	}
	return
}

type idEventFunc struct {
	id string
	ef EventFunc
}

type eventBody struct {
	EventFuncID EventFuncID `json:"eventFuncId,omitempty"`
	Event       Event       `json:"event,omitempty"`
}

type serverSideData struct {
	Schema interface{} `json:"schema,omitempty"`
}

func (p *PageBuilder) render(
	ssd *serverSideData,
	w http.ResponseWriter,
	r *http.Request,
	c context.Context,
	head *PageInjector,
) (pager *PageResponse, isRenderHTML bool) {

	if p.pageRenderFunc == nil {
		return
	}

	ctx := MustGetEventContext(c)

	ctx.Hub = p
	ctx.R = r
	ctx.W = w
	ctx.Injector = head

	pr, err := p.pageRenderFunc(ctx)
	if err != nil {
		panic(err)
	}
	pager = &pr

	if pager.Schema == nil {
		panic("page's RenderFunc returns nil schema, use pr.Schema = root to set it")
	}

	// fmt.Println("eventFuncs count: ", len(p.eventFuncs))
	if comp, ok := pager.Schema.(SchemaComponent); ok {
		b, err := comp.MarshalSchema(ctx)
		if err != nil {
			panic(err)
		}
		ssd.Schema = json.RawMessage(b)
	} else if comp, ok := pager.Schema.(h.HTMLComponent); ok {
		b, err := comp.MarshalHTML(c)
		if err != nil {
			panic(err)
		}
		isRenderHTML = true
		ssd.Schema = string(b)
	}

	return
}

func (p *PageBuilder) index(w http.ResponseWriter, r *http.Request) {
	var err error

	var ssd = &serverSideData{}
	var head = &PageInjector{}

	ctx := new(EventContext)
	c := WrapEventContext(r.Context(), ctx)
	pr, isRenderHTML := p.render(ssd, w, r, c, head)

	if len(pr.PageTitle) > 0 {
		head.Title(pr.PageTitle)
	}

	var schema = ssd.Schema
	body := bytes.NewBuffer(nil)

	if isRenderHTML {
		body.WriteString(schema.(string))
	} else {
		var serverSideDataJSON []byte
		serverSideDataJSON, err = json.MarshalIndent(ssd, "", "\t")
		if err != nil {
			panic(err)
		}
		err = h.Fprint(
			body,
			h.Script(fmt.Sprintf("window.__serverSideData__=%s\n", string(serverSideDataJSON))),
			c,
		)
		if err != nil {
			panic(err)
		}
	}

	var resp string
	resp, err = p.b.layoutFunc(r, head, body.String())
	if err != nil {
		panic(err)
	}
	_, err = fmt.Fprintln(w, resp)
	if err != nil {
		panic(err)
	}
}

func (p *PageBuilder) parseForm(r *http.Request) *multipart.Form {
	maxSize := p.maxFormSize
	if maxSize == 0 {
		maxSize = 128 << 20 // 128MB
	}

	err := r.ParseMultipartForm(maxSize)
	if err != nil {
		panic(err)
	}

	return r.MultipartForm
}

func (p *PageBuilder) eventBodyFromRequest(r *http.Request) *eventBody {
	var err error
	mf := p.parseForm(r)

	var eb eventBody
	err = json.NewDecoder(strings.NewReader(mf.Value["__event_data__"][0])).Decode(&eb)
	if err != nil {
		panic(err)
	}

	return &eb
}

func (p *PageBuilder) executeEvent(w http.ResponseWriter, r *http.Request) {
	// for server side restart and lost all the eventFuncs, but user keep clicking page without refresh page to call p.render to fill up eventFuncs

	ctx := new(EventContext)
	ctx.R = r
	ctx.W = w
	ctx.Hub = p
	ctx.Injector = &PageInjector{}

	c := WrapEventContext(r.Context(), ctx)

	eb := p.eventBodyFromRequest(r)
	ctx.Event = &eb.Event
	// because default added reload
	if len(p.eventFuncs) <= 1 && p.eventFuncById(eb.EventFuncID.ID) == nil {
		log.Println("Re-render because eventFuncs gone, might server restarted")
		ssd := &serverSideData{}
		head := &PageInjector{}
		p.render(ssd, w, r, c, head)
		_, err := json.Marshal(ssd) // to fill in event funcs that setup inside a component
		if err != nil {
			panic(err)
		}
	}

	ef := p.eventFuncById(eb.EventFuncID.ID)
	if ef == nil {
		panic(fmt.Errorf("event %s not found", eb.EventFuncID.ID))
	}

	eb.Event.Params = eb.EventFuncID.Params

	er, err := ef(ctx)
	if err != nil {
		panic(err)
	}

	if er.Reload {
		ssd := &serverSideData{}
		head := &PageInjector{}
		pr, _ := p.render(ssd, w, r, c, head)
		er.Schema = ssd.Schema
		if len(er.PageTitle) == 0 {
			er.PageTitle = pr.PageTitle
		}
	}

	eventResponseWithContext(ctx, c, &er)

	err = json.NewEncoder(w).Encode(er)
	if err != nil {
		panic(err)
	}
}

func reload(ctx *EventContext) (r EventResponse, err error) {
	r.Reload = true
	return
}

func eventResponseWithContext(ctx *EventContext, c context.Context, er *EventResponse) {
	if comp, ok := er.Schema.(SchemaComponent); ok {
		er.Schema = WithContext(ctx, comp)
	}

	if comp, ok := er.Schema.(h.HTMLComponent); ok {
		er.Schema = h.MustString(comp, c)
	}

	for _, up := range er.UpdatePortals {
		if comp, ok := up.Schema.(h.HTMLComponent); ok {
			up.Schema = h.MustString(comp, c)
		}
	}
}

func (p *PageBuilder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && strings.Index(r.URL.String(), "__execute_event__") >= 0 {
		p.executeEvent(w, r)
		return
	}
	p.index(w, r)
}
