package middleware

import (
	"balance/db"
	"encoding/gob"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func InitStore(key []byte) {
	if len(key) != 32 {
		log.Fatal("Key must be 32 bytes")
	}

	store = sessions.NewCookieStore(key)

	gob.Register(&db.User{})
	gob.Register(&Flash{})

	store.Options.Path = "/"
	store.Options.Secure = false
	store.Options.SameSite = http.SameSiteDefaultMode
}
