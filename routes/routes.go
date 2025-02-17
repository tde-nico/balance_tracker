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
	middleware.AuthHandleFunc("GET /in", inTransactionsGet)
	middleware.AuthHandleFunc("GET /in/{id}", inTransactionGet)
	middleware.AuthHandleFunc("GET /out", outTransactionsGet)
	middleware.AuthHandleFunc("GET /out/{id}", outTransactionGet)

	middleware.EditorHandleFunc("POST /in", inTransactionPost)
	middleware.EditorHandleFunc("PUT /in/{id}", inTransactionPut)
	middleware.EditorHandleFunc("DELETE /in/{id}", inTransactionDelete)
	middleware.EditorHandleFunc("POST /out", outTransactionPost)
	middleware.EditorHandleFunc("PUT /out/{id}", outTransactionPut)
	middleware.EditorHandleFunc("DELETE /out/{id}", outTransactionDelete)

	log.Notice("Serving on :8000")
	if err := http.ListenAndServe("0.0.0.0:8000", nil); err != nil {
		log.Fatalf("Error Serving: %v", err)
	}
}
