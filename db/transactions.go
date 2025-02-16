package db

import (
	"database/sql"
	"fmt"
)

func GetInTransactions(username string) ([]Transaction, error) {
	query, err := GetStatement("GetInTransactions")
	if err != nil {
		return nil, fmt.Errorf("error getting statement: %v", err)
	}

	rows, err := query.Query(username)
	if err != nil {
		return nil, fmt.Errorf("error querying transactions: %v", err)
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var t Transaction
		var note sql.NullString
		err = rows.Scan(&t.From, &t.To, &t.Amount, &t.IsArrived, &note)
		if err != nil {
			return nil, fmt.Errorf("error scanning transactions: %v", err)
		}
		if note.Valid {
			t.Note = note.String
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}

func GetOutTransactions(username string) ([]Transaction, error) {
	query, err := GetStatement("GetOutTransactions")
	if err != nil {
		return nil, fmt.Errorf("error getting statement: %v", err)
	}

	rows, err := query.Query(username, username)
	if err != nil {
		return nil, fmt.Errorf("error querying transactions: %v", err)
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var t Transaction
		var note sql.NullString
		err = rows.Scan(&t.From, &t.To, &t.Purpose, &t.Amount, &t.IsArrived, &note)
		if err != nil {
			return nil, fmt.Errorf("error scanning transactions: %v", err)
		}
		if note.Valid {
			t.Note = note.String
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}
