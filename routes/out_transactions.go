package routes

import (
	"balance/db"
	"balance/log"
	"balance/middleware"
	"net/http"
	"strconv"
)

type OutTransactionData struct {
	Data
	OutTransactions []db.Transaction
}

func outTransactionsGet(ctx *middleware.Ctx) {
	tmpl := getTemplates(ctx, "outs", "out_row")
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
	tmpl := getTemplate(ctx, "out")
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

func outTransactionPost(ctx *middleware.Ctx) {
	tmpl := getSingleTemplate(ctx, "out_row")
	if tmpl == nil {
		return
	}

	amountStr := ctx.FormValue("amount")

	amount, err := strconv.ParseFloat(amountStr, 32)
	if err != nil {
		ctx.InternalError(err)
		return
	}

	t := db.Transaction{
		From:      ctx.FormValue("from"),
		To:        ctx.FormValue("to"),
		Purpose:   ctx.FormValue("purpose"),
		Amount:    float32(amount),
		IsArrived: ctx.FormValue("is_arrived") == "1",
		Note:      ctx.FormValue("note"),
	}

	log.Noticef("Creating Out Transaction: %+v", t)

	err = db.InsertOutTransaction(&t)
	if err != nil {
		ctx.InternalError(err)
		return
	}

	executeSingleTemplate(ctx, tmpl, "row", t)
}

func outTransactionPut(ctx *middleware.Ctx) {
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

	ctx.WriteHeader(http.StatusNoContent)
}

func outTransactionDelete(ctx *middleware.Ctx) {
	idStr := ctx.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.InternalError(err)
		return
	}

	log.Warningf("Deleting Out Transaction: %d", id)

	err = db.DeleteOutTransaction(id)
	if err != nil {
		ctx.InternalError(err)
		return
	}

	ctx.WriteHeader(http.StatusNoContent)
}
