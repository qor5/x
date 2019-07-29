package presets_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sunfmin/bran/ui"

	"github.com/sunfmin/bran/presets/examples"
)

var cases = []struct {
	reqFunc            func() *http.Request
	eventResponseMatch func(er *ui.EventResponse, t *testing.T)
	pageMatch          func(body *bytes.Buffer, t *testing.T)
}{
	{
		reqFunc: func() *http.Request {
			r := httptest.NewRequest("POST", "/admin/users?__execute_event__=update", strings.NewReader(`
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="__event_data__"

{"eventFuncId":{"id":"update","params":[""],"pushState":null},"event":{}}
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="Bool1"

false
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="CreatedAt"

0001-01-01T00:00:00Z
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="Float1"

0
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="ID"

0
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="JobTitle"


------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="UpdatedAt"

0001-01-01T00:00:00Z
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="Int1"

01
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="Name"


------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8
Content-Disposition: form-data; name="Name"

Felix
------WebKitFormBoundaryOv2oq9YJ8tIG3xJ8--
`))
			r.Header.Add("Content-Type", `multipart/form-data; boundary=----WebKitFormBoundaryOv2oq9YJ8tIG3xJ8`)
			return r
		},
		eventResponseMatch: func(er *ui.EventResponse, t *testing.T) {
			t.Error(er)
			return
		},
	},
}

func TestPresets(t *testing.T) {

	p := examples.Preset1()

	for _, c := range cases {
		w := httptest.NewRecorder()
		r := c.reqFunc()
		p.ServeHTTP(w, r)

		if c.eventResponseMatch != nil {
			var er ui.EventResponse
			err := json.NewDecoder(w.Body).Decode(&er)
			if err != nil {
				panic(err)
			}
			c.eventResponseMatch(&er, t)
		}

		if c.pageMatch != nil {
			c.pageMatch(w.Body, t)
		}

		//diff := htmltestingutils.PrettyHtmlDiff(w.Body, "body", "abc")
		//if len(diff) > 0 {
		//	t.Error(diff)
		//}

	}
}
