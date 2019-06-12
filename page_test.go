package bran_test

import (
	"bytes"
	"context"
	"fmt"
	goji "goji.io"
	"goji.io/pat"
	"io/ioutil"
	"mime/multipart"
	"net/http/httptest"
	"testing"

	"github.com/sunfmin/bran"
	"github.com/sunfmin/bran/ui"
	h "github.com/theplant/htmlgo"
	"github.com/theplant/htmltestingutils"
	"github.com/theplant/testingutils"
)

type User struct {
	Name    string
	Address *Address
}

type Address struct {
	Zipcode string
	City    string
}

var userData = &User{
	Name:    "Felix",
	Address: &Address{"123123", "Hangzhou"},
}

var userZero *User
var userZero2 ****User

var zeroBody = `
{
	"schema": {}
}
`

var userBody = `
{
	"schema": {},
	"states": {
		"Address.City": [
			"Hangzhou"
		],
		"Address.Zipcode": [
			"123123"
		],
		"Name": [
			"Felix"
		]
	}
}
`
var pageStateCases = []struct {
	name       string
	state      interface{}
	schema     ui.Component
	body       string
	renderHTML bool
}{
	{
		name:  "empty",
		state: nil,
		body:  zeroBody,
	},
	{
		name:  "zero",
		state: userZero,
		body:  zeroBody,
	},
	{
		name:  "zero 2",
		state: userZero2,
		body:  zeroBody,
	},
	{
		name:  "valid 1",
		state: User{Name: "Felix", Address: &Address{"123123", "Hangzhou"}},
		body:  userBody,
	},
	{
		name:  "valid 2",
		state: userData,
		body:  userBody,
	},
	{
		name:  "valid 3",
		state: &userData,
		body:  userBody,
	},
	{
		name:   "html",
		state:  &userData,
		schema: ui.RawSchema("{}"),
		body: `
{
	"schema": {},
	"states": {
		"Address.City": [
			"Hangzhou"
		],
		"Address.Zipcode": [
			"123123"
		],
		"Name": [
			"Felix"
		]
	}
}
		`,
	},
	{
		name:       "html component",
		state:      &userData,
		schema:     h.RawHTML("<h1>Hello</h1>"),
		renderHTML: true,
		body: `<!DOCTYPE html>
<html>
<head>
<meta charset="utf8"/>
<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no"/>
</head>
<body class='front'>
<div id="app">
<h1>Hello</h1></div>
<script type='text/javascript'>
window.__serverSideData__={
	"states": {
		"Address.City": [
			"Hangzhou"
		],
		"Address.Zipcode": [
			"123123"
		],
		"Name": [
			"Felix"
		]
	}
}
</script>

</body>
</html>

`,
	},
}

func TestPageState(t *testing.T) {
	pb := bran.New()

	for _, c := range pageStateCases {
		p := pb.Page(func(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
			ctx.State = c.state
			pr.Schema = ui.RawSchema("{}")
			if c.schema != nil {
				pr.Schema = c.schema
			}
			pr.JSONOnly = !c.renderHTML
			return
		})

		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)
		p.ServeHTTP(w, r)

		diff := htmltestingutils.PrettyHtmlDiff(w.Body, "*", c.body)
		if len(diff) > 0 {
			t.Error(c.name, diff)
		}
	}
}

func runEvent(
	eventFunc ui.EventFunc,
	renderChanger func(ctx *ui.EventContext, pr *ui.PageResponse),
	eventFormChanger func(mw *multipart.Writer),
) (indexResp *bytes.Buffer, eventResp *bytes.Buffer) {
	pb := bran.New()

	var f = func(ctx *ui.EventContext) (r ui.EventResponse, err error) {
		r.Reload = true
		return
	}

	if eventFunc != nil {
		f = eventFunc
	}

	var p = pb.Page(func(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
		ctx.Hub.RefEventFunc("call", f)

		if renderChanger != nil {
			renderChanger(ctx, &pr)
		} else {
			ctx.StateOrInit(&User{})
			pr.Schema = ui.RawSchema("{}")
			pr.JSONOnly = true
		}
		return
	})

	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	p.ServeHTTP(w, r)

	indexResp = w.Body

	body := bytes.NewBuffer(nil)

	mw := multipart.NewWriter(body)
	mw.WriteField("__event_data__", `{"eventFuncId":{"id":"call","pushState":null},"event":{"value":""}}
	`)

	if eventFormChanger != nil {
		eventFormChanger(mw)
	}

	mw.Close()

	r = httptest.NewRequest("POST", "/?__execute_event__=call", body)
	r.Header.Add("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%s", mw.Boundary()))

	w = httptest.NewRecorder()
	p.ServeHTTP(w, r)

	eventResp = w.Body
	return
}

