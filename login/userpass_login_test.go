package login

import (
	"errors"
	"testing"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

type fakePassUser struct {
	ID uint
	UserPass
}

var fakePassUsers = map[string]*fakePassUser{}

func (*fakePassUser) FindUser(_ *gorm.DB, _ interface{}, account string) (interface{}, error) {
	u, ok := fakePassUsers[account]
	if !ok {
		return nil, gorm.ErrRecordNotFound
	}
	return u, nil
}

func TestHasUsablePasswordHash(t *testing.T) {
	valid, err := bcrypt.GenerateFromPassword([]byte("pw"), bcrypt.MinCost)
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		name   string
		user   interface{}
		expect bool
	}{
		{name: "valid hash", user: &fakePassUser{UserPass: UserPass{Password: string(valid)}}, expect: true},
		{name: "empty hash", user: &fakePassUser{}, expect: false},
		{name: "invalid hash", user: &fakePassUser{UserPass: UserPass{Password: "not-a-bcrypt-hash"}}, expect: false},
		{name: "model without hash getter", user: struct{}{}, expect: true},
	} {
		t.Run(c.name, func(t *testing.T) {
			if got := hasUsablePasswordHash(c.user); got != c.expect {
				t.Fatalf("hasUsablePasswordHash = %v, want %v", got, c.expect)
			}
		})
	}
}

// Every failing path of authUserPass must spend one bcrypt comparison, so that
// response timing does not reveal whether an account exists, is locked, or has
// no usable password (e.g. an OAuth-only account).
func TestAuthUserPassTimingEqualization(t *testing.T) {
	valid, err := bcrypt.GenerateFromPassword([]byte("right"), bcrypt.MinCost)
	if err != nil {
		t.Fatal(err)
	}
	now := time.Now()
	fakePassUsers = map[string]*fakePassUser{
		"normal": {ID: 1, UserPass: UserPass{Account: "normal", Password: string(valid)}},
		"locked": {ID: 2, UserPass: UserPass{Account: "locked", Password: string(valid), Locked: true, LockedAt: &now}},
		"oauth":  {ID: 3, UserPass: UserPass{Account: "oauth"}},
		"broken": {ID: 4, UserPass: UserPass{Account: "broken", Password: "garbage"}},
	}
	defer func() { fakePassUsers = map[string]*fakePassUser{} }()

	calls := 0
	orig := runDummyPasswordCompare
	runDummyPasswordCompare = func(password string) bool {
		calls++
		return false
	}
	defer func() { runDummyPasswordCompare = orig }()

	b := New().Secret("test-secret").UserModel(&fakePassUser{}).MaxRetryCount(0)
	for _, c := range []struct {
		account, password string
		expectErr         error
		expectDummy       int
	}{
		{account: "missing", password: "x", expectErr: ErrUserNotFound, expectDummy: 1},
		{account: "locked", password: "x", expectErr: ErrUserLocked, expectDummy: 1},
		{account: "oauth", password: "x", expectErr: ErrWrongPassword, expectDummy: 1},
		{account: "broken", password: "x", expectErr: ErrWrongPassword, expectDummy: 1},
		// real bcrypt comparison already spends the time
		{account: "normal", password: "x", expectErr: ErrWrongPassword, expectDummy: 0},
		{account: "normal", password: "right", expectErr: nil, expectDummy: 0},
	} {
		t.Run(c.account+"/"+c.password, func(t *testing.T) {
			calls = 0
			_, err := b.authUserPass(c.account, c.password)
			if err != c.expectErr {
				t.Fatalf("err = %v, want %v", err, c.expectErr)
			}
			if calls != c.expectDummy {
				t.Fatalf("dummy comparisons = %d, want %d", calls, c.expectDummy)
			}
		})
	}
}
