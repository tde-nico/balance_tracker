package routes

import (
	"balance/middleware"
	"net/http"
)

func logout(ctx *middleware.Ctx) {
	ctx.ExpireCookie()
	ctx.Redirect("/login", http.StatusSeeOther)
}
