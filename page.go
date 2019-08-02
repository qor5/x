package bran

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/sunfmin/bran/ui"
	h "github.com/theplant/htmlgo"
)

type PageBuilder struct {
	b              *Builder
	eventFuncRefs  map[string]ui.EventFunc
	pageRenderFunc ui.PageFunc
	maxFormSize    int64
}

func (b *Builder) Page(pslf ui.PageFunc) (p *PageBuilder) {
	p = &PageBuilder{}
	p.eventFuncRefs = make(map[string]ui.EventFunc)
	p.b = b
	p.pageRenderFunc = pslf

	return
}

func (p *PageBuilder) MaxFormSize(v int64) (r *PageBuilder) {
	p.maxFormSize = v
	r = p
	return
}

func (p *PageBuilder) RegisterEventFunc(eventFuncId string, ef ui.EventFunc) (key string) {
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
	EventFuncID ui.EventFuncID `json:"eventFuncId,omitempty"`
	Event       ui.Event       `json:"event,omitempty"`
	PageState   ui.PageState   `json:"pageState,omitempty"`
}

type serverSideData struct {
	Schema interface{} `json:"schema,omitempty"`
}

func (p *PageBuilder) render(
	ssd *serverSideData,
	w http.ResponseWriter,
	r *http.Request,
	c context.Context,
	head *DefaultPageInjector,
) (pager *ui.PageResponse, isRenderHTML bool) {

	if p.pageRenderFunc == nil {
		return
	}

	ctx := ui.MustGetEventContext(c)

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

	// fmt.Println("eventFuncRefs count: ", len(p.eventFuncRefs))
	if comp, ok := pager.Schema.(ui.SchemaComponent); ok {
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
	var head = &DefaultPageInjector{}

	ctx := new(ui.EventContext)
	c := ui.WrapEventContext(r.Context(), ctx)
	pr, isRenderHTML := p.render(ssd, w, r, c, head)

	var schema = ssd.Schema

	if isRenderHTML && !pr.JSONOnly {
		ssd.Schema = nil
	}

	var serverSideDataJSON []byte
	serverSideDataJSON, err = json.MarshalIndent(ssd, "", "\t")
	if err != nil {
		panic(err)
	}

	if pr.JSONOnly {
		_, err = fmt.Fprintln(w, string(serverSideDataJSON))
		if err != nil {
			panic(err)
		}
		return
	}

	body := bytes.NewBuffer(nil)

	//body.WriteString(head.MainStyles(true))

	if isRenderHTML {
		err = h.Fprint(
			body,
			h.Div(
				h.RawHTML(schema.(string)),
			).Id("app").Attr("v-cloak", true),
			c,
		)
		if err != nil {
			panic(err)
		}
	}

	err = h.Fprint(
		body,
		h.Script(fmt.Sprintf("window.__serverSideData__=%s\n", string(serverSideDataJSON))),
		c,
	)
	if err != nil {
		panic(err)
	}

	//body.WriteString(head.MainScripts(true))
	body.WriteString(head.TailHTML())
	body.WriteString("\n")

	var resp string
	resp, err = p.b.getLayoutMiddleFunc()(nil, head)(r, body.String())
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

	ctx := new(ui.EventContext)
	ctx.R = r
	ctx.W = w
	ctx.Hub = p
	ctx.Injector = &DefaultPageInjector{}

	c := ui.WrapEventContext(r.Context(), ctx)

	if len(p.eventFuncRefs) == 0 {
		log.Println("Re-render because eventFuncs gone, might server restarted")
		ssd := &serverSideData{}
		head := &DefaultPageInjector{}
		p.render(ssd, w, r, c, head)
		_, err := json.Marshal(ssd) // to fill in event funcs that setup inside a component
		if err != nil {
			panic(err)
		}
	}

	eb := p.eventBodyFromRequest(r)
	ctx.Event = &eb.Event

	ef, ok := p.eventFuncRefs[eb.EventFuncID.ID]
	if !ok {
		panic(fmt.Errorf("event %s not found", eb.EventFuncID.ID))
	}

	eb.Event.Params = eb.EventFuncID.Params

	er, err := ef(ctx)
	if err != nil {
		panic(err)
	}

	if er.Reload {
		// panic(fmt.Sprintf("er.State %#+v", er.State))
		ssd := &serverSideData{}
		head := &DefaultPageInjector{}
		p.render(ssd, w, r, c, head)
		er.Schema = ssd.Schema

	}

	eventResponseWithContext(ctx, c, &er)

	err = json.NewEncoder(w).Encode(er)
	if err != nil {
		panic(err)
	}
}

func eventResponseWithContext(ctx *ui.EventContext, c context.Context, er *ui.EventResponse) {
	if comp, ok := er.Alert.(ui.SchemaComponent); ok {
		er.Alert = ui.WithContext(ctx, comp)
	}
	if comp, ok := er.Confirm.(ui.SchemaComponent); ok {
		er.Confirm = ui.WithContext(ctx, comp)
	}
	if comp, ok := er.Schema.(ui.SchemaComponent); ok {
		er.Schema = ui.WithContext(ctx, comp)
	}
	if comp, ok := er.Dialog.(ui.SchemaComponent); ok {
		er.Dialog = ui.WithContext(ctx, comp)
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
