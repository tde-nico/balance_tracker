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
		err = rows.Scan(&t.ID, &t.From, &t.To, &t.Amount, &t.IsArrived, &note)
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

func GetAllInTransactions() ([]Transaction, error) {
	query, err := GetStatement("GetAllInTransactions")
	if err != nil {
		return nil, fmt.Errorf("error getting statement: %v", err)
	}

	rows, err := query.Query()
	if err != nil {
		return nil, fmt.Errorf("error querying transactions: %v", err)
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var t Transaction
		var note sql.NullString
		err = rows.Scan(&t.ID, &t.From, &t.To, &t.Amount, &t.IsArrived, &note)
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

func GetInTransaction(id string) (Transaction, error) {
	query, err := GetStatement("GetInTransaction")
	if err != nil {
		return Transaction{}, fmt.Errorf("error getting statement: %v", err)
	}

	var t Transaction
	var note sql.NullString
	err = query.QueryRow(id).Scan(&t.ID, &t.From, &t.To, &t.Amount, &t.IsArrived, &note)
	if err != nil {
		return Transaction{}, fmt.Errorf("error scanning transaction: %v", err)
	}
	if note.Valid {
		t.Note = note.String
	}
	return t, nil
}

func InsertInTransaction(t Transaction) error {
	query, err := GetStatement("InsertInTransaction")
	if err != nil {
		return fmt.Errorf("error getting statement: %v", err)
	}

	_, err = query.Exec(t.From, t.To, t.Amount, t.IsArrived, t.Note)
	if err != nil {
		return fmt.Errorf("error inserting transaction: %v", err)
	}
	return nil
}

func InsertInTransactionWithID(t Transaction) error {
	query, err := GetStatement("InsertInTransactionWithID")
	if err != nil {
		return fmt.Errorf("error getting statement: %v", err)
	}

	var note sql.NullString
	if t.Note == "" {
		note = sql.NullString{Valid: false}
	} else {
		note = sql.NullString{String: t.Note, Valid: true}
	}

	_, err = query.Exec(t.ID, t.From, t.To, t.Amount, t.IsArrived, note)
	if err != nil {
		return fmt.Errorf("error inserting transaction: %v", err)
	}
	return nil
}

func DeleteInTransaction(id int) error {
	query, err := GetStatement("DeleteInTransaction")
	if err != nil {
		return fmt.Errorf("error getting statement: %v", err)
	}

	_, err = query.Exec(id)
	if err != nil {
		return fmt.Errorf("error deleting transaction: %v", err)
	}
	return nil
}

func UpdateInTransaction(t Transaction) error {
	err := DeleteInTransaction(t.ID)
	if err != nil {
		return fmt.Errorf("error updating transaction: %v", err)
	}

	err = InsertInTransactionWithID(t)
	if err != nil {
		return fmt.Errorf("error updating transaction: %v", err)
	}
	return nil
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
		err = rows.Scan(&t.ID, &t.From, &t.To, &t.Purpose, &t.Amount, &t.IsArrived, &note)
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

func GetAllOutTransactions() ([]Transaction, error) {
	query, err := GetStatement("GetAllOutTransactions")
	if err != nil {
		return nil, fmt.Errorf("error getting statement: %v", err)
	}

	rows, err := query.Query()
	if err != nil {
		return nil, fmt.Errorf("error querying transactions: %v", err)
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var t Transaction
		var note sql.NullString
		err = rows.Scan(&t.ID, &t.From, &t.To, &t.Purpose, &t.Amount, &t.IsArrived, &note)
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

func GetOutTransaction(id string) (Transaction, error) {
	query, err := GetStatement("GetOutTransaction")
	if err != nil {
		return Transaction{}, fmt.Errorf("error getting statement: %v", err)
	}

	var t Transaction
	var note sql.NullString
	err = query.QueryRow(id).Scan(&t.ID, &t.From, &t.To, &t.Purpose, &t.Amount, &t.IsArrived, &note)
	if err != nil {
		return Transaction{}, fmt.Errorf("error scanning transaction: %v", err)
	}
	if note.Valid {
		t.Note = note.String
	}
	return t, nil
}

func InsertOutTransaction(t Transaction) error {
	query, err := GetStatement("InsertOutTransaction")
	if err != nil {
		return fmt.Errorf("error getting statement: %v", err)
	}

	_, err = query.Exec(t.From, t.To, t.Purpose, t.Amount, t.IsArrived, t.Note)
	if err != nil {
		return fmt.Errorf("error inserting transaction: %v", err)
	}
	return nil
}

func InsertOutTransactionWithID(t Transaction) error {
	query, err := GetStatement("InsertOutTransactionWithID")
	if err != nil {
		return fmt.Errorf("error getting statement: %v", err)
	}

	var note sql.NullString
	if t.Note == "" {
		note = sql.NullString{Valid: false}
	} else {
		note = sql.NullString{String: t.Note, Valid: true}
	}

	_, err = query.Exec(t.ID, t.From, t.To, t.Purpose, t.Amount, t.IsArrived, note)
	if err != nil {
		return fmt.Errorf("error inserting transaction: %v", err)
	}
	return nil
}

func DeleteOutTransaction(id int) error {
	query, err := GetStatement("DeleteOutTransaction")
	if err != nil {
		return fmt.Errorf("error getting statement: %v", err)
	}

	_, err = query.Exec(id)
	if err != nil {
		return fmt.Errorf("error deleting transaction: %v", err)
	}
	return nil
}

func UpdateOutTransaction(t Transaction) error {
	err := DeleteOutTransaction(t.ID)
	if err != nil {
		return fmt.Errorf("error updating transaction: %v", err)
	}

	err = InsertOutTransactionWithID(t)
	if err != nil {
		return fmt.Errorf("error updating transaction: %v", err)
	}
	return nil
}
