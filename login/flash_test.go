package login

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

// Input flashes are stored in client-side cookies, so they must never carry
// password values (VULN-18).
func TestInputFlashesDoNotCarryPasswords(t *testing.T) {
	for _, typ := range []reflect.Type{
		reflect.TypeOf(WrongLoginInputFlash{}),
		reflect.TypeOf(WrongForgetPasswordInputFlash{}),
		reflect.TypeOf(WrongResetPasswordInputFlash{}),
		reflect.TypeOf(WrongChangePasswordInputFlash{}),
	} {
		for i := 0; i < typ.NumField(); i++ {
			if f := typ.Field(i); strings.Contains(strings.ToLower(f.Name), "password") {
				t.Errorf("%s.%s must not be stored in a flash cookie", typ.Name(), f.Name)
			}
		}
	}
}

func TestWrongLoginInputFlashCookie(t *testing.T) {
	b := New().Secret("test-secret")
	w := httptest.NewRecorder()
	b.setWrongLoginInputFlash(w, WrongLoginInputFlash{Account: "alice"})

	var ck *http.Cookie
	for _, c := range w.Result().Cookies() {
		if c.Name == wrongLoginInputFlashCookieName {
			ck = c
		}
	}
	if ck == nil {
		t.Fatal("flash cookie not set")
	}
	raw, err := base64.StdEncoding.DecodeString(ck.Value)
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	if err := json.Unmarshal(raw, &m); err != nil {
		t.Fatal(err)
	}
	if len(m) != 1 || m["Account"] != "alice" {
		t.Fatalf("unexpected flash payload: %s", raw)
	}

	// The account is still refilled from the flash.
	r := httptest.NewRequest(http.MethodGet, "/auth/login", nil)
	r.AddCookie(ck)
	got := b.ViewHelper().GetWrongLoginInputFlash(httptest.NewRecorder(), r)
	if got.Account != "alice" {
		t.Fatalf("expected account to be refilled, got %+v", got)
	}
}
