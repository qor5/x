package login

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

// Input flashes are stored in client-side cookies, so password fields (kept
// only for source compatibility) must never be serialized (VULN-18).
func TestInputFlashPasswordFieldsAreNotSerialized(t *testing.T) {
	for _, typ := range []reflect.Type{
		reflect.TypeOf(WrongLoginInputFlash{}),
		reflect.TypeOf(WrongForgetPasswordInputFlash{}),
		reflect.TypeOf(WrongResetPasswordInputFlash{}),
		reflect.TypeOf(WrongChangePasswordInputFlash{}),
	} {
		for i := 0; i < typ.NumField(); i++ {
			f := typ.Field(i)
			if strings.Contains(strings.ToLower(f.Name), "password") && f.Tag.Get("json") != "-" {
				t.Errorf("%s.%s must be tagged json:\"-\"", typ.Name(), f.Name)
			}
		}
	}
}

func flashCookieValue(t *testing.T, w *httptest.ResponseRecorder, name string) (*http.Cookie, string) {
	t.Helper()
	for _, c := range w.Result().Cookies() {
		if c.Name == name {
			raw, err := base64.StdEncoding.DecodeString(c.Value)
			if err != nil {
				t.Fatal(err)
			}
			return c, string(raw)
		}
	}
	t.Fatalf("flash cookie %s not set", name)
	return nil, ""
}

func TestInputFlashCookiesNeverContainPasswords(t *testing.T) {
	b := New().Secret("test-secret")
	const secret = "s3cr3t-P4ss"

	// Even if a caller sets the deprecated password fields, they must not
	// reach the cookie.
	w := httptest.NewRecorder()
	b.setWrongLoginInputFlash(w, WrongLoginInputFlash{Account: "alice", Password: secret})
	ck, payload := flashCookieValue(t, w, wrongLoginInputFlashCookieName)
	if payload != `{"Account":"alice"}` {
		t.Fatalf("unexpected login flash payload: %s", payload)
	}

	// Account is still refilled; password is not.
	r := httptest.NewRequest(http.MethodGet, "/auth/login", nil)
	r.AddCookie(ck)
	got := b.ViewHelper().GetWrongLoginInputFlash(httptest.NewRecorder(), r)
	if got.Account != "alice" || got.Password != "" {
		t.Fatalf("unexpected refill: %+v", got)
	}

	w = httptest.NewRecorder()
	b.setWrongResetPasswordInputFlash(w, WrongResetPasswordInputFlash{Password: secret, ConfirmPassword: secret, TOTP: "123456"})
	_, payload = flashCookieValue(t, w, wrongResetPasswordInputFlashCookieName)
	if payload != `{"TOTP":"123456"}` {
		t.Fatalf("unexpected reset password flash payload: %s", payload)
	}

	w = httptest.NewRecorder()
	b.setWrongChangePasswordInputFlash(w, WrongChangePasswordInputFlash{OldPassword: secret, NewPassword: secret, ConfirmPassword: secret, TOTP: "123456"})
	_, payload = flashCookieValue(t, w, wrongChangePasswordInputFlashCookieName)
	if payload != `{"TOTP":"123456"}` {
		t.Fatalf("unexpected change password flash payload: %s", payload)
	}
}

// A legacy cookie written by an older version must not refill passwords.
func TestLegacyInputFlashCookieDoesNotRefillPassword(t *testing.T) {
	b := New().Secret("test-secret")
	r := httptest.NewRequest(http.MethodGet, "/auth/login", nil)
	r.AddCookie(&http.Cookie{
		Name:  wrongLoginInputFlashCookieName,
		Value: base64.StdEncoding.EncodeToString([]byte(`{"Account":"alice","Password":"old"}`)),
	})
	got := b.ViewHelper().GetWrongLoginInputFlash(httptest.NewRecorder(), r)
	if got.Account != "alice" || got.Password != "" {
		t.Fatalf("unexpected refill: %+v", got)
	}
}

// A cookie value is ASCII only, so a notice has to survive the round trip
// through one whatever language it is written in.
func TestNoticeFlashKeepsNonASCIIMessages(t *testing.T) {
	vh := New().Secret("test-secret").ViewHelper()

	for _, msg := range []string{
		"Password must contain at least one letter, one number and one symbol",
		"パスワードには英字・数字・記号をそれぞれ1文字以上含めてください",
		`quotes ", semicolons ; and # signs`,
	} {
		w := httptest.NewRecorder()
		setNoticeFlash(w, &NoticeError{Level: NoticeLevel_Error, Message: msg})

		r := httptest.NewRequest(http.MethodGet, "/auth/reset-password", nil)
		for _, c := range w.Result().Cookies() {
			r.AddCookie(c)
		}
		got := vh.GetNoticeFlash(httptest.NewRecorder(), r)
		if got == nil || got.Level != NoticeLevel_Error || got.Message != msg {
			t.Errorf("notice %q came back as %+v", msg, got)
		}
	}
}

// A notice set before messages were encoded is still shown.
func TestNoticeFlashReadsUnencodedCookie(t *testing.T) {
	vh := New().Secret("test-secret").ViewHelper()

	r := httptest.NewRequest(http.MethodGet, "/auth/login", nil)
	r.AddCookie(&http.Cookie{Name: noticeFlashCookieName, Value: "2#User locked. Try again later."})
	got := vh.GetNoticeFlash(httptest.NewRecorder(), r)
	if got == nil || got.Message != "User locked. Try again later." {
		t.Errorf("unexpected notice: %+v", got)
	}
}
