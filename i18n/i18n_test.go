package i18n_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goplaid/x/i18n"
	"golang.org/x/text/language"
)

type Messages struct {
	Update string
}

var Messages_zh_CN = &Messages{
	Update: "更新",
}

var Messages_en_US = &Messages{
	Update: "Update",
}

func TestLanguage(t *testing.T) {
	var mediaLibraryKey i18n.ModuleKey = "mediaLibraryKey"

	b := i18n.New().
		SupportLanguages(language.English, language.Japanese, language.SimplifiedChinese).
		RegisterForModule(language.SimplifiedChinese, mediaLibraryKey, Messages_zh_CN).
		RegisterForModule(language.English, mediaLibraryKey, Messages_en_US)

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := i18n.MustGetModuleMessages(r, mediaLibraryKey, Messages_en_US).(*Messages)
		_, _ = fmt.Fprint(w, msg.Update)
	})

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/?lang=zh", nil)
	b.EnsureLanguage(h).ServeHTTP(recorder, req)

	if recorder.Body.String() != "更新" {
		t.Errorf("response is wrong, %s", recorder.Body.String())
	}
	if len(recorder.Header().Get("Set-Cookie")) == 0 {
		t.Error("cookie not set")
	}

	recorder = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{Name: "lang", Value: "zh-Hans"})
	b.EnsureLanguage(h).ServeHTTP(recorder, req)

	if recorder.Body.String() != "更新" {
		t.Errorf("response is wrong, %s", recorder.Body.String())
	}

	recorder = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/", nil)
	req.Header.Add("Accept-Language", "zh")
	b.EnsureLanguage(h).ServeHTTP(recorder, req)

	if recorder.Body.String() != "更新" {
		t.Errorf("response is wrong, %s", recorder.Body.String())
	}

	recorder = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/", nil)
	b.EnsureLanguage(h).ServeHTTP(recorder, req)

	if recorder.Body.String() != "Update" {
		t.Errorf("response is wrong, %s", recorder.Body.String())
	}

}
