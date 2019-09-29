package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	examples2 "github.com/goplaid/x/presets/examples"
	"github.com/jinzhu/gorm"
	"github.com/theplant/gofixtures"
)

var customerData = gofixtures.Data(gofixtures.Sql(`
				insert into customers (id, name) values (11, 'Felix1');
			`, []string{"customers"}))

var productData = gofixtures.Data(gofixtures.Sql(`
				insert into products (id, name) values (12, 'Product 1');
			`, []string{"products"}))

var emptyCustomerData = gofixtures.Data(gofixtures.Sql(``, []string{"customers"}))
var creditCardData = gofixtures.Data(customerData, gofixtures.Sql(``, []string{"credit_cards"}))

var cases = []struct {
	name               string
	reqFunc            func(db *gorm.DB) *http.Request
	eventResponseMatch func(er *testEventResponse, db *gorm.DB, t *testing.T)
	pageMatch          func(body *bytes.Buffer, db *gorm.DB, t *testing.T)
}{
	{
		name: "Update",
		reqFunc: func(db *gorm.DB) *http.Request {
			customerData.TruncatePut(db)
			r := httptest.NewRequest("POST", "/admin/customers?__execute_event__=update", strings.NewReader(`
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="__event_data__"

{"eventFuncId":{"id":"update","params":["11"],"pushState":null},"event":{}}
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="Bool1"

true
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="ID"

11
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="Int1"

42
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="Name"

Felix11
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8--
`))
			r.Header.Add("Content-Type", `multipart/form-data; boundary=----WebKitFormBoundaryOv2oq9YJ8tIG3xJ8`)
			return r
		},
		eventResponseMatch: func(er *testEventResponse, db *gorm.DB, t *testing.T) {
			var u = &examples2.Customer{}
			err := db.Find(u, 11).Error
			if err != nil {
				t.Error(err)
			}
			if u.Name != "Felix11" {
				t.Error(u)
			}
			return
		},
	},
	{
		name: "Create",
		reqFunc: func(db *gorm.DB) *http.Request {
			emptyCustomerData.TruncatePut(db)
			r := httptest.NewRequest("POST", "/admin/customers?__execute_event__=update", strings.NewReader(`
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="__event_data__"

{"eventFuncId":{"id":"update","params":[""],"pushState":null},"event":{}}
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="Bool1"

true
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="ID"

0
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="Int1"

42
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="Name"

Felix
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8--
`))
			r.Header.Add("Content-Type", `multipart/form-data; boundary=----WebKitFormBoundaryOv2oq9YJ8tIG3xJ8`)
			return r
		},
		eventResponseMatch: func(er *testEventResponse, db *gorm.DB, t *testing.T) {
			var u = &examples2.Customer{}
			err := db.First(u).Error
			if err != nil {
				t.Error(err)
			}
			if u.Name != "Felix" {
				t.Error(u)
			}
			return
		},
	},

	{
		name: "New Form For Creating",
		reqFunc: func(db *gorm.DB) *http.Request {
			emptyCustomerData.TruncatePut(db)
			r := httptest.NewRequest("POST", "/admin/credit-cards?__execute_event__=DrawerNew", strings.NewReader(`
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="__event_data__"

{"eventFuncId":{"id":"DrawerNew","params":[""],"pushState":null},"event":{}}
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
`))
			r.Header.Add("Content-Type", `multipart/form-data; boundary=----WebKitFormBoundaryOv2oq9YJ8tIG3xJ8`)
			return r
		},
		eventResponseMatch: func(er *testEventResponse, db *gorm.DB, t *testing.T) {
			partial := er.UpdatePortals[0].Body
			if strings.Index(partial, "field-name='Number'") < 0 {
				t.Error("can't find field-name='Number'", partial)
			}
			return
		},
	},

	{
		name: "Create CreditCard",
		reqFunc: func(db *gorm.DB) *http.Request {
			creditCardData.TruncatePut(db)
			r := httptest.NewRequest("POST", "/admin/credit-cards?__execute_event__=update", strings.NewReader(`
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="__event_data__"

{"eventFuncId":{"id":"update","params":["", "11"],"pushState":null},"event":{}}
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="Number"

12345678
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8--

`))
			r.Header.Add("Content-Type", `multipart/form-data; boundary=----WebKitFormBoundaryOv2oq9YJ8tIG3xJ8`)
			return r
		},
		eventResponseMatch: func(er *testEventResponse, db *gorm.DB, t *testing.T) {
			var u = &examples2.CreditCard{}
			err := db.First(u).Error
			if err != nil {
				t.Error(err)
			}
			if u.Number != "12345678" {
				t.Error(u)
			}

			return
		},
	},

	{
		name: "Without Editing Config/Product Edit Form",
		reqFunc: func(db *gorm.DB) *http.Request {
			productData.TruncatePut(db)
			r := httptest.NewRequest("POST", "/admin/products?__execute_event__=DrawerEdit", strings.NewReader(`
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="__event_data__"

{"eventFuncId":{"id":"DrawerEdit","params":["12"],"pushState":null},"event":{}}
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
`))
			r.Header.Add("Content-Type", `multipart/form-data; boundary=----WebKitFormBoundaryOv2oq9YJ8tIG3xJ8`)
			return r
		},
		eventResponseMatch: func(er *testEventResponse, db *gorm.DB, t *testing.T) {
			partial := er.UpdatePortals[0].Body
			if strings.Index(partial, "field-name='OwnerName'") < 0 {
				t.Error("can't find field-name='OwnerName'", partial)
			}
			return
		},
	},

	{
		name: "Without Editing Config/Create Product",
		reqFunc: func(db *gorm.DB) *http.Request {
			productData.TruncatePut(db)
			r := httptest.NewRequest("POST", "/admin/products?__execute_event__=update", strings.NewReader(`
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="__event_data__"

{"eventFuncId":{"id":"update","params":["12"],"pushState":null},"event":{}}
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="OwnerName"

owner1
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8--

`))
			r.Header.Add("Content-Type", `multipart/form-data; boundary=----WebKitFormBoundaryOv2oq9YJ8tIG3xJ8`)
			return r
		},
		eventResponseMatch: func(er *testEventResponse, db *gorm.DB, t *testing.T) {
			var u = &examples2.Product{}
			err := db.First(u).Error
			if err != nil {
				t.Error(err)
			}
			if u.OwnerName != "owner1" {
				t.Error(u)
			}

			return
		},
	},

	{
		name: "formDrawerAction AgreeTerms",
		reqFunc: func(db *gorm.DB) *http.Request {
			customerData.TruncatePut(db)
			r := httptest.NewRequest("POST", "/admin/customers/11?__execute_event__=DrawerAction", strings.NewReader(`
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="__event_data__"

{"eventFuncId":{"id":"DrawerAction","params":["AgreeTerms", "11"],"pushState":null},"event":{}}
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
`))
			r.Header.Add("Content-Type", `multipart/form-data; boundary=----WebKitFormBoundaryOv2oq9YJ8tIG3xJ8`)
			return r
		},
		eventResponseMatch: func(er *testEventResponse, db *gorm.DB, t *testing.T) {
			partial := er.UpdatePortals[0].Body
			if strings.Index(partial, "field-name='Agree'") < 0 {
				t.Error("can't find field-name='Agree'", partial)
			}
			return
		},
	},

	{
		name: "doAction AgreeTerms",
		reqFunc: func(db *gorm.DB) *http.Request {
			customerData.TruncatePut(db)
			r := httptest.NewRequest("POST", "/admin/customers/11?__execute_event__=doAction", strings.NewReader(`
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="__event_data__"

{"eventFuncId":{"id":"doAction","params":["AgreeTerms", "11"],"pushState":null},"event":{}}
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="Agree"

true
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
`))
			r.Header.Add("Content-Type", `multipart/form-data; boundary=----WebKitFormBoundaryOv2oq9YJ8tIG3xJ8`)
			return r
		},
		eventResponseMatch: func(er *testEventResponse, db *gorm.DB, t *testing.T) {
			var u = &examples2.Customer{}
			err := db.First(u).Error
			if err != nil {
				t.Error(err)
			}
			if u.TermAgreedAt == nil {
				t.Error(fmt.Sprintf("%#+v", u))
			}
			return
		},
	},
}

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", os.Getenv("DBString"))
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	return db
}