func TestPageStateInitAndSet(t *testing.T) {

	var login = func(ctx *ui.EventContext) (r ui.EventResponse, err error) {
		add := ctx.SubStateOrInit("Address", &Address{}).(*Address)
		add.City = "hz"

		r.Reload = true
		return
	}

	indexResp, eventResp := runEvent(login, nil, nil)

	diff := testingutils.PrettyJsonDiff(`
{
	"schema": {},
	"states": {
		"Name": [
			""
		]
	}
}
	`, indexResp.String())
	if len(diff) > 0 {
		t.Error(diff)
	}

	diff = testingutils.PrettyJsonDiff(`
{
	"schema": {},
	"states": {
		"Address.City": [
			"hz"
		],
		"Address.Zipcode": [
			""
		],
		"Name": [
			""
		]
	},
	"reload": true
}
	`, eventResp.String())
	if len(diff) > 0 {
		t.Error(diff)
	}
}

func TestFileUpload(t *testing.T) {
	type mystate struct {
		File1 []*multipart.FileHeader `form:"-"`
	}

	var uploadFile = func(ctx *ui.EventContext) (r ui.EventResponse, err error) {
		r.Reload = true
		return
	}

	pb := bran.New()
	p := pb.Page(func(ctx *ui.EventContext) (pr ui.PageResponse, err error) {

		s := ctx.StateOrInit(&mystate{}).(*mystate)

		var data []byte
		if len(s.File1) > 0 {
			var mf multipart.File
			mf, err = s.File1[0].Open()
			if err != nil {
				panic(err)
			}
			data, err = ioutil.ReadAll(mf)
			if err != nil {
				panic(err)
			}
		}

		ctx.Hub.RefEventFunc("uploadFile", uploadFile)

		pr.Schema = ui.RawSchema(fmt.Sprintf(`{"__text__": "%s"}`, string(data)))
		pr.JSONOnly = true
		return
	})

	body := bytes.NewBuffer(nil)

	mw := multipart.NewWriter(body)
	mw.WriteField("__event_data__", `{"eventFuncId":{"id":"uploadFile","pushState":null},"event":{"value":""}}
	`)
	fw, _ := mw.CreateFormFile("File1", "myfile.txt")
	fw.Write([]byte("Hello"))

	mw.Close()

	r := httptest.NewRequest("POST", "/?__execute_event__=uploadFile", body)
	r.Header.Add("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%s", mw.Boundary()))

	w := httptest.NewRecorder()
	p.ServeHTTP(w, r)

	diff := testingutils.PrettyJsonDiff(`
{
	"schema": {
		"__text__": "Hello"
	},
	"states": {},
	"reload": true
}
	`, w.Body.String())
	if len(diff) > 0 {
		t.Error(diff)
	}
}

type DummyComp struct {
}

func (dc *DummyComp) MarshalHTML(ctx context.Context) (r []byte, err error) {
	r = []byte("<div>hello</div>")
	ui.Injector(ctx).PutScript(`
	function hello() {
		console.log("hello")
	}
`)

	ui.Injector(ctx).PutStyle(`
	div {
		background-color: red;
	}
`)
	return
}

var eventCases = []struct {
	name              string
	eventFunc         ui.EventFunc
	renderChanger     func(ctx *ui.EventContext, pr *ui.PageResponse)
	eventFormChanger  func(mw *multipart.Writer)
	expectedIndexResp string
	expectedEventResp string
}{
	{
		name: "run event reload states",
		renderChanger: func(ctx *ui.EventContext, pr *ui.PageResponse) {
			s := ctx.StateOrInit(&User{Address: &Address{}}).(*User)
			pr.Schema = h.Text(s.Name + " " + s.Address.City)
			s.Name = "Felix"
		},
		eventFunc: func(ctx *ui.EventContext) (r ui.EventResponse, err error) {
			s := ctx.State.(*User)
			r.Reload = true
			s.Name = "Felix1"
			s.Address = &Address{City: "Hangzhou"}
			return
		},
		expectedEventResp: `{
	"schema": "Felix1 Hangzhou",
	"states": {
		"Address.City": [
			"Hangzhou"
		],
		"Address.Zipcode": [
			""
		],
		"Name": [
			"Felix"
		]
	},
	"reload": true
}
`,
	},
	{
		name: "render schema in event func",
		eventFunc: func(ctx *ui.EventContext) (r ui.EventResponse, err error) {
			r.Schema = h.Div(
				h.H1("hello"),
			)
			return
		},
		expectedEventResp: `
	{
		"schema": "\n\u003cdiv\u003e\n\u003ch1\u003ehello\u003c/h1\u003e\n\u003c/div\u003e\n",
		"states": {
			"Name": [
				""
			]
		}
	}
			`,
	},

	{
		name: "case 1",
		renderChanger: func(ctx *ui.EventContext, pr *ui.PageResponse) {
			pr.Schema = h.RawHTML("<h1>Hello</h1>")
		},
		expectedEventResp: `
	{
		"schema": "\u003ch1\u003eHello\u003c/h1\u003e",
		"reload": true
	}
			`,
	},
	{
		name: "case 2",
		renderChanger: func(ctx *ui.EventContext, pr *ui.PageResponse) {
			ctx.Injector.PutTailHTML("<script src='/assets/main.js'></script>")
			pr.Schema = &DummyComp{}
		},
		expectedEventResp: `
{
	"schema": "\u003cdiv\u003ehello\u003c/div\u003e",
	"reload": true,
	"scripts": "\n\tfunction hello() {\n\t\tconsole.log(\"hello\")\n\t}\n",
	"styles": "\n\tdiv {\n\t\tbackground-color: red;\n\t}\n"
}
`,
		expectedIndexResp: `<!DOCTYPE html>
<html>
<head>
<meta charset="utf8"/>
<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no"/>
</head>
<body class='front'>
<style id="main_styles" type="text/css">

	div {
		background-color: red;
	}

</style>
<div id="app">
<div>hello</div></div>
<script type='text/javascript'>
window.__serverSideData__={}
</script>
<script id="main_scripts">

	function hello() {
		console.log("hello")
	}

</script>
<script src='/assets/main.js'></script>
</body>
</html>

`,
	},
}

