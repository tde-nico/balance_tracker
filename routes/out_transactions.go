package routes

import (
	"balance/db"
	"balance/middleware"
)

type OutTransactionData struct {
	Data
	OutTransactions []db.Transaction
}

func outTransactionsGet(ctx *middleware.Ctx) {
	tmpl := getTemplate(ctx, "out_transactions")
	if tmpl == nil {
		return
	}

	outTransactions, err := db.GetAllOutTransactions()
	if err != nil {
		ctx.InternalError(err)
		return
	}

	data := &OutTransactionData{
		Data:            Data{User: ctx.User},
		OutTransactions: outTransactions,
	}
	executeTemplate(ctx, tmpl, data)
}
