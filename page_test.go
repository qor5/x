package page_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http/httptest"
	"testing"

	"github.com/theplant/testingutils"

	ui "github.com/sunfmin/page"
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
		schema: ui.StringComponentJSON("{}"),
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
		schema:     ui.StringComponentHTML("<h1>Hello</h1>"),
		renderHTML: true,
		body: `

<!DOCTYPE html>
<html>
	<head>
	<meta charset="utf8"/>
<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no"/>

	</head>
	<body class="front">
	<h1>Hello</h1>


<script type="text/javascript">

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
	<script type="text/javascript" src="/main.js"></script>
	</body>
</html>
`,
	},
}

func TestPageState(t *testing.T) {
	pb := ui.New().NewPage()

	for _, c := range pageStateCases {
		pb.RenderFunc(func(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
			ctx.State = c.state
			pr.Schema = ui.StringComponentJSON("{}")
			if c.schema != nil {
				pr.Schema = c.schema
			}
			pr.JSONOnly = !c.renderHTML
			return
		})
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)
		pb.Handler().ServeHTTP(w, r)

		diff := testingutils.PrettyJsonDiff(c.body, w.Body.String())
		if len(diff) > 0 {
			t.Error(c.name, diff)
		}
	}
}

func TestPageStateInitAndSet(t *testing.T) {
	pb := ui.New().NewPage()

	var login = func(ctx *ui.EventContext) (r ui.EventResponse, err error) {
		add := ctx.SubStateOrInit("Address", &Address{}).(*Address)
		add.City = "hz"

		r.Reload = true
		return
	}

	pb.RenderFunc(func(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
		ctx.StateOrInit(&User{})

		ctx.Hub.RefEventFunc("login", login)

		pr.Schema = ui.StringComponentJSON("{}")
		pr.JSONOnly = true
		return
	})

	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	pb.Handler().ServeHTTP(w, r)

	diff := testingutils.PrettyJsonDiff(`
{
	"schema": {},
	"states": {
		"Name": [
			""
		]
	}
}
	`, w.Body.String())
	if len(diff) > 0 {
		t.Error(diff)
	}

	body := bytes.NewBuffer(nil)

	mw := multipart.NewWriter(body)
	mw.WriteField("__event_data__", `{"eventFuncId":{"id":"login","pushState":null},"event":{"value":""}}
	`)
	mw.Close()

	r = httptest.NewRequest("POST", "/__execute_event__/login", body)
	r.Header.Add("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%s", mw.Boundary()))

	w = httptest.NewRecorder()
	pb.Handler().ServeHTTP(w, r)

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
	`, w.Body.String())
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

	pb := ui.New().NewPage()
	pb.RenderFunc(func(ctx *ui.EventContext) (pr ui.PageResponse, err error) {

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

		pr.Schema = ui.StringComponentJSON(fmt.Sprintf(`{"__text__": "%s"}`, string(data)))
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

	r := httptest.NewRequest("POST", "/__execute_event__/uploadFile", body)
	r.Header.Add("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%s", mw.Boundary()))

	w := httptest.NewRecorder()
	pb.Handler().ServeHTTP(w, r)

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
