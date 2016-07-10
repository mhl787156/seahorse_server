package auth

import (
	"crypto/rand"
	"encoding/base64"
	"time"
        "github.com/mhl787156/seahorse_server/models"
        "github.com/mhl787156/seahorse_server/db"
)

//The size in bytes of the session key to be generated
const keySize = 512 / 8

//The length of time a session is valid for, before the user is logged out
const sessionDuration = 604800

//Generates a Cryptographically secure random string to be used as a session key
func getRandomkey() string {
	bs := make([]byte, keySize)
	_, err := rand.Read(bs)
	if err != nil {
		panic("Insufficient entropy. Abort. ABOOOORT!")
	}
	return base64.URLEncoding.EncodeToString(bs)
}

//Creates a new session with a randomly generated session key and a
//reasonable expiry time
func NewSession(uid string, rdb *db.DbConn) models.Session {
	var newSession models.Session
	newSession.UID = uid
	newSession.SessionKey = getRandomkey()
	newSession.Expires = time.Now().UTC().Unix() + sessionDuration

        rdb.WriteSession(newSession)
	return newSession
}