func TestEvents(t *testing.T) {
	for _, c := range eventCases {
		indexResp, eventResp := runEvent(c.eventFunc, c.renderChanger, c.eventFormChanger)
		var diff string
		if len(c.expectedIndexResp) > 0 {
			diff = htmltestingutils.PrettyHtmlDiff(indexResp, "*", c.expectedIndexResp)

			if len(diff) > 0 {
				t.Error(c.name, diff)
			}
		}

		if len(c.expectedEventResp) > 0 {
			diff = testingutils.PrettyJsonDiff(c.expectedEventResp, eventResp.String())
			if len(diff) > 0 {
				t.Error(c.name, diff)
			}
		}
	}
}

var mountCases = []struct {
	name     string
	method   string
	path     string
	bodyFunc func(w *multipart.Writer)
	expected string
}{
	{
		name:     "with param get",
		method:   "GET",
		path:     "/home/topics/xgb123",
		bodyFunc: nil,
		expected: `<div><a href="#" v-on:click='onclick({"id":"bookmark","pushState":null}, $event)'>xgb123</a></div>`,
	},
	{
		name:   "with param post",
		method: "POST",
		path:   "/home/topics/xgb123?__execute_event__",
		bodyFunc: func(w *multipart.Writer) {
			w.WriteField("__event_data__", `{"eventFuncId":{"id":"bookmark","pushState":null},"event":{"value":""}}`)
		},
		expected: `{"schema":"\n\u003ch1\u003exgb123 bookmarked\u003c/h1\u003e\n"}`,
	},
}

func TestMultiplePagesAndEvents(t *testing.T) {
	var topicIndex = func(ctx *ui.EventContext) (r ui.PageResponse, err error) {
		r.Schema = h.H1("Hello Topic List")
		return
	}

	var bookmark = func(ctx *ui.EventContext) (r ui.EventResponse, err error) {
		topicId := pat.Param(ctx.R, "topicID")
		r.Schema = h.H1(topicId + " bookmarked")
		return
	}

	var topicDetail = func(ctx *ui.EventContext) (r ui.PageResponse, err error) {
		topicId := pat.Param(ctx.R, "topicID")
		r.Schema = h.Div(
			ui.Bind(h.A().Href("#").Text(topicId)).
				OnClick(ctx.Hub, "bookmark", bookmark),
		)
		return
	}

	pb := bran.New()

	mux := goji.NewMux()
	mux.Handle(pat.New("/home/topics/:topicID"), pb.Page(topicDetail))
	mux.Handle(pat.New("/home/topics"), pb.Page(topicIndex))

	for _, c := range mountCases {

		buf := new(bytes.Buffer)
		var mw *multipart.Writer
		if c.bodyFunc != nil {
			mw = multipart.NewWriter(buf)
			c.bodyFunc(mw)
			_ = mw.Close()
		}

		r := httptest.NewRequest(c.method, c.path, buf)
		if mw != nil {
			r.Header.Add("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%s", mw.Boundary()))
		}
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, r)
		selector := "#app div"
		if mw != nil {
			selector = "*"
		}
		diff := htmltestingutils.PrettyHtmlDiff(w.Body, selector, c.expected)
		if len(diff) > 0 {
			t.Error(c.name, diff)
		}

	}

}
