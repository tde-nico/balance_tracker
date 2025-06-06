package routes

import (
	"balance/db"
	"balance/log"
	"balance/middleware"
	"net/http"
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
	tmpl := getTemplates(ctx, "ins", "in_row")
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
	tmpl := getTemplate(ctx, "in")
	if tmpl == nil {
		return
	}

	id := ctx.PathValue("id")
	inTransaction, err := db.GetInTransaction(id)
	if err != nil {
		log.Errorf("Error getting in transaction %s: %v", id, err)
		ctx.AddFlash("404 Page not found")
		ctx.Redirect("/in", http.StatusSeeOther)
		return
	}

	data := &InTransactionData{
		Data:        Data{User: ctx.User},
		Transaction: inTransaction,
	}
	executeTemplate(ctx, tmpl, data)
}

func inTransactionPost(ctx *middleware.Ctx) {
	tmpl := getSingleTemplate(ctx, "in_row")
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
		Amount:    float32(amount),
		IsArrived: ctx.FormValue("is_arrived") == "1",
		Note:      ctx.FormValue("note"),
	}

	log.Noticef("Creating In Transaction: %+v", t)

	err = db.InsertInTransaction(&t)
	if err != nil {
		ctx.InternalError(err)
		return
	}

	executeSingleTemplate(ctx, tmpl, "row", t)
}

func inTransactionPut(ctx *middleware.Ctx) {
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

	log.Noticef("Updating In Transaction: %+v", t)

	err = db.UpdateInTransaction(t)
	if err != nil {
		ctx.InternalError(err)
		return
	}

	ctx.WriteHeader(http.StatusNoContent)
}

func inTransactionDelete(ctx *middleware.Ctx) {
	idStr := ctx.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.InternalError(err)
		return
	}

	log.Warningf("Deleting In Transaction: %d", id)

	err = db.DeleteInTransaction(id)
	if err != nil {
		ctx.InternalError(err)
		return
	}

	ctx.WriteHeader(http.StatusNoContent)
}
