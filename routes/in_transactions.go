package routes

import (
	"balance/db"
	"balance/log"
	"balance/middleware"
	"strconv"
)

type InTransactionsData struct {
	Data
	InTransactions []db.Transaction
}

type InTransactionData struct {
	Data
	Transaction db.Transaction
}

func inTransactionsGet(ctx *middleware.Ctx) {
	tmpl := getTemplate(ctx, "in_transactions")
	if tmpl == nil {
		return
	}

	inTransactions, err := db.GetAllInTransactions()
	if err != nil {
		ctx.InternalError(err)
		return
	}

	data := &InTransactionsData{
		Data:           Data{User: ctx.User},
		InTransactions: inTransactions,
	}
	executeTemplate(ctx, tmpl, data)
}

func inTransactionGet(ctx *middleware.Ctx) {
	tmpl := getTemplate(ctx, "in_transaction")
	if tmpl == nil {
		return
	}

	id := ctx.PathValue("id")
	inTransaction, err := db.GetInTransaction(id)
	if err != nil {
		ctx.InternalError(err)
		return
	}

	data := &InTransactionData{
		Data:        Data{User: ctx.User},
		Transaction: inTransaction,
	}
	executeTemplate(ctx, tmpl, data)
}

func inTransactionPut(ctx *middleware.Ctx) {
	tmpl := getSingleTemplate(ctx, "in_transaction")
	if tmpl == nil {
		return
	}

	idStr := ctx.PathValue("id")
	amountStr := ctx.FormValue("amount")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.InternalError(err)
		return
	}

	amount, err := strconv.ParseFloat(amountStr, 32)
	if err != nil {
		ctx.InternalError(err)
		return
	}

	t := db.Transaction{
		ID:        id,
		From:      ctx.FormValue("from"),
		To:        ctx.FormValue("to"),
		Amount:    float32(amount),
		IsArrived: ctx.FormValue("is_arrived") == "1",
		Note:      ctx.FormValue("note"),
	}

	log.Noticef("Updating Transaction: %+v", t)

	err = db.UpdateInTransaction(t)
	if err != nil {
		ctx.InternalError(err)
		return
	}

	data := &InTransactionData{
		Data:        Data{User: ctx.User},
		Transaction: t,
	}
	executeSingleTemplate(ctx, tmpl, data)
}
