package presets_test

import (
	"net/http/httptest"
	"testing"

	"github.com/sunfmin/bran/presets/examples"
	"github.com/theplant/htmltestingutils"
)

type A struct {
}

func TestHello(t *testing.T) {

	p := examples.Preset1()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/admin/users?keyword=hello", nil)
	p.ServeHTTP(w, r)
	//panic(w.Body.String())
	diff := htmltestingutils.PrettyHtmlDiff(w.Body, "body", "abc")
	if len(diff) > 0 {
		t.Error(diff)
	}
}
