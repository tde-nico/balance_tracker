package routes

import (
	"balance/db"
	"balance/log"
	"balance/middleware"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

var USERNAME_REGEX = regexp.MustCompile(`[0-9a-zA-Z_!@#€\-&+]{4,32}`)

func loginUser(ctx *middleware.Ctx, username, password string) (string, error) {
	username = strings.TrimSpace(username)
	if !USERNAME_REGEX.MatchString(username) {
		ctx.AddFlash("Invalid username")
		return "", fmt.Errorf("invalid username")
	}

	apiKey, err := db.LoginUser(username, password)
	if err != nil {
		ctx.AddFlash("Invalid username or password")
		return "", err
	}

	return apiKey, nil
}

func login_get(ctx *middleware.Ctx) {
	tmpl := getTemplate(ctx, "login")
	if tmpl == nil {
		return
	}

	if ctx.User != nil {
		ctx.Redirect("/", http.StatusSeeOther)
		return
	}

	executeTemplate(ctx, tmpl, &Data{})
}

func login_post(ctx *middleware.Ctx) {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	if ctx.User != nil {
		ctx.Redirect("/", http.StatusSeeOther)
		return
	}

	apiKey, err := loginUser(ctx, username, password)
	if err != nil {
		log.Errorf("Error logging in: %v", err)
		ctx.Redirect("/login", http.StatusSeeOther)
		return
	}

	ctx.SetSessionValue("apikey", apiKey)

	ctx.Redirect("/", http.StatusSeeOther)
}
