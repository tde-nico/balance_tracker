package routes

import (
	"balance/log"
	"balance/middleware"
	"net/http"
)

func StartRouting(key []byte) {
	middleware.InitStore(key)

	static := http.FileServer(http.Dir("static"))
	http.Handle("GET /static/", http.StripPrefix("/static/", static))

	middleware.HandleFunc("GET /login", login_get)
	middleware.HandleFunc("POST /login", login_post)

	middleware.AuthHandleFunc("GET /", home)
	middleware.AuthHandleFunc("GET /logout", logout)
	middleware.AuthHandleFunc("GET /users", users)
	middleware.AuthHandleFunc("GET /user/{username}", userInfo)

	log.Notice("Serving on :8000")
	if err := http.ListenAndServe("0.0.0.0:8000", nil); err != nil {
		log.Fatalf("Error Serving: %v", err)
	}
}