type testPortalUpdate struct {
	Name        string `json:"name,omitempty"`
	Body        string `json:"body,omitempty"`
	AfterLoaded string `json:"afterLoaded,omitempty"`
}

type testEventResponse struct {
	PageTitle     string              `json:"pageTitle,omitempty"`
	Body          string              `json:"body,omitempty"`
	Reload        bool                `json:"reload,omitempty"`
	ReloadPortals []string            `json:"reloadPortals,omitempty"`
	UpdatePortals []*testPortalUpdate `json:"updatePortals,omitempty"`
	Data          interface{}         `json:"data,omitempty"`
}

func TestAll(t *testing.T) {
	db := ConnectDB()
	p := examples2.Preset1(db)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := c.reqFunc(db)
			p.ServeHTTP(w, r)

			if c.eventResponseMatch != nil {
				var er testEventResponse
				err := json.NewDecoder(w.Body).Decode(&er)
				if err != nil {
					panic(err)
				}
				c.eventResponseMatch(&er, db, t)
			}

			if c.pageMatch != nil {
				c.pageMatch(w.Body, db, t)
			}
		})

		//diff := htmltestingutils.PrettyHtmlDiff(w.Body, "body", "abc")
		//if len(diff) > 0 {
		//	t.Error(diff)
		//}

	}
}
