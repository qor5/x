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
