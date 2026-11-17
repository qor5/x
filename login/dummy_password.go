package login

import (
	"crypto/rand"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

// dummyPasswordCost mirrors the cost used by UserPass.EncryptPassword.
const dummyPasswordCost = 10

var dummyPasswordHash = sync.OnceValue(func() []byte {
	secret := make([]byte, 32)
	if _, err := rand.Read(secret); err != nil {
		panic(err)
	}
	h, err := bcrypt.GenerateFromPassword(secret, dummyPasswordCost)
	if err != nil {
		panic(err)
	}
	return h
})

// compareDummyPassword runs a bcrypt comparison against a random hash so that
// login attempts for unknown (or locked) accounts take roughly as long as
// attempts with a wrong password, closing an account enumeration side channel.
func compareDummyPassword(password string) bool {
	return bcrypt.CompareHashAndPassword(dummyPasswordHash(), []byte(password)) == nil
}

// runDummyPasswordCompare is what authUserPass calls; replaceable in tests.
var runDummyPasswordCompare = compareDummyPassword

// passwordHashGetter is optionally implemented by user models (UserPass does)
// to expose the stored password hash.
type passwordHashGetter interface {
	GetEncryptedPassword() string
}

// hasUsablePasswordHash reports whether checking a password against user will
// cost a real bcrypt comparison. It is false when the stored hash is empty or
// not a valid bcrypt hash (e.g. OAuth-only accounts), in which case bcrypt
// fails fast. Models that do not expose their hash are assumed usable.
func hasUsablePasswordHash(user interface{}) bool {
	g, ok := user.(passwordHashGetter)
	if !ok {
		return true
	}
	h := g.GetEncryptedPassword()
	if h == "" {
		return false
	}
	_, err := bcrypt.Cost([]byte(h))
	return err == nil
}
