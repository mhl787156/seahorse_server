package auth

import (
	"crypto/rand"
	"encoding/base64"
)

//The size in bytes of the session key to be generated
const keySize = 512 / 8

//The length of time a session is valid for, before the user is logged out
const sessionDuration = 604800

//Generates a Cryptographically secure random string
func getRandomString() string {
	bs := getRandomBytes()
	return base64.URLEncoding.EncodeToString(bs)
}

//Generates a Cryptographically secure random bytes to be used as a jwt code
func getRandomBytes() []byte {
	bs := make([]byte, keySize)
	_, err := rand.Read(bs)
	if err != nil {
		panic("Insufficient entropy. Abort. ABOOOORT!")
	}
	return bs
}
