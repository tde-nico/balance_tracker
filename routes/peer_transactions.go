package routes

import (
	"balance/db"
	"balance/log"
	"balance/middleware"
	"net/http"
	"strconv"
)

type PeerTransactionsData struct {
	Data
	PeerTransactions []db.Transaction
}

type PeerTransactionData struct {
	Data
	Transaction db.Transaction
}

func peerTransactionsGet(ctx *middleware.Ctx) {
	tmpl := getTemplates(ctx, "peers", "peer_row")
	if tmpl == nil {
		return
	}

	peerTransactions, err := db.GetAllPeerTransactions()
	if err != nil {
		ctx.InternalError(err)
		return
	}

	data := &PeerTransactionsData{
		Data:             Data{User: ctx.User},
		PeerTransactions: peerTransactions,
	}
	executeTemplate(ctx, tmpl, data)
}

func peerTransactionGet(ctx *middleware.Ctx) {
	tmpl := getTemplate(ctx, "peer")
	if tmpl == nil {
		return
	}

	id := ctx.PathValue("id")
	peerTransaction, err := db.GetPeerTransaction(id)
	if err != nil {
		log.Errorf("Error getting peer transaction %s: %v", id, err)
		ctx.AddFlash("404 Page not found")
		ctx.Redirect("/peer", http.StatusSeeOther)
		return
	}

	data := &PeerTransactionData{
		Data:        Data{User: ctx.User},
		Transaction: peerTransaction,
	}
	executeTemplate(ctx, tmpl, data)
}

func peerTransactionPost(ctx *middleware.Ctx) {
	tmpl := getSingleTemplate(ctx, "peer_row")
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

	log.Noticef("Creating Peer Transaction: %+v", t)

	err = db.InsertPeerTransaction(&t)
	if err != nil {
		ctx.InternalError(err)
		return
	}

	executeSingleTemplate(ctx, tmpl, "row", t)
}

func peerTransactionPut(ctx *middleware.Ctx) {
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

	log.Noticef("Updating Peer Transaction: %+v", t)

	err = db.UpdatePeerTransaction(t)
	if err != nil {
		ctx.InternalError(err)
		return
	}

	ctx.WriteHeader(http.StatusNoContent)
}

func peerTransactionDelete(ctx *middleware.Ctx) {
	idStr := ctx.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.InternalError(err)
		return
	}

	log.Warningf("Deleting peer Transaction: %d", id)

	err = db.DeletePeerTransaction(id)
	if err != nil {
		ctx.InternalError(err)
		return
	}

	ctx.WriteHeader(http.StatusNoContent)
}
