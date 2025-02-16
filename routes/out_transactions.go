package routes

import (
	"balance/db"
	"balance/log"
	"balance/middleware"
	"strconv"
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

func outTransactionGet(ctx *middleware.Ctx) {
	tmpl := getTemplate(ctx, "out_transaction")
	if tmpl == nil {
		return
	}

	id := ctx.PathValue("id")
	inTransaction, err := db.GetOutTransaction(id)
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

func outTransactionPut(ctx *middleware.Ctx) {
	tmpl := getSingleTemplate(ctx, "out_transaction")
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
		Purpose:   ctx.FormValue("purpose"),
		Amount:    float32(amount),
		IsArrived: ctx.FormValue("is_arrived") == "1",
		Note:      ctx.FormValue("note"),
	}

	log.Noticef("Updating Out Transaction: %+v", t)

	err = db.UpdateOutTransaction(t)
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
