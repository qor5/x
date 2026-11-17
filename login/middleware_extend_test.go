package login

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type extendTestUser struct {
	ID        uint `gorm:"primarykey"`
	DeletedAt gorm.DeletedAt
	OAuthInfo
}

// A session token issued by an older version carries email/name. When it is
// extended, the hook must receive the token exactly as the client sent it
// (downstream looks sessions up by token hash), while the new token must not
// carry the stripped PII.
func TestMiddlewareExtendLegacyOAuthToken(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{Logger: logger.Discard})
	if err != nil {
		t.Fatal(err)
	}
	if err := db.AutoMigrate(&extendTestUser{}); err != nil {
		t.Fatal(err)
	}
	if err := db.Create(&extendTestUser{ID: 7, OAuthInfo: OAuthInfo{OAuthProvider: "google", OAuthUserID: "oid"}}).Error; err != nil {
		t.Fatal(err)
	}

	for _, legacy := range []bool{true, false} {
		name := "current token"
		if legacy {
			name = "legacy token with PII"
		}
		t.Run(name, func(t *testing.T) {
			var gotOld, gotNew string
			var gotAvatar string
			b := New().Secret("test-secret").DB(db).UserModel(&extendTestUser{}).
				AutoExtendSession(true).
				AfterExtendSession(func(r *http.Request, user interface{}, extraVals ...interface{}) error {
					gotOld = extraVals[0].(string)
					gotNew = extraVals[1].(string)
					return nil
				})

			claims := UserClaims{
				Provider:         "google",
				UserID:           "7",
				AvatarURL:        "https://example.com/a.png",
				RegisteredClaims: b.genBaseSessionClaim("7"),
			}
			if legacy {
				claims.Email = "alice@example.com"
				claims.Name = "Alice"
			}
			// old enough to be extended
			claims.IssuedAt = jwt.NewNumericDate(time.Now().Add(-time.Duration(b.sessionMaxAge) * time.Second / 2))
			token := b.mustGetSessionToken(claims)

			h := b.Middleware()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotAvatar = GetCurrentUser(r).(OAuthUser).GetAvatar()
				w.WriteHeader(http.StatusOK)
			}))
			r := httptest.NewRequest(http.MethodGet, "/admin", nil)
			r.AddCookie(&http.Cookie{Name: b.authCookieName, Value: token})
			w := httptest.NewRecorder()
			h.ServeHTTP(w, r)

			if w.Code != http.StatusOK {
				t.Fatalf("status = %d", w.Code)
			}
			if gotOld != token {
				t.Fatalf("old session token passed to hook differs from the cookie value")
			}
			nc, err := parseClaims(&UserClaims{}, gotNew, b.secret)
			if err != nil {
				t.Fatal(err)
			}
			n := nc.(*UserClaims)
			if n.Email != "" || n.Name != "" {
				t.Fatalf("new token still carries PII: %+v", n)
			}
			if n.AvatarURL != claims.AvatarURL || n.UserID != "7" || n.Provider != "google" {
				t.Fatalf("new token lost identity/avatar: %+v", n)
			}
			if gotAvatar != claims.AvatarURL {
				t.Fatalf("avatar on current user = %q", gotAvatar)
			}
		})
	}
}
