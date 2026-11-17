package login

import (
	"errors"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestUserPassLoginFailCode(t *testing.T) {
	other := errors.New("db down")
	for _, c := range []struct {
		name     string
		unified  bool
		err      error
		expect   FailCode
		expectOK bool
	}{
		{name: "default not found", err: ErrUserNotFound, expect: FailCodeIncorrectAccountNameOrPassword, expectOK: true},
		{name: "default wrong password", err: ErrWrongPassword, expect: FailCodeIncorrectAccountNameOrPassword, expectOK: true},
		{name: "default locked", err: ErrUserLocked, expect: FailCodeUserLocked, expectOK: true},
		{name: "default get locked", err: ErrUserGetLocked, expect: FailCodeUserLocked, expectOK: true},
		{name: "default other", err: other, expectOK: false},
		{name: "unified not found", unified: true, err: ErrUserNotFound, expect: FailCodeIncorrectAccountNameOrPassword, expectOK: true},
		{name: "unified wrong password", unified: true, err: ErrWrongPassword, expect: FailCodeIncorrectAccountNameOrPassword, expectOK: true},
		{name: "unified locked", unified: true, err: ErrUserLocked, expect: FailCodeIncorrectAccountNameOrPassword, expectOK: true},
		{name: "unified get locked", unified: true, err: ErrUserGetLocked, expect: FailCodeIncorrectAccountNameOrPassword, expectOK: true},
		{name: "unified other", unified: true, err: other, expectOK: false},
	} {
		t.Run(c.name, func(t *testing.T) {
			b := New().Secret("test-secret").UnifiedLoginFailure(c.unified)
			code, ok := b.userPassLoginFailCode(c.err)
			if ok != c.expectOK || (ok && code != c.expect) {
				t.Fatalf("userPassLoginFailCode(%v) = (%v, %v), want (%v, %v)", c.err, code, ok, c.expect, c.expectOK)
			}
		})
	}
}

func TestDummyPasswordHash(t *testing.T) {
	h := dummyPasswordHash()
	cost, err := bcrypt.Cost(h)
	if err != nil {
		t.Fatal(err)
	}
	// Must match the cost used by UserPass.EncryptPassword so that a
	// not-found account takes as long as a wrong password.
	if cost != 10 {
		t.Fatalf("dummy hash cost = %d, want 10", cost)
	}
	if &dummyPasswordHash()[0] != &h[0] {
		t.Fatal("dummy hash must be generated once")
	}
	if compareDummyPassword("anything") {
		t.Fatal("dummy comparison must never succeed")
	}
}
