package routes

import (
	"balance/db"
	"balance/middleware"
)

type UsersData struct {
	Data
	Users []db.Player
}

func users(ctx *middleware.Ctx) {
	tmpl := getTemplate(ctx, "users")
	if tmpl == nil {
		return
	}

	players, err := db.GetPlayers()
	if err != nil {
		ctx.InternalError(err)
		return
	}

	data := &UsersData{
		Data:  Data{User: ctx.User},
		Users: players,
	}
	executeTemplate(ctx, tmpl, data)
}
