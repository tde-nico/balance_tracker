package db

import (
	"balance/crypto_utils"
	"fmt"
)

type NotedTransaction struct {
	Note   string
	Amount float32
}

func GetTotalBalance() (float32, float32, error) {
	query, err := GetStatement("GetTotalBalance")
	if err != nil {
		return 0, 0, fmt.Errorf("error getting statement: %v", err)
	}

	rows, err := query.Query()
	if err != nil {
		return 0, 0, fmt.Errorf("error querying balance: %v", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return 0, 0, fmt.Errorf("error scanning balance: %v", err)
	}

	var balance, fakeBalance float32
	err = rows.Scan(&balance, &fakeBalance)
	if err != nil {
		return 0, 0, fmt.Errorf("error scanning balance: %v", err)
	}
	return crypto_utils.RoundToTwoDecimals(balance), crypto_utils.RoundToTwoDecimals(fakeBalance), nil
}

func GetNotedTransactions() ([]NotedTransaction, error) {
	query, err := GetStatement("GetNotedTransactions")
	if err != nil {
		return nil, fmt.Errorf("error getting statement: %v", err)
	}

	rows, err := query.Query()
	if err != nil {
		return nil, fmt.Errorf("error querying noted transactions: %v", err)
	}
	defer rows.Close()

	var transactions []NotedTransaction
	for rows.Next() {
		var t NotedTransaction
		err = rows.Scan(&t.Note, &t.Amount)
		if err != nil {
			return nil, fmt.Errorf("error scanning noted transactions: %v", err)
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}
