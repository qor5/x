package presets_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/theplant/gofixtures"

	"github.com/jinzhu/gorm"

	"github.com/sunfmin/bran/ui"

	"github.com/sunfmin/bran/presets/examples"
)

var customerData = gofixtures.Data(gofixtures.Sql(`
				insert into customers (id, name) values (11, 'Felix1');
			`, []string{"customers"}))
var emptyCustomerData = gofixtures.Data(gofixtures.Sql(``, []string{"customers"}))
var creditCardData = gofixtures.Data(customerData, gofixtures.Sql(``, []string{"credit_cards"}))

var cases = []struct {
	name               string
	reqFunc            func(db *gorm.DB) *http.Request
	eventResponseMatch func(er *ui.EventResponse, db *gorm.DB, t *testing.T)
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
		eventResponseMatch: func(er *ui.EventResponse, db *gorm.DB, t *testing.T) {
			var u = &examples.Customer{}
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
		eventResponseMatch: func(er *ui.EventResponse, db *gorm.DB, t *testing.T) {
			var u = &examples.Customer{}
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
		name: "New Form For CloneForCreating",
		reqFunc: func(db *gorm.DB) *http.Request {
			emptyCustomerData.TruncatePut(db)
			r := httptest.NewRequest("POST", "/admin/credit-cards?__execute_event__=formDrawerNew", strings.NewReader(`
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="__event_data__"

{"eventFuncId":{"id":"formDrawerNew","params":[""],"pushState":null},"event":{}}
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
`))
			r.Header.Add("Content-Type", `multipart/form-data; boundary=----WebKitFormBoundaryOv2oq9YJ8tIG3xJ8`)
			return r
		},
		eventResponseMatch: func(er *ui.EventResponse, db *gorm.DB, t *testing.T) {
			partial := er.UpdatePortals[0].Schema.(string)
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
		eventResponseMatch: func(er *ui.EventResponse, db *gorm.DB, t *testing.T) {
			var u = &examples.CreditCard{}
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
}

func TestPresets(t *testing.T) {
	db, err := gorm.Open("postgres", os.Getenv("TEST_DB"))
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	p := examples.Preset1(db)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := c.reqFunc(db)
			p.ServeHTTP(w, r)

			if c.eventResponseMatch != nil {
				var er ui.EventResponse
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
