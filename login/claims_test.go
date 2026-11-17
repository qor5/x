package login

import (
	"reflect"
	"testing"

	"github.com/markbates/goth"
)

type testOAuthUser struct {
	ID uint
	OAuthInfo
}

func TestOAuthSessionClaims(t *testing.T) {
	ouser := goth.User{
		Provider:  "google",
		UserID:    "oid-1",
		Email:     "alice@example.com",
		Name:      "Alice",
		AvatarURL: "https://example.com/a.png",
	}

	t.Run("with user model, email/name are not carried in the session JWT", func(t *testing.T) {
		b := New().Secret("test-secret").UserModel(&testOAuthUser{})
		c := b.newOAuthSessionClaims(ouser, "42")
		if c.Provider != "google" || c.UserID != "42" || c.Subject != "42" {
			t.Fatalf("unexpected identity claims: %+v", c)
		}
		if c.Email != "" || c.Name != "" || c.Location != "" || c.IDToken != "" {
			t.Fatalf("PII leaked into session claims: %+v", c)
		}
		if c.AvatarURL != ouser.AvatarURL {
			t.Fatalf("avatar URL must be kept: %+v", c)
		}
	})

	t.Run("without user model, claims are the user and keep profile fields", func(t *testing.T) {
		b := New().Secret("test-secret")
		c := b.newOAuthSessionClaims(ouser, "oid-1")
		if c.Email != ouser.Email || c.Name != ouser.Name || c.AvatarURL != ouser.AvatarURL || c.UserID != "oid-1" {
			t.Fatalf("expected profile fields to be kept: %+v", c)
		}
	})
}

func TestUserClaimsStripProfile(t *testing.T) {
	c := UserClaims{Provider: "google", UserID: "1", Email: "e", Name: "n", AvatarURL: "a", Location: "l", IDToken: "t", PassUpdatedAt: "p", TOTPValidated: true}
	c.stripProfile()
	want := UserClaims{Provider: "google", UserID: "1", AvatarURL: "a", PassUpdatedAt: "p", TOTPValidated: true}
	if !reflect.DeepEqual(c, want) {
		t.Fatalf("got %+v, want %+v", c, want)
	}
}
