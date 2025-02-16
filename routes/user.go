package routes

import (
	"balance/db"
	"balance/log"
	"balance/middleware"
	"net/http"
)

type UserData struct {
	Data
	Player          *db.Player
	InTransactions  []db.Transaction
	OutTransactions []db.Transaction
}

func userInfo(ctx *middleware.Ctx) {
	tmpl := getTemplate(ctx, "user")
	if tmpl == nil {
		return
	}

	username := ctx.PathValue("username")

	player, err := db.GetPlayer(username)
	if err != nil {
		ctx.AddFlash("Error getting user")
		log.Errorf("Error getting user: %v", err)
		ctx.Redirect("/", http.StatusSeeOther)
		return
	}

	inTransactions, err := db.GetInTransactions(player.Username)
	if err != nil {
		ctx.InternalError(err)
		return
	}

	outTransactions, err := db.GetOutTransactions(player.Username)
	if err != nil {
		ctx.InternalError(err)
		return
	}

	data := &UserData{
		Data:            Data{User: ctx.User},
		Player:          player,
		InTransactions:  inTransactions,
		OutTransactions: outTransactions,
	}
	executeTemplate(ctx, tmpl, data)
}
