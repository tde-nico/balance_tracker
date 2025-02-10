package routes

import (
	"balance/db"
	"balance/middleware"
	"net/http"
)

type TotalBalanceData struct {
	Data
	Balance           float32
	FakeBalance       float32
	NotedTransactions []db.NotedTransaction
}

func home(ctx *middleware.Ctx) {
	if ctx.RequestURI != "/" {
		ctx.AddFlash("404 Page not found")
		ctx.Redirect("/", http.StatusSeeOther)
		return
	}

	tmpl := getTemplate(ctx, "home")
	if tmpl == nil {
		return
	}

	balance, fakeBalance, err := db.GetTotalBalance()
	if err != nil {
		ctx.InternalError(err)
		return
	}

	transactions, err := db.GetNotedTransactions()
	if err != nil {
		ctx.InternalError(err)
		return
	}

	data := &TotalBalanceData{
		Data:              Data{User: ctx.User},
		Balance:           balance,
		FakeBalance:       fakeBalance,
		NotedTransactions: transactions,
	}
	executeTemplate(ctx, tmpl, data)
}
